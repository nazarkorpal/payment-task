package middlewares

import (
	"net/http"
	"strconv"
	"time"
)

type responseWithHeaders struct {
	http.ResponseWriter

	isHeaderWritten bool
	start           time.Time
	host            string
}

func Headers(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(&responseWithHeaders{w, false, time.Now(), r.Host}, r)
	})
}

func (w *responseWithHeaders) WriteHeader(statusCode int) {
	duration := time.Since(w.start)
	us := int(duration.Truncate(1000*time.Millisecond).Microseconds() / 1000)
	w.Header().Set("X-Response-Time", strconv.Itoa(us)+" us")
	w.Header().Set("X-Server-Name", w.host)

	w.ResponseWriter.WriteHeader(statusCode)
	w.isHeaderWritten = true
}

func (w *responseWithHeaders) Write(b []byte) (int, error) {
	if !w.isHeaderWritten {
		w.WriteHeader(200)
	}

	return w.ResponseWriter.Write(b)
}
