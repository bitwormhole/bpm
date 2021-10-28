package po

import "github.com/bitwormhole/bpm/data/entity"

// PackageSourceList 软件包源列表
type PackageSourceList struct {
	Sources []*entity.PackSource
}
