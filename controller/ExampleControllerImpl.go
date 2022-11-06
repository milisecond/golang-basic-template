package controller

import (
	"encoding/json"
	"golang-basic-http/model/request"
	"io/ioutil"
	"net/http"
)

type ExampleControllerImpl struct{}

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
	message := "Oops the method is doesn exist"
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
		w.Write([]byte(message))
	}
}
