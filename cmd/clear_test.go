package cmd

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func TestClearCommandList(t *testing.T) {

	initialCommands := []ShortcutCommand{
		{Shortcut: "cmd1", Command: "echo 'Hello'"},
		{Shortcut: "cmd2", Command: "echo 'World'"},
	}

	commandsFilePath := getCommandsFilePath()
	err := saveCommands(initialCommands)
	if err != nil {
		t.Fatalf("Failed to save initial commands: %s", err)
	}
	defer os.Remove(commandsFilePath)

	cmd := &cobra.Command{}

	clearCommandList(cmd, []string{})

	data, err := os.ReadFile(commandsFilePath)
	if err != nil {
		t.Fatalf("Failed to read commands file: %s", err)
	}

	var loadedCommands []ShortcutCommand
	err = json.Unmarshal(data, &loadedCommands)
	if err != nil {
		t.Fatalf("Failed to unmarshal commands: %s", err)
	}

	if len(loadedCommands) != 0 {
		t.Errorf("Expected 0 commands, got %d", len(loadedCommands))
	}
}
