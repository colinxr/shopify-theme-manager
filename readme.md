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

## Usage

To use the tool, run the following command:

```bash
  stm
```

### List Themes (`stm list`)

List all themes for a specific store.

```bash
  stm list <alias>
```

### Dev Theme (`stm dev`)

Start theme development server.

```bash
  stm dev <themeId>
```

## Configuration

The tool stores configurations in:

```bash
~/.config/shopify-theme-manager/config.json
```

## Development

To contribute to this project:

1. Clone the repository
2. Install dependencies:
   ```bash
   npm install
   ```
3. Run tests:
   ```bash
   npm test
   ```
4. Start in development mode:
   ```bash
   npm start
   ```

## License

ISC

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
