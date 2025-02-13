package commands

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
	"github.com/colinxr/shopify-theme-manager/config"
)

func NewDevCommand(cfg *config.Manager) *cobra.Command {
	return &cobra.Command{
		Use:   "dev <theme-id>",
		Short: "Start theme development server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			themeID := args[0]

			// Build shopify CLI command
			shopifyCmd := exec.Command("shopify", "theme", "dev", "--theme", themeID)

			// Set output to current process
			shopifyCmd.Stdout = cmd.OutOrStdout()
			shopifyCmd.Stderr = cmd.ErrOrStderr()

			return shopifyCmd.Run()
		},
	}
} 