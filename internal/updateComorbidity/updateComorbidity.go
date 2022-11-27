package updateComorbidity

import "context"

type Repository interface {
	UpdateComorbidityRepo(ctx context.Context, Id int64, NameComorbidity string, DescriptionComorbidity string) (bool, error)
}

type Service interface {
	UpdateComorbiditySvc(ctx context.Context, Id string, NameComorbidity string, DescriptionComorbidity string) (UpdateComorbidityResponse, error)
}

type UpdateComorbidityRequest struct {
	Id                     string `json:"id"`
	NameComorbidity        string `json:"nameComorbidity"`
	DescriptionComorbidity string `json:"descriptionComorbidity"`
}

type UpdateComorbidityResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
