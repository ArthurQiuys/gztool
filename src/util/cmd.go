package util

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func MustCmd(cmdStr string) {
	fmt.Printf("%s\n", cmdStr)
	cmd := exec.Command("sh", "-c", cmdStr)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stdout
	err := cmd.Start()
	if err != nil {
		fmt.Printf("start command err:%+v\n", err)
		return
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("wait command err:%+v\n", err)
		return
	}
}

func MustCmdZshWithOutput(cmdStr string) string {
	fmt.Printf("%s\n", cmdStr)
	cmd := exec.Command("zsh", "-c", cmdStr)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	err := cmd.Start()
	if err != nil {
		fmt.Printf("start command err:%s\n", err)
		return ""
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("wait command err:%s\n", err)
		return ""
	}
	return buf.String()
}

func MustCmdSpawn(cmdStr string) {
	fmt.Printf("%s\n", cmdStr)
	cmd := exec.Command("sh", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("run cmd err:%+v\n", err)
		return
	}
}
