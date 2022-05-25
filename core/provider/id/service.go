package id

import (
	"github.com/rs/xid"
)

type ColonIDService struct {
}

func NewColonIDService(params ...interface{}) (interface{}, error) {
	return &ColonIDService{}, nil
}

func (s *ColonIDService) NewID() string {
	return xid.New().String()
}
