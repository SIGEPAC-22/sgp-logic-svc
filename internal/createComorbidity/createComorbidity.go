package createComorbidity

import "context"

type Repository interface {
	CreateComorbidityRepo(ctx context.Context, NameComorbidity string, DescriptionComorbidity string) (bool, error)
}

type Service interface {
	CreateComorbiditySvc(ctx context.Context, NameComorbidity string, DescriptionComorbidity string) (CreateComorbidityResponse, error)
}

type CreateComorbidityRequest struct {
	NameComorbidity        string `json:"nameComorbidity"`
	DescriptionComorbidity string `json:"descriptionComorbidity"`
}

type CreateComorbidityResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
