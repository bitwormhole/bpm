package tools

import "github.com/bitwormhole/starter/collection"

// InjectEnvToProperties 把所有的环境变量注入到目标属性表中（以“env.”为前缀）
func InjectEnvToProperties(dst collection.Properties, src collection.Environment) {
	const prefix = "env."
	all := src.Export(nil)
	for name, value := range all {
		dst.SetProperty(prefix+name, value)
	}
}
