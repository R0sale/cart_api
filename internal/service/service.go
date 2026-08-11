package service

type service struct {
	repository repository
}

func NewService(rep repository) *service {
	return &service{
		repository: rep,
	}
}
