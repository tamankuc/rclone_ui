// Test PikPak filesystem interface
package pikpak_test

import (
	"testing"

	"github.com/tamankuc/rclone_ui/backend/pikpak"
	"github.com/tamankuc/rclone_ui/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestPikPak:",
		NilObject:  (*pikpak.Object)(nil),
	})
}
