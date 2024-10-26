// Run the more functional vfstest package on the vfs

package vfs_test

import (
	"testing"

	_ "github.com/tamankuc/rclone_ui/backend/all" // import all the backends
	"github.com/tamankuc/rclone_ui/cmd/mountlib"
	"github.com/tamankuc/rclone_ui/fstest"
	"github.com/tamankuc/rclone_ui/vfs"
	"github.com/tamankuc/rclone_ui/vfs/vfscommon"
	"github.com/tamankuc/rclone_ui/vfs/vfstest"
)

// TestFunctional runs more functional tests all the tests against all the
// VFS cache modes
func TestFunctional(t *testing.T) {
	if *fstest.RemoteName != "" {
		t.Skip("Skip on non local")
	}
	vfstest.RunTests(t, true, vfscommon.CacheModeOff, true, func(VFS *vfs.VFS, mountpoint string, opt *mountlib.Options) (unmountResult <-chan error, unmount func() error, err error) {
		unmountResultChan := make(chan (error), 1)
		unmount = func() error {
			unmountResultChan <- nil
			return nil
		}
		return unmountResultChan, unmount, nil
	})
}
