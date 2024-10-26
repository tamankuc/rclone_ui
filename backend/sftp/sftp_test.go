// Test Sftp filesystem interface

//go:build !plan9

package sftp_test

import (
	"testing"

	"github.com/tamankuc/rclone_ui/backend/sftp"
	"github.com/tamankuc/rclone_ui/fstest"
	"github.com/tamankuc/rclone_ui/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSFTPOpenssh:",
		NilObject:  (*sftp.Object)(nil),
	})
}

func TestIntegration2(t *testing.T) {
	if *fstest.RemoteName != "" {
		t.Skip("skipping as -remote is set")
	}
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSFTPRclone:",
		NilObject:  (*sftp.Object)(nil),
	})
}

func TestIntegration3(t *testing.T) {
	if *fstest.RemoteName != "" {
		t.Skip("skipping as -remote is set")
	}
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSFTPRcloneSSH:",
		NilObject:  (*sftp.Object)(nil),
	})
}
