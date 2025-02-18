package commands

import (
	"fmt"
	"strings"

	"github.com/colinxr/shopify-theme-manager/config"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// Declare runPrompt at package level for production use
var runPrompt = func(p promptui.Prompt) (string, error) {
	return p.Run()
}

func NewAddCommand(cfg config.Manager) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a new Shopify store configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Store ID prompt
			storePrompt := promptui.Prompt{
				Label:    "Enter the Shopify store ID",
				Validate: notEmptyValidator,
			}
			storeID, err := runPrompt(storePrompt)
			if err != nil {
				return err
			}

			// Alias prompt
			aliasPrompt := promptui.Prompt{
				Label:   "Enter an alias for the store (optional)",
				Default: storeID,
			}
			alias, err := runPrompt(aliasPrompt)
			if err != nil {
				return err
			}
			if alias == "" {
				alias = storeID
			}

			// Project directory prompt
			dirPrompt := promptui.Prompt{
				Label:    "Enter the project directory path",
				Validate: notEmptyValidator,
			}
			projectDir, err := runPrompt(dirPrompt)
			if err != nil {
				return err
			}

			if err := cfg.AddStore(storeID, alias, projectDir); err != nil {
				return err
			}

			fmt.Printf("Store %s added successfully\n", alias)
			return nil
		},
	}
}

func notEmptyValidator(input string) error {
	if input == "" || len(strings.TrimSpace(input)) == 0 {
		return fmt.Errorf("value cannot be empty")
	}
	return nil
}
