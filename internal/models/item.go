package models

type Script struct {
	Exec     []string `json:"exec,omitempty"`
	Type     string   `json:"type,omitempty"`
	Packages []string `json:"packages,omitempty"`
}

type Event struct {
	Listen string `json:"listen,omitempty"`
	Script Script `json:"script,omitempty"`
}

type Header struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Body struct {
	Mode string `json:"mode,omitempty"`
	Raw  string `json:"raw,omitempty"`
}

type Request struct {
	Method string   `json:"method,omitempty"`
	Header []Header `json:"header,omitempty"`
	Body   Body     `json:"body,omitempty"`
	URL    string   `json:"url,omitempty"`
}

type Item struct {
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Events  []Event  `json:"events,omitempty"`
	Request Request  `json:"request,omitempty"`
	Reponse []string `json:"response,omitempty"`
	IsValid bool     `json:"is_valid"`
}
