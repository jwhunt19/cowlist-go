package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"encoding/json"

	"github.com/jackc/pgx/v5"
)

type cow struct {
	id int
	name string
	age int
	color string
	healthy bool
}



func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func AddCow(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing the form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	age := r.FormValue("age")
	color := r.FormValue("color")
	healthy := r.FormValue("healthy")

	fmt.Fprintf(w, "Received: %s, %s, %s, %s", name, age, color, healthy)
}

func GetAllCows(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	enableCors(&w)

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	rows, err := conn.Query(ctx, "select * from cows")
	if err != nil {
		fmt.Printf("Query error: %v", err)
		return
	}
	defer rows.Close()

	var cows []cow

	for rows.Next() {
		var id int
		var name string
		var age int
		var color string
		var healthy bool

		err = rows.Scan(&id, &name, &age, &color, &healthy)
		if err != nil {
			fmt.Printf("Scan error: %v", err)
			return
		}

		// fmt.Println(id, name, age, color, healthy)
		cows = append(cows, cow{id:id, name:name, age:age, color:color, healthy:healthy})
	}

	if rows.Err() != nil {
		fmt.Printf("rows error: %v", rows.Err())
		return
	}

	fmt.Println(cows)

	// todo - return cows to front end as json

	// return cows as json



}
