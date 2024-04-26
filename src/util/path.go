package util

import (
	"path/filepath"
	"runtime"
)

func GetFilePath() string {
	// Caller(0) 返回当前调用者的文件名和行号信息。
	// 0 代表当前函数，1 代表上一级调用者，以此类推。
	_, filename, _, ok := runtime.Caller(1) // 注意这里使用1获取调用者的信息
	if !ok {
		panic("获取文件所在目录失败")
	}
	return filepath.Dir(filename)
}

func GetRootPath() string {
	return filepath.Dir(filepath.Dir(GetFilePath()))
}
