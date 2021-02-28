package utils

func CreateResponse(response interface{}) *HttpResponse {
	var data *HttpResponse

	switch response.(type) {
	case error:
		err := response.(error)
		data = &HttpResponse{
			Error:  &HttpError{Message: err.Error()},
			Status: GenericError,
		}
	default:
		data = &HttpResponse{
			Status: Success,
			Data:   response,
		}
	}
	return data
}
