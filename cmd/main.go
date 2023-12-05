package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jwhunt19/cowlist-go/internal/server"
)

func main() {
	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// assign database url to variable
	url := os.Getenv("DATABASE_URL")
	fmt.Println(url)

	// connect to database
	conn, err := pgx.Connect(ctx, url) // todo - move to database directory
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

	// handle routes - todo - move to server.go
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

	// listen on port 8080
	http.ListenAndServe(":8080", nil)
}
