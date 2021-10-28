package entity

// BasePackageInfo 基本的包结构
type BasePackageInfo struct {
	ID string // PK

	Name         string // the package name
	URL          string
	Type         string // mime-type
	SHA256       string
	Version      string
	Dependencies string // 依赖的其它包名，以逗号“,”隔开
	Revision     int
	Size         int64
}

// AvailablePackageInfo 表示可安装的包
type AvailablePackageInfo struct {
	BasePackageInfo
}

// InstalledPackageInfo 表示已经安装的包
type InstalledPackageInfo struct {
	BasePackageInfo
	AutoUpgrade bool
	File        string // local-path
}
