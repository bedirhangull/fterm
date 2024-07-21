package cmd

import (
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fterm [command]",
	Short: "A CLI tool for managing shortcuts",
	Long:  `fTerm is a CLI tool for managing shortcuts.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			displayTitle()
			cmd.Help()
		} else {
			findCommand(cmd, args)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func displayTitle() {
	fTermTitle := figure.NewColorFigure("fTerm", "isometric3", "green", true)
	fTermTitle.Print()
	println("\n\n")
}
