#!/usr/bin/env node
/**
 * GitQ 跨平台构建脚本
 * 使用 Go 交叉编译生成 6 个平台的二进制文件
 * 同时生成每个平台包的 package.json
 */

const { execSync } = require("child_process");
const fs = require("fs");
const path = require("path");

// 读取主包的 package.json 获取版本号
const ROOT_DIR = path.resolve(__dirname, "..");
const NPM_DIR = path.join(ROOT_DIR, "npm");
const mainPkg = JSON.parse(fs.readFileSync(path.join(ROOT_DIR, "package.json"), "utf-8"));
const VERSION = mainPkg.version;

// 构建目标平台配置
// GOOS/GOARCH -> npm 平台名称的映射
const targets = [
  { goos: "darwin", goarch: "amd64", platform: "darwin-x64", os: "darwin", cpu: "x64", ext: "" },
  { goos: "darwin", goarch: "arm64", platform: "darwin-arm64", os: "darwin", cpu: "arm64", ext: "" },
  { goos: "linux", goarch: "amd64", platform: "linux-x64", os: "linux", cpu: "x64", ext: "" },
  { goos: "linux", goarch: "arm64", platform: "linux-arm64", os: "linux", cpu: "arm64", ext: "" },
  { goos: "windows", goarch: "amd64", platform: "win32-x64", os: "win32", cpu: "x64", ext: ".exe" },
  { goos: "windows", goarch: "arm64", platform: "win32-arm64", os: "win32", cpu: "arm64", ext: ".exe" },
];

/**
 * 确保目录存在
 */
function ensureDir(dir) {
  if (!fs.existsSync(dir)) {
    fs.mkdirSync(dir, { recursive: true });
  }
}

/**
 * 生成平台包的 package.json
 */
function generatePackageJson(target) {
  const { platform, os, cpu, ext } = target;
  const outputDir = path.join(NPM_DIR, platform);
  const binaryName = `gitq${ext}`;
  
  // 使用 gitq- 前缀命名（非 @gitq 组织）
  const packageJson = {
    name: `gitq-${platform}`,
    version: VERSION,
    description: `GitQ binary for ${os} ${cpu}`,
    os: [os],
    cpu: [cpu],
    main: binaryName,
    files: [binaryName],
    license: "MIT",
    repository: mainPkg.repository,
    publishConfig: {
      registry: "https://registry.npmjs.org/"
    }
  };

  const packageJsonPath = path.join(outputDir, "package.json");
  fs.writeFileSync(packageJsonPath, JSON.stringify(packageJson, null, 2) + "\n");
  console.log(`   📄 生成 ${platform}/package.json`);
}

/**
 * 执行 Go 交叉编译
 */
function buildTarget(target) {
  const { goos, goarch, platform, ext } = target;
  const outputDir = path.join(NPM_DIR, platform);
  const binaryName = `gitq${ext}`;
  const outputPath = path.join(outputDir, binaryName);

  console.log(`📦 构建 ${platform}...`);

  // 确保输出目录存在
  ensureDir(outputDir);

  // 生成 package.json
  generatePackageJson(target);

  // 设置环境变量并执行 Go 编译
  const env = {
    ...process.env,
    GOOS: goos,
    GOARCH: goarch,
    CGO_ENABLED: "0", // 禁用 CGO 以支持交叉编译
  };

  try {
    execSync(`go build -ldflags="-s -w" -o "${outputPath}" .`, {
      cwd: ROOT_DIR,
      env,
      stdio: "inherit",
    });
    console.log(`   ✅ 构建成功: ${binaryName}`);
  } catch (error) {
    console.error(`   ❌ 构建失败:`, error.message);
    process.exit(1);
  }
}

/**
 * 主函数
 */
function main() {
  console.log("🔧 GitQ 跨平台构建开始\n");
  console.log(`📁 项目目录: ${ROOT_DIR}`);
  console.log(`📁 输出目录: ${NPM_DIR}`);
  console.log(`📌 版本号: ${VERSION}\n`);

  // 构建所有目标平台
  for (const target of targets) {
    buildTarget(target);
    console.log(""); // 空行分隔
  }

  console.log("🎉 所有平台构建完成！");

  // 显示构建结果
  console.log("\n📊 构建结果:");
  for (const target of targets) {
    const binaryPath = path.join(NPM_DIR, target.platform, `gitq${target.ext}`);
    if (fs.existsSync(binaryPath)) {
      const stats = fs.statSync(binaryPath);
      const sizeMB = (stats.size / 1024 / 1024).toFixed(2);
      console.log(`   gitq-${target.platform}: ${sizeMB} MB`);
    }
  }
}

main();

