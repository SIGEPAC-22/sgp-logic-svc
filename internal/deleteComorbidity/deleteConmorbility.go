package deleteComorbidity

import "context"

type Repository interface {
	DeleteConmorbilityRepo(ctx context.Context, Id int64) (bool, error)
}

type Service interface {
	DeleteConmorbilitySvc(ctx context.Context, Id string) (DeleteConmorbilityResponse, error)
}

type DeleteConmorbilityRequest struct {
	Id string `json:"id"`
}

type DeleteConmorbilityResponse struct {
	ResponseCode int    `json:"response_code"`
	Messagge     string `json:"messagge"`
}
