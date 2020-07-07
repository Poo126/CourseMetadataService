package model

type Course struct {
	Name  string `json:"name"`
	AuthorId string `json:"authorId"`
	Version string `json:"version"`
	State string `json:"state"`
}
