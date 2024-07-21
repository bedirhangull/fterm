package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

func TestRunCommand(t *testing.T) {
	tmpDir := t.TempDir()

	oldHome := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHome)
	os.Setenv("HOME", tmpDir)

	testFilePath := filepath.Join(tmpDir, ".commands.json")

	testCommands := []ShortcutCommand{
		{Shortcut: "cmd1", Command: "echo Hello, World!", Description: "Test command 1"},
		{Shortcut: "cmd2", Command: "echo Goodbye, World!", Description: "Test command 2"},
	}

	jsonData, err := json.Marshal(testCommands)
	if err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(testFilePath, jsonData, 0644); err != nil {
		t.Fatal(err)
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cmd := &cobra.Command{}
	cmd.Flags().StringP("run", "r", "cmd1", "Which command to execute")
	findCommand(cmd, []string{"cmd1"})

	w.Close()
	os.Stdout = oldStdout

	output := make([]byte, 1024)
	n, _ := r.Read(output)
	output = output[:n]

	expectedOutput := "Executing command: Test command 1\nHello, World!\n"

	if string(output) != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
