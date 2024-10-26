// Package vfsflags implements command line flags to set up a vfs
package vfsflags

import (
	"github.com/tamankuc/rclone_ui/fs/config/flags"
	"github.com/tamankuc/rclone_ui/vfs/vfscommon"
	"github.com/spf13/pflag"
)

// AddFlags adds the non filing system specific flags to the command
func AddFlags(flagSet *pflag.FlagSet) {
	flags.AddFlagsFromOptions(flagSet, "", vfscommon.OptionsInfo)
}
