package webrest

//Result - holds result message of a non GET api
type Result struct {
	Operation string `json:"operation"`
	Message   string `json:"message"`
	Error     error  `json:"error"`
}
