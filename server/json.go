package server

import (
	"encoding/json"
	"net/http"
)

func jsonRes(w http.ResponseWriter, r *http.Request) {

	//w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Message sent"
	/*	jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp) */
	json.NewEncoder(w).Encode(resp)
}
