// GitQ - Git 用户切换命令行工具
// 使用 charmbracelet/huh 和 lipgloss 构建美观的终端界面
// 主要功能：显示 Git 仓库信息、切换和管理 Git 用户
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"gitq/config"
	"gitq/git"
	"gitq/i18n"
	"gitq/ui"
)

func main() {
	// 初始化 i18n，检测系统语言
	i18n.Init()
	t := i18n.T()

	// 检查当前目录是否为 Git 仓库
	if !git.IsGitRepository() {
		fmt.Println(ui.RenderError(t.ErrNotGitRepo))
		fmt.Println(t.ErrRunInGitRepo)
		os.Exit(1)
	}

	// 主循环
	for {
		// 清屏后显示 Git 信息
		clearScreen()
		showGitInfo()

		// 显示主菜单
		action, err := ui.ShowMainMenu()
		if err != nil {
			// 用户按 ESC 或 Ctrl+C 退出
			fmt.Println("\n" + t.Goodbye)
			return
		}

		// 处理用户选择
		switch action {
		case ui.ActionSwitchUser:
			handleSwitchUser()
		case ui.ActionManageUser:
			handleManageUser()
		case ui.ActionChangeLang:
			handleChangeLang()
			// 切换语言后需要重新获取翻译
			t = i18n.T()
		case ui.ActionRefresh:
			// 刷新信息，直接继续循环
			continue
		case ui.ActionExit:
			fmt.Println("\n" + t.Goodbye)
			return
		}
	}
}

// clearScreen 清除终端屏幕
// 根据操作系统选择不同的清屏命令
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// showGitInfo 显示当前 Git 仓库信息
func showGitInfo() {
	info := git.GetGitInfo()
	fmt.Println(ui.RenderGitInfo(info))
}

// handleSwitchUser 处理切换用户操作
func handleSwitchUser() {
	t := i18n.T()

	// 加载预设用户列表
	users, err := config.LoadUsers()
	if err != nil {
		fmt.Println(ui.RenderError(t.ErrLoadUsers + ": " + err.Error()))
		return
	}

	if len(users) == 0 {
		// 没有预设用户时，询问是否立即添加
		fmt.Println(ui.RenderWarning(t.UserNoPreset))
		if confirmAdd, err := ui.ShowConfirmAddUser(); err == nil && confirmAdd {
			handleAddUser()
		}
		return
	}

	// 显示用户选择
	index, err := ui.ShowUserSelect(users)
	if err != nil {
		return
	}

	// 用户选择返回
	if index < 0 {
		return
	}

	// 切换用户
	selectedUser := users[index]
	if err := git.SetUser(selectedUser.Name, selectedUser.Email, selectedUser.SigningKey, selectedUser.GithubName); err != nil {
		fmt.Println(ui.RenderError(t.ErrSwitchFailed + ": " + err.Error()))
		return
	}

	fmt.Println(ui.RenderSuccess(fmt.Sprintf(t.SuccessSwitched, selectedUser.Name, selectedUser.Email)))
}

// handleChangeLang 处理语言切换操作
func handleChangeLang() {
	lang, err := ui.ShowLanguageSelect()
	if err != nil {
		return
	}
	i18n.SetLang(lang)
}

// handleManageUser 处理用户管理操作
func handleManageUser() {
	for {
		action, err := ui.ShowManageMenu()
		if err != nil {
			return
		}

		switch action {
		case ui.ManageAdd:
			handleAddUser()
		case ui.ManageEdit:
			handleEditUser()
		case ui.ManageDelete:
			handleDeleteUser()
		case ui.ManageBack:
			return
		}
	}
}

// handleAddUser 处理添加用户
func handleAddUser() {
	t := i18n.T()

	name, email, signingKey, githubName, err := ui.ShowAddUserForm()
	if err != nil {
		return
	}

	if err := config.AddUser(name, email, signingKey, githubName); err != nil {
		fmt.Println(ui.RenderError(t.ErrAddFailed + ": " + err.Error()))
		return
	}

	fmt.Println(ui.RenderSuccess(fmt.Sprintf(t.SuccessAdded, name, email)))
}

// handleEditUser 处理编辑用户
func handleEditUser() {
	t := i18n.T()

	users, err := config.LoadUsers()
	if err != nil {
		fmt.Println(ui.RenderError(t.ErrLoadUsers + ": " + err.Error()))
		return
	}

	if len(users) == 0 {
		fmt.Println(ui.RenderWarning(t.UserNoEdit))
		return
	}

	// 选择要编辑的用户
	index, err := ui.ShowUserSelectForEdit(users, t.UserSelectEdit)
	if err != nil || index < 0 {
		return
	}

	// 显示编辑表单
	oldUser := users[index]
	name, email, signingKey, githubName, err := ui.ShowEditUserForm(oldUser.Name, oldUser.Email, oldUser.SigningKey, oldUser.GithubName)
	if err != nil {
		return
	}

	// 保存修改
	if err := config.UpdateUser(index, name, email, signingKey, githubName); err != nil {
		fmt.Println(ui.RenderError(t.ErrUpdateFailed + ": " + err.Error()))
		return
	}

	fmt.Println(ui.RenderSuccess(fmt.Sprintf(t.SuccessUpdated, name, email)))
}

// handleDeleteUser 处理删除用户
func handleDeleteUser() {
	t := i18n.T()

	users, err := config.LoadUsers()
	if err != nil {
		fmt.Println(ui.RenderError(t.ErrLoadUsers + ": " + err.Error()))
		return
	}

	if len(users) == 0 {
		fmt.Println(ui.RenderWarning(t.UserNoDelete))
		return
	}

	// 选择要删除的用户
	index, err := ui.ShowUserSelectForEdit(users, t.UserSelectDelete)
	if err != nil || index < 0 {
		return
	}

	// 确认删除
	confirm, err := ui.ShowDeleteConfirm(users[index])
	if err != nil || !confirm {
		return
	}

	// 执行删除
	if err := config.DeleteUser(index); err != nil {
		fmt.Println(ui.RenderError(t.ErrDeleteFailed + ": " + err.Error()))
		return
	}

	fmt.Println(ui.RenderSuccess(t.SuccessDeleted))
}
