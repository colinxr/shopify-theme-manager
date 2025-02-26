package commands

import (
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

func TestDevCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantCmd  string
		wantArgs []string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "optional theme ID",
			args:     []string{"dev"},
			wantCmd:  "shopify",
			wantArgs: []string{"theme", "dev"},
			wantErr:  false,
		},
		{
			name:     "valid theme ID",
			args:     []string{"dev", "123456"},
			wantCmd:  "shopify",
			wantArgs: []string{"theme", "dev", "--theme", "123456"},
			wantErr:  false,
		},
		{
			name:    "too many arguments",
			args:    []string{"dev", "123456", "extra"},
			wantErr: true,
			errMsg:  "accepts at most 1 arg(s), received 2",
		},
		{
			name:     "theme ID with flags",
			args:     []string{"dev", "123456", "--port", "9292"},
			wantCmd:  "shopify",
			wantArgs: []string{"theme", "dev", "--theme", "123456", "--port", "9292"},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newTestHelper(t)

			var executedCmd string
			var executedArgs []string

			cleanup := MockExecCommand(func(cmd string, args ...string) *exec.Cmd {
				executedCmd = cmd
				executedArgs = args
				return exec.Command("true")
			})
			defer cleanup()

			cmd := NewDevCommand(h.mock)
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

func TestDevCommand_ExecutionFailure(t *testing.T) {
	h := newTestHelper(t)

	cleanup := MockExecCommand(func(cmd string, args ...string) *exec.Cmd {
		// Return a command that will fail
		return exec.Command("false")
	})
	defer cleanup()

	cmd := NewDevCommand(h.mock)
	h.setupCommand(cmd)

	h.cmd.SetArgs([]string{"dev", "123456"})
	err := h.cmd.Execute()

	if err == nil {
		t.Error("expected error from failed command execution but got none")
	}
}
