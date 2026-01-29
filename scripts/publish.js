#!/usr/bin/env node
/**
 * GitQ 一键发布脚本
 * 自动发布所有平台包和主包到 npm
 */

const { execSync } = require('child_process');
const path = require('path');
const fs = require('fs');

const ROOT_DIR = path.resolve(__dirname, '..');
const NPM_DIR = path.join(ROOT_DIR, 'npm');

// 平台包列表
const platforms = [
  'darwin-x64',
  'darwin-arm64',
  'linux-x64',
  'linux-arm64',
  'win32-x64',
  'win32-arm64',
];

/**
 * 执行命令并打印输出
 */
function run(cmd, cwd) {
  console.log(`📍 ${cwd}`);
  console.log(`   $ ${cmd}`);
  try {
    execSync(cmd, { cwd, stdio: 'inherit' });
    return true;
  } catch (error) {
    console.error(`❌ 命令执行失败`);
    return false;
  }
}

/**
 * 检查平台包是否包含二进制文件
 */
function checkBinaries() {
  console.log('\n🔍 检查二进制文件...\n');
  let allExist = true;
  
  for (const platform of platforms) {
    const ext = platform.startsWith('win32') ? '.exe' : '';
    const binaryPath = path.join(NPM_DIR, platform, `gitq${ext}`);
    const exists = fs.existsSync(binaryPath);
    
    if (exists) {
      const stats = fs.statSync(binaryPath);
      const sizeMB = (stats.size / 1024 / 1024).toFixed(2);
      console.log(`   ✅ ${platform}: ${sizeMB} MB`);
    } else {
      console.log(`   ❌ ${platform}: 未找到`);
      allExist = false;
    }
  }
  
  return allExist;
}

/**
 * 发布平台包
 */
function publishPlatformPackages() {
  console.log('\n📦 发布平台包...\n');
  
  for (const platform of platforms) {
    const packageDir = path.join(NPM_DIR, platform);
    console.log(`\n--- ${platform} ---`);
    
    if (!run('npm publish --access public', packageDir)) {
      console.error(`\n❌ ${platform} 发布失败，终止发布流程`);
      return false;
    }
    
    console.log(`✅ gitq-${platform} 发布成功`);
  }
  
  return true;
}

/**
 * 发布主包
 */
function publishMainPackage() {
  console.log('\n📦 发布主包...\n');
  
  if (!run('npm publish --access public', ROOT_DIR)) {
    console.error('\n❌ 主包发布失败');
    return false;
  }
  
  console.log('✅ gitq 主包发布成功');
  return true;
}

/**
 * 主函数
 */
function main() {
  console.log('🚀 GitQ 一键发布\n');
  
  // 检查参数
  const args = process.argv.slice(2);
  const dryRun = args.includes('--dry-run');
  
  if (dryRun) {
    console.log('⚠️  Dry run 模式，不会实际发布\n');
  }
  
  // 1. 检查二进制文件
  if (!checkBinaries()) {
    console.error('\n❌ 缺少二进制文件，请先运行: npm run build:all');
    process.exit(1);
  }
  
  if (dryRun) {
    console.log('\n✅ Dry run 检查通过');
    console.log('   移除 --dry-run 参数以实际发布');
    return;
  }
  
  // 2. 发布平台包
  if (!publishPlatformPackages()) {
    process.exit(1);
  }
  
  // 3. 发布主包
  if (!publishMainPackage()) {
    process.exit(1);
  }
  
  console.log('\n🎉 所有包发布完成！');
  console.log('\n用户可以通过以下命令安装:');
  console.log('   npm install -g gitq');
}

main();
