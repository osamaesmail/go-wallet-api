package transaction

import "math"

type IService interface {
	Create(req CreateRequest) (ResponseDTO, error)
	List(req ListRequest) (ResponseListDTO, error)
}

type Service struct {
	mapper Mapper
	repo   IRepository
}

func NewService(repo IRepository, mapper Mapper) Service {
	return Service{repo: repo, mapper: mapper}
}

func (s Service) Create(req CreateRequest) (ResponseDTO, error) {
	model := s.mapper.CreateRequestToModel(req)
	resp, err := s.repo.Create(model)

	return s.mapper.ModelToDTO(resp), err
}

func (s Service) List(req ListRequest) (ResponseListDTO, error) {
	list, total, err := s.repo.List(req)

	resp := s.mapper.ModelsDTOToListDTO(list)
	resp.PerPage = req.GetPerPage()
	resp.Page = req.GetPage()
	resp.TotalPages = int(math.Ceil(float64(total) / float64(req.GetPerPage())))
	resp.TotalRecords = int(total)

	return resp, err
}
