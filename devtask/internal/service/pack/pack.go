package pack

import (
	"devtask/internal/model"
	"errors"
)

type CheckVariant interface {
	Check(weight int64) (int64, error)
}

type Service struct {
	checkVariants map[model.PackageType]CheckVariant
}

func NewService(packageVariants map[model.PackageType]CheckVariant) *Service {
	return &Service{checkVariants: packageVariants}
}

// CheckWeight checks the weight of the order and return the additional cost depends on package type
func (s Service) CheckWeight(weight int64, pack model.PackageType) (int64, error) {
	existence := s.checkVariants[pack]
	if existence == nil {
		return 0, errors.New("такого типа упаковки нет")
	}
	return s.checkVariants[pack].Check(weight)
}
