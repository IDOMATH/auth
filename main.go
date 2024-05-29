package main

import (
	"fmt"
	"github.com/IDOMATH/session/memorystore"
	"net/http"
)

type Repository struct {
	memstore *memorystore.MemoryStore
}

func NewRepository() *Repository {
	return &Repository{}
}

func main() {
	fmt.Println("hello world")
	memstore := memorystore.New()
	repo := NewRepository()
	repo.memstore = memstore
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Home"))
}
