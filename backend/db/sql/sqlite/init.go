package sqlite

import (
	"fmt"
	"github.com/DnsUnlock/Dpanel/backend/config"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

func Conn() (dB *gorm.DB, err error) {
	if err = isValidPath(config.Config.Sql.Connection); err != nil {
		return
	}
	//判断 config.Config.Sql.Connection 是否为正常的文件路径
	if _, err := os.Stat(config.Config.Sql.Connection); os.IsNotExist(err) {
		// 文件不存在，创建空的数据库文件
		file, err := os.Create(config.Config.Sql.Connection)
		if err != nil {
			fmt.Printf("无法创建数据库文件: %v\n", err)
			return nil, err
		}
		file.Close()
		fmt.Println("数据库文件创建成功")
	}
	dB, err = gorm.Open(
		sqlite.Open(config.Config.Sql.Connection),
		&gorm.Config{},
	)
	return
}

// 检查路径是否有效
func isValidPath(path string) error {
	// 获取路径的绝对值
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("路径无法解析为绝对路径: %w", err)
	}

	// 检查父目录是否存在，如果不存在则创建
	dir := filepath.Dir(absPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 创建所有必需的父目录
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("无法创建目录: %s, 错误: %w", dir, err)
		}
		fmt.Printf("已创建目录: %s\n", dir)
	}

	// 如果目录不可写，返回错误
	if !isWritable(dir) {
		return fmt.Errorf("路径的父目录不可写: %s", dir)
	}

	return nil
}

// 检查目录是否可写
func isWritable(path string) bool {
	testFile := filepath.Join(path, ".test_write")
	err := os.WriteFile(testFile, []byte{}, 0644)
	if err != nil {
		return false
	}
	_ = os.Remove(testFile) // 删除测试文件
	return true
}
