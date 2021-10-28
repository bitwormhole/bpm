package po

import "github.com/bitwormhole/bpm/data/entity"

// InstalledPackages 已安装的包列表
type InstalledPackages struct {
	Packages []*entity.InstalledPackageInfo
}

// AvailablePackages 可安装的包列表
type AvailablePackages struct {
	Packages []*entity.AvailablePackageInfo
}
