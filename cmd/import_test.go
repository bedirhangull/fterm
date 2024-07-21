package cmd

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func TestImportCommandList(t *testing.T) {

	tmpfile, err := os.CreateTemp("", "test-commands.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	testCommands := []ShortcutCommand{
		{Shortcut: "cmd1", Command: "command1"},
		{Shortcut: "cmd2", Command: "command2"},
	}
	testData, err := json.Marshal(testCommands)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Write(testData); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	cmd := &cobra.Command{}
	cmd.Flags().StringP("file", "f", "", "File to import")
	cmd.Flags().Set("file", tmpfile.Name())

	rootCmd = &cobra.Command{Use: "root"}
	importCmd = &cobra.Command{
		Use:   "import",
		Short: "Import a command from a file",
		Long:  `Import a command from a file and add it to the .commands.json file.`,
		Run:   importCommandList,
	}
	rootCmd.AddCommand(importCmd)

	importCommandList(cmd, []string{})

	commands, err := loadCommands()
	if err != nil {
		t.Fatal(err)
	}
	if len(commands) != len(testCommands) {
		t.Errorf("Expected %d commands, got %d", len(testCommands), len(commands))
	}
	for i, command := range commands {
		if command.Shortcut != testCommands[i].Shortcut || command.Command != testCommands[i].Command {
			t.Errorf("Expected command %d to be %v, got %v", i, testCommands[i], command)
		}
	}
}
