package createConmorbility

import "context"

type Repository interface {
	CreateConmorbilityRepo(ctx context.Context, NameConmorbility string, DescriptionConmorbility string) (bool, error)
}

type Service interface {
	CreateConmorbilitySvc(ctx context.Context, NameConmorbility string, DescriptionConmorbility string) (CreateConmorbilityResponse, error)
}

type CreateConmorbilityRequest struct {
	NameConmorbility        string `json:"name_conmorbility"`
	DescriptionConmorbility string `json:"description_conmorbility"`
}

type CreateConmorbilityResponse struct {
	ResponseCode int    `json:"response_code"`
	Messagge     string `json:"messagge"`
}
