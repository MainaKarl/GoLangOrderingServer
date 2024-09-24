package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"	
	"my_business/pkg/models"
)

// Initialize the database connection
func InitDB() {
    var err error
	
    models.NewConfig.DB, err = sql.Open("postgres", models.NewConfig.DBconn)
    if err != nil {
        log.Fatal(err)
    }

    err = models.NewConfig.DB.Ping()
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }
    fmt.Println("Connected to the database")
}