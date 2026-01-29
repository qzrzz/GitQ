#!/usr/bin/env node
/**
 * GitQ npm 可执行入口
 * 根据当前平台自动定位并执行对应的二进制文件
 */

const { spawn } = require('child_process');
const path = require('path');
const fs = require('fs');

// 平台映射：Node.js 平台/架构 -> npm 包名
const PLATFORM_MAP = {
  'darwin-x64': 'gitq-darwin-x64',
  'darwin-arm64': 'gitq-darwin-arm64',
  'linux-x64': 'gitq-linux-x64',
  'linux-arm64': 'gitq-linux-arm64',
  'win32-x64': 'gitq-win32-x64',
  'win32-arm64': 'gitq-win32-arm64',
};

/**
 * 获取当前平台的标识符
 */
function getPlatformKey() {
  const platform = process.platform;
  const arch = process.arch;
  return `${platform}-${arch}`;
}

/**
 * 获取二进制文件名
 */
function getBinaryName() {
  return process.platform === 'win32' ? 'gitq.exe' : 'gitq';
}

/**
 * 尝试查找二进制文件路径
 */
function findBinaryPath() {
  const platformKey = getPlatformKey();
  const packageName = PLATFORM_MAP[platformKey];
  const binaryName = getBinaryName();

  if (!packageName) {
    console.error(`❌ 不支持的平台: ${platformKey}`);
    console.error('支持的平台:', Object.keys(PLATFORM_MAP).join(', '));
    process.exit(1);
  }

  // 方法 1: 从 node_modules 中查找平台包
  const possiblePaths = [
    // 安装在项目 node_modules
    path.join(__dirname, '..', 'node_modules', packageName, binaryName),
    // 安装在全局或 pnpm/yarn 的特殊位置
    path.join(__dirname, '..', '..', packageName, binaryName),
    // 直接在 npm 目录中（开发模式）
    path.join(__dirname, '..', 'npm', platformKey.replace('-', '-'), binaryName),
  ];

  for (const binaryPath of possiblePaths) {
    if (fs.existsSync(binaryPath)) {
      return binaryPath;
    }
  }

  // 方法 2: 使用 require.resolve 查找
  try {
    const packagePath = require.resolve(`${packageName}/package.json`);
    const packageDir = path.dirname(packagePath);
    const binaryPath = path.join(packageDir, binaryName);
    if (fs.existsSync(binaryPath)) {
      return binaryPath;
    }
  } catch (e) {
    // 包未安装
  }

  console.error(`❌ 找不到平台二进制文件: ${packageName}`);
  console.error('请确保已正确安装 gitq 包');
  process.exit(1);
}

/**
 * 执行二进制文件
 */
function runBinary(binaryPath) {
  // 传递命令行参数（跳过 node 和脚本路径）
  const args = process.argv.slice(2);

  const child = spawn(binaryPath, args, {
    stdio: 'inherit', // 继承标准输入/输出/错误
    windowsHide: true,
  });

  child.on('error', (error) => {
    console.error(`❌ 执行失败: ${error.message}`);
    process.exit(1);
  });

  child.on('close', (code) => {
    process.exit(code || 0);
  });
}

// 主入口
const binaryPath = findBinaryPath();
runBinary(binaryPath);
