package main

import (
	"chat/internal/client/app"
	clienthttp "chat/internal/client/http"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "chat-client",
	Short: "Chat app client",
	Long:  "Chat app client",
}

var startClient = &cobra.Command{
	Use:   "start",
	Short: "Serve client app files",
	Long:  `Serve client app files`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrapClient()
	},
}

func bootstrapClient() {
	fmt.Println("Starting client...")
	app := app.NewApplication()
	clienthttp.BootstrapServer(app)
}

func main() {

	rootCmd.AddCommand(startClient)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
