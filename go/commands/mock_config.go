package commands

import (
	"fmt"
	"strings"

	"github.com/colinxr/shopify-theme-manager/config"
)

// MockConfig implements config.Manager for testing
type MockConfig struct {
	stores    []config.Store
	workspace string
}

func NewMockConfig() config.Manager {
	return &MockConfig{
		stores: make([]config.Store, 0),
	}
}

func (m *MockConfig) AddStore(storeID, alias, projectDir string) error {
	m.stores = append(m.stores, config.Store{
		StoreID:    storeID,
		Alias:      alias,
		ProjectDir: projectDir,
	})
	return nil
}

func (m *MockConfig) GetStore(alias string) *config.Store {
	for _, store := range m.stores {
		if store.Alias == alias {
			return &store
		}
	}
	return nil
}

func (m *MockConfig) SetWorkspace(path string) error {
	// Check for null bytes in path
	if strings.Contains(path, "\x00") {
		return fmt.Errorf("invalid workspace path: contains null byte")
	}
	m.workspace = path
	return nil
}

func (m *MockConfig) GetWorkspace() string {
	return m.workspace
}
