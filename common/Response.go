package common

const (
	GENERAL      = 0
	EXCEPTION    = 1
	UNAUTHORIZED = 401
)

type ResponseCode int

type ResponseModel[I any, O any] struct {
	Input          I `json:"input"`
	Output         O `json:"output"`
	ResponseCode   `json:"responseCode"`
	ExceptionModel `json:"exceptionModel"`
}

type ExceptionModel struct {
	Message       string `json:"message"`
	ExceptionCode int    `json:"code"`
	StackTrace    string `json:"stackTrace"`
}

func NewResponseModel[I any, O any]() *ResponseModel[I, O] {
	return &ResponseModel[I, O]{}
}

func NewExceptionModel(message string, exceptionCode int, stackTrace string) *ExceptionModel {
	return &ExceptionModel{message, exceptionCode, stackTrace}
}
