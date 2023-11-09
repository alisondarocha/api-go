package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	Birthdate time.Time `json:"birthdate"`
	Email     string    `json:"email"`
}

func (client Client) GetFullInfoClient() string {
	return fmt.Sprintf("Id: %s, Name: %s, LastName: %s, Birthdate: %s, Email: %s",
		client.Id.String(), client.Name, client.LastName, client.Birthdate, client.Email)
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	client := Client{
		Id:        uuid.New(),
		Name:      "João",
		LastName:  "Pereira",
		Birthdate: time.Date(1999, time.October, 7, 0, 0, 0, 0, time.UTC),
		Email:     "joaopereira@email.com",
	}

	json.NewEncoder(w).Encode(client)
}
