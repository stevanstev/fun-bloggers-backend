package model

/*
	Minimal Response of API
*/
type BaseResponse struct {
	StatusCode int `json: "statusCode"`
	StatusMessage string `json: "statusMessage"`
}

/*
	Call this function to set default value of BaseResponse
	@message string, "messages to set on default API's response" 
*/
func (baseResponse *BaseResponse) GetDefault(message string) {
	baseResponse.StatusCode = 200
	baseResponse.StatusMessage = message
}