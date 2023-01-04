package main

import (
	"bufio"
	"bytes"
	"gek_exec"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type appsList []string

// getAppsFromFile 从文件中读取APP列表
func getAppsFromFile(appFile string) (apps appsList, err error) {
	// 打开文件
	file, err := os.Open(appFile)
	if err != nil {
		return nil, err
	}
	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	// 按行读取文件中数据
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 跳过注释
		if strings.Contains(scanner.Text(), "//") {
			continue
		}

		// 截取每行 package: 后半的包名,并存储到apps切片中
		apps = append(apps, strings.ReplaceAll(scanner.Text(), "package:", ""))
	}
	return apps, nil
}

// getAppListFromADB 使用adb获取全部应用列表
func getAppListFromADB() (apps appsList, err error) {
	cmd := exec.Command("adb", "shell", "pm", "list", "packages")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	bytesReader := bytes.NewReader(output)
	// 按行读取数据
	scanner := bufio.NewScanner(bytesReader)
	for scanner.Scan() {
		// 截取每行 package: 后半的包名,并存储到apps切片中
		apps = append(apps, strings.ReplaceAll(scanner.Text(), "package:", ""))
	}
	return apps, nil
}

// ShowAppsList Options are:
//      -f: see their associated file
//      -a: all known packages (but excluding APEXes)
//      -d: filter to only show disabled packages
//      -e: filter to only show enabled packages
//      -s: filter to only show system packages
//      -3: filter to only show third party packages
//      -i: see the installer for the packages
//      -l: ignored (used for compatibility with older releases)
//      -U: also show the package UID
//      -u: also include uninstalled packages
//      --show-versioncode: also show the version code
//      --apex-only: only show APEX packages
//      --uid UID: filter to only show packages with the given UID
//      --user USER_ID: only list packages belonging to the given user
func showAppsList(options string) (output string, err error) {
	out, err := gek_exec.Output("adb shell pm list package " + options)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func (apps appsList) clear() {
	for _, app := range apps {
		cmd := exec.Command("adb", "shell", "pm", "clear", app)

		err := gek_exec.Run(cmd)
		if err != nil {
			log.Println(err)
		}
	}
}

func (apps appsList) uninstall() {
	for _, app := range apps {
		cmd := exec.Command("adb", "shell", "pm", "uninstall", app)

		err := gek_exec.Run(cmd)
		if err != nil {
			log.Println(err)
		}
	}
}

func (apps appsList) uninstallUser(uid int) {
	for _, app := range apps {
		cmd := exec.Command("adb", "shell", "pm", "uninstall", "--user", strconv.Itoa(uid), app)

		err := gek_exec.Run(cmd)
		if err != nil {
			log.Println(err)
		}
	}
}

func (apps appsList) reinstall() {
	for _, app := range apps {
		cmd := exec.Command("adb", "shell", "cmd", "package", "install-existing", app)

		err := gek_exec.Run(cmd)
		if err != nil {
			log.Println(err)
		}
	}
}

func (apps appsList) disable() {
	for _, app := range apps {
		cmd := exec.Command("adb", "shell", "pm", "disable", app)

		err := gek_exec.Run(cmd)
		if err != nil {
			log.Println(err)
		}
	}
}

func (apps appsList) disableUser() {
	for _, app := range apps {
		cmd := exec.Command("adb", "shell", "pm", "disable-user", app)

		err := gek_exec.Run(cmd)
		if err != nil {
			log.Println(err)
		}
	}
}

func (apps appsList) enable() {
	for _, app := range apps {
		cmd := exec.Command("adb", "shell", "pm", "enable", app)

		err := gek_exec.Run(cmd)
		if err != nil {
			log.Println(err)
		}
	}
}
