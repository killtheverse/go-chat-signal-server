package ports

import "net/http"

type UserService interface {
    Login(username string, password string) error
    Register(username string, password string, name string) error
}

type UserRepository interface {
    Login(username string, password string) error
    Register(username string, password string, name string) error
}

type UserHandlers interface {
    Login(rw http.ResponseWriter, r *http.Request)
    Register(rw http.ResponseWriter, r *http.Request)
}

