package withLogs

import (
	"context"
	"io"
	"math/rand"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type loggingHandler struct {
	writer  io.Writer
	handler http.Handler
}

type LogRecord struct {
	http.ResponseWriter
	status int
}

func (r *LogRecord) Write(p []byte) (int, error) {
	return r.ResponseWriter.Write(p)
}

func (r *LogRecord) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

type requestIDKey string

// RequestID is the context key
var RequestID = requestIDKey("requestID")

// withContext is a middleware that adds context to a request with a unique requestID
func LogWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, RequestID, id)

		record := &LogRecord{
			ResponseWriter: w,
		}

		log.WithFields(log.Fields{
			"requestID": id,
			"method":    r.Method,
			"url":       r.URL,
		}).Info("Request Recived")

		next.ServeHTTP(record, r.WithContext(ctx))

		toLog := log.Fields{
			"requestID": id,
			"elapsed":   time.Now().Sub(start),
			"status":    record.status,
		}

		if record.status == http.StatusBadRequest {
			log.WithFields(toLog).Error("Request Errored")
		} else {
			log.WithFields(toLog).Info("Request Succeded")
		}

	})
}

// withLogging logs the inital request and the response
// func withLogging(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		ctx := r.Context()
// 		id, ok := ctx.Value(requestID).(int64)
// 		if !ok {
// 			id = 0
// 		}

// 		log.WithFields(log.Fields{
// 			"requestID": id,
// 			"method":    r.Method,
// 			"url":       r.URL,
// 		}).Info("Request Recived")

// 		go func() {
// 			select {
// 			case <-ctx.Done():
// 				log.WithFields(log.Fields{
// 					"requestID": id,
// 					"elapsed":   time.Now().Sub(start),
// 				}).Info("Request Finished")
// 			}
// 		}()

// 		next.ServeHTTP(w, r)
// 	})
// }

// func WrapHandler(f http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// record := &LogRecord{
// 		ResponseWriter: w,
// }

// 			f.ServeHTTP(record, r)

// log.Println("Bad Request ", record.status)

// if record.status == http.StatusBadRequest {
// 		log.Println("Bad Request ", r)
// }
// 	}
// }
