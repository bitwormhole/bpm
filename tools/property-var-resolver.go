package tools

import "github.com/bitwormhole/starter/collection"

// ResolveConfig 解析配置中的变量
func ResolveConfig(props collection.Properties) error {
	const (
		t1 = "${"
		t2 = "}"
	)
	return collection.ResolvePropertiesVarWithTokens(t1, t2, props)
}
