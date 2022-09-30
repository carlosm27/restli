/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/carlosm27/restli/cmd/methods"
	"github.com/spf13/cobra"
)

// PUTCmd represents the PUT command
var PUTCmd = &cobra.Command{
	Use:   "PUT",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("PUT called")

		json.NewEncoder(os.Stdout).Encode(bodyPut)

		methods.Post(urlPut, bodyPut)

	},
}

var (
	urlPut  string
	bodyPut string
)

func init() {
	rootCmd.AddCommand(PUTCmd)

	PUTCmd.Flags().StringVarP(&urlPut, "url", "u", "", "URL of API you want to request")
	PUTCmd.Flags().StringVarP(&bodyPut, "body", "b", "", "The body you want to put")
}
