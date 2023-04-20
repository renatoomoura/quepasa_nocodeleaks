package models

import (
	. "github.com/nocodeleaks/quepasa/whatsapp"
)

// Parameters to be acessed/passed on Views (receive.tmpl)
type QPFormReceiveData struct {
	PageTitle           string
	ErrorMessage        string
	Number              string
	Token               string
	DownloadPrefix      string
	FormAccountEndpoint string
	Messages            []WhatsappMessage
}
