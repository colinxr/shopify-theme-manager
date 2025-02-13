package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/colinxr/shopify-theme-manager/config"
	"github.com/colinxr/shopify-theme-manager/commands"
)

func main() {
	cfg, err := config.NewManager()
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	var rootCmd = &cobra.Command{
		Use:     "stm",
		Version: "1.0.0",
		Short:   "Shopify Theme Manager - A CLI tool to manage Shopify themes",
	}

	// Add commands
	rootCmd.AddCommand(
		commands.NewAddCommand(cfg),
		commands.NewListCommand(cfg),
		commands.NewDevCommand(cfg),
		commands.NewSetWorkspaceCommand(cfg),
		commands.NewCdCommand(cfg),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
} 