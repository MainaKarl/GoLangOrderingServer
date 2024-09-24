package services

import (
	"encoding/json"
	"net/http"
	"my_business/pkg/models"
)

// Customer struct
type Customer struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Code string `json:"code"`
}

// Handler to add a customer
func AddCustomer(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var customer Customer
    err := json.NewDecoder(r.Body).Decode(&customer)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    sqlStatement := `INSERT INTO customers (name, code) VALUES ($1, $2) RETURNING id`
    err = models.NewConfig.DB.QueryRow(sqlStatement, customer.Name, customer.Code).Scan(&customer.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(customer)
}