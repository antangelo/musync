package cmd

import (
	"github.com/antangelo/musync/spotify"
	"github.com/spf13/cobra"
)

var (
	spotifyRootCmd = &cobra.Command{
		Use:   "spotify",
		Short: "Tools for managing spotify playlists",
	}
	spotifyAuthCmd = &cobra.Command{
		Use:   "login",
		Short: "Retrieve a login token from spotify",
		RunE: func(_ *cobra.Command, _ []string) error {
			auth := spotify.NewAuthenticator()
			return auth.Login()
		},
	}
	spotifyGetPlaylistsCmd = &cobra.Command{
		Use:   "ls-playlist",
		Short: "Retrieve a list of playlists from shopify",
		RunE: func(_ *cobra.Command, _ []string) error {
			return spotify.GetPlaylists()
		},
	}
)

func init() {
	spotifyRootCmd.AddCommand(spotifyAuthCmd)
	spotifyRootCmd.AddCommand(spotifyGetPlaylistsCmd)

	rootCmd.AddCommand(spotifyRootCmd)
}
