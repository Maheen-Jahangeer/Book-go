package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) jsonParser(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)

	if err != nil {
		app.logger.Print("error :", err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) jsonError(w http.ResponseWriter, status int, err error) {
	type ErrorMessage struct {
		Message string `json:"message"`
	}

	theError := ErrorMessage{
		Message: err.Error(),
	}
	app.logger.Print("error from function ", theError)
	app.jsonParser(w, http.StatusBadRequest, theError, "error")
}
