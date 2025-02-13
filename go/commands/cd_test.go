package commands

import (
	"strings"
	"testing"
)

func TestCdCommand(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		setupMock func(*testHelper)
		wantPath  string
		wantErr   bool
		errMsg    string
	}{
		{
			name: "change to valid store directory",
			args: []string{"cd", "test-alias"},
			setupMock: func(h *testHelper) {
				h.mock.AddStore("test-store", "test-alias", "test-dir")
				h.mock.SetWorkspace("/test/workspace")
			},
			wantPath: "/test/workspace/test-dir",
			wantErr:  false,
		},
		{
			name: "store not found",
			args: []string{"cd", "invalid-alias"},
			setupMock: func(h *testHelper) {
				h.mock.SetWorkspace("/test/workspace") // Add this line
			},
			wantErr: true,
			errMsg:  "store with alias \"invalid-alias\" not found",
		},
		{
			name: "workspace not set",
			args: []string{"cd", "test-alias"},
			setupMock: func(h *testHelper) {
				h.mock.AddStore("test-store", "test-alias", "test-dir")
			},
			wantErr: true,
			errMsg:  "workspace not set",
		},
		{
			name:    "missing store alias argument",
			args:    []string{"cd"},
			wantErr: true,
			errMsg:  "accepts 1 arg(s), received 0",
		},
		{
			name:    "too many arguments",
			args:    []string{"cd", "test-alias", "extra-arg"},
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

			var changedDir string
			cleanup := MockChdir(func(dir string) error {
				// Just store the directory without actually changing to it
				changedDir = dir
				return nil
			})
			defer cleanup()

			cmd := NewCdCommand(h.mock)
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

			if tt.wantPath != "" && changedDir != tt.wantPath {
				t.Errorf("changed to directory = %v, want %v", changedDir, tt.wantPath)
			}
		})
	}
}
