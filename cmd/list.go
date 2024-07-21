package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all commands",
	Long:  `List all commands and their shortcuts.`,
	Run:   listCommands,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listCommands(cmd *cobra.Command, args []string) {
	commands, err := loadCommands()
	if err != nil {
		fmt.Println("Error loading commands:", err)
		return
	}

	maxShortcutLen := len("Shortcut")
	maxCommandLen := len("Command")
	maxDescriptionLen := len("Description")

	for _, command := range commands {
		if len(command.Shortcut) > maxShortcutLen {
			maxShortcutLen = len(command.Shortcut)
		}
		if len(command.Command) > maxCommandLen {
			maxCommandLen = len(command.Command)
		}
		if len(command.Description) > maxDescriptionLen {
			maxDescriptionLen = len(command.Description)
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	fmt.Fprintf(w, "+-%s-+-%s-+-%s-+\n",
		strings.Repeat("-", maxShortcutLen),
		strings.Repeat("-", maxCommandLen),
		strings.Repeat("-", maxDescriptionLen))
	fmt.Fprintf(w, "| %-*s | %-*s | %-*s |\n",
		maxShortcutLen, "Shortcut",
		maxCommandLen, "Command",
		maxDescriptionLen, "Description")
	fmt.Fprintf(w, "+-%s-+-%s-+-%s-+\n",
		strings.Repeat("-", maxShortcutLen),
		strings.Repeat("-", maxCommandLen),
		strings.Repeat("-", maxDescriptionLen))

	for _, command := range commands {
		fmt.Fprintf(w, "| %-*s | %-*s | %-*s |\n",
			maxShortcutLen, command.Shortcut,
			maxCommandLen, command.Command,
			maxDescriptionLen, command.Description)
		fmt.Fprintf(w, "+-%s-+-%s-+-%s-+\n",
			strings.Repeat("-", maxShortcutLen),
			strings.Repeat("-", maxCommandLen),
			strings.Repeat("-", maxDescriptionLen))
	}

	w.Flush()
}
