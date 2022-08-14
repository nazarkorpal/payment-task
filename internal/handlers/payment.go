package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func PaymentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parsedTemplate, _ := template.ParseFiles("templates/payment.html")
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			log.Println("Error executing template :", err)
			return
		}
	}
}
