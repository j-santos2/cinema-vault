package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/j-santos2/cinema-vault/internal/assert"
)

func TestRateLimit(t *testing.T) {
	t.Run("Test request under limit ok", func(t *testing.T) {
		app := application{
			config: config{
				limiter: struct {
					rps     float64
					burst   int
					enabled bool
				}{
					rps:     2,
					burst:   3,
					enabled: true,
				},
			},
		}
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		app.rateLimit(next).ServeHTTP(rr, r)

		assert.Equal(t, rr.Code, 200)
	})

	t.Run("Test request rate limiting disabled ok", func(t *testing.T) {
		app := application{
			config: config{
				limiter: struct {
					rps     float64
					burst   int
					enabled bool
				}{
					rps:     2,
					burst:   3,
					enabled: false,
				},
			},
		}
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		app.rateLimit(next).ServeHTTP(rr, r)

		assert.Equal(t, rr.Code, 200)
	})
}

func TestRateLimitIntegration(t *testing.T) {
	t.Run("Test request rate limiting disabled ok", func(t *testing.T) {
		app := application{
			config: config{
				limiter: struct {
					rps     float64
					burst   int
					enabled bool
				}{
					rps:     2,
					burst:   3,
					enabled: false,
				},
			},
		}

		server := httptest.NewServer(
			app.rateLimit(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})),
		)
		defer server.Close()

		status := []int{}

		for i := 0; i < 3; i++ {
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatalf("Failed to make GET request: %v", err)
			}
			defer resp.Body.Close()

			status = append(status, resp.StatusCode)
		}

		assert.EqualSlice(t, status, []int{http.StatusOK, http.StatusOK, http.StatusOK})
	})

	t.Run("Test burst requests get StatusTooManyRequests", func(t *testing.T) {
		app := application{
			config: config{
				limiter: struct {
					rps     float64
					burst   int
					enabled bool
				}{
					rps:     2,
					burst:   3,
					enabled: true,
				},
			},
		}

		server := httptest.NewServer(
			app.rateLimit(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})),
		)
		defer server.Close()

		status := []int{}

		for i := 0; i < 4; i++ {
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatalf("Failed to make GET request: %v", err)
			}
			defer resp.Body.Close()

			status = append(status, resp.StatusCode)
		}

		assert.EqualSlice(
			t,
			status,
			[]int{http.StatusOK, http.StatusOK, http.StatusOK, http.StatusTooManyRequests},
		)
	})
}
