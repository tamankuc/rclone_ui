// Test Uptobox filesystem interface
package uptobox_test

import (
	"testing"

	"github.com/tamankuc/rclone_ui/backend/uptobox"
	"github.com/tamankuc/rclone_ui/fstest"
	"github.com/tamankuc/rclone_ui/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	if *fstest.RemoteName == "" {
		*fstest.RemoteName = "TestUptobox:"
	}
	fstests.Run(t, &fstests.Opt{
		RemoteName: *fstest.RemoteName,
		NilObject:  (*uptobox.Object)(nil),
	})
}
