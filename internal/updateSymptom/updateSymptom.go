package updateSymptom

import "context"

type Repository interface {
	UpdateSymptomRepo(ctx context.Context, Id int64, NameSymptom string, DescriptionSymptom string) (bool, error)
}

type Service interface {
	UpdateSymptomSvc(ctx context.Context, Id string, NameSymptom string, DescriptionSymptom string) (UpdateSymptomResponse, error)
}

type UpdateSymptomRequest struct {
	Id                 string `json:"id"`
	NameSymptom        string `json:"nameSymptom"`
	DescriptionSymptom string `json:"descriptionSymptom"`
}

type UpdateSymptomResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
