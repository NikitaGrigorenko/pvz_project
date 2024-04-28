package order

import (
	"devtask/internal/model"
	"errors"
	"time"
)

type storage interface {
	TakeOrder(todo model.OrderInput) error
	ReturnOrder(id int64) error
	GiveOrder(slice []int64) error
	ListOrders(slice []int64, lastN int64) ([]model.Order, error)
	ClientRefund(orderId int64) error
	ListRefunds(slice []int64, pageNumber int64, onPage int64) ([]model.Order, error)
	ListAll() map[int64]model.Order
}

type packService interface {
	CheckWeight(weight int64, pack model.PackageType) (int64, error)
}

type Service struct {
	s     storage
	packS packService
}

func NewService(s storage, packS packService) *Service {
	return &Service{s: s, packS: packS}
}

// TakeOrder creates order in the storage
func (s Service) TakeOrder(input model.OrderInput) error {
	all := s.s.ListAll()
	for _, order := range all {
		if order.IdOrder == input.IdOrder && !order.IsReturned {
			return errors.New("невозможно принять заказ, так как такой заказ уже существует")
		}
	}

	costAdd, err := s.packS.CheckWeight(input.WeightGrams, model.PackageType(input.PackageType))
	if err != nil {
		return err
	}

	input.PriceKopeck = input.PriceKopeck + costAdd

	if input.IdClient < 1 || input.IdOrder < 1 {
		return errors.New("id заказа или клиента не может быть меньше 1")
	}

	if input.ExpireDate.Before(time.Now()) {
		return errors.New("невозможно принять заказ, потому что срок хранения в прошлом")
	}

	return s.s.TakeOrder(input)
}

// ReturnOrder sets IsReturned flag true
func (s Service) ReturnOrder(id int64) error {
	all := s.s.ListAll()
	if id < 1 {
		return errors.New("id заказа не может быть меньше 1")
	}

	existOrder := false
	for _, order := range all {
		if order.IdOrder != id {
			continue
		}
		if order.IsReturned {
			return errors.New("невозможно вернуть заказ курьеру, который уже был возвращен")
		}
		// We can return an order to courier only if it is not given to client and has not expired
		if order.IsGiven || !order.ExpireDate.Before(time.Now()) {
			return errors.New("невозможно вернуть заказ курьеру, который был получен или его срок хранения не истек")
		}
		existOrder = true
		break
	}

	if !existOrder {
		return errors.New("невозможно вернуть заказ курьеру, так как данного заказа нет")
	}
	return s.s.ReturnOrder(id)
}

// GiveOrder gives an order (slice) to the client
func (s Service) GiveOrder(slice []int64) error {
	all := s.s.ListAll()
	for _, num := range slice {
		if num < 1 {
			return errors.New("id заказа не может быть меньше 1")
		}
	}

	var existInFile bool
	var existOrderClient int64
	// Check if the all orders in a slice belong to one client
	for idx, orderSlice := range slice {
		existInFile = false
		for _, order := range all {
			if orderSlice != order.IdOrder {
				continue
			}
			if idx != 0 && order.IdClient != existOrderClient {
				return errors.New("невозможно выдать заказ клиенту. заказы относятся к разным клиентам")
			}
			existOrderClient = order.IdClient
			existInFile = true
		}

		if !existInFile {
			return errors.New("невозможно выдать заказ клиенту. данного заказа не существует")
		}
	}

	ordersToGive := make([]int64, 0)
	// Check the expiry date and ability to give an order
	for _, orderSlice := range slice {
		for idx, order := range all {
			if orderSlice != order.IdOrder {
				continue
			}

			if order.IsReturned || order.IsGiven {
				return errors.New("невозможно выдать заказ клиенту. заказ уже выдан или возвращен курьеру")
			}
			if order.ExpireDate.Before(time.Now()) {
				return errors.New("невозможно выдать заказ кленту. срок хранения истек")
			}
			ordersToGive = append(ordersToGive, idx)
		}
	}

	return s.s.GiveOrder(ordersToGive)
}

// ListOrders returns all todos from storage
func (s Service) ListOrders(clientId int64, lastN int64, inPP bool) ([]model.Order, error) {
	all := s.s.ListAll()
	if clientId < 0 {
		return nil, errors.New("отрицательный id клиента")
	}
	if lastN < 0 {
		return nil, errors.New("отрицательный параметр для n последних заказов")
	}

	sliceId := make([]int64, 0)
	var existInFile bool
	for idx, order := range all {
		if clientId != order.IdClient {
			continue
		}

		if inPP {
			// Condition to check if the order in Pickup Point
			if !order.IsReturned && !order.IsGiven && !order.ExpireDate.Before(time.Now()) {
				sliceId = append(sliceId, idx)
			}
		} else {
			sliceId = append(sliceId, idx)
		}
		existInFile = true
	}
	if !existInFile {
		return nil, errors.New("невозможно вывести список заказов. данного клиента не существует")
	}

	orders, err := s.s.ListOrders(sliceId, lastN)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// ClientRefund make refund of the client order
func (s Service) ClientRefund(orderId int64, clientId int64) error {
	all := s.s.ListAll()
	if orderId < 0 {
		return errors.New("отрицательный id заказа")
	}
	if clientId < 0 {
		return errors.New("отрицательный id клиента")
	}

	var orderToRefund int64
	existOrder := false
	for indx, order := range all {
		if order.IdOrder != orderId || order.IdClient != clientId {
			continue
		}
		if order.IsRefunded {
			return errors.New("невозможно вернуть заказ, который уже был возвращен")
		} else if order.IsGiven && time.Since(order.ExpireDate) < 48*time.Hour {
			orderToRefund = indx
		} else {
			return errors.New("невозможно вернуть заказ в пвз, потому что прошло два дня с момента выдачи, либо данный заказ не выдавался в данном пвз")
		}
		existOrder = true
		break
	}
	// If such an order does not exist -> error
	if !existOrder {
		return errors.New("невозможно оформить возврат, так как данного заказа не существует")
	}

	return s.s.ClientRefund(orderToRefund)
}

// ListRefunds returns all Refunds from storage
func (s Service) ListRefunds(pageNumber int64, onPage int64) ([]model.Order, error) {
	all := s.s.ListAll()
	if pageNumber < 1 {
		return nil, errors.New("номер страницы не может быть меньше 1")
	}

	// Collect all refunds from storage
	sliceId := make([]int64, 0)
	for _, order := range all {
		if order.IsRefunded {
			sliceId = append(sliceId, order.IdOrder)
		}
	}

	orders, err := s.s.ListRefunds(sliceId, pageNumber, onPage)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
