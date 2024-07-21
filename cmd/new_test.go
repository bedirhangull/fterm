package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestAddNewCommand(t *testing.T) {
	commandsFilePath := getCommandsFilePath()
	originalContent, _ := os.ReadFile(commandsFilePath)
	defer os.WriteFile(commandsFilePath, originalContent, 0644)

	tests := []struct {
		name        string
		shortcut    string
		command     string
		description string
		expectError bool
	}{
		{"ValidCommand", "test", "echo 'test'", "Test command", false},
		{"DuplicateCommand", "test", "echo 'test'", "Test command", true},
		{"DuplicateShortcut", "test", "echo 'new test'", "New test command", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd := &cobra.Command{Use: "root"}
			rootCmd.AddCommand(addNewCommandCmd)

			args := []string{
				"new",
				"-c", tt.command,
				"-s", tt.shortcut,
				"-d", tt.description,
			}
			rootCmd.SetArgs(args)

			var out bytes.Buffer
			rootCmd.SetOut(&out)
			rootCmd.SetErr(&out)

			err := rootCmd.Execute()
			output := out.String()
			t.Logf("Output: %s", output)

			if tt.expectError {
				assert.Error(t, err, output)
			} else {
				assert.NoError(t, err, output)
				assert.Contains(t, output, "Command added successfully", output)

				commands, err := loadCommands()
				assert.NoError(t, err)
				found := false
				for _, command := range commands {
					if command.Shortcut == tt.shortcut && command.Command == tt.command && command.Description == tt.description {
						found = true
						break
					}
				}
				assert.True(t, found)
			}
		})
	}
}

func TestCheckUniqCommand(t *testing.T) {
	existingCommands := []ShortcutCommand{
		{Shortcut: "test", Command: "echo 'test'", Description: "Test command"},
	}

	tests := []struct {
		name        string
		newCommand  ShortcutCommand
		expectError bool
	}{
		{"UniqueCommand", ShortcutCommand{Shortcut: "new", Command: "echo 'new'", Description: "New command"}, false},
		{"DuplicateCommand", ShortcutCommand{Shortcut: "test", Command: "echo 'test'", Description: "Duplicate command"}, true},
		{"DuplicateShortcut", ShortcutCommand{Shortcut: "test", Command: "echo 'new test'", Description: "New test command"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkUniqCommand(existingCommands, tt.newCommand)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
