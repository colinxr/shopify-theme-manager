# Shopify Theme Manager (stm)

A CLI tool to simplify working with Shopify themes. This tool helps manage multiple Shopify stores and their themes through an easy-to-use command line interface.

## Prerequisites

Before using this tool, ensure you have:

- Node.js 14 or higher
- Shopify CLI installed globally:

  ```bash
  npm install -g @shopify/cli @shopify/theme

  ```

## Installation

Install the package globally:

```bash
  npm install -g shopify-theme-manager
```

## Commands

### Add Store (`stm add`)

Add a new Shopify store configuration. The command will prompt for:

- Store ID (required)
- Store alias (optional,defaults to store ID)
- Project directory path (required)

```bash
stm add
```

### Set Workspace (`stm set-workspace`)

Set the workspace directory for all projects. This is the root directory where all store projects are located.

```bash
# Set to specific directory
stm set-workspace /path/to/workspace

# Set to current directory
stm set-workspace
```

### List Themes (`stm list`)

List all themes for a specific store.

```bash
stm list <store-alias> [--name <theme-name>]
```

### Development Server (`stm dev`)

Start theme development server for a specific theme.

```bash
stm dev <theme-id>
```

### Change Directory (`stm cd`)

Change to a store's project directory within the workspace.

```bash
stm-cd store-alias
```

## Configuration

The tool stores configurations in:

```bash
~/.config/shopify-theme-manager/config.json
```

Configuration includes:

- Workspace directory
- Store configurations (ID, alias, project directory)

```

```
