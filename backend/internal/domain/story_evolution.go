package domain

type StoryEvolutionCondition struct {
	PaginateRequest
	ProjectID uint64 `json:"project_id"`
}
