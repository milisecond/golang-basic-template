package main

import (
	"fmt"
	"golang-basic-http/conf"
	"golang-basic-http/router"
	"net/http"
)

func init() {
	router.App()
}

func main() {
	fmt.Printf("Server running at %s", conf.ADDR)

	err := http.ListenAndServe(conf.ADDR, nil)
	if err != nil {
		panic(err)
	}
}
