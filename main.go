package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// root GET "/"
func root(w http.ResponseWriter, r *http.Request) {
	log := fmt.Sprintf("%s %s %s /", r.Proto, r.Method, r.Host)
	fmt.Println(log, http.StatusOK)

	fmt.Fprintf(w, "Hello World!")
}

// routes
func routes() {
	r := mux.NewRouter() //.StrictSlash(true)
	r.HandleFunc("/", root).Methods("GET")
	r.HandleFunc("/posts", Index).Methods("GET")
	r.HandleFunc("/posts", Create).Methods("POST")
	r.HandleFunc("/posts/{id}", Update).Methods("PUT")
	r.HandleFunc("/posts/{id}", Show).Methods("GET")
	r.HandleFunc("/posts/{id}", Delete).Methods("DELETE")

	// This could be bad CORS, don't put into prod without looking into these settings!
	corsObj := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(corsObj, headersOk)(r)))
}

// getEnv key to search; value returned if key is not found
func getEnv(key, value string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return value
	}
	return val
}

var db *gorm.DB
var err error

func main() {
	// Secrets
	e := godotenv.Load()
	if e != nil {
		panic("no dotenv")
	}

	// Postgres connection
	pgUser := getEnv("PGUSER", "postgres")
	pgPass := getEnv("PGPASS", "password")
	dbName := getEnv("DBNAME", "blog")
	dbHost := getEnv("DBHOST", "localhost")
	uri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, pgUser, dbName, pgPass)
	db, err = gorm.Open("postgres", uri)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	db.Debug().AutoMigrate(&Post{})

	// Handle routes
	fmt.Println("* Running on http://localhost:8081/ (Press CTRL+C to quit)")
	routes()
}
