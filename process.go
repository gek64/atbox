package main

import (
	"flag"
	"log"
)

func clear() {
	if cliClear {
		if cliAll {
			apps, err := getAppListFromADB()
			if err != nil {
				log.Panicln(err)
			}
			apps.clear()

		} else if flag.NArg() != 0 {
			apps, err := getAppsFromFile(flag.Args()[0])
			if err != nil {
				log.Panicln(err)
			}
			apps.clear()
		}
	}
}
func uninstall() {
	if cliUninstall {
		if cliAll {
			apps, err := getAppListFromADB()
			if err != nil {
				log.Panicln(err)
			}
			apps.uninstall()

		} else if flag.NArg() != 0 {
			apps, err := getAppsFromFile(flag.Args()[0])
			if err != nil {
				log.Panicln(err)
			}
			apps.uninstall()
		}
	}
}
func uninstallUser() {
	if cliUninstallUser {
		if cliAll {
			apps, err := getAppListFromADB()
			if err != nil {
				log.Panicln(err)
			}
			apps.uninstallUser(0)

		} else if flag.NArg() != 0 {
			apps, err := getAppsFromFile(flag.Args()[0])
			if err != nil {
				log.Panicln(err)
			}
			apps.uninstallUser(0)
		}
	}
}
func reinstall() {
	if cliReinstall {
		if cliAll {
			apps, err := getAppListFromADB()
			if err != nil {
				log.Panicln(err)
			}
			apps.reinstall()

		} else if flag.NArg() != 0 {
			apps, err := getAppsFromFile(flag.Args()[0])
			if err != nil {
				log.Panicln(err)
			}
			apps.reinstall()
		}
	}
}
func disable() {
	if cliDisable {
		if cliAll {
			apps, err := getAppListFromADB()
			if err != nil {
				log.Panicln(err)
			}
			apps.disable()

		} else if flag.NArg() != 0 {
			apps, err := getAppsFromFile(flag.Args()[0])
			if err != nil {
				log.Panicln(err)
			}
			apps.disable()
		}
	}
}
func disableUser() {
	if cliDisableUser {
		if cliAll {
			apps, err := getAppListFromADB()
			if err != nil {
				log.Panicln(err)
			}
			apps.disable()

		} else if flag.NArg() != 0 {
			apps, err := getAppsFromFile(flag.Args()[0])
			if err != nil {
				log.Panicln(err)
			}
			apps.disable()
		}
	}
}
func enable() {
	if cliEnable {
		if cliAll {
			apps, err := getAppListFromADB()
			if err != nil {
				log.Panicln(err)
			}
			apps.enable()

		} else if flag.NArg() != 0 {
			apps, err := getAppsFromFile(flag.Args()[0])
			if err != nil {
				log.Panicln(err)
			}
			apps.enable()
		}
	}
}
