// Test Opendrive filesystem interface
package opendrive_test

import (
	"testing"

	"github.com/tamankuc/rclone_ui/backend/opendrive"
	"github.com/tamankuc/rclone_ui/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestOpenDrive:",
		NilObject:  (*opendrive.Object)(nil),
	})
}
