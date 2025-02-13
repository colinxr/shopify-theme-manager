package commands

import (
	"github.com/spf13/cobra"
	"github.com/colinxr/shopify-theme-manager/config"
)

func NewRootCommand(cfg config.Manager) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "stm",
		Version: "1.0.0",
		Short:   "Shopify Theme Manager - A CLI tool to manage Shopify themes",
	}

	// Add commands
	rootCmd.AddCommand(
		NewAddCommand(cfg),
		NewListCommand(cfg),
		NewDevCommand(cfg),
		NewSetWorkspaceCommand(cfg),
		NewCdCommand(cfg),
	)

	return rootCmd
} 