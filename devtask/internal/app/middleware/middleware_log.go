package middleware

import (
	"bytes"
	"devtask/internal/model"
	"io"
	"log"
	"net/http"
	"time"
)

type SenderKafka interface {
	SendEvent(model.EventMessage) error
}

func Logger(next http.Handler, sender SenderKafka) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		currentTime := time.Now()
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return
		}
		req.Body = io.NopCloser(io.Reader(bytes.NewBuffer(bodyBytes)))

		err = sender.SendEvent(model.EventMessage{
			Method:    req.Method,
			Body:      string(bodyBytes),
			TimeStamp: currentTime,
		})

		if err != nil {
			log.Fatal(err)
			return
		}
		next.ServeHTTP(w, req)
	}
}
