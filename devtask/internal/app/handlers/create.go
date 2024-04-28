package handlers

import (
	"context"
	"devtask/internal/metrics"
	"devtask/internal/model"
	"encoding/json"
	"net/http"
)

func Create(service StoragePVZ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var unm model.PVZRequest
		err := json.NewDecoder(req.Body).Decode(&unm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		pvzJson, status := Add(service, req.Context(), unm)
		if status != http.StatusOK {
			w.WriteHeader(status)
		}

		_, err = w.Write(pvzJson)
		metrics.TotalNumberOfPVZMetric.Inc()
		defer metrics.AddAttemptMetric.Add(1)

		if err != nil {
			return
		}
	})
}
func Add(s StoragePVZ, ctx context.Context, unm model.PVZRequest) ([]byte, int) {
	pvzRepo := &model.PVZ{
		Name:    unm.Name,
		Address: unm.Address,
		Contact: unm.Contact,
	}
	response, err := s.AddInfo(ctx, *pvzRepo)
	if err != nil {
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
