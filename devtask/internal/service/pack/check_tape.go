package pack

import (
	"fmt"
)

type TapeVariant struct {
}

func (v TapeVariant) Check(_ int64) (int64, error) {
	fmt.Println("заказ упакован в пленку")
	return TapeCostKopecks, nil
}
