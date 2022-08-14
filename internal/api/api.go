package api

import (
	"net/http"

	"github.com/nazarkorpal/payment-task/internal/handlers"
	"github.com/nazarkorpal/payment-task/internal/middlewares"
)

func Start() {
	router := http.NewServeMux()

	router.HandleFunc("/payments/", handlers.PaymentHandler())
	router.HandleFunc("/products/", handlers.ProductHandler())

	http.ListenAndServe(":8080", middlewares.Headers(router))
}
