package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-dep/repo"
	"github.com/go-dep/usecase"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Server struct {
	usecase usecase.UserService
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	u, err := s.usecase.GetUser(r.Context(), name)
	if err != nil {
		fmt.Println(err)
	}

	if err := json.NewEncoder(w).Encode(u); err != nil {
		// TODO: handle error
	}

}

func (s *Server) Post(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	err := s.usecase.AddUser(r.Context(), name)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	err := s.usecase.DeleteUser(r.Context(), name)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) Update(w http.ResponseWriter, r *http.Request) {
	oldName := r.FormValue("oldName")
	newName := r.FormValue("newName")

	err := s.usecase.UpdateUser(r.Context(), oldName, newName)

	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) Start() error {
	connStr := "user=uraulasevic password=postgres dbname=gotest sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	repo := repo.Repo{Conn: db}
	useCase := usecase.UserUsecase{UserRepo: repo}
	serv := Server{usecase: useCase}

	router := mux.NewRouter()
	router.HandleFunc("/users", serv.Get).
		Methods("GET")
	router.HandleFunc("/user/new", serv.Post).
		Methods("POST")
	router.HandleFunc("/user/update", serv.Update).
		Methods("PUT")
	router.HandleFunc("/user/delete", serv.Delete).
		Methods("DELETE")
	fmt.Println("starting server at :8080")
	return http.ListenAndServe(":8080", router)
}
