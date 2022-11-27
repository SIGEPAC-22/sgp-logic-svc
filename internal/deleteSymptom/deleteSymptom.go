package deleteSymptom

import "context"

type Repository interface {
	DeleteSymptomRepo(ctx context.Context, Id int64) (bool, error)
}

type Service interface {
	DeleteSymptomSvc(ctx context.Context, Id string) (DeleteSymptomResponse, error)
}

type DeleteSymptomRequest struct {
	Id string `json:"id"`
}

type DeleteSymptomResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
