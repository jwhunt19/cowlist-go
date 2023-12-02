package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

type cow struct {
	Id      int
	Name    string
	Age     int
	Color   string
	Healthy bool
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// run a query to get all cows against the database
	rows, err := conn.Query(ctx, "select * from cows")
	if err != nil {
		fmt.Printf("Query error: %v", err)
		return
	}
	defer rows.Close()

	// create cows variable which is a slice containing type cow
	var cows []cow

	// interate over the rows returned by the query
	for rows.Next() {
		var id int
		var name string
		var age int
		var color string
		var healthy bool

		// Set the variables to the corresponding values found in the row
		err = rows.Scan(&id, &name, &age, &color, &healthy)
		if err != nil {
			fmt.Printf("Scan error: %v", err)
			return
		}

		// append this cow to the cows slice
		cows = append(cows, cow{Id: id, Name: name, Age: age, Color: color, Healthy: healthy})

		// todo - sort cows by id. when updating cows in front end, updated cow moves
		// to bottom. I think sorting it here should fix it. If not, sort it in the
		// getCows() function in react.
	}

	// handle any error from interating the rows?
	if rows.Err() != nil {
		fmt.Printf("rows error: %v", rows.Err())
		return
	}

	// encode the cows data into encoded json
	res, err := json.Marshal(cows)
	if err != nil {
		fmt.Println("error:", err)
	}

	// sets the content type to json in the header
	w.Header().Set("Content-Type", "application/json")

	// write the encoded json to the http response
	_, err = w.Write(res)
	if err != nil {
		fmt.Println("error:", err)

	}
}

func UpdateCow(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	enableCors(&w)

	// checks if request is expected method
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// parses form, not really sure what that means
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing the form", http.StatusBadRequest)
		return
	}

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// set values to variables
	name := r.FormValue("Name")
	age := r.FormValue("Age")
	color := r.FormValue("Color")
	healthy := r.FormValue("Healthy")
	id := r.FormValue("Id")

	// create update query with variables
	query := fmt.Sprintf("UPDATE cows SET Name='%s', Age='%s', Color='%s', Healthy='%s' WHERE id='%s'", name, age, color, healthy, id)

	// execute query
	_, err = conn.Exec(ctx, query)
	if err != nil {
		fmt.Println("error:", err)
	}

	// printing received data - todo delete
	fmt.Fprintf(w, "Received: %s, %s, %s, %s, %s", name, age, color, healthy, id)
}
