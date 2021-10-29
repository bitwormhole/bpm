package entity

// ManifestItem 文件清单项目
type ManifestItem struct {
	ID     string
	Name   string // 文件名
	Path   string
	SHA256 string
	Size   int64
	IsDir  bool
}

// ManifestMeta 文件清单元数据（主要是包信息）
type ManifestMeta struct {
	BasePackageInfo
	MainPath string
}
