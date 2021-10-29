package po

import "github.com/bitwormhole/bpm/data/entity"

// Manifest BPM 包项目清单
type Manifest struct {
	Meta  entity.ManifestMeta
	Items []*entity.ManifestItem
}

// Signature BPM 包签名信息
type Signature struct {
	Info entity.SignatureInfo
}
