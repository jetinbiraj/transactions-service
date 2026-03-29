package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"transactions-service/domain"
)

func SuccessJson(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int, logEnabled bool) {
	jsonMsg, err := json.Marshal(data)
	if err != nil {
		Error(w, r, fmt.Errorf("serialising response failed: %w", err), 500, logEnabled)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	Success(w, r, jsonMsg, statusCode, logEnabled)
}

func Success(w http.ResponseWriter, r *http.Request, jsonMsg []byte, statusCode int, logEnabled bool) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if statusCode == 0 {
		statusCode = http.StatusOK
	}

	w.WriteHeader(statusCode)
	if _, err := w.Write(jsonMsg); err != nil {
		log.Printf("Error writing response: %v", err)
	}

	if logEnabled {
		logSuccessResponse(r, statusCode)
	}
}

func Error(w http.ResponseWriter, r *http.Request, err error, code int, logEnabled bool) {
	if code == 0 {
		code = toHTTPStatusCode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	if err == nil {
		err = fmt.Errorf("nil err")
	}
	logErr := err
	errorMsgJSON, err := json.Marshal(domain.MessageResponse{Message: err.Error()})
	if err != nil {
		log.Println(err)
	} else {
		if _, err = w.Write(errorMsgJSON); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	}

	if logEnabled {
		logErrorResponse(r, code, logErr)
	}
}

func toHTTPStatusCode(err error) int {
	switch {
	case errors.Is(err, domain.ErrInvalidRequestBody):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func logSuccessResponse(r *http.Request, statusCode int) {
	log.Printf(
		"%s %s %s %d",
		r.Method,
		r.RequestURI,
		r.RemoteAddr,
		statusCode,
	)
}

func logErrorResponse(r *http.Request, statusCode int, logErr error) {
	log.Printf(
		"%s %s %s %d %s",
		r.Method,
		r.RequestURI,
		r.RemoteAddr,
		statusCode,
		logErr.Error(),
	)
}
