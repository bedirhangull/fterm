package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCommandCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command based on a shortcut",
	Long:  `Run a command based on a shortcut provided in the .commands.json file.`,
	Run: func(cmd *cobra.Command, args []string) {
		findCommand(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(runCommandCmd)
	runCommandCmd.Flags().StringP("run", "r", "", "Which command to execute")
}

func findCommand(cmd *cobra.Command, args []string) {
	shortcut := getShortcut(cmd, args)
	if shortcut == "" {
		fmt.Println("No shortcut provided")
		return
	}

	commands, err := loadCommands()
	if err != nil {
		fmt.Println("Error loading commands:", err)
		return
	}

	for _, command := range commands {
		if command.Shortcut == shortcut {
			fmt.Println("Executing command:", command.Description)
			executeCommand(command.Command)
			return
		}
	}
	fmt.Println("Command not found")
}

func getShortcut(cmd *cobra.Command, args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return cmd.Flag("run").Value.String()
}

func executeCommand(command string) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
}
