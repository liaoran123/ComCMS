package apply

import (
	"os"
	"path/filepath"
)

// 获取程序绝对路径目录
func GetCurrentAbPath() string {
	exePath, err := os.Executable()
	if err != nil {
		//log.Fatal(err)
		return ""
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res + "\\"
}
