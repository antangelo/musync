package spotify

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/zmb3/spotify"
)

func GetPlaylists() error {
	auth := NewAuthenticator(spotify.ScopePlaylistReadPrivate)
	token := auth.authenticate()
	client := auth.Auth.NewClient(token)

	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		return errors.Wrap(err, "could not get playlists")
	}

	playlistNames := []string{}
	for _, playlist := range playlists.Playlists {
		playlistNames = append(playlistNames, playlist.Name)
	}

	bytes, err := json.Marshal(playlistNames)
	if err != nil {
		return errors.Wrap(err, "unable to pretty print")
	}
	fmt.Println(string(bytes))

	return nil
}
