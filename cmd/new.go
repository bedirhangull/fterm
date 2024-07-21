package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type ShortcutCommand struct {
	Shortcut    string `json:"shortcut"`
	Command     string `json:"command"`
	Description string `json:"description"`
}

var addNewCommandCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new shortcut command",
	Long:  `Add a new shortcut command to the .commands.json file.`,
	RunE:  addNewCommand,
}

func init() {
	rootCmd.AddCommand(addNewCommandCmd)

	addNewCommandCmd.Flags().StringP("command", "c", "", "Command to execute")
	addNewCommandCmd.Flags().StringP("description", "d", "", "Description of the command")
	addNewCommandCmd.Flags().StringP("shortcut", "s", "", "Shortcut for the command")

	addNewCommandCmd.MarkFlagRequired("command")
	addNewCommandCmd.MarkFlagRequired("shortcut")
}

func addNewCommand(cmd *cobra.Command, args []string) error {
	command, err := cmd.Flags().GetString("command")
	if err != nil {
		return err
	}
	shortcut, err := cmd.Flags().GetString("shortcut")
	if err != nil {
		return err
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return err
	}

	commands, err := loadCommands()
	if err != nil {
		return fmt.Errorf("error loading commands: %v", err)
	}

	newCommand := ShortcutCommand{Shortcut: shortcut, Command: command, Description: description}
	if err := checkUniqCommand(commands, newCommand); err != nil {
		return err
	}

	commands = append(commands, newCommand)
	if err := saveCommands(commands); err != nil {
		return fmt.Errorf("error saving commands: %v", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), "Command added successfully")
	return nil
}

func checkUniqCommand(commands []ShortcutCommand, newCommand ShortcutCommand) error {
	for _, cmd := range commands {
		if cmd.Command == newCommand.Command {
			return fmt.Errorf("command already exists")
		} else if cmd.Shortcut == newCommand.Shortcut {
			return fmt.Errorf("shortcut already exists")
		}
	}
	return nil
}
