package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all commands",
	Long:  `Clear all commands from the .commands.json ile.`,
	Run:   clearCommandList,
}

func init() {
	rootCmd.AddCommand(clearCmd)
}

func clearCommandList(cmd *cobra.Command, args []string) {
	err := saveCommands([]ShortcutCommand{})
	if err != nil {
		fmt.Println("Error clearing commands:", err)
		return
	}

	fmt.Println("Commands cleared successfully.")
}
