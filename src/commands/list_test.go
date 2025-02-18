package commands

import (
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

func TestListCommand(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		setupMock func(*testHelper)
		wantCmd   string
		wantArgs  []string
		wantErr   bool
		errMsg    string
	}{
		{
			name: "list themes for valid store",
			args: []string{"list", "test-alias"},
			setupMock: func(h *testHelper) {
				h.mock.AddStore("test-store", "test-alias", "test-dir")
			},
			wantCmd:  "shopify",
			wantArgs: []string{"theme", "list", "--store", "test-store"},
			wantErr:  false,
		},
		{
			name: "list themes with name filter",
			args: []string{"list", "test-alias", "--name", "dawn"},
			setupMock: func(h *testHelper) {
				h.mock.AddStore("test-store", "test-alias", "test-dir")
			},
			wantCmd:  "shopify",
			wantArgs: []string{"theme", "list", "--store", "test-store", "--name", "dawn"},
			wantErr:  false,
		},
		{
			name:    "store not found",
			args:    []string{"list", "invalid-store"},
			wantErr: true,
			errMsg:  "store with alias \"invalid-store\" not found",
		},
		{
			name:    "missing store alias",
			args:    []string{"list"},
			wantErr: true,
			errMsg:  "accepts 1 arg(s), received 0",
		},
		{
			name:    "too many arguments",
			args:    []string{"list", "test-alias", "extra"},
			wantErr: true,
			errMsg:  "accepts 1 arg(s), received 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newTestHelper(t)

			if tt.setupMock != nil {
				tt.setupMock(h)
			}

			var executedCmd string
			var executedArgs []string

			cleanup := MockExecCommand(func(cmd string, args ...string) *exec.Cmd {
				executedCmd = cmd
				executedArgs = args
				return exec.Command("true") // Use 'true' command which always succeeds
			})
			defer cleanup()

			cmd := NewListCommand(h.mock)
			h.setupCommand(cmd)

			h.cmd.SetArgs(tt.args)
			err := h.cmd.Execute()

			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got none")
				} else if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("error = %v, want error containing %v", err, tt.errMsg)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if tt.wantCmd != "" && executedCmd != tt.wantCmd {
				t.Errorf("executed command = %s, want %s", executedCmd, tt.wantCmd)
			}

			if tt.wantArgs != nil && !reflect.DeepEqual(executedArgs, tt.wantArgs) {
				t.Errorf("executed args = %v, want %v", executedArgs, tt.wantArgs)
			}
		})
	}
}

func TestListCommand_ExecutionFailure(t *testing.T) {
	h := newTestHelper(t)

	// Setup mock store
	h.mock.AddStore("test-store", "test-alias", "test-dir")

	cleanup := MockExecCommand(func(cmd string, args ...string) *exec.Cmd {
		// Return a command that will fail
		return exec.Command("nonexistent-command")
	})
	defer cleanup()

	cmd := NewListCommand(h.mock)
	h.setupCommand(cmd)

	h.cmd.SetArgs([]string{"list", "test-alias"})
	err := h.cmd.Execute()

	if err == nil {
		t.Error("expected error from failed command execution but got none")
	}
}

func TestListCommand_OutputCapture(t *testing.T) {
	h := newTestHelper(t)

	// Setup mock store
	h.mock.AddStore("test-store", "test-alias", "test-dir")

	expectedOutput := "mocked theme list output"
	cleanup := MockExecCommand(func(cmd string, args ...string) *exec.Cmd {
		return exec.Command("echo", expectedOutput)
	})
	defer cleanup()

	cmd := NewListCommand(h.mock)
	h.setupCommand(cmd)

	h.cmd.SetArgs([]string{"list", "test-alias"})
	err := h.cmd.Execute()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	output := h.output.String()
	if !strings.Contains(output, expectedOutput) {
		t.Errorf("output = %q, want to contain %q", output, expectedOutput)
	}
}
