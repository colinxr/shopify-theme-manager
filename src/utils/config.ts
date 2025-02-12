import { homedir } from 'os';
import { join } from 'path';
import { existsSync, mkdirSync, readFileSync, writeFileSync } from 'fs';

interface StoreConfig {
  storeId: string;
  alias: string;
}

interface Config {
  stores: StoreConfig[];
}

export class ConfigManager {
  private configDir: string;
  private configPath: string;
  private config: Config;

  constructor() {
    this.configDir = join(homedir(), '.config', 'shopify-theme-manager');
    this.configPath = join(this.configDir, 'config.json');
    this.ensureConfigExists();
    this.config = this.loadConfig();
  }

  private ensureConfigExists(): void {
    if (!existsSync(this.configDir)) {
      mkdirSync(this.configDir, { recursive: true });
    }
    
    if (!existsSync(this.configPath)) {
      this.saveConfig({ stores: [] });
    }
  }

  private loadConfig(): Config {
    const configData = readFileSync(this.configPath, 'utf-8');
    return JSON.parse(configData);
  }

  private saveConfig(config: Config): void {
    writeFileSync(this.configPath, JSON.stringify(config, null, 2));
  }

  addStore(storeId: string, alias: string): void {
    const store = { storeId, alias };
    this.config.stores.push(store);
    this.saveConfig(this.config);
  }

  getStore(alias: string): StoreConfig | undefined {
    return this.config.stores.find(store => store.alias === alias);
  }

  listStores(): StoreConfig[] {
    return this.config.stores;
  }
} 