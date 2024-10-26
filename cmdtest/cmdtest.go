// Package cmdtest creates a testable interface to rclone main
//
// The interface is used to perform end-to-end test of
// commands, flags, environment variables etc.
package cmdtest

// The rest of this file is a 1:1 copy from rclone.go

import (
	_ "github.com/tamankuc/rclone_ui/backend/all" // import all backends
	"github.com/tamankuc/rclone_ui/cmd"
	_ "github.com/tamankuc/rclone_ui/cmd/all"    // import all commands
	_ "github.com/tamankuc/rclone_ui/lib/plugin" // import plugins
)

func main() {
	cmd.Main()
}
