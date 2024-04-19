package main

import (
	"chat/internal/app"
	chat_http "chat/internal/http"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat app",
	Long:  "Chat app",
}

var start = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long:  `Start server`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap()
	},
}

func bootstrap() {
	log.Println("Starting app...")
	app := app.NewApplication()
	chat_http.BootstrapServer(app)
}

func main() {

	rootCmd.AddCommand(start)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
