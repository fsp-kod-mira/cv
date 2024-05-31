package domain

type CV struct {
	Id           string
	Status       string
	FileId       string
	UploadedById string
	TotalScore   string
	Features     map[string]interface{}
}
