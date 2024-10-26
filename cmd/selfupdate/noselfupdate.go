//go:build noselfupdate

package selfupdate

import (
	"github.com/tamankuc/rclone_ui/lib/buildinfo"
)

func init() {
	buildinfo.Tags = append(buildinfo.Tags, "noselfupdate")
}
