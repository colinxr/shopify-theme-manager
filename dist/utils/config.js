"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConfigManager = void 0;
const os_1 = require("os");
const path_1 = require("path");
const fs_1 = require("fs");
class ConfigManager {
    constructor() {
        this.configDir = (0, path_1.join)((0, os_1.homedir)(), '.config', 'shopify-theme-manager');
        this.configPath = (0, path_1.join)(this.configDir, 'config.json');
        this.ensureConfigExists();
        this.config = this.loadConfig();
    }
    ensureConfigExists() {
        if (!(0, fs_1.existsSync)(this.configDir)) {
            (0, fs_1.mkdirSync)(this.configDir, { recursive: true });
        }
        if (!(0, fs_1.existsSync)(this.configPath)) {
            this.saveConfig({ stores: [] });
        }
    }
    loadConfig() {
        const configData = (0, fs_1.readFileSync)(this.configPath, 'utf-8');
        return JSON.parse(configData);
    }
    saveConfig(config) {
        (0, fs_1.writeFileSync)(this.configPath, JSON.stringify(config, null, 2));
    }
    addStore(storeId, alias) {
        const store = { storeId, alias };
        this.config.stores.push(store);
        this.saveConfig(this.config);
    }
    getStore(alias) {
        return this.config.stores.find(store => store.alias === alias);
    }
    listStores() {
        return this.config.stores;
    }
}
exports.ConfigManager = ConfigManager;
