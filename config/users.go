// Package config 提供 GitQ 用户配置管理功能
// 用于管理预设的 Git 用户列表，存储在 ~/.gitq/users.json 文件中
package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// User 表示一个 Git 用户配置
type User struct {
	Name       string `json:"name"`                 // 用户名
	Email      string `json:"email"`                // 邮箱
	SigningKey string `json:"signingKey,omitempty"` // GPG 签名密钥 (可选)
	GithubName string `json:"githubName,omitempty"` // GitHub 用户名 (可选，用于 SSH 密钥切换)
}

// 配置文件相关常量
const (
	configDirName  = ".gitq"      // 配置目录名
	configFileName = "users.json" // 用户配置文件名
)

// getConfigPath 获取配置文件的完整路径
// 返回 ~/.gitq/users.json 的绝对路径
func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configDirName, configFileName), nil
}

// ensureConfigDir 确保配置目录存在
// 如果 ~/.gitq 目录不存在，则创建它
func ensureConfigDir() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configDir := filepath.Join(homeDir, configDirName)
	return os.MkdirAll(configDir, 0755)
}

// LoadUsers 从配置文件加载用户列表
// 如果配置文件不存在，返回空列表而不是错误
func LoadUsers() ([]User, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		// 文件不存在时返回空列表
		if os.IsNotExist(err) {
			return []User{}, nil
		}
		return nil, err
	}

	// 解析 JSON
	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// SaveUsers 将用户列表保存到配置文件
// 会自动创建配置目录（如果不存在）
func SaveUsers(users []User) error {
	// 确保配置目录存在
	if err := ensureConfigDir(); err != nil {
		return err
	}

	// 获取配置文件路径
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	// 序列化为 JSON（带缩进，方便阅读）
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	// 写入文件
	return os.WriteFile(configPath, data, 0644)
}

// AddUser 添加新用户到配置
func AddUser(name, email, signingKey, githubName string) error {
	users, err := LoadUsers()
	if err != nil {
		return err
	}

	// 添加新用户
	users = append(users, User{Name: name, Email: email, SigningKey: signingKey, GithubName: githubName})
	return SaveUsers(users)
}

// DeleteUser 根据索引删除用户
// index 为用户在列表中的位置（从 0 开始）
func DeleteUser(index int) error {
	users, err := LoadUsers()
	if err != nil {
		return err
	}

	// 检查索引是否有效
	if index < 0 || index >= len(users) {
		return nil // 索引无效时静默忽略
	}

	// 删除指定位置的用户
	users = append(users[:index], users[index+1:]...)
	return SaveUsers(users)
}

// UpdateUser 更新指定索引的用户信息
func UpdateUser(index int, name, email, signingKey, githubName string) error {
	users, err := LoadUsers()
	if err != nil {
		return err
	}

	// 检查索引是否有效
	if index < 0 || index >= len(users) {
		return nil
	}

	// 更新用户信息
	users[index] = User{Name: name, Email: email, SigningKey: signingKey, GithubName: githubName}
	return SaveUsers(users)
}
