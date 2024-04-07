package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "chat-server",
	Short: "Chat app server",
	Long:  "Chat app server",
}

var startServer = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long:  `Start server`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrapServer()
	},
}

func bootstrapServer() {
	fmt.Println("Start server!")
}

func main() {

	rootCmd.AddCommand(startServer)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
