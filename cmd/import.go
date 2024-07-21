package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a command from a file",
	Long:  `Import a command from a file and add it to the .commands.json file.`,
	Run:   importCommandList,
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().StringP("file", "f", "", "File to import")
}

func importCommandList(cmd *cobra.Command, args []string) {
	file := cmd.Flag("file").Value.String()
	if file == "" {
		fmt.Println("No file provided")
		return
	}

	currentCommands, err := loadCommands()
	if err != nil {
		fmt.Println("Error loading commands:", err)
		return
	}

	newCommands, err := loadCommandsFromFile(file)
	if err != nil {
		fmt.Println("Error loading commands from file:", err)
		return
	}

	mergedCommands := mergeCommands(currentCommands, newCommands)

	err = saveCommands(mergedCommands)
	if err != nil {
		fmt.Println("Error saving commands:", err)
		return
	}

	fmt.Println("Commands imported successfully.")
}

func loadCommandsFromFile(filePath string) ([]ShortcutCommand, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var commands []ShortcutCommand
	err = json.Unmarshal(byteValue, &commands)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return commands, nil
}

func mergeCommands(currentCommands, newCommands []ShortcutCommand) []ShortcutCommand {
	commandMap := make(map[string]ShortcutCommand)
	for _, command := range currentCommands {
		commandMap[command.Shortcut] = command
	}

	for _, command := range newCommands {
		commandMap[command.Shortcut] = command
	}

	var mergedCommands []ShortcutCommand
	for _, command := range commandMap {
		mergedCommands = append(mergedCommands, command)
	}

	return mergedCommands
}

func saveCommands(commands []ShortcutCommand) error {
	filePath := getCommandsFilePath()

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating .commands.json file: %w", err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(commands, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling commands: %w", err)
	}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing to .commands.json file: %w", err)
	}

	return nil
}
