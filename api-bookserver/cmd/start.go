package cmd

import (
	"github.com/souravbiswassanto/client-go/api-bookserver/Deployment"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var (
	// Port stores port number for starting a connection
	Port     int
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start cmd starts the server on a port",
		Long: `It starts the server on a given port number, 
				Port number will be given in the cmd`,
		Run: func(cmd *cobra.Command, args []string) {
			Deployment.BookServerDeployment()
		},
	}
)

func init() {

	startCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8081, "Port number for starting server")
	rootCmd.AddCommand(startCmd)
}
