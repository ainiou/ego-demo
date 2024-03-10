package example

type Option struct {
}

type Service struct {
	*Option
}

func NewService(opt *Option) *Service {
	return &Service{opt}
}
