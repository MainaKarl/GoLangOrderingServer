package services

import (
	"encoding/json"
	"net/http"
	"time"
	"my_business/pkg/models"
)

// Order struct
type Order struct {
    ID         int       `json:"id"`
    CustomerID int       `json:"customer_id"`
    Item       string    `json:"item"`
    Amount     float64   `json:"amount"`
    Time       time.Time `json:"time"`
}

// Handler to add an order
func AddOrder(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var order Order
    err := json.NewDecoder(r.Body).Decode(&order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    sqlStatement := `INSERT INTO orders (customer_id, item, amount) VALUES ($1, $2, $3) RETURNING id, time`
    err = models.NewConfig.DB.QueryRow(sqlStatement, order.CustomerID, order.Item, order.Amount).Scan(&order.ID, &order.Time)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(order)
}