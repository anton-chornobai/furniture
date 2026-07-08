package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/furniture/config"
	"github.com/furniture/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("couldnt log the enviroment vars: %v", err)
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cfg loaded")

	_, err = db.ConnectToDB(&cfg.DB) 
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	fmt.Println("DB loaded")


	http.Handle("/", my_middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving HTML file")
		w.Write([]byte("HELLO WORLD"))
		w.WriteHeader(http.StatusAccepted)
	})))
	

	if err := http.ListenAndServe(":" + strconv.Itoa(cfg.Server.Port), nil); err != nil {
		log.Fatalf("server went down %v", err)
	}
}

func my_middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Im middleware")

		next.ServeHTTP(w, r)
	})
}
