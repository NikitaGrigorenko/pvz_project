package pack

import (
	"errors"
	"fmt"
)

type BoxVariant struct {
}

func (v BoxVariant) Check(orderWeight int64) (int64, error) {
	if orderWeight >= BoxWeightThreshold {
		return 0, errors.New("пакет не может быть использован при весе более 10кг")
	}
	fmt.Println("заказ упакован в коробку")
	return BoxCostKopecks, nil
}
