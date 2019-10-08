package main

import (
	"net/http"
)

func (serv *Server) HandleGetIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("!!! HELLO FROM REMOTE !!!"))
}
