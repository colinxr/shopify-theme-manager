"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
Object.defineProperty(exports, "__esModule", { value: true });
const commander_1 = require("commander");
const project_1 = require("../../commands/project");
const config_1 = require("../../utils/config");
const child_process_1 = require("child_process");
const inquirer = __importStar(require("inquirer"));
jest.mock('../../utils/config');
jest.mock('child_process');
jest.mock('inquirer', () => ({
    prompt: jest.fn()
}));
describe('Project Commands', () => {
    let program;
    let mockConfigManager;
    const mockPrompt = inquirer.prompt;
    beforeEach(() => {
        jest.clearAllMocks();
        // Create a mock instance before each test
        mockConfigManager = {
            addStore: jest.fn(),
            getStore: jest.fn(),
            listStores: jest.fn(),
        };
        config_1.ConfigManager.mockImplementation(() => mockConfigManager);
        program = new commander_1.Command();
        (0, project_1.setupProjectCommands)(program);
    });
    describe('add command', () => {
        it('should prompt for store details and add store configuration', async () => {
            // Setup
            const mockAnswers = {
                storeId: 'test-store',
                alias: 'test-alias'
            };
            mockPrompt.mockResolvedValue(mockAnswers);
            mockConfigManager.addStore.mockImplementation(() => { });
            // Execute
            await program.parseAsync(['node', 'test', 'add']);
            // Assert
            expect(mockPrompt).toHaveBeenCalledWith([
                expect.objectContaining({
                    type: 'input',
                    name: 'storeId',
                    message: 'Enter the Shopify store ID:'
                }),
                expect.objectContaining({
                    type: 'input',
                    name: 'alias',
                    message: 'Enter an alias for the store (optional):'
                })
            ]);
            expect(mockConfigManager.addStore).toHaveBeenCalledWith(mockAnswers.storeId, mockAnswers.alias);
        });
        it('should use store ID as alias when no alias is provided', async () => {
            // Setup
            const storeId = 'test-store';
            mockPrompt.mockResolvedValue({
                storeId,
                alias: storeId // Default value when user doesn't input an alias
            });
            // Execute
            await program.parseAsync(['node', 'test', 'add']);
            // Assert
            expect(mockConfigManager.addStore).toHaveBeenCalledWith(storeId, storeId);
        });
        it('should validate required store ID', async () => {
            // Setup
            const prompt = [
                {
                    type: 'input',
                    name: 'storeId',
                    message: 'Enter the Shopify store ID:',
                    validate: expect.any(Function)
                },
                expect.any(Object)
            ];
            mockPrompt.mockImplementation(async (questions) => {
                const storeIdQuestion = questions[0];
                // Test validation
                const emptyResult = storeIdQuestion.validate('');
                expect(emptyResult).toBe('Store ID is required');
                const validResult = storeIdQuestion.validate('valid-store');
                expect(validResult).toBe(true);
                // Return mock answers after validation
                return {
                    storeId: 'valid-store',
                    alias: 'test-alias'
                };
            });
            // Execute
            await program.parseAsync(['node', 'test', 'add']);
            // Assert
            expect(mockPrompt).toHaveBeenCalledWith(expect.arrayContaining(prompt));
        });
    });
    describe('list command', () => {
        it('should execute shopify CLI command with correct store ID', async () => {
            // Setup
            mockConfigManager.getStore.mockReturnValue({ storeId: 'test-store', alias: 'test-alias' });
            child_process_1.execSync.mockReturnValue(Buffer.from('themes list'));
            // Execute
            await program.parseAsync(['node', 'test', 'list', 'test-alias']);
            // Assert
            expect(child_process_1.execSync).toHaveBeenCalledWith('shopify theme list --store test-store', { encoding: 'utf-8' });
        });
        it('should include name filter when provided', async () => {
            // Setup
            mockConfigManager.getStore.mockReturnValue({ storeId: 'test-store', alias: 'test-alias' });
            child_process_1.execSync.mockReturnValue(Buffer.from('themes list'));
            // Execute
            await program.parseAsync(['node', 'test', 'list', 'test-alias', '--name', 'test-theme']);
            // Assert
            expect(child_process_1.execSync).toHaveBeenCalledWith('shopify theme list --store test-store --name test-theme', { encoding: 'utf-8' });
        });
        it('should handle store not found error', async () => {
            // Setup
            mockConfigManager.getStore.mockReturnValue(undefined);
            const consoleSpy = jest.spyOn(console, 'error').mockImplementation();
            // Execute
            await program.parseAsync(['node', 'test', 'list', 'non-existent']);
            // Assert
            expect(consoleSpy).toHaveBeenCalledWith('Store with alias "non-existent" not found');
            expect(child_process_1.execSync).not.toHaveBeenCalled();
        });
    });
});
