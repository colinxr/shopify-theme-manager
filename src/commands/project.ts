import { Command } from 'commander';
import { ConfigManager } from '../utils/config';
import inquirer from 'inquirer';
import { ensureShopifyCLI } from '../utils/cli-check';

export function setupProjectCommands(program: Command): void {
  const config = new ConfigManager();

  program
    .command('add')
    .description('Add a new Shopify store configuration')
    .action(async () => {
      const answers = await inquirer.prompt([
        {
          type: 'input',
          name: 'storeId',
          message: 'Enter the Shopify store ID:',
          validate: (input: string) => {
            if (!input.trim()) {
              return 'Store ID is required';
            }
            return true;
          }
        },
        {
          type: 'input',
          name: 'alias',
          message: 'Enter an alias for the store (optional):',
          default: (answers: { storeId: string }) => answers.storeId
        },
        {
          type: 'input',
          name: 'projectDirectory',
          message: 'Enter the project directory path:',
          default: process.cwd(),
          validate: (input: string) => {
            if (!input.trim()) {
              return 'Project directory is required';
            }
            return true;
          }
        }
      ]);

      config.addStore(answers.storeId, answers.alias, answers.projectDirectory);
      console.log(`Store ${answers.alias} added successfully`);
    });

  program
    .command('list')
    .description('List themes for a store')
    .argument('<alias>', 'Store alias')
    .option('-n, --name <name>', 'Filter by theme name')
    .action((alias, options) => {
      ensureShopifyCLI();
      const store = config.getStore(alias);
      console.log(store);
      
      if (!store) {
        console.error(`Store with alias "${alias}" not found`);
        return;
      }

      const command = ['shopify', 'theme', 'list', '--store', store.storeId];
      if (options.name) {
        command.push('--name', options.name);
      }
      
      // Execute command using child_process
      const { execSync } = require('child_process');
      try {
        const output = execSync(command.join(' '), { encoding: 'utf-8' });
        console.log(output);
      } catch (error) {
        console.error('Error executing Shopify CLI command:', error);
      }
    });

  program
    .command('dev')
    .description('Start theme development server')
    .argument('<themeId>', 'Theme ID to develop against')
    .action(async (themeId) => {
      ensureShopifyCLI();

      const command = ['shopify', 'theme', 'dev', '--theme', themeId];
      
      // Execute command using child_process
      const { spawn } = require('child_process');
      try {
        const process = spawn(command[0], command.slice(1), { 
          stdio: 'inherit',
          shell: true
        });

        process.on('error', (error: Error) => {
          console.error('Error executing Shopify CLI command:', error);
        });
      } catch (error) {
        console.error('Error executing Shopify CLI command:', error);
      }
    });
} 