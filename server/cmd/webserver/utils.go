package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/vicesoftware/vice-go-boilerplate/pkg/log"
	"go.uber.org/zap"

	"github.com/vicesoftware/vice-go-boilerplate/cmd/webserver/models"
	"github.com/vicesoftware/vice-go-boilerplate/pkg/database"
)

func Ok(w http.ResponseWriter, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func handler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// start time for time_taken calculation
		var (
			start  = time.Now()
			status int
			err    error
		)

		// use defer/recover; if the handler panics we can log the details
		defer func() {
			if perr := recover(); perr != nil {
				// panics are certainly going to be 500's
				//
				// panics are NOT used like exceptions in .NET/Java, they are usually
				// the result of incorrect code like a nil pointer dereference, for example.
				//
				// in contrast, unable to connect to a database is a normal fact of life and is
				// represented by an "error", not a forced unwinding of the stack.
				//
				// Go encourages you to handle errors (network interruptions and other facts of life)
				// and panics are reserved for correctness problems.
				status = http.StatusInternalServerError
				w.WriteHeader(status)

				// code that panics usually panics with an error, but not necessarily.
				// fun fact, in .NET you can "throw" anything, it doesn't need to be a type
				// that derives from System.Exception.
				panicErr, ok := perr.(error)
				if !ok {
					// %v formats the value in a default format; it's useful when you're
					// not sure (or don't care) what the type is
					panicErr = fmt.Errorf("%v", perr)
				}
				err = panicErr
			}

			duration := time.Since(start)
			go writeHTTPLog(r, duration, status, err)
		}()

		// for the API server we'll always use application/json
		w.Header().Set("content-type", "application/json")

		// call the handler
		err = f(w, r)

		// determine http status based on the type of error (if any) returned
		status = httpStatus(err)

		// if an error was returned write it to the client
		// TODO: for security reasons you may not always want to return to raw error
		// TODO: to the client
		if err != nil {
			http.Error(w, errToJSON(err), status)
		}

		// the defer function handles writing log output
	}
}

func notFoundHandler(_ http.ResponseWriter, _ *http.Request) error {
	return &notFound{}
}

func httpStatus(err error) int {
	if err == nil {
		return 200
	}
	if isInvalidRequest(err) || database.IsInvalidRequest(err) {
		return 400
	}
	if isNotFound(err) || database.IsNotFound(err) {
		return 404
	}
	return 500
}

func isInvalidRequest(err error) bool {
	_, ok := err.(*invalidRequest)
	return ok
}

func isNotFound(err error) bool {
	_, ok := err.(*notFound)
	return ok
}

func errToJSON(err error) string {
	b, _ := json.Marshal(models.ErrorResponse{Error: err.Error()})
	return string(b)
}

// writeHTTPLog writes the following keys to the log entry:
//
//   http_status          The HTTP status code returned.
//   ip                   The remote IP address. X-Real-IP and X-Forwarded-For aware.
//   method               GET, POST, PUT, DELETE, etc
//   time_taken           The time taken to complete the request in milliseconds.
//   uri                  The request URI.
//
// The log level is determined by the status code:
//
//   status < 400          Info
//   400 <= status < 500   Warning
//   status >= 500         Error
func writeHTTPLog(r *http.Request, duration time.Duration, status int, err error) {
	timeTakenSecs := float64(duration) / 1e9

	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		forwardedFor := r.Header.Get("X-Forwarded-For")
		ip = strings.SplitN(forwardedFor, ",", 2)[0]
		if ip == "" {
			var splitErr error
			ip, _, splitErr = net.SplitHostPort(r.RemoteAddr)
			if splitErr != nil {
				ip = r.RemoteAddr
			}
		}
	}

	fields := []zap.Field{
		zap.Int("http_status", status),
		zap.String("ip", ip),
		zap.String("method", r.Method),
		zap.Int64("time_taken", int64(timeTakenSecs*1000)),
		zap.String("uri", r.RequestURI),
	}

	msg := http.StatusText(status)
	if err != nil {
		fields = append(fields, zap.Error(err))
	}

	if status >= 400 && status < 500 {
		log.Warn(msg, fields...)
	} else if status >= 500 {
		log.Error(msg, fields...)
	} else {
		log.Info(msg, fields...)
	}
}
