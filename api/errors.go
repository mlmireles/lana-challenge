package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// BadRequest is handled by setting the status code in the reply to StatusBadRequest.
type BadRequest struct{ error }

// NotFound is handled by setting the status code in the reply to StatusNotFound.
type NotFound struct{ error }

// NotAuthorized is handled by setting the status code in the reply to StatusUnauthorized
type NotAuthorized struct{ error }

// StatusUnprocessableEntity is handled by setting the status code in the reply to StatusUnprocessableEntity
type UnprocessableEntity struct{ error }

// DuplicateItem is handled by setting the status code in the reply to
type DuplicateItem struct{ error }

// Handler wraps a function returning an error by handling the error and
// returning a http.Handler.
// If the error is of the one of the types defined above, it is handled as
// described for every type.
// If the error is of another type, it is considered as an internal error and
// its message is logged.
func HttpHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			log.Printf("(%d) %s: %s", 200, r.Method, r.URL)
			return
		}

		message, _ := json.Marshal(err)

		var status int
		switch err.(type) {
		case BadRequest:
			status = http.StatusBadRequest
			http.Error(w, string(message), http.StatusBadRequest)
			break
		case NotFound:
			status = http.StatusNotFound
			http.Error(w, string(message), http.StatusNotFound)
			break
		case NotAuthorized:
			status = http.StatusUnauthorized
			http.Error(w, string(message), http.StatusUnauthorized)
			break
		case UnprocessableEntity:
			status = http.StatusUnprocessableEntity
			http.Error(w, string(message), http.StatusUnprocessableEntity)
			break
		case DuplicateItem:
			status = http.StatusNotAcceptable
			http.Error(w, string(message), http.StatusNotAcceptable)
			break
		default:
			status = http.StatusInternalServerError
			http.Error(w, err.Error(), http.StatusInternalServerError)
			break
		}

		log.Printf("[Error] (%d) %s %s \n%s", status, r.Method, r.URL, err.Error())
	}
}
