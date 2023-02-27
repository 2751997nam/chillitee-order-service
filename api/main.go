package main

import (
	"fmt"
	"log"
	"logger-service/internal/driver"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var webPort = os.Getenv("PORT")

func main() {
	loadEnv()
	if webPort == "" {
		webPort = "8080"
	}

	db, err := run()
	if err != nil {
		log.Fatal("cannot open connection with db")
	}
	d, err := db.DB()
	if err != nil {
		log.Fatal("cannot open connection with db")
	}
	defer d.Close()

	router := routes()

	router.Run(fmt.Sprintf("localhost:%s", webPort))

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func run() (*gorm.DB, error) {
	db, err := driver.ConnectSQL()
	if err != nil {
		log.Fatal("cannot open connection with db")
	}

	return db, err
}
