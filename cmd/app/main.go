package main

import (
	"fmt"
	"net/http"

	"github.com/reoverduguez/map-app/backend/configs"
	"github.com/reoverduguez/map-app/backend/pkg/db"
)

func main() {
	config, err := configs.ParsedConfig()

	if err != nil {
		fmt.Println(err)
	}

	mongoClient, err := db.Connect(db.ConfigDB{
		Username: config.Database.USERNAME,
		Password: config.Database.PASSWORD,
	})

	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	http.ListenAndServe(":8080", nil)
}
