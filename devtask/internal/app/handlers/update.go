package handlers

import (
	"context"
	"devtask/internal/metrics"
	"devtask/internal/model"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func Update(service StoragePVZ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
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

		var unm model.PVZRequest
		if err = json.Unmarshal(body, &unm); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		pvzJson, status := Upd(service, req.Context(), unm, keyInt)
		w.WriteHeader(status)
		_, err = w.Write(pvzJson)
		defer metrics.UpdateAttemptMetric.Add(1)
		if err != nil {
			return
		}
	})
}

func Upd(s StoragePVZ, ctx context.Context, unm model.PVZRequest, keyInt int64) ([]byte, int) {
	pvzRepo := &model.PVZ{
		Name:    unm.Name,
		Address: unm.Address,
		Contact: unm.Contact,
	}
	response, err := s.UpdateInfo(ctx, pvzRepo, keyInt)
	if err != nil {
		if errors.Is(err, model.ErrNoRowsInResultSet) {
			return nil, http.StatusNotFound
		}
		return nil, http.StatusInternalServerError
	}

	resp := &model.PVZ{
		ID:      response,
		Name:    pvzRepo.Name,
		Address: pvzRepo.Address,
		Contact: pvzRepo.Contact,
	}
	pvzJson, _ := json.Marshal(resp)
	return pvzJson, http.StatusOK
}
