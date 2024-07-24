package model

type DeleteRequest struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}
