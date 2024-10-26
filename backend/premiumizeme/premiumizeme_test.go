// Test filesystem interface
package premiumizeme_test

import (
	"testing"

	"github.com/tamankuc/rclone_ui/backend/premiumizeme"
	"github.com/tamankuc/rclone_ui/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestPremiumizeMe:",
		NilObject:  (*premiumizeme.Object)(nil),
	})
}
