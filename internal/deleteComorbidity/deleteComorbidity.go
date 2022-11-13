package deleteComorbidity

import "context"

type Repository interface {
	DeleteComorbidityRepo(ctx context.Context, Id int64) (bool, error)
}

type Service interface {
	DeleteComorbiditySvc(ctx context.Context, Id string) (DeleteComorbidityResponse, error)
}

type DeleteComorbidityRequest struct {
	Id string `json:"id"`
}

type DeleteComorbidityResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
