package cmd_test

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/spf13/cobra"
	"github.com/vldcreation/helpme/cmd"
)

var (
	once sync.Once
)

// Helper function to create a new app instance and execute a command
func executeCommand(app *cobra.Command, args ...string) (string, string, error) {
	cmd := app
	cmd.SetArgs(args)
	var stdout, stderr bytes.Buffer
	cmd.SetOut(&stdout)
	cmd.SetErr(&stderr)
	err := cmd.Execute()
	return stdout.String(), stderr.String(), err
}

func Test_ExecuteParticularCommandDynamically(t *testing.T) {
	app := cmd.NewApp()

	// Iterate through the root commands
	for _, child := range app.Root().Commands() {
		testCommandInitialization(t, app.Root(), child.Name())

		if child.HasSubCommands() {
			for _, grandChild := range child.Commands() {
				testCommandInitialization(t, app.Root(), child.Name(), grandChild.Name())
			}
		}
	}
}

func testCommandInitialization(t *testing.T, rootApp *cobra.Command, commandNames ...string) {
	commandPath := strings.Join(commandNames, " ")
	fmt.Printf("Testing command initialization: %s\n", commandPath)

	// Attempt to execute the command (without specific arguments for now)
	_, _, err := executeCommand(rootApp, commandNames...)

	// Check for the flag redeclaration error message
	if err != nil && strings.Contains(err.Error(), "unable to redefine") {
		t.Errorf("Error during command initialization '%s': %s", commandPath, err.Error())
	} else if err != nil {
		// If there's an error but not the flag redeclaration, you might want to investigate further
		fmt.Printf("Command '%s' executed with error (not flag redeclaration): %s\n", commandPath, err.Error())
	} else {
		fmt.Printf("Command '%s' initialized successfully.\n", commandPath)
	}
}
