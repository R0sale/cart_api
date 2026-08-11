package service

type cartService struct {
	repository repository
}

func NewService(rep repository) *cartService {
	return &cartService{
		repository: rep,
	}
}
