package main

import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

type Payment struct {
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
}

func (p Payment) String() string {
	return fmt.Sprintf("%s: %.2f", p.Name, p.Amount)
}

type Payments struct {
	Id    string    `json:"id"`
	Title string    `json:"title"`
	Items []Payment `json:"items"`
}

func (p Payments) String() string {
	return fmt.Sprintf("%s\n%s\n", p.Title, p.Items)
}

type PaymentsHandler struct {
	Storage PaymentStorage
}

func generateNewId() string {
	return uuid.NewV4().String()
}

func NewPaymentsHandler(storage PaymentStorage) *PaymentsHandler {
	return &PaymentsHandler{
		Storage: storage,
	}
}

func (ph *PaymentsHandler) Update(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (ph *PaymentsHandler) Get(id string) (Payments, error) {
	return ph.Storage.Get(id)
}
func (ph *PaymentsHandler) Create(p Payments) (Payments, error) {
	id := generateNewId()
	p.Id = id
	if _, err := ph.Storage.Set(id, p); err != nil {
		return p, err
	}
	return p, nil
}
