#!/usr/bin/env node
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const commander_1 = require("commander");
const project_1 = require("./commands/project");
const program = new commander_1.Command();
program
    .name('shopify-theme-manager')
    .description('CLI tool to manage Shopify themes')
    .version('1.0.0');
(0, project_1.setupProjectCommands)(program);
program.parse(process.argv);
