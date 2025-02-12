import { execSync } from 'child_process';

export function checkShopifyCLI(): boolean {
  try {
    execSync('shopify version', { stdio: 'ignore' });
    return true;
  } catch (error) {
    return false;
  }
}

export function ensureShopifyCLI(): void {
  if (!checkShopifyCLI()) {
    console.error('Shopify CLI is not installed. Please install it first:');
    console.error('npm install -g @shopify/cli @shopify/theme');
    process.exit(1);
  }
} 