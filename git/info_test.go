package git

import (
	"os"
	"os/exec"
	"testing"
)

// TestSetRemoteURL 测试设置远程仓库地址功能
// 验证在没有 origin 远程仓库和已有 origin 远程仓库时都能成功设置
func TestSetRemoteURL(t *testing.T) {
	// 1. 创建临时测试目录
	tempDir, err := os.MkdirTemp("", "gitq_test_*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 获取当前工作目录，以便测试结束后恢复
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("获取当前工作目录失败: %v", err)
	}
	defer func() {
		_ = os.Chdir(originalWd)
	}()

	// 切换到临时目录
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("切换到临时目录失败: %v", err)
	}

	// 2. 初始化 Git 仓库
	initCmd := exec.Command("git", "init")
	if err := initCmd.Run(); err != nil {
		t.Fatalf("初始化 Git 仓库失败: %v", err)
	}

	// 配置临时的 git 用户，防止某些环境下 git 命令报错
	_ = exec.Command("git", "config", "user.name", "Test User").Run()
	_ = exec.Command("git", "config", "user.email", "test@example.com").Run()

	// 3. 验证当前没有 origin
	_, err = GetRemoteURL()
	if err == nil {
		t.Fatal("预期 GetRemoteURL 在无 origin 时返回错误，但实际没有")
	}

	// 4. 测试添加新的远程地址 (此时 origin 不存在)
	testURL1 := "git@github.com:qzrzz/GitQ.git"
	if err := SetRemoteURL(testURL1); err != nil {
		t.Fatalf("在 origin 不存在时 SetRemoteURL 失败: %v", err)
	}

	// 5. 验证是否成功添加
	url, err := GetRemoteURL()
	if err != nil {
		t.Fatalf("获取远程仓库地址失败: %v", err)
	}
	if url != testURL1 {
		t.Errorf("预期远程地址为 %s, 实际为 %s", testURL1, url)
	}

	// 6. 测试更新已存在的远程地址 (此时 origin 已经存在)
	testURL2 := "https://github.com/qzrzz/GitQ.git"
	if err := SetRemoteURL(testURL2); err != nil {
		t.Fatalf("在 origin 已存在时 SetRemoteURL 失败: %v", err)
	}

	// 7. 验证是否成功更新
	url, err = GetRemoteURL()
	if err != nil {
		t.Fatalf("获取远程仓库地址失败: %v", err)
	}
	if url != testURL2 {
		t.Errorf("预期更新后的远程地址为 %s, 实际为 %s", testURL2, url)
	}
}
