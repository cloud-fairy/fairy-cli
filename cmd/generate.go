/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"cloudfairy/templates"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate scaffold files for a cloudfairy module",
	Example: `fairy generate component acme/managed-container
fairy generate connector acme/conn-container-database
	`,
	Aliases: []string{"g", "gen"},
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cType := args[0]
		cOut := args[1]
		switch cType {
		case "component", "c", "comp":
			fmt.Println("Generating component in", cOut)
			templates.Write("component", cOut)
			break
		case "connector", "cn", "con":
			fmt.Println("Generating connector in", cOut)
			templates.Write("connector", cOut)
			break
		default:
			fmt.Println("Args", args)
			break
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
