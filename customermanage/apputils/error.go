package apputils

import (
	"encoding/json"
	"net/http"
)

type (
	appError struct {
		Error          string `json:"error"`
		HttpStatusCode int    `json:"status"`
	}
	errResource struct {
		ErrData appError
	}
)

func DisplayAppError(w http.ResponseWriter, handlererror error, errorCode int) {
	errObj := appError{Error: handlererror.Error(), HttpStatusCode: errorCode}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	if j, err := json.Marshal(errResource{ErrData: errObj}); err == nil {
		w.Write(j)
	}

}
