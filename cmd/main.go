package main

import (
	"context"
	"evolvecredit-test/api/v1/handlers"
	"evolvecredit-test/internal/common/db"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
)

var DBConn *bun.DB

func main() {
	// load env values
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error in loading environment values: %v", err)

	}
	var wait time.Duration
	flag.DurationVar(
		&wait,
		"graceful-timeout",
		time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish",
	)
	flag.Parse()

	// get access to db
	DBConn = db.GetDBConnection()

	// get routes here and display it.
	userHandler, err := handlers.NewUserHandler(DBConn)
	if err != nil {
		log.Fatalf("error in loading one of the routes: %v", err)
	}

	router := mux.NewRouter()
	//routes
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	})
	router.HandleFunc("/api/v1/users", userHandler.ListAll).Methods("GET")
	router.HandleFunc("/api/v1/users/search", userHandler.Search).Methods("GET")

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	// block until we receive a signal (Ctrl+c)
	<-c

	//
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// only shutdown if there are no connections and if they exist just wait for the given time.
	server.Shutdown(ctx)

	// close db connection
	if DBConn != nil {
		DBConn.Close()
	}

	log.Println("server is shutting down")
	os.Exit(0)
}
