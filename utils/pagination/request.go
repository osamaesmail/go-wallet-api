package pagination

type Request struct {
	Page    int `json:"page" validate:"omitempty,gte=1"`
	PerPage int `json:"per_page" validate:"omitempty,gte=1,lte=100"`
}

func (req Request) GetPage() int {
	if req.Page == 0 {
		return 1
	}
	return req.Page
}

func (req Request) GetPerPage() int {
	if req.PerPage == 0 {
		return 10
	}
	return req.PerPage
}

func (req Request) GetOffset() int {
	return req.GetPerPage() * (req.GetPage() - 1)
}
