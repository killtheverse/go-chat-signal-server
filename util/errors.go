package util


type FieldError struct {
    Field           string          `json:"field"`
    Message         string          `json:"message"`
}

func (fe *FieldError) Error() string {
    return fe.Message
}

type FieldErrors struct {
    InvalidFields   []FieldError    `json:"invalidFields"`
}


type DataBaseError struct {
    Message         string          `json:"message"`
}

func (dbe *DataBaseError) Error() string {
    return dbe.Message
}

type UsernameNotAvailableError struct {
    Message         string          `json:"message"`
}

func (e *UsernameNotAvailableError) Error() string {
    return e.Message
}

type ServerError struct {
    Message         string          `json:"message"`
}

func (e *ServerError) Error() string {
    return "Internal Server Error"
}
