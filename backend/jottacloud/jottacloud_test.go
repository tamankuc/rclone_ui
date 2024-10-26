// Test Box filesystem interface
package jottacloud_test

import (
	"testing"

	"github.com/tamankuc/rclone_ui/backend/jottacloud"
	"github.com/tamankuc/rclone_ui/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestJottacloud:",
		NilObject:  (*jottacloud.Object)(nil),
	})
}
