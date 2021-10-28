package po

import "github.com/bitwormhole/bpm/data/entity"

// Manifest BPM 包项目清单
type Manifest struct {
	Items []*entity.ManifestItem
}
