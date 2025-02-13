package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Store struct {
	StoreID      string `json:"storeId"`
	Alias        string `json:"alias"`
	ProjectDir   string `json:"projectDir"`
}

type Config struct {
	Stores    []Store `json:"stores"`
	Workspace string  `json:"workspace"`
}

type Manager struct {
	configDir  string
	configPath string
	config     *Config
}

func NewManager() (*Manager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(homeDir, ".config", "shopify-theme-manager")
	configPath := filepath.Join(configDir, "config.json")

	m := &Manager{
		configDir:  configDir,
		configPath: configPath,
	}

	if err := m.ensureConfigExists(); err != nil {
		return nil, err
	}

	if err := m.loadConfig(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Manager) ensureConfigExists() error {
	if err := os.MkdirAll(m.configDir, 0755); err != nil {
		return err
	}

	if _, err := os.Stat(m.configPath); os.IsNotExist(err) {
		config := Config{
			Stores: []Store{},
		}
		return m.saveConfig(&config)
	}

	return nil
}

func (m *Manager) loadConfig() error {
	data, err := os.ReadFile(m.configPath)
	if err != nil {
		return err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	m.config = &config
	return nil
}

func (m *Manager) saveConfig() error {
	data, err := json.MarshalIndent(m.config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(m.configPath, data, 0644)
}

func (m *Manager) AddStore(storeID, alias, projectDir string) error {
	store := Store{
		StoreID:    storeID,
		Alias:      alias,
		ProjectDir: projectDir,
	}
	m.config.Stores = append(m.config.Stores, store)
	return m.saveConfig()
}

func (m *Manager) GetStore(alias string) *Store {
	for _, store := range m.config.Stores {
		if store.Alias == alias {
			return &store
		}
	}
	return nil
}

func (m *Manager) SetWorkspace(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	m.config.Workspace = absPath
	return m.saveConfig()
}

func (m *Manager) GetWorkspace() string {
	return m.config.Workspace
} 