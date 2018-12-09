package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use:   "testurtle",
	Short: "testurtle is the micro infrastructure tester.",
	Long:  "testurtle is the micro infrastructure tester. That's all.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
            cmd.Help()
            os.Exit(0)
        }
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command {
	Use:   "version",
	Short: "Print the version number of testurtle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testurtle v0.0.1")
	},
}
