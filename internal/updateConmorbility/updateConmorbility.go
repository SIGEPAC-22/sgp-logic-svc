package updateConmorbility

import "context"

type Repository interface {
	UpdateConmorbilityRepo(ctx context.Context, Id int64, NameConmorbility string, DescriptionConmorbility string) (bool, error)
}

type Service interface {
	UpdateConmorbilitySvc(ctx context.Context, Id string, NameConmorbility string, DescriptionConmorbility string) (UpdateUpdateConmorbilityResponse, error)
}

type UpdateUpdateConmorbilityRequest struct {
	Id                      string `json:"id"`
	NameConmorbility        string `json:"name_conmorbility"`
	DescriptionConmorbility string `json:"descripcion_conmorbility"`
}

type UpdateUpdateConmorbilityResponse struct {
	ResponseCode int    `json:"response_code"`
	Messagge     string `json:"messagge"`
}
