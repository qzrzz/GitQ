// Package ui 提供 GitQ 的用户界面功能
// forms.go 使用 huh 库实现交互式表单
package ui

import (
	"fmt"

	"gitq/config"
	"gitq/i18n"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

// MainMenuAction 主菜单操作类型
type MainMenuAction string

const (
	ActionSwitchUser MainMenuAction = "switch"  // 切换用户
	ActionManageUser MainMenuAction = "manage"  // 管理用户
	ActionChangeLang MainMenuAction = "lang"    // 切换语言
	ActionRefresh    MainMenuAction = "refresh" // 刷新信息
	ActionExit       MainMenuAction = "exit"    // 退出
)

// ManageAction 用户管理操作类型
type ManageAction string

const (
	ManageAdd    ManageAction = "add"    // 添加用户
	ManageEdit   ManageAction = "edit"   // 编辑用户
	ManageDelete ManageAction = "delete" // 删除用户
	ManageBack   ManageAction = "back"   // 返回
)

// ShowMainMenu 显示主菜单，返回用户选择的操作
// 按 ESC 或 Ctrl+C 可退出
func ShowMainMenu() (MainMenuAction, error) {
	t := i18n.T()
	var action MainMenuAction

	// 自定义 KeyMap，添加 ESC 键到 Quit 绑定
	keymap := huh.NewDefaultKeyMap()
	keymap.Quit = key.NewBinding(key.WithKeys("esc", "ctrl+c"))

	// 创建主菜单表单
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[MainMenuAction]().
				Title(t.MenuTitle).
				Options(
					huh.NewOption(t.MenuSwitchUser, ActionSwitchUser),
					huh.NewOption(t.MenuManageUser, ActionManageUser),
					huh.NewOption(t.MenuChangeLang, ActionChangeLang),
					huh.NewOption(t.MenuRefresh, ActionRefresh),
					huh.NewOption(t.MenuExit, ActionExit),
				).
				Value(&action),
		),
	).WithTheme(huh.ThemeCatppuccin()).
		WithKeyMap(keymap)

	if err := form.Run(); err != nil {
		return ActionExit, err
	}

	return action, nil
}

// ShowLanguageSelect 显示语言选择菜单
func ShowLanguageSelect() (i18n.Language, error) {
	t := i18n.T()
	var selected i18n.Language

	// 构建语言选项
	options := make([]huh.Option[i18n.Language], len(i18n.AvailableLanguages))
	for idx, lang := range i18n.AvailableLanguages {
		options[idx] = huh.NewOption(lang.Name, lang.Code)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[i18n.Language]().
				Title(t.ChangeLangTitle).
				Options(options...).
				Value(&selected),
		),
	).WithTheme(huh.ThemeCatppuccin())

	if err := form.Run(); err != nil {
		return i18n.GetLang(), err
	}

	return selected, nil
}

// ShowUserSelect 显示用户选择列表
// 返回选择的用户索引，如果取消则返回 -1
func ShowUserSelect(users []config.User) (int, error) {
	t := i18n.T()
	if len(users) == 0 {
		return -1, nil
	}

	// 构建选项列表
	options := make([]huh.Option[int], len(users)+1)
	for i, user := range users {
		label := fmt.Sprintf("%s <%s>", user.Name, user.Email)
		options[i] = huh.NewOption(label, i)
	}
	// 添加返回选项
	options[len(users)] = huh.NewOption(t.Back, -1)

	var selected int
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title(t.UserSelectSwitch).
				Options(options...).
				Value(&selected),
		),
	).WithTheme(huh.ThemeCatppuccin())

	if err := form.Run(); err != nil {
		return -1, err
	}

	return selected, nil
}

// ShowManageMenu 显示用户管理菜单
func ShowManageMenu() (ManageAction, error) {
	t := i18n.T()
	var action ManageAction

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[ManageAction]().
				Title(t.UserManage).
				Options(
					huh.NewOption(t.UserAdd, ManageAdd),
					huh.NewOption(t.UserEdit, ManageEdit),
					huh.NewOption(t.UserDelete, ManageDelete),
					huh.NewOption(t.Back, ManageBack),
				).
				Value(&action),
		),
	).WithTheme(huh.ThemeCatppuccin())

	if err := form.Run(); err != nil {
		return ManageBack, err
	}

	return action, nil
}

// ShowAddUserForm 显示添加用户表单
// 返回用户输入的 name, email, signingKey 和 githubName
func ShowAddUserForm() (string, string, string, string, error) {
	t := i18n.T()
	var name, email, signingKey, githubName string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(t.UserName).
				Placeholder(t.UserNameExample).
				Value(&name).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf(t.UserNameEmpty)
					}
					return nil
				}),
			huh.NewInput().
				Title(t.UserEmail).
				Placeholder(t.UserEmailExample).
				Value(&email).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf(t.UserEmailEmpty)
					}
					return nil
				}),
			huh.NewInput().
				Title(t.UserSigningKey).
				Placeholder(t.UserSigningKeyEx).
				Value(&signingKey),
			huh.NewInput().
				Title(t.UserGithubName).
				Placeholder(t.UserGithubNameEx).
				Value(&githubName),
		),
	).WithTheme(huh.ThemeCatppuccin())

	if err := form.Run(); err != nil {
		return "", "", "", "", err
	}

	return name, email, signingKey, githubName, nil
}

// ShowEditUserForm 显示编辑用户表单
// 接收当前的 name, email, signingKey 和 githubName 作为默认值
func ShowEditUserForm(currentName, currentEmail, currentSigningKey, currentGithubName string) (string, string, string, string, error) {
	t := i18n.T()
	name := currentName
	email := currentEmail
	signingKey := currentSigningKey
	githubName := currentGithubName

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(t.UserName).
				Value(&name).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf(t.UserNameEmpty)
					}
					return nil
				}),
			huh.NewInput().
				Title(t.UserEmail).
				Value(&email).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf(t.UserEmailEmpty)
					}
					return nil
				}),
			huh.NewInput().
				Title(t.UserSigningKey).
				Placeholder(t.UserSigningKeyEx).
				Value(&signingKey),
			huh.NewInput().
				Title(t.UserGithubName).
				Placeholder(t.UserGithubNameEx).
				Value(&githubName),
		),
	).WithTheme(huh.ThemeCatppuccin())

	if err := form.Run(); err != nil {
		return "", "", "", "", err
	}

	return name, email, signingKey, githubName, nil
}

// ShowDeleteConfirm 显示删除确认对话框
func ShowDeleteConfirm(user config.User) (bool, error) {
	t := i18n.T()
	var confirm bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(fmt.Sprintf(t.UserConfirmDelete, user.Name, user.Email)).
				Affirmative(t.Confirm).
				Negative(t.Cancel).
				Value(&confirm),
		),
	).WithTheme(huh.ThemeCatppuccin())

	if err := form.Run(); err != nil {
		return false, err
	}

	return confirm, nil
}

// ShowUserSelectForEdit 显示用户选择列表（用于编辑/删除）
func ShowUserSelectForEdit(users []config.User, title string) (int, error) {
	t := i18n.T()
	if len(users) == 0 {
		return -1, nil
	}

	// 构建选项列表
	options := make([]huh.Option[int], len(users)+1)
	for i, user := range users {
		label := fmt.Sprintf("%s <%s>", user.Name, user.Email)
		options[i] = huh.NewOption(label, i)
	}
	options[len(users)] = huh.NewOption(t.Back, -1)

	var selected int
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title(title).
				Options(options...).
				Value(&selected),
		),
	).WithTheme(huh.ThemeCatppuccin())

	if err := form.Run(); err != nil {
		return -1, err
	}

	return selected, nil
}

// ShowConfirmAddUser 询问用户是否立即添加预设用户
func ShowConfirmAddUser() (bool, error) {
	t := i18n.T()
	var confirm bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(t.UserAddNow).
				Affirmative(t.Yes).
				Negative(t.No).
				Value(&confirm),
		),
	).WithTheme(huh.ThemeCatppuccin())

	if err := form.Run(); err != nil {
		return false, err
	}

	return confirm, nil
}
