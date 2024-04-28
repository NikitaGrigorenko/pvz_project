package file

import (
	"context"
	"devtask/internal/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestStoragePVZ_AddPVZ(t *testing.T) {
	tempFile := "temp_pvz_data.json"
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatal(err)
		}
		err = os.Remove("pvz_data.json")
		if err != nil {
			t.Fatal(err)
		}
	}(tempFile)

	storage, err := NewStoragePVZ(tempFile)
	assert.NoError(t, err)

	pvz := &model.PVZ{ID: 1, Name: "Test PVZ"}

	id, err := storage.AddPVZ(context.Background(), pvz)
	assert.NoError(t, err)
	assert.Equal(t, pvz.ID, id)

	retrievedPVZ, err := storage.GetPVZ(context.Background(), pvz.ID)
	assert.NoError(t, err)
	assert.Equal(t, pvz, retrievedPVZ)
}

func TestStoragePVZ_GetPVZ(t *testing.T) {
	tempFile := "temp_pvz_data.json"
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatal(err)
		}
		err = os.Remove("pvz_data.json")
		if err != nil {
			t.Fatal(err)
		}
	}(tempFile)

	storage, err := NewStoragePVZ(tempFile)
	assert.NoError(t, err)

	pvz := &model.PVZ{ID: 1, Name: "Test PVZ"}

	_, err = storage.AddPVZ(context.Background(), pvz)
	assert.NoError(t, err)

	retrievedPVZ, err := storage.GetPVZ(context.Background(), pvz.ID)
	assert.NoError(t, err)
	assert.Equal(t, pvz, retrievedPVZ)
}

func TestNewStoragePVZ(t *testing.T) {
	tempFile := "non_existing_file.json"
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatal(err)
		}
	}(tempFile)
	nonExistingFile := "non_existing_file.json"
	storage, err := NewStoragePVZ(nonExistingFile)
	assert.NoError(t, err)
	assert.Empty(t, storage.pvzs)
}
