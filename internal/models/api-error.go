package models

type ApiError struct {
	Code    int    `json:"errCode"`
	Message string `json:"errMessage"`
}

func (apiError *ApiError) Error() string {
	return apiError.Message
}
