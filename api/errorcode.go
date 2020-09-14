package main

type ErrorCode struct {
	StatusCode int        `json:"statusCode"`
	Message    string     `json:"message"`
	Data       string     `json:"data"`
	InnerError *ErrorCode `json:"innerError"`
}
