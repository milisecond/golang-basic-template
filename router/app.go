package router

import (
	"golang-basic-http/controller"
	"net/http"
)

func App() {
	controller := controller.NewExampleController()

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/body", controller.Store)
	http.HandleFunc("/query", controller.Find)
	http.HandleFunc("/multipart", controller.Multipart)
}
