package controllers

import (
	"fmt"
	"net/http"

	models "github.com/nocodeleaks/quepasa/models"
)

//region CONTROLLER - Message

func RevokeController(w http.ResponseWriter, r *http.Request) {

	// setting default reponse type as json
	w.Header().Set("Content-Type", "application/json")

	response := &models.QpResponse{}

	server, err := GetServer(r)
	if err != nil {
		response.ParseError(err)
		RespondInterface(w, response)
		return
	}

	// Default parameters
	messageid := GetMessageId(r)

	if len(messageid) == 0 {
		err = fmt.Errorf("empty message id")
		response.ParseError(err)
		RespondInterface(w, response)
		return
	} else {

		err = server.Revoke(messageid)
		if err != nil {
			response.ParseError(err)
			RespondInterface(w, response)
			return
		}

		response.ParseSuccess("revoked with success")
		RespondSuccess(w, response)
	}
	return
}

//endregion
