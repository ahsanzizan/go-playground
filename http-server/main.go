package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users map[string]User
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateId(n int) string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Responding to GET /")

	usersList := make([]User, 0, len(users))

	for _, user := range users {
		usersList = append(usersList, user)
	}

	res, err := json.Marshal(usersList)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Responding to POST /")

	var newUser User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	newUser.ID = generateId(12)
	users[newUser.ID] = newUser

	res, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func main() {
	users = make(map[string]User)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUsers(w, r)
		case http.MethodPost:
			createUser(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	const port string = ":8080"
	fmt.Printf("Server started on %s\n", port)
	err := http.ListenAndServe(port, mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
