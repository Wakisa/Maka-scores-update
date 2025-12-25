/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"www.github.com/Wakisa/maka/internal/api"
	"www.github.com/Wakisa/maka/internal/config"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the Fiber API server",

	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load()
		api := api.NewAPI(cfg)
		api.Start(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
