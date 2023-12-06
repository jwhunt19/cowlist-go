package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

// create cors enabling middleware
type corsHandler struct{
	handler func(http.ResponseWriter, *http.Request)
}

func (c corsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.handler(w, r)
}

func EnableCors(handler func(http.ResponseWriter, *http.Request)) corsHandler {
	return corsHandler{handler}
}

// TODO: turn into method of a struct
// add new cow to database
func AddCow(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {

	// handles preflight options request
	if r.Method != http.MethodPost { // TODO: switch to && - if/else - repeat for other instances
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		} else {
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
			return
		}
	}

	// set data variable and decode request body
	var data cow
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		msg := fmt.Sprintf("failed to decode request: %s", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	// TODO: - r.Body.Close() - need to close body to reset reader


	// TODO: move db stuff to file

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// create update query with variables
	query := "INSERT INTO cows (name, age, color, healthy) VALUES ($1, $2, $3, $4);"

	// execute query
	_, err := conn.Exec(ctx, query, data.Name, data.Age, data.Color, data.Healthy)
	if err != nil {
		fmt.Println("error:", err) // TODO: return error to client - http.Error ?
	}

	w.WriteHeader(http.StatusCreated)
}

// TODO: turn into method of a struct
// return all cows in database
func GetAllCows(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// run a query to get all cows against the database
	rows, err := conn.Query(ctx, "select * from cows order by id")
	if err != nil {
		fmt.Printf("Query error: %v", err)
		return
	}
	defer rows.Close()

	// create cows variable which is a slice containing type cow
	var cows []cow

	// interate over the rows returned by the query
	for rows.Next() {
		var cow cow

		// Set the variables to the corresponding values found in the row
		err = rows.Scan(&cow.Id, &cow.Name, &cow.Age, &cow.Color, &cow.Healthy)
		if err != nil {
			fmt.Printf("Scan error: %v", err)
			return
		}

		// append this cow to the cows slice
		cows = append(cows, cow)
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

// TODO: turn into method of a struct
// update cow func
func UpdateCow(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {

	// handles preflight options request
	if r.Method != http.MethodPut { //TODO: 
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		} else {
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
			return
		}
	}

	// set data variable and decode request body
	var data cow
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		msg := fmt.Sprintf("failed to decode request: %s", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// query to update cow
	query := "UPDATE cows SET Name=$1, Age=$2, Color=$3, Healthy=$4 WHERE id=$5"

	// execute query
	_, err := conn.Exec(ctx, query, data.Name, data.Age, data.Color, data.Healthy, data.Id)
	if err != nil {
		fmt.Println("error:", err) // TODO: return error to client
	}

	w.WriteHeader(http.StatusAccepted)
}

// TODO: turn into method of a struct
// delete cow func
func DeleteCow(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {

	// handles preflight options request
	if r.Method != http.MethodDelete {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		} else {
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
			return
		}
	}

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// convert url to string, get id (numbers after last /)
	url := string(r.URL.Path)
	lastSlash := strings.LastIndex(url, "/")
	id := url[lastSlash+1:]

	// TODO: handle edge cases (empty id, NaN)

	// execute query
	_, err := conn.Exec(ctx, "delete from cows where id = $1", id)
	if err != nil {
		fmt.Println("error:", err) // TODO: return error to  client
	}

	w.WriteHeader(http.StatusNoContent)
}
