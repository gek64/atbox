package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	cliClear         bool
	cliUninstall     bool
	cliUninstallUser bool
	cliReinstall     bool
	cliDisable       bool
	cliDisableUser   bool
	cliEnable        bool
	cliAll           bool
	cliHelp          bool
	cliVersion       bool
)

func init() {
	flag.BoolVar(&cliClear, "clear", false, "pm clear")
	flag.BoolVar(&cliUninstall, "uninstall", false, "pm uninstall")
	flag.BoolVar(&cliUninstallUser, "uninstall-user", false, "pm uninstall --user 0")
	flag.BoolVar(&cliReinstall, "reinstall", false, "cmd package install-existing")
	flag.BoolVar(&cliDisable, "disable", false, "pm disable")
	flag.BoolVar(&cliDisableUser, "disable-user", false, "pm disable-user")
	flag.BoolVar(&cliEnable, "enable", false, "pm enable")
	flag.BoolVar(&cliAll, "all", false, "use pm list packages get all apps")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
  atbox [Command] [Options] [Apps_File]

Options:
  -all              : use pm list packages get all apps

Command:
  -clear            : pm clear
  -uninstall        : pm uninstall
  -uninstall-user   : pm uninstall --user 0
  -reinstall        : cmd package install-existing
  -disable          : pm disable
  -disable-user     : pm disable-user
  -enable           : pm enable
  -h                : show help
  -v                : show version

Example:
1) atbox -clear -all             : clear data from all apps
2) atbox -clear /root/apps.txt   : clear data from app that list in /root/apps.txt`
		fmt.Println(helpInfo)
	}

	// 如果无 args 或者 指定 h 参数,打印用法后退出
	if len(os.Args) == 1 || cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		fmt.Println("v1.00")
		os.Exit(0)
	}
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release`
	fmt.Println(versionInfo)
}

func main() {
	clear()
	uninstall()
	uninstallUser()
	reinstall()
	disable()
	disableUser()
	enable()
}
