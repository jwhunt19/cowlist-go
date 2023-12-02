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
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	// setup schema and initial data
	_, err = conn.Exec(ctx, `
	CREATE TABLE cows (
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

	// handle routes
	http.HandleFunc("/addcow", server.AddCow)
	http.HandleFunc("/getallcows", func(w http.ResponseWriter, r *http.Request) {
		server.GetAllCows(w, r, conn)
	})
	http.HandleFunc("/updatecow", func(w http.ResponseWriter, r *http.Request) {
		server.UpdateCow(w, r, conn)
	})

	// listen on port 8080
	http.ListenAndServe(":8080", nil)

	/*
			=========================
			Test query - todo: delete
			=========================
	*/
	
	rows, err := conn.Query(ctx, "select * from cows")
	if err != nil {
		fmt.Printf("Query error: %v", err)
		return
	}
	defer rows.Close()

	fmt.Println(rows)

	// Iterate through the result set
	for rows.Next() {
		var name string
		var age int
		var color string
		var healthy bool
		var id int

		err = rows.Scan(&id, &name, &age, &color, &healthy)
		if err != nil {
			fmt.Printf("Scan error: %v", err)
			return
		}

		fmt.Println(name, age, color, healthy)
	}

	if rows.Err() != nil {
		fmt.Printf("rows error: %v", rows.Err())
		return
	}

	/*
			Test query end
	*/

}
