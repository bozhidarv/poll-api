package models

type PollFilters struct {
	UserIds  *string `json:"userId,omitempty"`
	Category *string `json:"category,omitempty"`
}

func (a PollFilters) Empty() bool {
	return a.Category == nil && a.UserIds == nil
}
