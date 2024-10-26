// Test Mega filesystem interface
package mega_test

import (
	"testing"

	"github.com/tamankuc/rclone_ui/backend/mega"
	"github.com/tamankuc/rclone_ui/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestMega:",
		NilObject:  (*mega.Object)(nil),
	})
}
