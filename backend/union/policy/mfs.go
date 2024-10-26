package policy

import (
	"context"

	"github.com/tamankuc/rclone_ui/backend/union/upstream"
	"github.com/tamankuc/rclone_ui/fs"
)

func init() {
	registerPolicy("mfs", &Mfs{})
}

// Mfs stands for most free space
// Search category: same as epmfs.
// Action category: same as epmfs.
// Create category: Pick the drive with the most free space.
type Mfs struct {
	EpMfs
}

// Create category policy, governing the creation of files and directories
func (p *Mfs) Create(ctx context.Context, upstreams []*upstream.Fs, path string) ([]*upstream.Fs, error) {
	if len(upstreams) == 0 {
		return nil, fs.ErrorObjectNotFound
	}
	upstreams = filterNC(upstreams)
	if len(upstreams) == 0 {
		return nil, fs.ErrorPermissionDenied
	}
	u, err := p.mfs(upstreams)
	return []*upstream.Fs{u}, err
}
