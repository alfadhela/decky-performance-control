package main

import (
	"backend/src/pkg/api"
	"net/http"
	"strings"

	"github.com/yousuf64/shift"
)

func main() {
	router := shift.New()
	router.Use(func(next shift.HandlerFunc) shift.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request, route shift.Route) error {
			w.Header().Add("Access-Control-Allow-Origin", "https://steamloopback.host")
			w.Header().Add("Access-Control-Allow-Headers", strings.Join([]string{"Origin", "Content-Type", "Accept"}, ", "))
			w.Header().Add("Access-Control-Request-Method", strings.Join([]string{http.MethodGet, http.MethodPut}, ", "))
			return next(w, r, route)
		}
	})
	router.Group("/api", api.Register)
	http.ListenAndServe(":33238", router.Serve())
}
