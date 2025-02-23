package collection

import "github.com/devshansharma/luke/internal/request"

type Collection struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Folders     []Folder               `json:"folders"`
	Requests    []request.Request      `json:"requests"`
	Variables   map[string]interface{} `json:"variables"`
}

type Folder struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Requests    []request.Request `json:"requests"`
}
