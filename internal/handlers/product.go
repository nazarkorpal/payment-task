package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type response struct {
	Url string `json:"url"`
}

func ProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		productID := strings.TrimPrefix(r.URL.Path, "/products/")

		time.Sleep(time.Second * 2)

		data := response{
			Url: fmt.Sprintf("http://localhost:8080/payments/%s", productID),
		}

		json.NewEncoder(w).Encode(data)
	}
}
