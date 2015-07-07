package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/redis.v2"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func renderAsJson(w http.ResponseWriter, v interface{}) {
	json, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func MainHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "main")
}

func setupPaymentsRoutes(r *mux.Router) {
	client := redis.NewTCPClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	store := &RedisStorage{
		Client: client,
	}
	ph := NewPaymentsHandler(store)

	r.
		Path("/api/payments/{id}").
		Methods("GET").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		p, err := ph.Get(id)
		if err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		renderAsJson(w, p)
	})

	r.
		Path("/api/payments").
		Methods("POST").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var p Payments
		err := decoder.Decode(&p)
		if err != nil {
			http.Error(w, "Bad data", http.StatusBadRequest)
			return
		}
		p, err = ph.Create(p)
		if err != nil {
			http.Error(w, "Err", http.StatusInternalServerError)
			return
		}
		renderAsJson(w, p)
	})
}

func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return port
	}
	return "3000"
}

func main() {
	mainRouter := mux.NewRouter().StrictSlash(true)

	setupPaymentsRoutes(mainRouter)
	n := negroni.Classic()
	n.UseHandler(mainRouter)
	n.Run(":" + getPort())
}
