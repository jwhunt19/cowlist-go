package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jwhunt19/cowlist-go/internal/server"
)

func main() {
	// TODO: Move everything from here through line 55 to new database package
	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// assign database url to variable
	url := os.Getenv("DATABASE_URL")
	fmt.Println(url)

	// connect to database
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	// setup schema and initial data
	_, err = conn.Exec(ctx, `
	CREATE TABLE if not exists cows (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50),
		age INTEGER,
		color VARCHAR(50),
		healthy BOOLEAN
	);
	
	INSERT INTO cows (name, age, color, healthy) 
	VALUES ('Miltank', 24, 'Pink', true);

	INSERT INTO cows (name, age, color, healthy) 
	VALUES ('Cow', 2, 'Black', true);

	INSERT INTO cows (name, age, color, healthy) 
	VALUES ('Dairy', 999, 'White', false);
	
	`)
	if err != nil {
		fmt.Printf("Unable to setup example schema and data: %v", err)
		return
	}

	// TODO: Move all http.Handle functions to a SetupRoutes func in server.go
	// TODO: will no longer need to wrap the handler parameter in http.Handle
	// handle routes
	http.Handle("/addcow", server.EnableCors(func(w http.ResponseWriter, r *http.Request) {
		server.AddCow(w, r, conn) 
	}))

	http.Handle("/getallcows", server.EnableCors(func(w http.ResponseWriter, r *http.Request) {
		server.GetAllCows(w, r, conn)
	}))

	http.Handle("/updatecow", server.EnableCors(func(w http.ResponseWriter, r *http.Request) {
		server.UpdateCow(w, r, conn)
	}))

	http.Handle("/deletecow/", server.EnableCors(func(w http.ResponseWriter, r *http.Request) {
		server.DeleteCow(w, r, conn)
	}))

	// TODO: find a way to initalize database via call in main
	// TODO: call a func to setup routes from server.go

	// listen on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}
