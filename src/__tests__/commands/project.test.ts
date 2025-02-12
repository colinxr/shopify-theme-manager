import { Command } from 'commander';
import { setupProjectCommands } from '../../commands/project';
import { ConfigManager } from '../../utils/config';
import { execSync } from 'child_process';
import * as inquirer from 'inquirer';
import { spawn } from 'child_process';
import { checkShopifyCLI } from '../../utils/cli-check';

jest.mock('../../utils/config');
jest.mock('child_process', () => ({
  execSync: jest.fn(),
  spawn: jest.fn()
}));
jest.mock('inquirer', () => ({
  prompt: jest.fn()
}));
jest.mock('../../utils/cli-check', () => ({
  checkShopifyCLI: jest.fn().mockReturnValue(true),
  ensureShopifyCLI: jest.fn()
}));

describe('Project Commands', () => {
  let program: Command;
  let mockConfigManager: jest.Mocked<ConfigManager>;
  const mockPrompt = (inquirer.prompt as unknown) as jest.Mock;
  
  beforeEach(() => {
    jest.clearAllMocks();
    
    // Create a mock instance before each test
    mockConfigManager = {
      addStore: jest.fn(),
      getStore: jest.fn(),
      listStores: jest.fn(),
    } as any;
    
    (ConfigManager as jest.Mock).mockImplementation(() => mockConfigManager);
    
    program = new Command();
    setupProjectCommands(program);
  });

  describe('add command', () => {
    it('should prompt for store details and project directory', async () => {
      // Setup
      const mockAnswers = {
        storeId: 'test-store',
        alias: 'test-alias',
        projectDir: 'test-store-dir'
      };
      
      mockPrompt.mockResolvedValue(mockAnswers);
      mockConfigManager.addStore.mockImplementation(() => {});

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
        }),
        expect.objectContaining({
          type: 'input',
          name: 'projectDir',
          message: 'Enter the project directory path:',
        })
      ]);
      
      expect(mockConfigManager.addStore).toHaveBeenCalledWith(
        mockAnswers.storeId,
        mockAnswers.alias,
        mockAnswers.projectDir
      );
    });

    it('should use store ID as alias when no alias is provided', async () => {
      // Setup
      const storeId = 'test-store';
      const alias = 'test-alias';
      const projectDir = 'test-store-dir';
      mockPrompt.mockResolvedValue({
        storeId, alias, projectDir
      });

      // Execute
      await program.parseAsync(['node', 'test', 'add']);

      // Assert
      expect(mockConfigManager.addStore).toHaveBeenCalledWith(
        storeId, 
        alias,
        projectDir
      );
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
        expect.any(Object),
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
          alias: 'test-alias',
          projectDir: 'test-store-dir'
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
      mockConfigManager.getStore.mockReturnValue({ 
        storeId: 'test-store', 
        alias: 'test-alias',
        projectDir: 'test-store-dir'
      });
      (execSync as jest.Mock).mockReturnValue(Buffer.from('themes list'));

      // Execute
      await program.parseAsync(['node', 'test', 'list', 'test-alias']);

      // Assert
      expect(execSync).toHaveBeenCalledWith(
        'shopify theme list --store test-store',
        { encoding: 'utf-8' }
      );
    });

    it('should include name filter when provided', async () => {
      // Setup
      mockConfigManager.getStore.mockReturnValue({ 
        storeId: 'test-store', 
        alias: 'test-alias',
        projectDir: 'test-store-dir'
      });
      (execSync as jest.Mock).mockReturnValue(Buffer.from('themes list'));

      // Execute
      await program.parseAsync(['node', 'test', 'list', 'test-alias', '--name', 'test-theme']);

      // Assert
      expect(execSync).toHaveBeenCalledWith(
        'shopify theme list --store test-store --name test-theme',
        { encoding: 'utf-8' }
      );
    });

    it('should handle store not found error', async () => {
      // Setup
      mockConfigManager.getStore.mockReturnValue(undefined);
      const consoleSpy = jest.spyOn(console, 'error').mockImplementation();

      // Execute
      await program.parseAsync(['node', 'test', 'list', 'non-existent']);

      // Assert
      expect(consoleSpy).toHaveBeenCalledWith(
        'Store with alias "non-existent" not found'
      );
      expect(execSync).not.toHaveBeenCalled();
    });
  });

  describe('dev command', () => {
    it('should start theme development server with correct parameters', async () => {
      // Setup
      mockConfigManager.getStore.mockReturnValue({ 
        storeId: 'test-store', 
        alias: 'test-alias',
        projectDir: 'test-store-dir'
      });
      (spawn as jest.Mock).mockReturnValue({
        on: jest.fn()
      });

      // Execute
      await program.parseAsync(['node', 'test', 'dev', '123456789']);

      // Assert
      expect(spawn).toHaveBeenCalledWith(
        'shopify',
        ['theme', 'dev', '--theme', '123456789'],
        expect.objectContaining({
          stdio: 'inherit',
          shell: true
        })
      );
    });
  });
}); 