package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("faster one wins ", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)

		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("return an erro if a server doesn;t response within 10s", func(t *testing.T) {
		aServer := makeDelayedServer(25 * time.Millisecond)
		bServer := makeDelayedServer(26 * time.Millisecond)

		defer aServer.Close()
		defer bServer.Close()

		_, err := ConfigurableRacer(aServer.URL, bServer.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
