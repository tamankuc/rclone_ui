package netstorage_test

import (
	"testing"

	"github.com/tamankuc/rclone_ui/backend/netstorage"
	"github.com/tamankuc/rclone_ui/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestnStorage:",
		NilObject:  (*netstorage.Object)(nil),
	})
}
