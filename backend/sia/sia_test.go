// Test Sia filesystem interface
package sia_test

import (
	"testing"

	"github.com/tamankuc/rclone_ui/backend/sia"

	"github.com/tamankuc/rclone_ui/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSia:",
		NilObject:  (*sia.Object)(nil),
	})
}
