package middlewares

import (
	"net/http"
	"time"
)

type timingResponseWriter struct {
	http.ResponseWriter
	start    time.Time
	wroteHdr bool
}

// Wrapper to WriteHeader
func (tw *timingResponseWriter) WriteHeader(statusCode int) {
	if !tw.wroteHdr {
		duration := time.Since(tw.start)
		tw.Header().Set("X-MS-Duration-Timing", duration.String())
		tw.wroteHdr = true
	}
	tw.ResponseWriter.WriteHeader(statusCode)
}

// Wrapper to Write, because I need to call WriteHeader
func (tw *timingResponseWriter) Write(b []byte) (int, error) {
	if !tw.wroteHdr {
		tw.WriteHeader(http.StatusOK)
	}
	return tw.ResponseWriter.Write(b)
}

func DurationTiming(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tw := &timingResponseWriter{
			ResponseWriter: w,
			start:          time.Now(),
		}

		next.ServeHTTP(tw, r)
	})
}
