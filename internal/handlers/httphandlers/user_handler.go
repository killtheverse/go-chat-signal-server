package httphandlers

import (
	"encoding/json"
	"net/http"

	"github.com/killtheverse/go-chat-signal-server/internal/core/ports"
	"github.com/killtheverse/go-chat-signal-server/internal/handlers/httphandlers/request"
)

type UserHandler struct {
    userService ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandler {
    return &UserHandler{
        userService: userService,
    }
}

func (h *UserHandler) Login(rw http.ResponseWriter, r *http.Request) error {
    return nil
}

func (h *UserHandler) Register(rw http.ResponseWriter, r *http.Request) error {
    // Parse request body and extract the user
    var userRequest request.UserRegisterRequest
    err := json.NewDecoder(r.Body).Decode(&userRequest)
    if err != nil {
        WriteResponse(rw, http.StatusBadRequest, "Invalid JSON body", nil)
    }

    err = h.userService.Register(userRequest.Username, userRequest.Password, userRequest.Name)
    if err != nil {
        // TODO: Implement error handling
        WriteResponse(rw, http.StatusBadRequest, "Error", err)
    }
    return nil
}
