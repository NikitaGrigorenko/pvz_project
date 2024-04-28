package handlers

import (
	"context"
	"devtask/internal/metrics"
	"devtask/internal/model"
	"encoding/json"
	"errors"
	"net/http"
)

func List(service StoragePVZ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		pvzJson, status := Lst(service, req.Context())
		w.WriteHeader(status)
		_, err := w.Write(pvzJson)

		defer metrics.ListAttemptMetric.Add(1)
		if err != nil {
			return
		}
	})
}

func Lst(s StoragePVZ, ctx context.Context) ([]byte, int) {
	pvzInfo, err := s.ListInfo(ctx)
	if err != nil {
		if errors.Is(err, model.ErrObjectNotFound) {
			return nil, http.StatusNotFound
		}
		return nil, http.StatusInternalServerError
	}
	pvzJson, _ := json.Marshal(pvzInfo)

	return pvzJson, http.StatusOK
}
