"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const config_1 = require("../../utils/config");
const fs_1 = require("fs");
const path_1 = require("path");
const os_1 = require("os");
jest.mock('fs');
jest.mock('os');
describe('ConfigManager', () => {
    const mockHomedir = '/mock/home';
    const mockConfigDir = (0, path_1.join)(mockHomedir, '.config', 'shopify-theme-manager');
    const mockConfigPath = (0, path_1.join)(mockConfigDir, 'config.json');
    beforeEach(() => {
        // Reset all mocks
        jest.clearAllMocks();
        // Setup default mock implementations
        os_1.homedir.mockReturnValue(mockHomedir);
        fs_1.existsSync.mockReturnValue(true);
        fs_1.readFileSync.mockReturnValue(JSON.stringify({ stores: [] }));
        fs_1.writeFileSync.mockImplementation(() => { });
        fs_1.mkdirSync.mockImplementation(() => { });
    });
    describe('constructor', () => {
        it('should create config directory if it does not exist', () => {
            // Setup
            fs_1.existsSync
                .mockReturnValueOnce(false) // config dir doesn't exist
                .mockReturnValueOnce(true); // config file exists
            // Execute
            new config_1.ConfigManager();
            // Assert
            expect(fs_1.mkdirSync).toHaveBeenCalledWith(mockConfigDir, { recursive: true });
        });
        it('should create config file if it does not exist', () => {
            // Setup
            fs_1.existsSync
                .mockReturnValueOnce(true) // config dir exists
                .mockReturnValueOnce(false); // config file doesn't exist
            // Execute
            new config_1.ConfigManager();
            // Assert
            expect(fs_1.writeFileSync).toHaveBeenCalledWith(mockConfigPath, JSON.stringify({ stores: [] }, null, 2));
        });
    });
    describe('addStore', () => {
        it('should add a new store to the configuration', () => {
            // Setup
            const config = new config_1.ConfigManager();
            const writeSpy = jest.spyOn(config_1.ConfigManager.prototype, 'saveConfig');
            // Execute
            config.addStore('test-store-id', 'test-alias');
            // Assert
            expect(writeSpy).toHaveBeenCalledWith({
                stores: [{ storeId: 'test-store-id', alias: 'test-alias' }]
            });
        });
    });
    describe('getStore', () => {
        it('should return store config when alias exists', () => {
            // Setup
            fs_1.readFileSync.mockReturnValue(JSON.stringify({
                stores: [{ storeId: 'test-store-id', alias: 'test-alias' }]
            }));
            const config = new config_1.ConfigManager();
            // Execute
            const store = config.getStore('test-alias');
            // Assert
            expect(store).toEqual({ storeId: 'test-store-id', alias: 'test-alias' });
        });
        it('should return undefined when alias does not exist', () => {
            // Setup
            const config = new config_1.ConfigManager();
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
                { storeId: 'store-1', alias: 'alias-1' },
                { storeId: 'store-2', alias: 'alias-2' }
            ];
            fs_1.readFileSync.mockReturnValue(JSON.stringify({
                stores: mockStores
            }));
            const config = new config_1.ConfigManager();
            // Execute
            const stores = config.listStores();
            // Assert
            expect(stores).toEqual(mockStores);
        });
    });
});
