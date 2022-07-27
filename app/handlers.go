package app

import (
	"api/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// type Customer struct {
// 	ID      int    `json:"id" xml:"id"`
// 	Name    string `json:"name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	ZipCode string `json:"zip_code" xml:"zipcode"`
// }

// var customers []Customer = []Customer{
// 	{1, "User1", "Jakarta", "12345"},
// 	{2, "User2", "Surabaya", "67890"},
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello Celerates!")
// }

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "get customer endpoint\n")

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

// func getCustomer(w http.ResponseWriter, r *http.Request) {

// 	// * get route variable
// 	vars := mux.Vars(r)

// 	customerId := vars["customer_id"]

// 	// * convert string to int
// 	id, err := strconv.Atoi(customerId)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "invalid customer id")
// 		return
// 	}

// 	// * searching customer data
// 	var cust Customer

// 	for _, data := range customers {
// 		if data.ID == id {
// 			cust = data
// 		}
// 	}

// 	if cust.ID == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "customer data not found")
// 		return
// 	}

// 	// * return customer data
// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(cust)
// }

// func addCustomer(w http.ResponseWriter, r *http.Request) {
// 	// * decode request body
// 	var cust Customer
// 	json.NewDecoder(r.Body).Decode(&cust)

// 	// * generate new id
// 	nextID := getNextID()
// 	cust.ID = nextID

// 	// * save data to array
// 	customers = append(customers, cust)

// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprintln(w, "customer successfully created")
// }

// func getNextID() int {
// 	lastIndex := len(customers) - 1
// 	lastCustomer := customers[lastIndex]
// 	// cust := customers[len(customers)-1]

// 	return lastCustomer.ID + 1
// }

// func updateCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	customerId := vars["customer_id"]

// 	// * convert string to int
// 	id, err := strconv.Atoi(customerId)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "invalid customer id")
// 		return
// 	}

// 	// * searching customer data
// 	var cust Customer

// 	for index, data := range customers {
// 		if data.ID == id {
// 			cust = data

// 			var newCust Customer
// 			json.NewDecoder(r.Body).Decode(&newCust)
// 			customers[index].Name = newCust.Name
// 			customers[index].City = newCust.City
// 			customers[index].ZipCode = newCust.ZipCode

// 			w.WriteHeader(http.StatusOK)
// 			fmt.Fprint(w, "Customer data updated")
// 			return
// 		}
// 	}
// 	//eror data not find
// 	if cust.ID == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "customer data not found")
// 		return
// 	}
// }

// func deleteCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	customerId := vars["customer_id"]
// 	id, delCust := strconv.Atoi(customerId)

// 	if delCust != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "invalid customer id")
// 		return
// 	}
// 	for i, delCust := range customers {
// 		if delCust.ID == id {
// 			customers = append(customers[:i], customers[i+1:]...)
// 			fmt.Fprintf(w, "ID %v has been deleted successfully", customerId)
// 		}
// 	}
// }
