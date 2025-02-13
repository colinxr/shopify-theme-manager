package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/manifoldco/promptui"
	"github.com/colinxr/shopify-theme-manager/config"
)

func NewAddCommand(cfg *config.Manager) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a new Shopify store configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Store ID prompt
			storePrompt := promptui.Prompt{
				Label:    "Enter the Shopify store ID",
				Validate: notEmptyValidator,
			}
			storeID, err := storePrompt.Run()
			if err != nil {
				return err
			}

			// Alias prompt
			aliasPrompt := promptui.Prompt{
				Label:    "Enter an alias for the store (optional)",
				Default:  storeID,
			}
			alias, err := aliasPrompt.Run()
			if err != nil {
				return err
			}

			// Project directory prompt
			dirPrompt := promptui.Prompt{
				Label:    "Enter the project directory path",
				Validate: notEmptyValidator,
			}
			projectDir, err := dirPrompt.Run()
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
	if input == "" {
		return fmt.Errorf("value cannot be empty")
	}
	return nil
} 