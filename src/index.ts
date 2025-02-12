#!/usr/bin/env node

import { Command } from 'commander';
import { setupProjectCommands } from './commands/project';

const program = new Command();

program
  .name('shopify-theme-manager')
  .description('CLI tool to manage Shopify themes')
  .version('1.0.0');

setupProjectCommands(program);

program.parse(process.argv); 