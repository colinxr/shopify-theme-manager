package commands

import (
	"github.com/colinxr/shopify-theme-manager/config"
	"github.com/spf13/cobra"
)

func NewRootCommand(cfg config.Manager) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "stm",
		Version: "0.0.9",
		Short:   "Shopify Theme Manager - A CLI tool to manage Shopify themes",
	}

	// Add commands
	rootCmd.AddCommand(
		NewAddCommand(cfg),
		NewListCommand(cfg),
		NewDevCommand(cfg),
		NewSetWorkspaceCommand(cfg),
	)

	return rootCmd
}
