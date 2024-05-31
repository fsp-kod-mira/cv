package domain

type CV struct {
	Id           string `json:"id"`
	Status       string `json:"status"`
	FileId       string `json:"fileId"`
	UploadedById string `json:"uploadedById"`
	TotalScore   string `json:"totalScore"`
	Features     []*FullfieldFeature
}
