package schemas

type DatabasesErrorSchema struct {
	Type string
	Code int
}

type ErrorResponseSchema struct {
	StatusCode int         `json:"code"`
	Error      interface{} `json:"error"`
}

type UnauthorizedErrorSchema struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
