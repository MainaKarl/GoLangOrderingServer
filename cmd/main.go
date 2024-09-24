package main

import (
    "fmt"
    "log"
    "net/http"
    _ "github.com/lib/pq"
    "my_business/pkg/models"
    "my_business/pkg/db"
    "github.com/joho/godotenv"
    "my_business/pkg/services"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    err = models.InitConfig()
    if err != nil {
        log.Fatal(err)
    }
}


func main() {
    // Initialize the database
    db.InitDB()

    // Setup routes
    http.HandleFunc("/customers", services.AddCustomer)
    http.HandleFunc("/orders", services.AddOrder)

    // Start the server
    fmt.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
