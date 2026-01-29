// Package ui 提供 GitQ 的用户界面功能
// display.go 负责使用 lipgloss 美化信息展示
package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gitq/git"
	"gitq/i18n"

	"github.com/charmbracelet/lipgloss"
)

// 定义颜色主题（使用 Catppuccin Mocha 配色）
var (
	// 主色调
	primaryColor   = lipgloss.Color("#cba6f7") // Mauve - 紫色
	secondaryColor = lipgloss.Color("#89b4fa") // Blue - 蓝色
	accentColor    = lipgloss.Color("#a6e3a1") // Green - 绿色
	warningColor   = lipgloss.Color("#fab387") // Peach - 橙色
	textColor      = lipgloss.Color("#cdd6f4") // Text - 浅色文字
	subtleColor    = lipgloss.Color("#6c7086") // Overlay0 - 灰色
	surfaceColor   = lipgloss.Color("#313244") // Surface0 - 深色背景

	// 标题样式
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primaryColor).
			MarginBottom(1)

	// 标签样式（字段名）
	labelStyle = lipgloss.NewStyle().
			Foreground(subtleColor).
			Width(10)

	// 值样式
	valueStyle = lipgloss.NewStyle().
			Foreground(textColor)

	// 高亮值样式
	highlightStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(accentColor)

	// URL 样式
	urlStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).
			Italic(true)

	// 用户信息样式
	userNameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(warningColor)

	userEmailStyle = lipgloss.NewStyle().
			Foreground(secondaryColor)

	// 边框样式
	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(1, 2).
			MarginTop(1).
			MarginBottom(1)

	// 分隔线样式
	dividerStyle = lipgloss.NewStyle().
			Foreground(subtleColor)
)

// RenderGitInfo 渲染 Git 仓库信息为美观的终端输出
// 使用 lipgloss 样式美化显示效果
func RenderGitInfo(info *git.GitInfo) string {
	t := i18n.T()
	var builder strings.Builder

	// 标题
	// 获取当前工作目录名称作为标题
	dir, _ := os.Getwd()
	folderName := filepath.Base(dir)
	title := titleStyle.Render("GitQ - " + folderName)
	builder.WriteString(title + "\n")

	// 远程分支信息
	remoteBranch := info.RemoteBranch
	if remoteBranch == "" {
		remoteBranch = t.GitNotSet
	}
	builder.WriteString(renderLine(t.GitRemote, valueStyle.Render(remoteBranch)) + "\n")

	// 远程 URL
	remoteURL := info.RemoteURL
	if remoteURL == "" || remoteURL == "N/A" {
		remoteURL = t.GitNA
	}
	builder.WriteString(renderLine(t.GitURL, urlStyle.Render(remoteURL)) + "\n")

	// 分隔线
	divider := dividerStyle.Render(strings.Repeat("─", 40))
	builder.WriteString(divider + "\n")

	// 用户名
	userName := info.UserName
	if userName == "" {
		userName = t.GitNotSet
	}
	builder.WriteString(renderLine(t.GitUser, userNameStyle.Render(userName)) + "\n")

	// 邮箱
	userEmail := info.UserEmail
	if userEmail == "" {
		userEmail = t.GitNotSet
	}
	builder.WriteString(renderLine(t.GitEmail, userEmailStyle.Render(userEmail)))

	// 使用边框包裹
	content := builder.String()
	return boxStyle.Render(content)
}

// renderLine 渲染一行信息（标签 + 值）
func renderLine(label, value string) string {
	return fmt.Sprintf("%s  %s", labelStyle.Render(label), value)
}

// RenderSuccess 渲染成功消息
func RenderSuccess(message string) string {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(accentColor)
	return style.Render("✅ " + message)
}

// RenderError 渲染错误消息
func RenderError(message string) string {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#f38ba8")) // Red
	return style.Render("❌ " + message)
}

// RenderWarning 渲染警告消息
func RenderWarning(message string) string {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(warningColor)
	return style.Render("⚠️  " + message)
}
