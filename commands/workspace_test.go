package commands

import (
	"strings"
	"testing"
)

func TestSetWorkspaceCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantPath string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "set workspace with explicit path",
			args:     []string{"set-workspace", "/test/workspace"},
			wantPath: "/test/workspace",
			wantErr:  false,
		},
		{
			name:     "set workspace with current directory",
			args:     []string{"set-workspace"},
			wantPath: ".",
			wantErr:  false,
		},
		{
			name:    "too many arguments",
			args:    []string{"set-workspace", "/test/workspace", "extra"},
			wantErr: true,
			errMsg:  "accepts at most 1 arg(s), received 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newTestHelper(t)

			cmd := NewSetWorkspaceCommand(h.mock)
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

			// Verify the workspace was set correctly
			if got := h.mock.GetWorkspace(); got != tt.wantPath {
				t.Errorf("workspace = %v, want %v", got, tt.wantPath)
			}
		})
	}
}

func TestSetWorkspaceCommand_InvalidPath(t *testing.T) {
	h := newTestHelper(t)

	// Use a path with null bytes which is invalid on all operating systems
	invalidPath := "test\x00path"

	cmd := NewSetWorkspaceCommand(h.mock)
	h.setupCommand(cmd)

	h.cmd.SetArgs([]string{"set-workspace", invalidPath})
	err := h.cmd.Execute()

	if err == nil {
		t.Error("expected error for invalid path but got none")
	} else if !strings.Contains(err.Error(), "invalid workspace path") {
		t.Errorf("expected error about invalid workspace path, got: %v", err)
	}
}
