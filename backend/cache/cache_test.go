// Test Cache filesystem interface

//go:build !plan9 && !js && !race

package cache_test

import (
	"testing"

	"github.com/tamankuc/rclone_ui/backend/cache"
	_ "github.com/tamankuc/rclone_ui/backend/local"
	"github.com/tamankuc/rclone_ui/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName:                      "TestCache:",
		NilObject:                       (*cache.Object)(nil),
		UnimplementableFsMethods:        []string{"PublicLink", "OpenWriterAt", "OpenChunkWriter", "DirSetModTime", "MkdirMetadata"},
		UnimplementableObjectMethods:    []string{"MimeType", "ID", "GetTier", "SetTier", "Metadata", "SetMetadata"},
		UnimplementableDirectoryMethods: []string{"Metadata", "SetMetadata", "SetModTime"},
		SkipInvalidUTF8:                 true, // invalid UTF-8 confuses the cache
	})
}
