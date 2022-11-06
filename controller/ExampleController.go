package controller

import "net/http"

type exampleInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
}

func NewExampleController() exampleInterface {
	return ExampleControllerImpl{}
}
