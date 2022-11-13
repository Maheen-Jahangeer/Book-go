package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) statusHelper(w http.ResponseWriter, r *http.Request) {
	app.logger.Print(app.config)
	currentStatus := Appstatus{
		Status:      "avaiable",
		Environment: app.config.env,
	}
	responseJson, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		app.logger.Print("error :", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
