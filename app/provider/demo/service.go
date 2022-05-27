package demo

import "github.com/go-colon/colon/core"

type Service struct {
	container core.Container
}

func NewService(params ...interface{}) (interface{}, error) {
	container := params[0].(core.Container)
	return &Service{container: container}, nil
}

func (s *Service) GetAllStudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "foo",
		},
		{
			ID:   2,
			Name: "bar",
		},
	}
}
