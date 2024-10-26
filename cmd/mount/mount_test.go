//go:build linux

package mount

import (
	"testing"

	"github.com/tamankuc/rclone_ui/vfs/vfscommon"
	"github.com/tamankuc/rclone_ui/vfs/vfstest"
)

func TestMount(t *testing.T) {
	vfstest.RunTests(t, false, vfscommon.CacheModeOff, true, mount)
}
