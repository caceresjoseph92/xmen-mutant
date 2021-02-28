package utils

type HttpResponse struct {
	Environment string      `json:"environment"`
	Status      StatusType  `json:"status"`
	Data        interface{} `json:"data"`
	Error       *HttpError  `json:"error"`
}

type HttpError struct {
	Message string `json:"message"`
}

type StatusType string

const (
	GenericError   = "Error"
	Warning        = "Warning"
	Success        = "Success"
	BusinessError  = "BusinessError"
	TechnicalError = "TechnicalError"
	InvalidFormat  = "InvalidFormat"
)
