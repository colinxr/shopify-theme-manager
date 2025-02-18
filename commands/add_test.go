package commands

import (
	"fmt"
	"strings"
	"testing"

	"github.com/manifoldco/promptui"
)

func TestAddCommand(t *testing.T) {
	tests := []struct {
		name            string
		promptResponses map[string]string
		wantErr         bool
		errMsg          string
		verify          func(t *testing.T, h *testHelper)
	}{
		{
			name: "successful add",
			promptResponses: map[string]string{
				"Enter the Shopify store ID":              "test-store",
				"Enter an alias for the store (optional)": "test-alias",
				"Enter the project directory path":        "test-dir",
			},
			wantErr: false,
			verify: func(t *testing.T, h *testHelper) {
				store := h.mock.GetStore("test-alias")
				if store == nil {
					t.Error("store was not added")
					return
				}
				if store.StoreID != "test-store" {
					t.Errorf("store ID = %s, want %s", store.StoreID, "test-store")
				}
				if store.ProjectDir != "test-dir" {
					t.Errorf("project dir = %s, want %s", store.ProjectDir, "test-dir")
				}
			},
		},
		{
			name: "empty store ID",
			promptResponses: map[string]string{
				"Enter the Shopify store ID": "",
			},
			wantErr: true,
			errMsg:  "value cannot be empty",
		},
		{
			name: "empty project directory",
			promptResponses: map[string]string{
				"Enter the Shopify store ID":              "test-store",
				"Enter an alias for the store (optional)": "test-alias",
				"Enter the project directory path":        "",
			},
			wantErr: true,
			errMsg:  "value cannot be empty",
		},
		{
			name: "default alias",
			promptResponses: map[string]string{
				"Enter the Shopify store ID":              "test-store",
				"Enter an alias for the store (optional)": "", // Should use store ID as default
				"Enter the project directory path":        "test-dir",
			},
			wantErr: false,
			verify: func(t *testing.T, h *testHelper) {
				store := h.mock.GetStore("test-store") // Use store ID as alias
				if store == nil {
					t.Error("store was not added")
					return
				}
				if store.StoreID != "test-store" {
					t.Errorf("store ID = %s, want %s", store.StoreID, "test-store")
				}
				if store.ProjectDir != "test-dir" {
					t.Errorf("project dir = %s, want %s", store.ProjectDir, "test-dir")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newTestHelper(t)

			cleanup := MockPrompt(func(p promptui.Prompt) (string, error) {
				if response, ok := tt.promptResponses[p.Label.(string)]; ok {
					if p.Validate != nil {
						if err := p.Validate(response); err != nil {
							return "", err
						}
					}
					return response, nil
				}
				return "", fmt.Errorf("unexpected prompt: %s", p.Label)
			})
			defer cleanup()

			cmd := NewAddCommand(h.mock)
			h.setupCommand(cmd)

			h.cmd.SetArgs([]string{"add"})
			err := h.cmd.Execute()

			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got none")
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("error = %v, want error containing %v", err, tt.errMsg)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if tt.verify != nil {
				tt.verify(t, h)
			}
		})
	}
}

func TestNotEmptyValidator(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid input",
			input:   "test",
			wantErr: false,
		},
		{
			name:    "empty input",
			input:   "",
			wantErr: true,
		},
		{
			name:    "whitespace input",
			input:   "   ",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := notEmptyValidator(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("notEmptyValidator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
