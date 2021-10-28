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
