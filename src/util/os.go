package util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func AppExist(name string) bool {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("which %s", name))
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

func AssetExist(asset string) bool {
	_, err := os.Lstat(asset)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func GetSubDirectories(dir string) []string {
	subDirs, err := os.ReadDir(dir)
	if err != nil {
		panic(fmt.Sprintf("read dir err:%+v path:%s", err, dir))
	}
	var res []string
	for _, subDir := range subDirs {
		if subDir.IsDir() {
			res = append(res, subDir.Name())
		}
	}
	return res
}

func HomeDir() string {
	home := func() string {
		if runtime.GOOS == "windows" {
			return "USERPROFILE"
		}
		return "HOME"
	}()
	home = os.Getenv(home)
	return home
}
