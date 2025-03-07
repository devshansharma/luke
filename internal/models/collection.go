package models

type Info struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Folder struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Item []Item `json:"item"`
}

type Collection struct {
	Info    Info     `json:"info"`
	Folders []Folder `json:"folders,omitempty"`
	Items   []Item   `json:"items,omitempty"`
}
