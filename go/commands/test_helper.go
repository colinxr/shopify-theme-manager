package commands

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/colinxr/shopify-theme-manager/config"
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

// Mock system calls
var (
	osChdir = os.Chdir
	runPrompt = func(p promptui.Prompt) (string, error) {
		return "", nil
	}
)

// Reset mocks after tests
func resetMocks() {
	execCommand = exec.Command
	osChdir = os.Chdir
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

// MockChdir replaces os.Chdir with a mock
func MockChdir(mockChdir func(string) error) func() {
	oldChdir := osChdir
	osChdir = mockChdir
	return func() {
		osChdir = oldChdir
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