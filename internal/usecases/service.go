package usecases

type CvsService struct {
	CvsRepository
}

func NewCvsService(cvsRepo CvsRepository) CvsUsecases {
	return &CvsService{
		cvsRepo,
	}
}
