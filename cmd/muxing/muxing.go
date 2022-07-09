package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{param}", getNameParamHandler).Methods("GET")
	router.HandleFunc("/bad", getBadHandler).Methods("GET")
	router.HandleFunc("/data", postDataHandler).Methods("POST")
	router.HandleFunc("/headers", postHeadersHandler).Methods("POST")
	router.HandleFunc("/", notDefinedHandler)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func getNameParamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	response := "Hello, " + vars["param"] + "!"
	w.Write([]byte(response))
}

func getBadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func postDataHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err == nil {
		response := "I got message:\n" + string(b)
		w.Write([]byte(response))
	}
}

func postHeadersHandler(w http.ResponseWriter, r *http.Request) {
	h := r.Header

	if a, ok := h["A"]; ok {
		if b, ok := h["B"]; ok {
			ai, err := strconv.Atoi(a[0])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			bi, err := strconv.Atoi(b[0])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			w.Header().Set("a+b", strconv.Itoa(ai+bi))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
func notDefinedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
