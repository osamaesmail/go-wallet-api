package api

type ErrorResponse struct {
	Errors  []Error     `json:"errors"`
	Links   interface{} `json:"links,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Jsonapi interface{} `json:"jsonapi,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Source  string `json:"source,omitempty"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}
