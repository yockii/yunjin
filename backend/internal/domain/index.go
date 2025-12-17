package domain

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type PaginateRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PaginateResponse struct {
	Response
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}

func NewPaginateResponse(code int, data any, msg string, total int64, page, limit int) PaginateResponse {
	return PaginateResponse{
		Response: Response{
			Code: code,
			Data: data,
			Msg:  msg,
		},
		Total: total,
		Page:  page,
		Limit: limit,
	}
}

func NewResponse(code int, data any, msg string) Response {
	return Response{
		Code: code,
		Data: data,
		Msg:  msg,
	}
}
