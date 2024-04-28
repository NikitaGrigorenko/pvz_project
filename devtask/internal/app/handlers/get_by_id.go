package handlers

import (
	"context"
	"devtask/internal/metrics"
	"devtask/internal/model"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetByID(service StoragePVZ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		key, ok := mux.Vars(req)[QueryParamKey]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		keyInt, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		data, status := Get(service, req.Context(), keyInt)
		w.WriteHeader(status)
		_, err = w.Write(data)

		defer metrics.GetAttemptMetric.Add(1)

		if err != nil {
			return
		}
	})
}

func ValidateGetByID(keyInt int64) bool {
	if keyInt <= 0 {
		return false
	}
	return true
}

func Get(s StoragePVZ, ctx context.Context, keyInt int64) ([]byte, int) {
	if !ValidateGetByID(keyInt) {
		return nil, http.StatusBadRequest
	}
	response, err := s.GetInfo(ctx, keyInt)
	if err != nil {
		if errors.Is(err, model.ErrObjectNotFound) {
			return nil, http.StatusNotFound
		}
		return nil, http.StatusInternalServerError
	}
	resp := &model.PVZ{
		ID:      response.ID,
		Name:    response.Name,
		Address: response.Address,
		Contact: response.Contact,
	}
	pvzJson, _ := json.Marshal(resp)
	return pvzJson, http.StatusOK
}
