package entity

// MainHead 主配置文件中的头部
type MainHead struct {
	Name        string // name of app
	Title       string // title  of app
	Description string // desc of app
	Package     string // name of package
	Script      string // the default script name
}

// MainScript 主配置文件中的脚本
type MainScript struct {
	Name             string // the name of script
	Arguments        string // the cli args
	Executable       string // path to exe
	WorkingDirectory string // path to WD
}
