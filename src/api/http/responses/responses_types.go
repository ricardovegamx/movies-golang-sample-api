package responses

import "github.com/ricardovegamx/movies-api/src/api/domain"

/*
* Responses based on https://github.com/omniti-labs/jsend specification
 */
const (
	StatusSuccess = "success"
	StatusFail    = "fail"
	StatusError   = "error"
)

type SuccessResponseGetAll struct {
	Status string `json:"status"`
	Data   struct {
		Movies *[]domain.Movie `json:"movies"`
	} `json:"data"`
}

type SuccessResponseGet struct {
	Status string `json:"status"`
	Data   struct {
		Movie domain.Movie `json:"movie"`
	} `json:"data"`
}

type FailResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
