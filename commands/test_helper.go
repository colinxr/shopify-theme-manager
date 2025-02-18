package commands

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/colinxr/shopify-theme-manager/config"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// testHelper provides common test utilities
type testHelper struct {
	cmd    *cobra.Command
	output *bytes.Buffer
	mock   config.Manager
}

func newTestHelper(t *testing.T) *testHelper {
	output := &bytes.Buffer{}
	cmd := &cobra.Command{Use: "test"}
	cmd.SetOut(output)
	cmd.SetErr(output)
	return &testHelper{
		cmd:    cmd,
		output: output,
		mock:   NewMockConfig(),
	}
}

// setupCommand is a helper to set up a command for testing
func (h *testHelper) setupCommand(cmd *cobra.Command) {
	h.cmd.AddCommand(cmd)
}

// Reset mocks after tests
func resetMocks() {
	execCommand = exec.Command
	runPrompt = func(p promptui.Prompt) (string, error) {
		return "", nil
	}
}

// MockPrompt replaces the runPrompt function with a mock
func MockPrompt(mockRun func(promptui.Prompt) (string, error)) func() {
	oldRun := runPrompt
	runPrompt = mockRun
	return func() {
		runPrompt = oldRun
	}
}

// MockExecCommand replaces exec.Command with a mock
func MockExecCommand(mockExec func(string, ...string) *exec.Cmd) func() {
	oldExec := execCommand
	execCommand = mockExec
	return func() {
		execCommand = oldExec
	}
}
