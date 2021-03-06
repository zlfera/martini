package martini

import (
	"log"
	"net/http"
	"time"
)

// Logger returns a middleware handler that logs the request as it goes in and the response as it goes out.
func Logger() Handler {
	return func(res http.ResponseWriter, req *http.Request, c Context, log *log.Logger) {
		start := time.Now()
		log.Printf("Started %s %s", req.Method, req.URL.Path)

		rw := NewResponseWriter(res)
		c.MapTo(rw, (*http.ResponseWriter)(nil))

		c.Next()

		log.Printf("Completed %v %s in %v\n", rw.Status(), http.StatusText(rw.Status()), time.Now().Sub(start))
	}
}
