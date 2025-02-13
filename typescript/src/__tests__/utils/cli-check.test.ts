import { checkShopifyCLI, ensureShopifyCLI } from '../../utils/cli-check';
import { execSync } from 'child_process';

jest.mock('child_process');

describe('CLI Check Utils', () => {
  const mockExit = jest.spyOn(process, 'exit').mockImplementation((number) => { throw new Error('process.exit: ' + number); });
  const mockConsoleError = jest.spyOn(console, 'error').mockImplementation();

  beforeEach(() => {
    jest.clearAllMocks();
  });

  describe('checkShopifyCLI', () => {
    it('should return true when Shopify CLI is installed', () => {
      (execSync as jest.Mock).mockImplementation(() => {});
      expect(checkShopifyCLI()).toBe(true);
    });

    it('should return false when Shopify CLI is not installed', () => {
      (execSync as jest.Mock).mockImplementation(() => {
        throw new Error('command not found');
      });
      expect(checkShopifyCLI()).toBe(false);
    });
  });

  describe('ensureShopifyCLI', () => {
    it('should not exit when Shopify CLI is installed', () => {
      (execSync as jest.Mock).mockImplementation(() => {});
      expect(() => ensureShopifyCLI()).not.toThrow();
      expect(mockExit).not.toHaveBeenCalled();
    });

    it('should exit with error when Shopify CLI is not installed', () => {
      (execSync as jest.Mock).mockImplementation(() => {
        throw new Error('command not found');
      });
      
      expect(() => ensureShopifyCLI()).toThrow('process.exit: 1');
      expect(mockConsoleError).toHaveBeenCalledWith('Shopify CLI is not installed. Please install it first:');
      expect(mockConsoleError).toHaveBeenCalledWith('npm install -g @shopify/cli @shopify/theme');
      expect(mockExit).toHaveBeenCalledWith(1);
    });
  });
}); 