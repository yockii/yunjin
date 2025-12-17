package domain

type ProjectCondition struct {
	PaginateRequest
	Title   string `query:"title"`
	OwnerID uint64 `query:"owner_id"`
}
