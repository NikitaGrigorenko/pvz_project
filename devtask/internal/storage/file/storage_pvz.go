package file

import (
	"bufio"
	"context"
	"devtask/internal/model"
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"
)

type StoragePVZ struct {
	pvzs  map[int64]model.PVZ
	mutex sync.RWMutex
}

func NewStoragePVZ(filePath string) (StoragePVZ, error) {
	file, err := os.Open(filePath)
	if err != nil {
		file, err = os.Create(filePath)
		if err != nil {
			return StoragePVZ{}, err
		}
		rawBytes, err := json.MarshalIndent(map[int64]model.PVZ{}, "", "    ")
		if err != nil {
			return StoragePVZ{}, err
		}

		err = os.WriteFile(filePath, rawBytes, 0777)
		if err != nil {
			return StoragePVZ{}, err
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
		return StoragePVZ{}, err
	}
	var pvzsStor map[int64]model.PVZ
	err = json.Unmarshal(rawBytes, &pvzsStor)
	if err != nil {
		return StoragePVZ{}, err
	}

	storage := make(map[int64]model.PVZ)
	for _, pvz := range pvzsStor {
		storage[pvz.ID] = pvz
	}

	return StoragePVZ{
		pvzs: storage,
	}, nil
}

func (p *StoragePVZ) AddPVZ(_ context.Context, pvz *model.PVZ) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.pvzs[pvz.ID] = *pvz
	err := p.saveToFile("pvz_data.json")
	if err != nil {
		return pvz.ID, errors.New("ошибка при записи данных в файл")
	}
	return pvz.ID, nil
}

func (p *StoragePVZ) GetPVZ(_ context.Context, id int64) (*model.PVZ, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	pvz, ok := p.pvzs[id]
	if !ok {
		return &pvz, errors.New("данный пвз отсутствует")
	}
	return &pvz, nil
}

func (p *StoragePVZ) Update(_ context.Context, _ *model.PVZ, _ int64) (int64, error) {
	// TODO: implement method Update for this file storage
	return 0, nil
}

func (p *StoragePVZ) Delete(_ context.Context, _ int64) error {
	// TODO: implement method Delete for this file storage
	return nil
}

func (p *StoragePVZ) List(_ context.Context) ([]model.PVZ, error) {
	// TODO: implement method List for this file storage
	return nil, nil
}

func (p *StoragePVZ) saveToFile(filename string) error {
	jsonData, err := json.Marshal(p.pvzs)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}
