// Test filefabric filesystem interface
package filefabric_test

import (
	"testing"

	"github.com/tamankuc/rclone_ui/backend/filefabric"
	"github.com/tamankuc/rclone_ui/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestFileFabric:",
		NilObject:  (*filefabric.Object)(nil),
	})
}
