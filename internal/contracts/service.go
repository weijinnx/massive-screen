package contracts

type Service interface {
	Name() string
}

type AService struct {
	name string
}

func NewService(name string) *AService {
	return &AService{name: name}
}

func (s *AService) Name() string {
	return s.name
}
