package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8083",
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
