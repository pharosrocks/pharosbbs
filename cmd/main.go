package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/pharosrocks/pharosbbs/bbs"
	"github.com/spf13/cobra"
)

var (
	// Version contains the application version number. It's set via ldflags
	// when building.
	Version = "dev"

	// CommitSHA contains the SHA of the commit that this application was built
	// against. It's set via ldflags when building.
	CommitSHA = ""

	description = "PharosBBS: a bbs over websockets/ssh, a mastodon compatible instance."

	ssh = -1
	ws  = -1
	wss = -1

	rootCmd = &cobra.Command{
		Use:   "help",
		Short: "help",
		Long:  description,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Long:  description,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%v\n", Version)
		},
	}

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start server",
		Long:  description,
		RunE: func(cmd *cobra.Command, args []string) error {
			bbs := bbs.NewServer()
			godotenv.Load()
			return bbs.ListenAndServe(":8080")
		},
	}
)

func main() {
	startCmd.Flags().IntVarP(&ssh, "ssh", "s", 2222, "port for ssh service.")
	startCmd.Flags().IntVarP(&ws, "ws", "w", 8080, "port for websockets service.")
	startCmd.Flags().IntVarP(&wss, "wss", "W", 8443, "port for websockets over tls service.")

	rootCmd.CompletionOptions.HiddenDefaultCmd = false
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.Execute()
}
