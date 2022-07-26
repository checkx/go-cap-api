package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	mux := mux.NewRouter()

	// * defining routes
	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customers_id}", getCustomers).Methods(http.MethodGet)

	// * starting the server
	http.ListenAndServe(":8080", mux)
}
