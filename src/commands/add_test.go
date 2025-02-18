package commands

import (
	"errors"
	"strings"
	"testing"

	"github.com/manifoldco/promptui"
)

func TestAddCommand(t *testing.T) {
	tests := []struct {
		name            string
		promptResponses map[string]string
		mockErrors      map[string]error
		configError     error
		wantErr         bool
		errMsg          string
		verify          func(t *testing.T, h *testHelper)
	}{
		{
			name: "successful store addition",
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
			name: "error on store ID prompt",
			mockErrors: map[string]error{
				"Enter the Shopify store ID": errors.New("store ID prompt failed"),
			},
			wantErr: true,
			errMsg:  "store ID prompt failed",
		},
		{
			name: "error on alias prompt",
			promptResponses: map[string]string{
				"Enter the Shopify store ID": "test-store",
			},
			mockErrors: map[string]error{
				"Enter an alias for the store (optional)": errors.New("alias prompt failed"),
			},
			wantErr: true,
			errMsg:  "alias prompt failed",
		},
		{
			name: "error on directory prompt",
			promptResponses: map[string]string{
				"Enter the Shopify store ID":              "test-store",
				"Enter an alias for the store (optional)": "test-alias",
			},
			mockErrors: map[string]error{
				"Enter the project directory path": errors.New("directory prompt failed"),
			},
			wantErr: true,
			errMsg:  "directory prompt failed",
		},
		{
			name: "error adding store to config",
			promptResponses: map[string]string{
				"Enter the Shopify store ID":              "test-store",
				"Enter an alias for the store (optional)": "test-alias",
				"Enter the project directory path":        "test-dir",
			},
			configError: errors.New("failed to add store to config"),
			wantErr:     true,
			errMsg:      "failed to add store to config",
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

			// Mock the prompt responses and errors
			cleanup := MockPrompt(func(p promptui.Prompt) (string, error) {
				label := p.Label.(string)
				if err, ok := tt.mockErrors[label]; ok {
					return "", err
				}
				if response, ok := tt.promptResponses[label]; ok {
					if p.Validate != nil {
						if err := p.Validate(response); err != nil {
							return "", err
						}
					}
					return response, nil
				}
				return "", errors.New("unexpected prompt")
			})
			defer cleanup()

			// Set up mock config error if specified
			if tt.configError != nil {
				h.mock = &mockConfigWithError{
					MockConfig: h.mock.(*MockConfig),
					addError:   tt.configError,
				}
			}

			cmd := NewAddCommand(h.mock)
			h.setupCommand(cmd)

			h.cmd.SetArgs([]string{"add"})
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

// mockConfigWithError wraps MockConfig to simulate config errors
type mockConfigWithError struct {
	*MockConfig
	addError error
}

func (m *mockConfigWithError) AddStore(storeID, alias, projectDir string) error {
	return m.addError
}
