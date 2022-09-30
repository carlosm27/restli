/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"encoding/json"

	"github.com/carlosm27/restli/cmd/methods"
	"github.com/spf13/cobra"
)

// POSTCmd represents the POST command
var POSTCmd = &cobra.Command{
	Use:   "POST",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("POST called")

		json.NewEncoder(os.Stdout).Encode(urlPost)

		methods.Post(urlPost, bodyPost)
	},
}

var (
	urlPost  string
	bodyPost string
)

func init() {
	rootCmd.AddCommand(POSTCmd)

	POSTCmd.Flags().StringVarP(&urlPost, "url", "u", "", "URL of API you want to request")
	POSTCmd.Flags().StringVarP(&bodyPost, "body", "b", "", "The body you want to post")

	
}
