import { ConfigManager } from '../../utils/config';
import { existsSync, readFileSync, writeFileSync, mkdirSync } from 'fs';
import { join } from 'path';
import { homedir } from 'os';

jest.mock('fs');
jest.mock('os');

describe('ConfigManager', () => {
  const mockHomedir = '/mock/home';
  const mockConfigDir = join(mockHomedir, '.config', 'shopify-theme-manager');
  const mockConfigPath = join(mockConfigDir, 'config.json');
  
  beforeEach(() => {
    // Reset all mocks
    jest.clearAllMocks();
    
    // Setup default mock implementations
    (homedir as jest.Mock).mockReturnValue(mockHomedir);
    (existsSync as jest.Mock).mockReturnValue(true);
    (readFileSync as jest.Mock).mockReturnValue(JSON.stringify({ stores: [] }));
    (writeFileSync as jest.Mock).mockImplementation(() => {});
    (mkdirSync as jest.Mock).mockImplementation(() => {});
  });

  describe('constructor', () => {
    it('should create config directory if it does not exist', () => {
      // Setup
      (existsSync as jest.Mock)
        .mockReturnValueOnce(false) // config dir doesn't exist
        .mockReturnValueOnce(true); // config file exists

      // Execute
      new ConfigManager();

      // Assert
      expect(mkdirSync).toHaveBeenCalledWith(mockConfigDir, { recursive: true });
    });

    it('should create config file if it does not exist', () => {
      // Setup
      (existsSync as jest.Mock)
        .mockReturnValueOnce(true) // config dir exists
        .mockReturnValueOnce(false); // config file doesn't exist

      // Execute
      new ConfigManager();

      // Assert
      expect(writeFileSync).toHaveBeenCalledWith(
        mockConfigPath,
        JSON.stringify({ stores: [] }, null, 2)
      );
    });
  });

  describe('addStore', () => {
    it('should add a new store to the configuration', () => {
      // Setup
      const config = new ConfigManager();
      const writeSpy = jest.spyOn(ConfigManager.prototype as any, 'saveConfig');

      // Execute
      config.addStore('test-store-id', 'test-alias', 'test-store-dir');

      // Assert
      expect(writeSpy).toHaveBeenCalledWith({
        stores: [{ storeId: 'test-store-id', alias: 'test-alias', projectDir: 'test-store-dir' }]
      });
    });
  });

  describe('getStore', () => {
    it('should return store config when alias exists', () => {
      // Setup
      (readFileSync as jest.Mock).mockReturnValue(JSON.stringify({
        stores: [{ 
          storeId: 'test-store-id', 
          alias: 'test-alias',
          projectDir: 'test-store-dir'
        }]
      }));
      
      const config = new ConfigManager();

      // Execute
      const store = config.getStore('test-alias');

      // Assert
      expect(store).toEqual({ 
        storeId: 'test-store-id', 
        alias: 'test-alias',
        projectDir: 'test-store-dir'
      });
    });

    it('should return undefined when alias does not exist', () => {
      // Setup
      const config = new ConfigManager();

      // Execute
      const store = config.getStore('non-existent');

      // Assert
      expect(store).toBeUndefined();
    });
  });

  describe('listStores', () => {
    it('should return all stored configurations', () => {
      // Setup
      const mockStores = [
        { storeId: 'store-1', alias: 'alias-1', projectDir: 'dir-1' },
        { storeId: 'store-2', alias: 'alias-2', projectDir: 'dir-2' }
      ];
      
      (readFileSync as jest.Mock).mockReturnValue(JSON.stringify({
        stores: mockStores
      }));
      
      const config = new ConfigManager();

      // Execute
      const stores = config.listStores();

      // Assert
      expect(stores).toEqual(mockStores);
    });
  });

  describe('workspace', () => {
    it('should set workspace in config', () => {
      // Setup
      const config = new ConfigManager();
      const workspace = '/path/to/workspace';
      const writeSpy = jest.spyOn(ConfigManager.prototype as any, 'saveConfig');

      // Execute
      config.setWorkspace(workspace);

      // Assert
      expect(writeSpy).toHaveBeenCalledWith(
        expect.objectContaining({
          workspace: workspace,
          stores: expect.any(Array)
        })
      );
    });

    it('should normalize the directory path', () => {
      // Setup
      const config = new ConfigManager();
      const writeSpy = jest.spyOn(ConfigManager.prototype as any, 'saveConfig');

      // Execute
      config.setWorkspace('./relative/path');

      // Assert
      expect(writeSpy).toHaveBeenCalledWith(
        expect.objectContaining({
          workspace: expect.stringMatching(/^\/.*\/relative\/path$/),
          stores: expect.any(Array)
        })
      );
    });

    it('should get workspace from config', () => {
      // Setup
      const workspace = '/path/to/workspace';
      (readFileSync as jest.Mock).mockReturnValue(JSON.stringify({
        stores: [],
        workspace: workspace
      }));
      
      const config = new ConfigManager();

      // Execute
      const result = config.getWorkspace();

      // Assert
      expect(result).toBe(workspace);
    });

    it('should return undefined if no workspace is set', () => {
      // Setup
      (readFileSync as jest.Mock).mockReturnValue(JSON.stringify({
        stores: []
      }));
      
      const config = new ConfigManager();

      // Execute
      const result = config.getWorkspace();

      // Assert
      expect(result).toBeUndefined();
    });
  });

}); 