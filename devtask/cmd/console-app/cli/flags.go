package cli

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseTakeOrderFlags parses the command-line flags for take_order command
func ParseTakeOrderFlags(args []string) (int64, int64, time.Time, int64, int64, int64, error) {
	takeCmd := flag.NewFlagSet("take_order", flag.ExitOnError)
	orderID := takeCmd.Int64("order_id", 0, "id of an order")
	clientID := takeCmd.Int64("client_id", 0, "id of a client")
	shelfLife := takeCmd.String("date_exp", "2006-01-02", "shelf life of an order")
	weight := takeCmd.Float64("weight", 0.0, "weight of the order, e.g. 2.64kg")
	price := takeCmd.Float64("price", 0.0, "price of the order, e.g. 200.01rub")
	packType := takeCmd.Int64("package", 0, "type of the package")

	err := takeCmd.Parse(args)

	priceKopecks := int64(*price * 100)
	weightGrams := int64(*weight * 1000)

	if err != nil {
		return 0, 0, time.Time{}, 0, 0, 0, err
	}

	if *orderID == 0 {
		return 0, 0, time.Time{}, 0, 0, 0, errors.New("необходимо указать номер заказа. Номер заказа должен быть > 0")
	}

	if *clientID == 0 {
		return 0, 0, time.Time{}, 0, 0, 0, errors.New("необходимо указать номер клиента. Номер клиента должен быть > 0")
	}

	if *shelfLife == "2006-01-02" {
		return 0, 0, time.Time{}, 0, 0, 0, errors.New("необходимо указать актуальный срок хранения заказа")
	}

	if *weight == 0.0 {
		return 0, 0, time.Time{}, 0, 0, 0, errors.New("необходимо указать вес заказа. Вес заказа должен быть > 0")
	}

	if *price == 0.0 {
		return 0, 0, time.Time{}, 0, 0, 0, errors.New("необходимо указать стоимость заказа. Стоимость заказа должна быть > 0")
	}

	if *packType < 0 || *packType > 2 {
		return 0, 0, time.Time{}, 0, 0, 0, errors.New("необходимо указать тип упаковки. Тип упаковки может быть 0, 1 или 2")
	}

	date, err := time.Parse("2006-01-02", *shelfLife)
	if err != nil {
		return 0, 0, time.Time{}, 0, 0, 0, err
	}

	return *orderID, *clientID, date, weightGrams, priceKopecks, *packType, nil
}

// ParseReturnOrderFlags parses the command-line flags for return_order command
func ParseReturnOrderFlags(args []string) (int64, error) {
	returnCmd := flag.NewFlagSet("return_order", flag.ExitOnError)
	orderID := returnCmd.Int64("order_id", 0, "id of an order")

	err := returnCmd.Parse(args)
	if err != nil {
		return 0, err
	}

	if *orderID == 0 {
		return 0, errors.New("необходимо указать номер заказа. Номер заказа должен быть > 0")
	}

	return *orderID, nil
}

// ParseGiveOrderFlags parses the command-line flags for give_order command
func ParseGiveOrderFlags(args []string) ([]int64, error) {
	giveCmd := flag.NewFlagSet("give_order", flag.ExitOnError)
	sliceOrders := giveCmd.String("slice", "1", "slice of orders")

	err := giveCmd.Parse(args)
	if err != nil {
		return nil, err
	}

	stringSlice := strings.Split(*sliceOrders, ",")
	intSlice := make([]int64, len(stringSlice))
	for i, s := range stringSlice {
		value, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("ошибка во время конвертации слайса. Вводите номер заказов через запятую (1,2,3):")
			return nil, err
		}
		intSlice[i] = int64(value)
	}
	return intSlice, nil
}

// ParseListOrdersFlags parses the command-line flags for list_orders command
func ParseListOrdersFlags(args []string) (int64, int64, bool, error) {
	listOrdersCmd := flag.NewFlagSet("list_orders", flag.ExitOnError)
	clientID := listOrdersCmd.Int64("client_id", 0, "id of a client")
	lastN := listOrdersCmd.Int64("last_n", 0, "number of last n orders")
	inPP := listOrdersCmd.Bool("inpp", false, "orders in Pickup Point or not")

	err := listOrdersCmd.Parse(args)
	if err != nil {
		return 0, 0, false, err
	}
	if *clientID == 0 {
		return 0, 0, false, errors.New("необходимо указать номер клиента. Номер клиента должен быть > 0")
	}

	return *clientID, *lastN, *inPP, nil
}

// ParseClientRefundFlags parses the command-line flags for client_refund command
func ParseClientRefundFlags(args []string) (int64, int64, error) {
	refundCmd := flag.NewFlagSet("client_refund", flag.ExitOnError)
	orderID := refundCmd.Int64("order_id", 0, "id of an order")
	clientID := refundCmd.Int64("client_id", 0, "id of a client")

	err := refundCmd.Parse(args)
	if err != nil {
		return 0, 0, err
	}

	if *orderID == 0 {
		return 0, 0, errors.New("необходимо указать номер заказа. Номер заказа должен быть > 0")
	}
	if *clientID == 0 {
		return 0, 0, errors.New("необходимо указать номер клиента. Номер клиента должен быть > 0")
	}

	return *orderID, *clientID, nil
}

// ParseListRefundsFlags parses the command-line flags for list_refunds command
func ParseListRefundsFlags(args []string) (int64, int64, error) {
	listRefundsCmd := flag.NewFlagSet("list_refunds", flag.ExitOnError)
	pageNumber := listRefundsCmd.Int64("page_number", 1, "number of page to show")
	onPage := listRefundsCmd.Int64("on_page", 1, "how many items to show on a page")

	err := listRefundsCmd.Parse(args)
	if err != nil {
		return 0, 0, err
	}

	return *pageNumber, *onPage, nil
}
