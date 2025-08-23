package middleware

import (
	"fmt"
	"net/http"

	"golang/api"

	"golang/internal/tools"

	log "github.com/sirupsen/logrus"
)

var ErrUnauthorized = fmt.Errorf("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" {
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(ErrUnauthorized)
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

	})
}
