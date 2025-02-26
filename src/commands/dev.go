package commands

import (
	"github.com/colinxr/shopify-theme-manager/config"
	"github.com/spf13/cobra"
)

func NewDevCommand(cfg config.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dev <theme-id>",
		Short: "Start theme development server",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var themeID string
			if len(args) > 0 {
				themeID = args[0]
			}

			// Build base command args
			cmdArgs := []string{"theme", "dev"}
			if themeID != "" {
				cmdArgs = append(cmdArgs, "--theme", themeID)
			}

			// Add port flag if provided
			if port, _ := cmd.Flags().GetString("port"); port != "" {
				cmdArgs = append(cmdArgs, "--port", port)
			}

			// Create command with all arguments
			shopifyCmd := execCommand("shopify", cmdArgs...)

			// Set output to current process
			shopifyCmd.Stdout = cmd.OutOrStdout()
			shopifyCmd.Stderr = cmd.ErrOrStderr()
			shopifyCmd.Stdin = cmd.InOrStdin()

			return shopifyCmd.Run()
		},
	}

	// Add flags
	cmd.Flags().String("port", "", "Port to use")
	cmd.Flags().Bool("live-reload", true, "Enable live reload")

	return cmd
}
