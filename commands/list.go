package commands

import (
	"fmt"
	"os"

	"github.com/colinxr/shopify-theme-manager/config"
	"github.com/spf13/cobra"
)

func NewListCommand(cfg config.Manager) *cobra.Command {
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

			// Build base shopify CLI command
			args = []string{"theme", "list", "--store", store.StoreID}

			// Add name filter if provided
			if themeName != "" {
				args = append(args, "--name", themeName)
			}

			// Create command with all arguments
			shopifyCmd := execCommand("shopify", args...)

			// Set up the command to use the current terminal
			shopifyCmd.Stdin = cmd.InOrStdin()
			shopifyCmd.Stdout = cmd.OutOrStdout()
			shopifyCmd.Stderr = cmd.ErrOrStderr()

			// Inherit the parent environment
			shopifyCmd.Env = append([]string{
				"TERM=" + os.Getenv("TERM"),
				"HOME=" + os.Getenv("HOME"),
				"PATH=" + os.Getenv("PATH"),
			}, os.Environ()...)

			return shopifyCmd.Run()
		},
	}

	cmd.Flags().StringVarP(&themeName, "name", "n", "", "Filter themes by name")
	return cmd
}
