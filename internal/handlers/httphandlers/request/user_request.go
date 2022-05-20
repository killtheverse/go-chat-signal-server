package request

// UserRegisterRequest represents the http request format for registering a new user
type UserRegisterRequest struct {
    Username        string          `json:"username"`
    Password        string          `json:"password"`
    Name            string          `json:"name"`
}

// UserLoginRequest represents the http request format for logging a user in 
type UserLoginRequest struct {
    Username        string          `json:"username"`
    Password        string          `json:"password"`
}
