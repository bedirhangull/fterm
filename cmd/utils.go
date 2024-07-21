package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func getCommandsFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, ".commands.json")
}

func loadCommands() ([]ShortcutCommand, error) {
	filePath := getCommandsFilePath()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, []byte("[]"), 0644); err != nil {
			return nil, fmt.Errorf("error creating .commands.json file: %w", err)
		}
	}

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var commands []ShortcutCommand
	err = json.Unmarshal(byteValue, &commands)
	if err != nil {
		return nil, err
	}
	return commands, nil
}
