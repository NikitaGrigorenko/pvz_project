package handlers

import (
	"context"
	"devtask/internal/metrics"
	"devtask/internal/model"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Delete(service StoragePVZ) http.Handler {
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

		status := Del(service, req.Context(), keyInt)
		w.WriteHeader(status)
		metrics.TotalNumberOfPVZMetric.Dec()
		defer metrics.DeleteAttemptMetric.Add(1)
	})
}

func ValidateDelete(keyInt int64) bool {
	if keyInt <= 0 {
		return false
	}
	return true
}

func Del(s StoragePVZ, ctx context.Context, keyInt int64) int {
	if !ValidateDelete(keyInt) {
		return http.StatusBadRequest
	}
	err := s.DeleteInfo(ctx, keyInt)
	if err != nil {
		if errors.Is(err, model.ErrObjectNotFound) {
			return http.StatusNotFound
		}
		if errors.Is(err, model.ErrNoRowsInResultSet) {
			return http.StatusNotFound
		}
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
