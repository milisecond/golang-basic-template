package router

import (
	"golang-basic-http/controller"
	"net/http"
)

func App() {
	controller := controller.NewExampleController()

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/person", controller.Store)
}
