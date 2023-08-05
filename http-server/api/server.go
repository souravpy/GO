package api
import (
    "github.com/google/uuid"
    "github.com/gorilla/mux"
   
)

type Item struct {
  ID uuid.UUID 'json:"id"'
  Name string 'json:"name"'
}

type Server struct {
  *mux.Router

  shoppinItems []Item
}

func Newserver() *Server {
  s := &Server {
    Router: mux.NewRouter(),
    shoppinItems: []Item{},
  }
  return s
}

func ()
