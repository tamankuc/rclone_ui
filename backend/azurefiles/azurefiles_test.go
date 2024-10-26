//go:build !plan9 && !js

package azurefiles

import (
	"testing"

	"github.com/tamankuc/rclone_ui/fstest/fstests"
)

func TestIntegration(t *testing.T) {
	var objPtr *Object
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestAzureFiles:",
		NilObject:  objPtr,
	})
}
