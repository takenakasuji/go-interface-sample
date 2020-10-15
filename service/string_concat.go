package service

type ConcatService interface {
	DoSomething() string
}

type concatService struct {
	a string
	b string
}

func NewConcatService(a, b string) ConcatService {
	return &concatService{a: a, b: b}
}

func (x *concatService) DoSomething() string {
	return x.a + x.b
}
