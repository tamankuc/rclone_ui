// Package rcflags implements command line flags to set up the remote control
package rcflags

import (
	"github.com/tamankuc/rclone_ui/fs/config/flags"
	"github.com/tamankuc/rclone_ui/fs/rc"
	"github.com/spf13/pflag"
)

// FlagPrefix is the prefix used to uniquely identify command line flags.
const FlagPrefix = "rc-"

// AddFlags adds the remote control flags to the flagSet
func AddFlags(flagSet *pflag.FlagSet) {
	flags.AddFlagsFromOptions(flagSet, "", rc.OptionsInfo)
}
