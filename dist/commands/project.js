"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.setupProjectCommands = setupProjectCommands;
const config_1 = require("../utils/config");
const inquirer_1 = __importDefault(require("inquirer"));
function setupProjectCommands(program) {
    const config = new config_1.ConfigManager();
    program
        .command('add')
        .description('Add a new Shopify store configuration')
        .action(async () => {
        const answers = await inquirer_1.default.prompt([
            {
                type: 'input',
                name: 'storeId',
                message: 'Enter the Shopify store ID:',
                validate: (input) => {
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
                default: (answers) => answers.storeId // Use store ID as default alias
            }
        ]);
        config.addStore(answers.storeId, answers.alias);
        console.log(`Store ${answers.alias} added successfully`);
    });
    program
        .command('list')
        .description('List themes for a store')
        .argument('<alias>', 'Store alias')
        .option('-n, --name <name>', 'Filter by theme name')
        .action((alias, options) => {
        const store = config.getStore(alias);
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
        }
        catch (error) {
            console.error('Error executing Shopify CLI command:', error);
        }
    });
}
