package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetCommandsFilePath(t *testing.T) {
	expectedPath := filepath.Join(os.Getenv("HOME"), ".commands.json")
	actualPath := getCommandsFilePath()

	if actualPath != expectedPath {
		t.Errorf("Expected path: %s, but got: %s", expectedPath, actualPath)
	}
}

func TestLoadCommands(t *testing.T) {

	tmpDir := t.TempDir()

	oldHome := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHome)
	os.Setenv("HOME", tmpDir)

	testFilePath := filepath.Join(tmpDir, ".commands.json")

	testData := []ShortcutCommand{
		{Command: "command1", Shortcut: "cmd1", Description: "Test command 1"},
		{Command: "command2", Shortcut: "cmd2", Description: "Test command 2"},
	}

	jsonData, err := json.Marshal(testData)
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	if err := os.WriteFile(testFilePath, jsonData, 0644); err != nil {
		t.Fatalf("Failed to write test data to test file: %v", err)
	}

	commands, err := loadCommands()
	if err != nil {
		t.Fatalf("Failed to load commands: %v", err)
	}

	expectedCommands := []ShortcutCommand{
		{Command: "command1", Shortcut: "cmd1", Description: "Test command 1"},
		{Command: "command2", Shortcut: "cmd2", Description: "Test command 2"},
	}

	if !reflect.DeepEqual(commands, expectedCommands) {
		t.Errorf("Loaded commands do not match expected commands.\nExpected: %v\nGot: %v", expectedCommands, commands)
	}
}
