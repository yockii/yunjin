package domain

type TagCondition struct {
	PaginateRequest
	Name string `query:"name"`
}
