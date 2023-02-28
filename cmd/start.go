/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"api-server/api"
	"github.com/spf13/cobra"
)

var Port int

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the server",
	Long:  `choose port by -p flag or default port is 8080`,
	Run: func(cmd *cobra.Command, args []string) {
		api.StartServer(Port)
	},
}

func init() {
	startCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8080, "default port for http server")
	rootCmd.AddCommand(startCmd)
}
