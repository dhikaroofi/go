package helper

import (
	"encoding/json"
	"net/http"
)

type defaultResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Warning interface{} `json:"warning"`
	Data    interface{} `json:"data"`
}

func RespondSuccess(w http.ResponseWriter, status int, msg string, payload interface{}) {
	res := defaultResponse{}
	res.Status = true
	res.Message = msg
	res.Data = payload
	res.respondJSON(w, status, payload)
}
func RespondError(w http.ResponseWriter, status int, msg error, payload interface{}) {
	res := defaultResponse{}
	res.Status = false
	res.Message = msg.Error()
	res.Data = payload
	res.respondJSON(w, status, payload)
}

func (res *defaultResponse) respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func CorsHelper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,DELETE,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-CSRF-Token,Authorization")
	w.Header().Set("Access-Control-Max-Age", "7200")
}
