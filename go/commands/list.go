package commands

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
	"github.com/colinxr/shopify-theme-manager/config"
)

func NewListCommand(cfg *config.Manager) *cobra.Command {
	var themeName string

	cmd := &cobra.Command{
		Use:   "list <store-alias>",
		Short: "List themes for a store",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			alias := args[0]
			store := cfg.GetStore(alias)
			if store == nil {
				return fmt.Errorf("store with alias %q not found", alias)
			}

			// Build shopify CLI command
			shopifyCmd := exec.Command("shopify", "theme", "list", "--store", store.StoreID)
			if themeName != "" {
				shopifyCmd.Args = append(shopifyCmd.Args, "--name", themeName)
			}

			// Set output to current process
			shopifyCmd.Stdout = cmd.OutOrStdout()
			shopifyCmd.Stderr = cmd.ErrOrStderr()

			return shopifyCmd.Run()
		},
	}

	cmd.Flags().StringVarP(&themeName, "name", "n", "", "Filter themes by name")
	return cmd
} 