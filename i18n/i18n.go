// Package i18n 提供国际化支持
// 支持中文、英文、日文三种语言
// 根据系统环境变量自动检测语言，或从配置文件读取
package i18n

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Language 语言类型
type Language string

const (
	LangAuto Language = "auto" // 自动检测
	LangZH   Language = "zh"   // 中文
	LangEN   Language = "en"   // 英文
	LangJA   Language = "ja"   // 日文
)

// LanguageInfo 语言显示信息
type LanguageInfo struct {
	Code Language
	Name string
}

// AvailableLanguages 可用语言列表
var AvailableLanguages = []LanguageInfo{
	{LangAuto, "🌐 Auto"},
	{LangZH, "🇨🇳 中文"},
	{LangEN, "🇺🇸 English"},
	{LangJA, "🇯🇵 日本語"},
}

// 当前语言
var currentLang Language = LangZH

// 配置文件路径
func getLangConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".gitq", "lang")
}

// Texts 包含所有 UI 文本的翻译
type Texts struct {
	// 通用
	AppTitle string
	Goodbye  string
	Yes      string
	No       string
	Back     string
	Confirm  string
	Cancel   string
	Success  string
	Error    string
	Warning  string

	// 主菜单
	MenuTitle       string
	MenuSwitchUser  string
	MenuManageUser  string
	MenuSetRemote   string
	MenuChangeLang  string
	MenuRefresh     string
	MenuExit        string
	ChangeLangTitle string

	// Git 信息
	GitBranch string
	GitRemote string
	GitURL    string
	GitUser   string
	GitEmail  string
	GitNotSet string
	GitNA     string

	// 错误消息
	ErrNotGitRepo      string
	ErrRunInGitRepo    string
	ErrLoadUsers       string
	ErrSwitchFailed    string
	ErrAddFailed       string
	ErrUpdateFailed    string
	ErrDeleteFailed    string
	ErrSetRemoteFailed string

	// 用户管理
	UserNoPreset         string
	UserAddNow           string
	UserSelectSwitch     string
	UserManage           string
	UserAdd              string
	UserEdit             string
	UserDelete           string
	UserSelectEdit       string
	UserSelectDelete     string
	UserNoEdit           string
	UserNoDelete         string
	UserConfirmDelete    string
	UserName             string
	UserEmail            string
	UserSigningKey       string
	UserGithubName       string
	UserNameEmpty        string
	UserEmailEmpty       string
	UserNameExample      string
	UserEmailExample     string
	UserSigningKeyEx     string
	UserGithubNameEx     string
	SetRemoteTitle       string
	SetRemotePlaceholder string
	ErrRemoteEmpty       string

	// 成功消息
	SuccessSwitched  string
	SuccessAdded     string
	SuccessUpdated   string
	SuccessDeleted   string
	SuccessSetRemote string
}

// 中文翻译
var zhTexts = Texts{
	AppTitle: "GitQ",
	Goodbye:  "再见! 👋",
	Yes:      "是",
	No:       "否",
	Back:     "↩️  返回",
	Confirm:  "确定",
	Cancel:   "取消",
	Success:  "成功",
	Error:    "错误",
	Warning:  "警告",

	MenuTitle:       "请选择操作 (ESC 退出)",
	MenuSwitchUser:  "🔄 切换 Git 用户",
	MenuManageUser:  "⚙️  管理预设用户",
	MenuSetRemote:   "🔗 设置远程地址",
	MenuChangeLang:  "🌐 切换语言",
	MenuRefresh:     "🔃 刷新信息",
	MenuExit:        "🚪 退出",
	ChangeLangTitle: "选择语言",

	GitBranch: "🔗 远程",
	GitRemote: "🔗 远程",
	GitURL:    "🌐 地址",
	GitUser:   "👤 用户",
	GitEmail:  "📧 邮箱",
	GitNotSet: "未设置",
	GitNA:     "N/A",

	ErrNotGitRepo:      "当前目录不是 Git 仓库",
	ErrRunInGitRepo:    "请在 Git 仓库目录中运行此工具",
	ErrLoadUsers:       "加载用户列表失败",
	ErrSwitchFailed:    "切换用户失败",
	ErrAddFailed:       "添加用户失败",
	ErrUpdateFailed:    "更新用户失败",
	ErrDeleteFailed:    "删除用户失败",
	ErrSetRemoteFailed: "设置远程地址失败",

	UserNoPreset:         "没有预设用户",
	UserAddNow:           "是否立即添加预设用户？",
	UserSelectSwitch:     "选择要切换的用户",
	UserManage:           "用户管理",
	UserAdd:              "➕ 添加用户",
	UserEdit:             "✏️  编辑用户",
	UserDelete:           "🗑️  删除用户",
	UserSelectEdit:       "选择要编辑的用户",
	UserSelectDelete:     "选择要删除的用户",
	UserNoEdit:           "没有可编辑的用户",
	UserNoDelete:         "没有可删除的用户",
	UserConfirmDelete:    "确定删除用户 %s <%s>？",
	UserName:             "用户名",
	UserEmail:            "邮箱",
	UserSigningKey:       "签名密钥 (可选)",
	UserGithubName:       "GitHub 用户名 (可选)",
	UserNameEmpty:        "用户名不能为空",
	UserEmailEmpty:       "邮箱不能为空",
	UserNameExample:      "例如: John Doe",
	UserEmailExample:     "例如: john@example.com",
	UserSigningKeyEx:     "GPG 密钥 ID (可留空)",
	UserGithubNameEx:     "用于 SSH 密钥切换 (可留空)",
	SetRemoteTitle:       "设置远程仓库地址 (origin)",
	SetRemotePlaceholder: "例如: git@github.com:user/repo.git",
	ErrRemoteEmpty:       "远程地址不能为空",

	SuccessSwitched:  "已切换到用户: %s <%s>",
	SuccessAdded:     "已添加用户: %s <%s>",
	SuccessUpdated:   "已更新用户: %s <%s>",
	SuccessDeleted:   "已删除用户",
	SuccessSetRemote: "已设置远程地址: %s",
}

// 英文翻译
var enTexts = Texts{
	AppTitle: "GitQ",
	Goodbye:  "Goodbye! 👋",
	Yes:      "Yes",
	No:       "No",
	Back:     "↩️  Back",
	Confirm:  "Confirm",
	Cancel:   "Cancel",
	Success:  "Success",
	Error:    "Error",
	Warning:  "Warning",

	MenuTitle:       "Select action (ESC to quit)",
	MenuSwitchUser:  "🔄 Switch Git User",
	MenuManageUser:  "⚙️  Manage Users",
	MenuSetRemote:   "🔗 Set Remote URL",
	MenuChangeLang:  "🌐 Change Language",
	MenuRefresh:     "🔃 Refresh",
	MenuExit:        "🚪 Exit",
	ChangeLangTitle: "Select Language",

	GitBranch: "🔗 Remote",
	GitRemote: "🔗 Remote",
	GitURL:    "🌐 URL",
	GitUser:   "👤 User",
	GitEmail:  "📧 Email",
	GitNotSet: "Not set",
	GitNA:     "N/A",

	ErrNotGitRepo:      "Current directory is not a Git repository",
	ErrRunInGitRepo:    "Please run this tool in a Git repository",
	ErrLoadUsers:       "Failed to load users",
	ErrSwitchFailed:    "Failed to switch user",
	ErrAddFailed:       "Failed to add user",
	ErrUpdateFailed:    "Failed to update user",
	ErrDeleteFailed:    "Failed to delete user",
	ErrSetRemoteFailed: "Failed to set remote URL",

	UserNoPreset:         "No preset users",
	UserAddNow:           "Add a user now?",
	UserSelectSwitch:     "Select user to switch",
	UserManage:           "User Management",
	UserAdd:              "➕ Add User",
	UserEdit:             "✏️  Edit User",
	UserDelete:           "🗑️  Delete User",
	UserSelectEdit:       "Select user to edit",
	UserSelectDelete:     "Select user to delete",
	UserNoEdit:           "No users to edit",
	UserNoDelete:         "No users to delete",
	UserConfirmDelete:    "Delete user %s <%s>?",
	UserName:             "Name",
	UserEmail:            "Email",
	UserSigningKey:       "Signing Key (optional)",
	UserGithubName:       "GitHub Username (optional)",
	UserNameEmpty:        "Name cannot be empty",
	UserEmailEmpty:       "Email cannot be empty",
	UserNameExample:      "e.g. John Doe",
	UserEmailExample:     "e.g. john@example.com",
	UserSigningKeyEx:     "GPG key ID (optional)",
	UserGithubNameEx:     "For SSH key switching (optional)",
	SetRemoteTitle:       "Set Remote URL (origin)",
	SetRemotePlaceholder: "e.g. git@github.com:user/repo.git",
	ErrRemoteEmpty:       "Remote URL cannot be empty",

	SuccessSwitched:  "Switched to: %s <%s>",
	SuccessAdded:     "Added user: %s <%s>",
	SuccessUpdated:   "Updated user: %s <%s>",
	SuccessDeleted:   "User deleted",
	SuccessSetRemote: "Remote URL set: %s",
}

// 日文翻译
var jaTexts = Texts{
	AppTitle: "GitQ",
	Goodbye:  "さようなら! 👋",
	Yes:      "はい",
	No:       "いいえ",
	Back:     "↩️  戻る",
	Confirm:  "確認",
	Cancel:   "キャンセル",
	Success:  "成功",
	Error:    "エラー",
	Warning:  "警告",

	MenuTitle:       "操作を選択 (ESC で終了)",
	MenuSwitchUser:  "🔄 Git ユーザーを切替",
	MenuManageUser:  "⚙️  ユーザー管理",
	MenuSetRemote:   "🔗 リモート設定",
	MenuChangeLang:  "🌐 言語切替",
	MenuRefresh:     "🔃 更新",
	MenuExit:        "🚪 終了",
	ChangeLangTitle: "言語を選択",

	GitBranch: "🔗 リモート",
	GitRemote: "🔗 リモート",
	GitURL:    "🌐 URL",
	GitUser:   "👤 ユーザー",
	GitEmail:  "📧 メール",
	GitNotSet: "未設定",
	GitNA:     "N/A",

	ErrNotGitRepo:      "現在のディレクトリは Git リポジトリではありません",
	ErrRunInGitRepo:    "Git リポジトリで実行してください",
	ErrLoadUsers:       "ユーザーリストの読み込みに失敗しました",
	ErrSwitchFailed:    "ユーザーの切替に失敗しました",
	ErrAddFailed:       "ユーザーの追加に失敗しました",
	ErrUpdateFailed:    "ユーザーの更新に失敗しました",
	ErrDeleteFailed:    "ユーザーの削除に失敗しました",
	ErrSetRemoteFailed: "リモートの設定に失敗しました",

	UserNoPreset:         "プリセットユーザーがありません",
	UserAddNow:           "今すぐユーザーを追加しますか？",
	UserSelectSwitch:     "切り替えるユーザーを選択",
	UserManage:           "ユーザー管理",
	UserAdd:              "➕ ユーザー追加",
	UserEdit:             "✏️  ユーザー編集",
	UserDelete:           "🗑️  ユーザー削除",
	UserSelectEdit:       "編集するユーザーを選択",
	UserSelectDelete:     "削除するユーザーを選択",
	UserNoEdit:           "編集できるユーザーがありません",
	UserNoDelete:         "削除できるユーザーがありません",
	UserConfirmDelete:    "ユーザー %s <%s> を削除しますか？",
	UserName:             "名前",
	UserEmail:            "メール",
	UserSigningKey:       "署名キー (任意)",
	UserGithubName:       "GitHub ユーザー名 (任意)",
	UserNameEmpty:        "名前は必須です",
	UserEmailEmpty:       "メールは必須です",
	UserNameExample:      "例: 山田太郎",
	UserEmailExample:     "例: taro@example.com",
	UserSigningKeyEx:     "GPG キー ID (任意)",
	UserGithubNameEx:     "SSH キー切替用 (任意)",
	SetRemoteTitle:       "リモート URL 設定 (origin)",
	SetRemotePlaceholder: "例: git@github.com:user/repo.git",
	ErrRemoteEmpty:       "リモート URL は必須です",

	SuccessSwitched:  "ユーザーを切り替えました: %s <%s>",
	SuccessAdded:     "ユーザーを追加しました: %s <%s>",
	SuccessUpdated:   "ユーザーを更新しました: %s <%s>",
	SuccessDeleted:   "ユーザーを削除しました",
	SuccessSetRemote: "リモート URL を設定しました: %s",
}

// Init 初始化 i18n
// 优先从配置文件读取，否则检测系统语言
func Init() {
	// 1. 尝试从配置文件读取
	configPath := getLangConfigPath()
	if data, err := os.ReadFile(configPath); err == nil {
		savedLang := Language(strings.TrimSpace(string(data)))
		// 如果保存的是 auto，使用系统检测
		if savedLang == LangAuto {
			currentLang = detectSystemLanguage()
			return
		}
		// 如果是有效的语言代码，直接使用
		if savedLang == LangZH || savedLang == LangEN || savedLang == LangJA {
			currentLang = savedLang
			return
		}
	}

	// 2. 检测系统语言
	currentLang = detectSystemLanguage()
}

// detectSystemLanguage 检测系统语言
// 支持 macOS (AppleLanguages) 和 Linux/Unix (LANG 环境变量)
func detectSystemLanguage() Language {
	// macOS: 使用 defaults 命令读取 AppleLanguages
	if output, err := runCommand("defaults", "read", "-g", "AppleLanguages"); err == nil {
		lang := strings.ToLower(output)
		if strings.Contains(lang, "zh") {
			return LangZH
		}
		if strings.Contains(lang, "ja") {
			return LangJA
		}
	}

	// Linux/Unix: 检测 LANG 环境变量
	// 优先级: LANG > LC_ALL > LC_MESSAGES
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}
	if lang == "" {
		lang = os.Getenv("LC_MESSAGES")
	}

	// 解析语言代码
	lang = strings.ToLower(lang)
	switch {
	case strings.HasPrefix(lang, "zh"):
		return LangZH
	case strings.HasPrefix(lang, "ja"):
		return LangJA
	default:
		return LangEN
	}
}

// runCommand 执行命令并返回输出
func runCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// T 获取当前语言的翻译文本
func T() *Texts {
	switch currentLang {
	case LangJA:
		return &jaTexts
	case LangEN:
		return &enTexts
	default:
		return &zhTexts
	}
}

// SetLang 手动设置语言并保存到配置文件
func SetLang(lang Language) {
	// 如果选择自动，检测系统语言
	if lang == LangAuto {
		currentLang = detectSystemLanguage()
	} else {
		currentLang = lang
	}

	// 保存到配置文件（保存用户选择，包括 auto）
	configPath := getLangConfigPath()
	// 确保目录存在
	os.MkdirAll(filepath.Dir(configPath), 0755)
	os.WriteFile(configPath, []byte(string(lang)), 0644)
}

// GetLang 获取当前语言
func GetLang() Language {
	return currentLang
}
