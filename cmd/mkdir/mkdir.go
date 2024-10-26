// Package mkdir provides the mkdir command.
package mkdir

import (
	"context"
	"strings"

	"github.com/tamankuc/rclone_ui/cmd"
	"github.com/tamankuc/rclone_ui/fs"
	"github.com/tamankuc/rclone_ui/fs/operations"
	"github.com/spf13/cobra"
)

func init() {
	cmd.Root.AddCommand(commandDefinition)
}

var commandDefinition = &cobra.Command{
	Use:   "mkdir remote:path",
	Short: `Make the path if it doesn't already exist.`,
	Annotations: map[string]string{
		"groups": "Important",
	},
	Run: func(command *cobra.Command, args []string) {
		cmd.CheckArgs(1, 1, command, args)
		fdst := cmd.NewFsDir(args)
		if !fdst.Features().CanHaveEmptyDirectories && strings.Contains(fdst.Root(), "/") {
			fs.Logf(fdst, "Warning: running mkdir on a remote which can't have empty directories does nothing")
		}
		cmd.Run(true, false, command, func() error {
			return operations.Mkdir(context.Background(), fdst, "")
		})
	},
}
