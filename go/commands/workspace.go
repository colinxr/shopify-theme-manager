package commands

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/colinxr/shopify-theme-manager/config"
)

func NewSetWorkspaceCommand(cfg *config.Manager) *cobra.Command {
	return &cobra.Command{
		Use:   "set-workspace [directory]",
		Short: "Set the workspace directory for all projects",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Use current directory if no argument provided
			workspacePath := "."
			if len(args) > 0 {
				workspacePath = args[0]
			}

			if err := cfg.SetWorkspace(workspacePath); err != nil {
				return fmt.Errorf("failed to set workspace: %w", err)
			}

			fmt.Printf("Workspace set to: %s\n", cfg.GetWorkspace())
			return nil
		},
	}
} 