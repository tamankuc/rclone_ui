// Sync files and directories to and from local and remote object stores
//
// Nick Craig-Wood <nick@craig-wood.com>
package main

import (
	_ "github.com/tamankuc/rclone_ui/backend/all" // import all backends
	"github.com/tamankuc/rclone_ui/cmd"
	_ "github.com/tamankuc/rclone_ui/cmd/all"    // import all commands
	_ "github.com/tamankuc/rclone_ui/lib/plugin" // import plugins
)

func main() {
	cmd.Main()
}
