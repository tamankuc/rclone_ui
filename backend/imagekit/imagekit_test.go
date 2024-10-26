package imagekit

import (
	"testing"

	"github.com/tamankuc/rclone_ui/fstest"
	"github.com/tamankuc/rclone_ui/fstest/fstests"
)

func TestIntegration(t *testing.T) {
	debug := true
	fstest.Verbose = &debug
	fstests.Run(t, &fstests.Opt{
		RemoteName:      "TestImageKit:",
		NilObject:       (*Object)(nil),
		SkipFsCheckWrap: true,
	})
}
