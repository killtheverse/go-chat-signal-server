package response

// Response to be returned to a http request
type Response struct {
    // Status of the Response
    Status          int         `json:"status"`

    // Message in the Response
    Message         string      `json:"message"`

    // Content to be displayed
    Content         interface{}   `json:"content"`
}

