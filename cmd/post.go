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

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		json.NewEncoder(os.Stdout).Encode(body)
		fmt.Printf("The body is: ", body)
		methods.Post(urlPost, body)

	},
}

var urlPost string

var body string

func init() {
	rootCmd.AddCommand(postCmd)

	postCmd.Flags().StringVarP(&urlPost, "url", "u", "", "URL of API you want to request")
	postCmd.Flags().StringVarP(&body, "body", "b", "", "The body you want to post")

}
