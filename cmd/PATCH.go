/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// PATCHCmd represents the PATCH command
var PATCHCmd = &cobra.Command{
	Use:   "PATCH",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("PATCH called")

		json.NewEncoder(os.Stdout).Encode(urlPatch)

		methods.Patch(urlPatch, bodyPatch)
	},
}


var (
	urlPatch  string
	bodyPatch string
)

func init() {
	rootCmd.AddCommand(PATCHCmd)

	PATCHCmd.Flags().StringVarP(&urlPatch, "url", "u", "", "URL of API you want to request")

	PATCHCmd.Flags().StringVarP(&bodyPatch, "body", "b", "", "The body you want to patch")


}
