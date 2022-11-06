package controller

import (
	"encoding/json"
	"golang-basic-http/model/request"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ExampleControllerImpl struct{}

const (
	MESSAGE = "Oops the method is doesn exist"
)

func (e ExampleControllerImpl) Index(w http.ResponseWriter, r *http.Request) {
	message := "Oops the method is doesn exist"
	switch r.Method {
	case "GET":
		var response map[string]interface{}

		response = map[string]interface{}{
			"success": true,
			"message": "Ok",
			"data":    []string{},
		}

		w.WriteHeader(200)

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(message))

	}

}

func (e ExampleControllerImpl) Store(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var exampleModel request.ExampleRequest
		bodyReq, err := ioutil.ReadAll(r.Body)

		json.Unmarshal(bodyReq, &exampleModel)

		response := map[string]interface{}{
			"success": true,
			"message": "Ok",
			"data":    exampleModel,
		}

		w.WriteHeader(200)

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(MESSAGE))
	}
}

func (e ExampleControllerImpl) Find(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")

		response := map[string]interface{}{
			"success": true,
			"message": "Ok",
			"data": map[string]string{
				"id": id,
			},
		}

		w.WriteHeader(200)

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(MESSAGE))
	}

}

func (e ExampleControllerImpl) Multipart(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var data request.ExampleRequest
		data = request.ExampleRequest{}
		r.ParseMultipartForm(0)

		data.Name = r.FormValue("name")
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			panic(err.Error())
		}

		data.Age = age

		response := map[string]interface{}{
			"success": true,
			"message": "Ok",
			"data":    data,
		}

		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(response)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(MESSAGE))
	}
}
