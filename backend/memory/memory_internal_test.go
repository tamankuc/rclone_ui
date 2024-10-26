package memory

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/tamankuc/rclone_ui/backend/local"
	"github.com/tamankuc/rclone_ui/fs/operations"
	"github.com/tamankuc/rclone_ui/fstest"
	"github.com/tamankuc/rclone_ui/fstest/fstests"
	"github.com/stretchr/testify/require"
)

var t1 = fstest.Time("2001-02-03T04:05:06.499999999Z")

// InternalTest dispatches all internal tests
func (f *Fs) InternalTest(t *testing.T) {
	t.Run("PurgeListDeadlock", func(t *testing.T) {
		testPurgeListDeadlock(t)
	})
}

// test that Purge fallback does not result in deadlock from concurrently listing and removing
func testPurgeListDeadlock(t *testing.T) {
	ctx := context.Background()
	r := fstest.NewRunIndividual(t)
	r.Mkdir(ctx, r.Fremote)
	r.Fremote.Features().Disable("Purge") // force fallback-purge

	// make a lot of files to prevent it from finishing too quickly
	for i := 0; i < 100; i++ {
		dst := "file" + fmt.Sprint(i) + ".txt"
		r.WriteObject(ctx, dst, "hello", t1)
	}

	require.NoError(t, operations.Purge(ctx, r.Fremote, ""))
}

var _ fstests.InternalTester = (*Fs)(nil)
