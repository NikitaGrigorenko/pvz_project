package file

import (
	"bufio"
	"devtask/internal/model"
	"encoding/json"
	"errors"
	"io"
	"os"
	"sort"
	"time"
)

type Storage struct {
	storage map[int64]OrderDTO
}

func NewStorageOrder(filePath string) (Storage, error) {
	file, err := os.Open(filePath)
	if err != nil {
		file, err = os.Create(filePath)
		if err != nil {
			return Storage{}, err
		}
		rawBytes, err := json.MarshalIndent([]OrderDTO{}, "", "    ")
		if err != nil {
			return Storage{}, err
		}

		err = os.WriteFile(filePath, rawBytes, 0777)
		if err != nil {
			return Storage{}, err
		}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := bufio.NewReader(file)
	rawBytes, err := io.ReadAll(reader)
	if err != nil {
		return Storage{}, err
	}

	var orders []OrderDTO
	err = json.Unmarshal(rawBytes, &orders)
	if err != nil {
		return Storage{}, err
	}

	storage := make(map[int64]OrderDTO)
	for _, order := range orders {
		storage[order.IdOrder] = order
	}

	return Storage{storage: storage}, nil

}

// TakeOrder creates order in the storage
func (s *Storage) TakeOrder(input model.OrderInput) error {
	all := s.storage

	// Fill all fields DTO to take new order from courier
	newOrder := OrderDTO{
		IdOrder:     input.IdOrder,
		IdClient:    input.IdClient,
		IsGiven:     false,
		IsReturned:  false,
		IsRefunded:  false,
		ExpireDate:  input.ExpireDate,
		WeightGrams: input.WeightGrams,
		PriceKopeck: input.PriceKopeck,
		PackageType: input.PackageType,
	}

	all[newOrder.IdOrder] = newOrder
	err := writeBytes("storage.json", all)
	if err != nil {
		return err
	}
	return nil
}

// ReturnOrder sets IsReturned flag true
func (s *Storage) ReturnOrder(id int64) error {
	all := s.storage

	order := all[id]
	order.IsReturned = true
	all[id] = order

	err := writeBytes("storage.json", all)
	if err != nil {
		return err
	}
	return nil
}

// GiveOrder gives an order (slice) to the client
func (s *Storage) GiveOrder(slice []int64) error {
	all := s.storage

	// Check the expiry date and ability to give an order
	for _, idx := range slice {
		order := all[idx]
		order.IsGiven = true
		order.ExpireDate = time.Now()
		all[idx] = order
	}

	err := writeBytes("storage.json", all)
	if err != nil {
		return err
	}
	return nil
}

// ListOrders returns orders from storage
func (s *Storage) ListOrders(slice []int64, lastN int64) ([]model.Order, error) {
	all := s.storage
	onlyNeedable := make([]model.Order, 0, len(all))
	// If the number of orders to be shown are 0 then we want to show all orders
	if lastN == 0 {
		lastN = int64(len(all))
	}
	counter := int64(0)

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})

	for _, idx := range slice {
		order := all[idx]
		onlyNeedable = append(onlyNeedable, writeOrder(order))
		counter++
		if counter > lastN-1 {
			break
		}
	}

	return onlyNeedable, nil
}

// ClientRefund make refund of the client order
func (s *Storage) ClientRefund(orderId int64) error {
	all := s.storage

	// Iterate through all orders and find with order_id and client_id
	order := all[orderId]
	order.IsGiven = false
	order.IsRefunded = true
	all[orderId] = order

	err := writeBytes("storage.json", all)
	if err != nil {
		return err
	}
	return nil
}

// ListRefunds returns all Refunds from storage
func (s *Storage) ListRefunds(slice []int64, pageNumber int64, onPage int64) ([]model.Order, error) {
	all := s.storage

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})

	// Collect all refunds from storage
	allRefunds := make([]model.Order, 0)
	for _, order := range slice {
		allRefunds = append(allRefunds, writeOrder(all[order]))
	}

	startShowPos := pageNumber*onPage - onPage
	if cap(allRefunds)-1 < int(startShowPos) {
		return nil, errors.New("данной страницы не существует")
	}

	if cap(allRefunds)-1 < int(startShowPos+onPage) {
		return allRefunds[startShowPos:], nil
	}

	return allRefunds[startShowPos : startShowPos+onPage], nil
}

// ListAll To list all orders from json file
func (s *Storage) ListAll() map[int64]model.Order {
	reader := s.storage
	ordersSlice := make(map[int64]model.Order)
	for idx, order := range reader {
		ordersSlice[idx] = writeOrder(order)
	}

	return ordersSlice
}

// To write order or orders into Json file
func writeBytes(filePath string, toDos map[int64]OrderDTO) error {
	var ordersSlice []OrderDTO
	for _, order := range toDos {
		ordersSlice = append(ordersSlice, order)
	}
	rawBytes, err := json.MarshalIndent(ordersSlice, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, rawBytes, 0777)
	if err != nil {
		return err
	}
	return nil
}

// To write order into OrderDto model
func writeOrder(order OrderDTO) model.Order {
	return model.Order{
		IdOrder:     order.IdOrder,
		IdClient:    order.IdClient,
		IsGiven:     order.IsGiven,
		IsReturned:  order.IsReturned,
		IsRefunded:  order.IsRefunded,
		ExpireDate:  order.ExpireDate,
		WeightGrams: order.WeightGrams,
		PriceKopeck: order.PriceKopeck,
		PackageType: order.PackageType,
	}
}
