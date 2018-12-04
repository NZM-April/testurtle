package cmd

import (
	"fmt"

	"github.com/NZM-April/testurtle/turtle"
	"github.com/spf13/cobra"
)

type Options struct {
    config string
}

var (
    o = &Options{}
)

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&o.config, "file", "f", "./testurtle.json", "config file")
}

var runCmd = &cobra.Command{
    Use:   "run",
    Short: "Start Testing.",
    Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config file: %s", o.config)
		turtle.Start()
	},
}