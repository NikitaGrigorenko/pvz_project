package cli

import (
	"context"
	"devtask/internal/model"
	"devtask/internal/pkg/utils"
	"devtask/internal/service/order"
	"devtask/internal/service/pvz"
	"fmt"
)

// ServicesCases handles commands from user
func ServicesCases(ctx context.Context, args []string, orderService *order.Service, pvzService *pvz.Service) error {
	switch args[0] {
	case "take_order":
		orderID, clientID, date, weight, expense, packetType, err := ParseTakeOrderFlags(args[1:])
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = orderService.TakeOrder(model.OrderInput{
			IdOrder:     orderID,
			IdClient:    clientID,
			ExpireDate:  date,
			WeightGrams: weight,
			PriceKopeck: expense,
			PackageType: model.PackageType(packetType),
		})
		if err != nil {
			return err
		}

		fmt.Println("Заказ успешно принят!")
	case "return_order":
		orderID, err := ParseReturnOrderFlags(args[1:])
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = orderService.ReturnOrder(orderID)
		if err != nil {
			return err
		}
		fmt.Println("Заказ успешно возвращен курьеру!")
	case "give_order":
		sliceOrders, err := ParseGiveOrderFlags(args[1:])
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = orderService.GiveOrder(sliceOrders)

		if err != nil {
			return err
		}
		fmt.Println("Заказ успешно выдан клиенту!")
	case "list_orders":
		clientID, lastN, inPP, err := ParseListOrdersFlags(args[1:])
		if err != nil {
			fmt.Println(err)
			return err
		}
		list, err := orderService.ListOrders(clientID, lastN, inPP)
		if err != nil {
			return err
		}
		for _, orderInfo := range list {
			fmt.Printf("%+v\n", orderInfo)
		}
	case "client_refund":
		orderID, clientID, err := ParseClientRefundFlags(args[1:])
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = orderService.ClientRefund(orderID, clientID)
		if err != nil {
			return err
		}
		fmt.Println("Заказ успешно возвращен на пвз!")
	case "list_refunds":
		pageNumber, onPage, err := ParseListRefundsFlags(args[1:])
		if err != nil {
			fmt.Println(err)
			return err
		}

		list, err := orderService.ListRefunds(pageNumber, onPage)
		if err != nil {
			return err
		}
		for _, orderInfo := range list {
			fmt.Printf("%+v\n", orderInfo)
		}
	case "pvz":
		err := Run(ctx, pvzService)
		if err != nil {
			return err
		}
	case "help":
		utils.PrintHelp()
	default:
		fmt.Println("Неизвестная команда!")
	}
	return nil
}
