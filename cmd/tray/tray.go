// cmd/tray/tray.go
package tray

import (
	"context"
	"log"
	"os"

	"github.com/tamankuc/rclone_ui/cmd"
	"github.com/tamankuc/rclone_ui/fs"
	"github.com/tamankuc/rclone_ui/fs/rc"
	"github.com/tamankuc/rclone_ui/fs/rc/rcserver"
	"github.com/spf13/cobra"
	"github.com/tamankuc/rclone_ui/lib/tray" // Обновленная ссылка на пакет
)

func init() {
	cmd.Root.AddCommand(commandDefinition)
}

var commandDefinition = &cobra.Command{
	Use:   "tray",
	Short: "Run Rclone in system tray mode with remote control capabilities",
	Long: `This command runs Rclone in system tray mode. It allows you to control Rclone 
through a system tray icon and receive notifications on sync events. The tray 
integration includes options for starting and stopping sync tasks directly from the tray menu.`,
	Run: func(command *cobra.Command, args []string) {
		// Запуск RC-сервера для управления через API
		ctx := context.Background()
		rc.Opt.Enabled = true

		s, err := rcserver.Start(ctx, &rc.Opt)
		if err != nil {
			fs.Fatalf(nil, "Failed to start RC server: %v", err)
		}
		if s == nil {
			fs.Fatal(nil, "RC server not configured")
		}

		// Запуск трея
		log.Println("Starting Rclone in tray mode...")
		go func() {
			if err := tray.Run(); err != nil {
				log.Fatalf("Tray error: %v", err)
			}
		}()

		// Завершаем RC-сервер при выходе
		defer func() {
			s.Wait()
			os.Exit(0)
		}()
	},
}
