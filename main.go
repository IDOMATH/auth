package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IDOMATH/auth/db"
	"github.com/IDOMATH/session/memorystore"
)

type Repository struct {
	memstore  *memorystore.MemoryStore
	userStore *db.UserStore
}

func NewRepository() *Repository {
	return &Repository{}
}

func main() {

	dbHost := "localhost"
	dbPort := "5432"
	dbName := "portfolio"
	dbUser := "postgres"
	dbPass := "postgres"
	dbSsl := "disable"

	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", dbHost, dbPort, dbName, dbUser, dbPass, dbSsl)
	fmt.Println("Connecting to Postgres")
	postgresDb, err := db.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Postgres")

	fmt.Println("hello world")
	memstore := memorystore.New()
	repo := NewRepository()
	repo.memstore = memstore
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Home"))
}

func (repo *Repository) auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAuth, err := repo.userStore.Authenticate(r.Form.Get("username"), r.Form.Get("password"))
		if err != nil {
			fmt.Errorf("unable to authenticate: ", err)
		}
		if isAuth > 0 {
			fmt.Println("Authenticated")
		}

		next(w, r)

	}
}
