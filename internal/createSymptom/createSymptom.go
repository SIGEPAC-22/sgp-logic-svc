package createSymptom

import "context"

type Repository interface {
	CreateSymptomRepo(ctx context.Context, NameSymptom string, DescriptionSymptom string) (bool, error)
}

type Service interface {
	CreateSymptomSvc(ctx context.Context, NameSymptom string, DescriptionSymptom string) (CreateSymptomResponse, error)
}

type CreateSymptomRequest struct {
	NameSymptom        string `json:"nameSymptom"`
	DescriptionSymptom string `json:"descriptionSymptom"`
}

type CreateSymptomResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
