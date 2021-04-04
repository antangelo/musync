package spotify

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

type SpotifyAuthenticator struct {
	Auth      *spotify.Authenticator
	TokenChan chan *oauth2.Token
}

func NewAuthenticator(scopes ...string) *SpotifyAuthenticator {
	callback := os.Getenv("SPOTIFY_CALLBACK")
	log.Printf("Using callback URL: %v", callback)
	auth := spotify.NewAuthenticator(callback, scopes...)
	tokenChan := make(chan *oauth2.Token, 1)

	return &SpotifyAuthenticator{
		Auth:      &auth,
		TokenChan: tokenChan,
	}
}

func (s *SpotifyAuthenticator) redirectCallback(w http.ResponseWriter, r *http.Request) {
	token, err := s.Auth.Token("state", r)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	s.TokenChan <- token
	fmt.Fprintf(w, "Authentication complete, safe to close window")
}

func (s *SpotifyAuthenticator) runAuthCallbackServer(done <-chan bool) {
	srv := &http.Server{Addr: ":8080"}
	http.HandleFunc("/callback", s.redirectCallback)
	log.Print("Starting callback server")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("fatal error serving: %v", err)
		}
	}()

	<-done
	log.Print("Stopping callback server")
	srv.Shutdown(context.Background())
}

func (s *SpotifyAuthenticator) authenticate() *oauth2.Token {
	done := make(chan bool)
	go s.runAuthCallbackServer(done)
	url := s.Auth.AuthURL("state")
	log.Printf("Go to this URL: %v", url)

	token := <-s.TokenChan
	done <- true
	return token
}

func (s *SpotifyAuthenticator) Login() error {
	token := s.authenticate()
	log.Printf("Token: %v", token)

	return nil
}
