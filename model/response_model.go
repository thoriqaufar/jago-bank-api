package model

type ResponseWithData struct {
	StatusCode int    `json:"code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type ResponseWithoutData struct {
	StatusCode int    `json:"code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

type ResponseParams struct {
	StatusCode int
	Message    string
	Data       any
}

func Response(params ResponseParams) any {
	var response any
	var status string

	if params.StatusCode >= 200 && params.StatusCode <= 299 {
		status = "success"
	} else {
		status = "failed"
	}

	if params.Data != nil {
		response = &ResponseWithData{
			StatusCode: params.StatusCode,
			Status:     status,
			Message:    params.Message,
			Data:       params.Data,
		}
	} else {
		response = &ResponseWithoutData{
			StatusCode: params.StatusCode,
			Status:     status,
			Message:    params.Message,
		}
	}

	return response
}
