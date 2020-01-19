package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"log"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/dig"
)

type Config struct {
	Enabled bool
	DatabasePath string
	Port string
}



type Person struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

type PersonRepository struct {
	database *sql.DB
}

func (r *PersonRepository) FindAll() []*Person {
	rows, err := r.database.Query(`Select id, name, age FROM people;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	people := []*Person{}

	for rows.Next() {
		var (
			id int
			name string
			age int
		)

		rows.Scan(&id, &name, &age)

		people = append(people, &Person{
			Id: id,
			Name: name,
			Age: age,
		})
	}

	return people
}


type PersonService struct {
	config *Config
	repository *PersonRepository
}

func (service *PersonService) FindAll() []*Person {
	if service.config.Enabled {
		return service.repository.FindAll()
	}

	return []*Person{}
}


type Server struct {
	config *Config
	personService *PersonService
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/people", s.findPeople)

	return mux
}

func (s *Server) Hello() {
	fmt.Println("Hello")
}

func (s *Server) Run() {
	httpServer := &http.Server {
		Addr: ":" + s.config.Port,
		Handler: s.Handler(),
	}

	httpServer.ListenAndServe()
}

func (s *Server) findPeople(w http.ResponseWriter, r *http.Request) {
	people := s.personService.FindAll()
	bytes, _ := json.Marshal(people)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func NewConfig() *Config {
	return &Config{
		Enabled: true,
		DatabasePath: "./example.db",
		Port: "8000",
	}
}

func NewDatabase(c *Config) (*sql.DB, error) {
	return sql.Open("sqlite3", c.DatabasePath)
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{database: db}
}

func NewPersonService(c *Config, r *PersonRepository) *PersonService {
	return &PersonService{config: c, repository: r}
}

func NewServer(c *Config, p *PersonService) *Server {
	return &Server{
		config: c,
		personService: p,
	}
}

func BuildContainer() *dig.Container {
	c := dig.New()

	c.Provide(NewServer)
	c.Provide(NewPersonService)
	c.Provide(NewPersonRepository)
	c.Provide(NewDatabase)
	c.Provide(NewConfig)

	return c
}

func main() {
	c := BuildContainer()

	err := c.Invoke(func(s *Server) {
		s.Hello()
	})

	if err != nil {
		log.Fatal(err)
	}

	err = c.Invoke(func(s *Server) {
		s.Run()
	})

	if err != nil {
		log.Fatal(err)
	}
}