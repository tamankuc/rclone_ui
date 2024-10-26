// lib/tray/tray.go
package tray

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/getlantern/systray"
)

// Run запускает системный трей
func Run() error {
	systray.Run(onReady, onExit)
	return nil
}

// onReady инициализирует элементы трея
func onReady() {
	systray.SetIcon([]byte{ /* icon data */ }) // Замените на иконку Rclone в формате []byte
	systray.SetTitle("Rclone")
	systray.SetTooltip("Manage Rclone tasks")

	// Элементы меню
	mStart := systray.AddMenuItem("Start Sync", "Start a synchronization task")
	mStop := systray.AddMenuItem("Stop Sync", "Stop the synchronization task")
	mQuit := systray.AddMenuItem("Quit", "Exit the application")

	// Запуск горутины для обработки кликов по элементам
	go func() {
		for {
			select {
			case <-mStart.ClickedCh:
				startSync()
			case <-mStop.ClickedCh:
				stopSync()
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

// onExit обрабатывает действия при выходе из приложения
func onExit() {
	// Закройте открытые ресурсы, если нужно
}

// startSync запускает синхронизацию через команду rclone sync
func startSync() {
	cmd := exec.Command("rclone", "sync", "source", "dest") // Замените на нужные параметры
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start sync: %v", err)
		showNotification("Rclone Error", "Failed to start synchronization")
		return
	}
	showNotification("Rclone", "Synchronization started")
}

// stopSync завершает процесс синхронизации через команду API или команду CLI
func stopSync() {
	// Можно использовать API, если Rclone запущен с rc
	url := "http://localhost:5572/core/quit"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Failed to stop sync: %v", err)
		showNotification("Rclone Error", "Failed to stop synchronization")
		return
	}
	showNotification("Rclone", "Synchronization stopped")
}

// showNotification показывает уведомление пользователю
func showNotification(title, message string) {
	switch runtime.GOOS {
	case "darwin":
		exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "%s"`, message, title)).Run()
	case "linux":
		exec.Command("notify-send", title, message).Run()
	case "windows":
		// Пример для Windows (необходимо go-toast)
		// toast := toast.Notification{AppID: "Rclone", Title: title, Message: message}
		// toast.Push()
	}
}
