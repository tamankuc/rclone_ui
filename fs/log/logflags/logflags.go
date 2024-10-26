// Package logflags implements command line flags to set up the log
package logflags

import (
	"github.com/tamankuc/rclone_ui/fs/config/flags"
	"github.com/tamankuc/rclone_ui/fs/log"
	"github.com/spf13/pflag"
)

// AddFlags adds the log flags to the flagSet
func AddFlags(flagSet *pflag.FlagSet) {
	flags.AddFlagsFromOptions(flagSet, "", log.OptionsInfo)
}
