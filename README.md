# Shopify Theme Manager (stm)

A CLI tool to simplify working with Shopify themes. This tool helps manage multiple Shopify stores and their themes through an easy-to-use command line interface.

## Prerequisites

Before using this tool, ensure you have:

- Shopify CLI installed:
  ```bash
  brew install shopify-cli
  ```

## Installation

Install via Homebrew:

```bash
brew tap colinxr/shopify-theme-manager
brew install shopify-theme-manager
```

## Updating

To update to the latest version:

```bash
brew upgrade shopify-theme-manager
```

You can verify the installed version with:

```bash
stm -V
```

## Commands

### Set Workspace (`stm set-workspace`)

Set the workspace directory for all projects. This is the root directory where all store projects are located.

```bash
# Set to specific directory
stm set-workspace /path/to/workspace

# Set to current directory
stm set-workspace
```

### Add Store (`stm add`)

Add a new Shopify store configuration. The command will prompt for:

- Store ID (required) - Your Shopify store ID (e.g., my-store.myshopify.com)
- Store alias (optional, defaults to store ID) - A shorthand name for the store
- Project directory path (required) - The directory containing your theme files

```bash
stm add
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
stm cd store-alias
```

## Configuration

The tool stores configurations in:

```bash
~/.config/shopify-theme-manager/config.json
```

Configuration includes:

- Workspace directory - Root directory for all projects
- Store configurations:
  - Store ID - Shopify store identifier
  - Alias - Custom name for the store
  - Project directory - Path to theme files (relative to workspace)

## Example Workflow

1. Set up workspace:

   ```bash
   stm set-workspace ~/shopify-projects
   ```

2. Add a store:

   ```bash
   stm add
   # > Enter store ID: my-store.myshopify.com
   # > Enter alias: store1
   # > Enter project directory: store1-theme
   ```

3. Navigate to store directory:

   ```bash
   stm cd store1
   ```

4. List themes:

   ```bash
   stm list store1
   ```

5. Start development:
   ```bash
   stm dev <theme-id>
   ```
