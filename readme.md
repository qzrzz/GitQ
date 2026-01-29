# GitQ

🔧 Git 用户切换命令行工具 - 使用 [charmbracelet/huh](https://github.com/charmbracelet/huh) 构建的美观终端界面

> 之所以制作这个工具是因为 [Git-User-Switch](https://github.com/geongeorge/Git-User-Switch) 不知什么原因有时候无法切换用户

- 📌 显示当前 Git 仓库信息（分支、远程地址、用户）
- 🔄 快速切换 Git 用户
  - `user.name`
  - `user.email`
  - `user.signingkey`
  - GitHub 用户名，可以切换 git remote 地址中的用户名例如 `git@github.com-qzrzz:qzrzz/GitQ.git`，可以引导 SSH 使用指定的用户密钥([详细方法](https://gist.github.com/oanhnn/80a89405ab9023894df7))。
- ⚙️ 管理预设用户（添加、编辑、删除）

## 安装

### 通过 npm 安装

```bash
npm install -g @qzrzz/gitq
```

## 使用

```bash
# 在 Git 仓库目录中运行
gitq
```

### 截图

```
╭────────────────────────────────────────────╮
│                                            │
│  GitQ - newpkg                             │
│                                            │
│  🔗 Remote   Not set                       │
│  🌐 URL      N/A                           │
│  ────────────────────────────────────────  │
│  👤 User     Qzrzz                         │
│  📧 Email    qzrz256@gmail.com             │
│                                            │
╰────────────────────────────────────────────╯

┃ Select action (ESC to quit)
┃ > 🔄 Switch Git User
┃   ⚙️  Manage Users
┃   🌐 Change Language
┃   🔃 Refresh
┃   🚪 Exit

↑ up • ↓ down • / filter • enter submit
```

## 配置文件

预设用户存储在 `~/.gitq/users.json`：

```json
[
  { "name": "Work Account", "email": "work@company.com" },
  { "name": "Personal", "email": "personal@example.com" }
]
```

## 开发

### 从源码构建

```bash
# 安装依赖
go mod download

# 编译当前平台
go build -o dist/gitq .

# 编译所有平台（用于 npm 发布）
node scripts/build.js

# 调试
go run .

```

### npm 发布

### 准备工作

1. 登录 npm：

   ```bash
   npm login
   ```

### 一键发布

```bash
# 1. 构建所有平台
npm run build:all

# 2. 检查发布（dry run）
npm run publish:all -- --dry-run

# 3. 正式发布
npm run publish:all
```

发布脚本会自动按顺序发布 6 个平台包 → 主包。

### 技术栈

- [charmbracelet/huh](https://github.com/charmbracelet/huh) - 终端表单库
- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) - 终端样式
- Catppuccin Mocha 配色主题
