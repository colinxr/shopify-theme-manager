package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/colinxr/shopify-theme-manager/config"
)

func NewCdCommand(cfg *config.Manager) *cobra.Command {
	return &cobra.Command{
		Use:   "cd <store-alias>",
		Short: "Change to store directory",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			alias := args[0]
			
			workspace := cfg.GetWorkspace()
			if workspace == "" {
				return fmt.Errorf("workspace not set. Please run 'stm set-workspace' first")
			}

			store := cfg.GetStore(alias)
			if store == nil {
				return fmt.Errorf("store with alias %q not found", alias)
			}

			targetDir := filepath.Join(workspace, store.ProjectDir)
			if err := os.Chdir(targetDir); err != nil {
				return fmt.Errorf("failed to change directory: %w", err)
			}

			fmt.Printf("Changed directory to: %s\n", targetDir)
			return nil
		},
	}
} 