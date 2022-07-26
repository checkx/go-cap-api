package main

import "api/app"

// import (
// 	"encoding/json"
// 	"encoding/xml"
// 	"fmt"
// 	"net/http"
// )

// type Customer struct {
// 	Name    string `json:"name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	Zipcode string `json:"zip_code" xml:"zipcode"`
// }

// var customers []Customer = []Customer{
// 	{"User1", "Jakarta", "12345"},
// 	{"User2", "Surabaya", "67890"},
// }

func main() {
	app.Start()

}

// 	// // * defining routes
// 	// http.HandleFunc("/greet", greet)
// 	// http.HandleFunc("/customers", getCustomers)

// 	// // * starting the server
// 	// http.ListenAndServe(":8080", nil)
// }

// // func greet(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprint(w, "Hello Celerates!")
// // }

// // func getCustomers(w http.ResponseWriter, r *http.Request) {
// // 	// fmt.Fprint(w, "Get Customers Endpoint!")
// // 	if r.Header.Get("Content-Type") == "application/xml" {
// // 		w.Header().Add("Content-Type", "application/xml")
// // 		xml.NewEncoder(w).Encode(customers)
// // 	} else {
// // 		w.Header().Add("Content-Type", "application/json")
// // 		json.NewEncoder(w).Encode(customers)
// // 	}

// }
