/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/carlosm27/restli/cmd/methods"
	"github.com/spf13/cobra"
)

// DELETECmd represents the DELETE command
var DELETECmd = &cobra.Command{
	Use:   "DELETE",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DELETE called")

		methods.Delete(url)
	},
}

var urlDelete string

func init() {
	rootCmd.AddCommand(DELETECmd)

	DELETECmd.Flags().StringVarP(&url, "url", "u", "", "URL of API you want to request")

}
