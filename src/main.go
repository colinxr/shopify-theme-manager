package main

import (
	"log"

	"github.com/colinxr/shopify-theme-manager/commands"
	"github.com/colinxr/shopify-theme-manager/config"
)

func main() {
	cfg, err := config.NewManager()
	if err != nil {
		log.Fatal(err)
	}

	rootCmd := commands.NewRootCommand(cfg)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
} 