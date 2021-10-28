// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	app0xdd1446 "github.com/bitwormhole/bpm/app"
	cli0xf30272 "github.com/bitwormhole/starter-cli/cli"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComMainLoop struct {
	instance *app0xdd1446.MainLoop
	 markup0x23084a.Component `class:"looper"`
	ClientFactory cli0xf30272.ClientFactory `inject:"#cli-client-factory"`
	Context application0x67f6c5.Context `inject:"context"`
}

