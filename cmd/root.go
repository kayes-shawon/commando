package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use: "commando",
	Short: "A Command Line Application",
	Long: `Command Line Application that will make life easier`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		_ = cmd.Help()
		//cmd.AddCommand(testCmd)
		os.Exit(0)
	},
}

var testCmd = &cobra.Command{
	Use: "test",
	Short: "Testing",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(cmd.Flags())
		fmt.Println("Testing cobra command, name: ", cmd.Flag("author").Value)
	},
}

var timeCmd = &cobra.Command{
	Use: "time",
	Short: "Time Now",
	Long: `Representing local time`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "utc" {
			fmt.Println(time.Now())
			fmt.Println(time.Now().UTC().Format(time.RFC3339Nano))
		}
	},
}




func initConfig() {
	fmt.Println("inside initConfig")
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(timeCmd)
	var author string

	//cmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")

	testCmd.Flags().StringVarP(&author, "author", "a", "", "use")
	testCmd.MarkFlagRequired("author")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}