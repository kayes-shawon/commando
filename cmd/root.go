package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "Commando",
	Short: "A Command Line Application",
	Long: `Application that will make life easier`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	//rootCmd.AddCommand(rootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}