package response

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, headers map[string]string, status int, response any) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(status)

	if response == nil {
		return
	}
	body, err := json.Marshal(response)
	if err != nil {
		return
	}

	_, err = w.Write(body)
	if err != nil {
		return
	}
}
