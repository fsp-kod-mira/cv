package usecases

import (
	"context"
	document_parser "cv/api/document-parser"
	"cv/api/features"
	"cv/internal/domain"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"github.com/nats-io/nats.go"
	"github.com/samber/lo"
)

type CvsService struct {
	CvsRepository
	featuresClient features.FeatureClient
	nats           *nats.Conn
	dpClient       document_parser.DocumentClient
}

func NewCvsService(cvsRepo CvsRepository, featuresClient features.FeatureClient, nats *nats.Conn, dp document_parser.DocumentClient) CvsUsecases {
	return &CvsService{
		CvsRepository:  cvsRepo,
		featuresClient: featuresClient,
		nats:           nats,
		dpClient:       dp,
	}
}

func (s *CvsService) Upload(ctx context.Context, uploadedById string, fileId string) (string, error) {
	js, _ := s.nats.JetStream()
	cvId, err := s.CvsRepository.Upload(ctx, uploadedById, fileId)
	if err != nil {
		slog.Error("cannot upload cvs service err while exec repo")

		return "", err
	}
	stream, err := s.dpClient.ExtractTextFromDocument(ctx, &document_parser.ExtractRequest{
		FileUrl: fileId,
	})
	var data []byte
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			stream.CloseSend()
			break
		}
		if err != nil {
			return "", err
		}
		data = append(data, req.FileContent...)
	}

	t := struct {
		CvID string `json:"cvId"`
		Text string `json:"text"`
	}{
		CvID: cvId,
		Text: string(data),
	}
	jdata, err := json.Marshal(t)
	if err != nil {
		slog.Error("cannot marshal json")
		return "", err
	}

	_, err = js.Publish("cv.new", jdata)
	if err != nil {
		slog.Error("cannot publish message to cv.new")
		return "", err
	}

	return cvId, nil
}

func (s *CvsService) GetOne(ctx context.Context, id string) (*domain.CV, error) {
	cv, err := s.CvsRepository.GetOne(ctx, id)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("error while get cv by id %d", err)
	}

	fs, err := s.GetFeaturesByCvId(ctx, id)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("error while get cv features by id %d", err)
	}

	domainCv := &domain.CV{
		Id:           id,
		Status:       cv.Status,
		FileId:       cv.FileId,
		UploadedById: cv.UploadedById,
		TotalScore:   cv.TotalScore,
	}

	for _, f := range fs {
		parsedFeature, err := s.featuresClient.GetFeaturesById(ctx, &features.IdStruct{
			Id: uint64(f.FeatureId),
		})
		if err != nil {
			continue
		}

		domainCv.Features = append(domainCv.Features, &domain.FullfieldFeature{
			FeatureName:  parsedFeature.FeatureName,
			PriorityName: parsedFeature.PriorityName,
			Coefficient:  parsedFeature.Coefficient,
		})
	}

	return domainCv, nil
}

func (s *CvsService) GetAll(ctx context.Context, limit, offset int) ([]*domain.CV, error) {
	cvsWithoutFields, err := s.CvsRepository.GetAll(ctx, limit, offset)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("errro while get all cvs")
	}

	return lo.Map(cvsWithoutFields, func(item *domain.CV, index int) *domain.CV {
		fs, err := s.GetFeaturesByCvId(ctx, item.Id)
		if err != nil {
			return nil
		}

		item.Features = make([]*domain.FullfieldFeature, 0)

		for _, f := range fs {
			parsedFeature, err := s.featuresClient.GetFeaturesById(ctx, &features.IdStruct{
				Id: uint64(f.FeatureId),
			})
			if err != nil {
				continue
			}

			item.Features = append(item.Features, &domain.FullfieldFeature{
				FeatureName:  parsedFeature.FeatureName,
				PriorityName: parsedFeature.PriorityName,
				Coefficient:  parsedFeature.Coefficient,
			})
		}

		return item
	}), nil
}
