package pack

import (
	"errors"
	"fmt"
)

type PacketVariant struct {
}

func (v PacketVariant) Check(orderWeight int64) (int64, error) {
	if orderWeight >= PacketWeightThreshold {
		return 0, errors.New("пакет не может быть использован при весе более 10кг")
	}
	fmt.Println("заказ упакован в пакет")
	return PacketCostKopecks, nil
}
