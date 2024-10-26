// cmd/tray.go
package cmd

import (
	"flag"
	"log"
	"os"

	// "github.com/tamankuc/rclone_ui/cmd/tray"
	"github.com/tamankuc/rclone_ui/fs"
	"github.com/tamankuc/rclone_ui/tray"
)

// trayFlag определяет, включен ли режим системного трея
var trayFlag = flag.Bool("tray", false, "Enable system tray mode")

// init инициализирует флаг для системного трея и добавляет его в обработку команд
func init() {
	flag.Parse()
	if *trayFlag {
		go startTrayMode()
	}
}

// startTrayMode запускает модуль трея
func startTrayMode() {
	if err := tray.Run(); err != nil {
		fs.Errorf(nil, "Failed to start tray mode: %v", err)
		os.Exit(1)
	}
	log.Println("Rclone tray mode started")
}
