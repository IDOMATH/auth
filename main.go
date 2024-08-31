package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IDOMATH/auth/db"
	"github.com/IDOMATH/auth/types"
	"github.com/IDOMATH/auth/util"
	"github.com/IDOMATH/session/memorystore"
	"golang.org/x/crypto/bcrypt"
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
	repo.userStore = db.NewUserStore(postgresDb.SQL)

	router := http.NewServeMux()
	router.HandleFunc("GET /", handleHome)
	router.HandleFunc("POST /user/", repo.handlePostUser)
	router.HandleFunc("GET /user/", repo.handleGetUser)

}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Home"))
}

func (repo *Repository) handlePostUser(w http.ResponseWriter, r *http.Request) {
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		w.Write([]byte("could not generate password hash"))
	}

	_, err = repo.userStore.InsertUser(types.User{Username: username, Email: email, PasswordHash: string(passwordHash)})
	if err != nil {
		util.Render(w, r, "error.go.html", &types.TemplateData{PageName: "Error"})
	}
	util.Render(w, r, "new-user.go.html", &types.TemplateData{PageName: "Success"})
}

func (repo *Repository) handleGetUser(w http.ResponseWriter, r *http.Request) {
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	_, err := repo.userStore.Authenticate(username, password)
	if err != nil {
		util.Render(w, r, "error.go.html", &types.TemplateData{PageName: "Error"})
	}
	util.Render(w, r, "success.go.html", &types.TemplateData{PageName: "Success"})

}

func (repo *Repository) auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAuth, err := repo.userStore.Authenticate(r.Form.Get("username"), r.Form.Get("password"))
		if err != nil {
			fmt.Println("unable to authenticate: ", err)
		}
		if isAuth > 0 {
			fmt.Println("Authenticated")
		}

		next(w, r)

	}
}
