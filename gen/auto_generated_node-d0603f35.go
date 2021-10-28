// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	command0xf0f741 "github.com/bitwormhole/bpm/command"
	service0xa5f732 "github.com/bitwormhole/bpm/service"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComBpmFetch struct {
	instance *command0xf0f741.BpmFetch
	 markup0x23084a.Component `class:"cli-handler"`
	Service service0xa5f732.FetchService `inject:"#bpm-fetch-service"`
}


type pComBpmHelp struct {
	instance *command0xf0f741.BpmHelp
	 markup0x23084a.Component `class:"cli-handler"`
}


type pComBpmInstall struct {
	instance *command0xf0f741.BpmInstall
	 markup0x23084a.Component `class:"cli-handler"`
	Service service0xa5f732.InstallService `inject:"#bpm-install-service"`
}


type pComBpmRun struct {
	instance *command0xf0f741.BpmRun
	 markup0x23084a.Component `class:"cli-handler"`
	Service service0xa5f732.RunService `inject:"#bpm-run-service"`
}


type pComBpmUpdate struct {
	instance *command0xf0f741.BpmUpdate
	 markup0x23084a.Component `class:"cli-handler"`
	Service service0xa5f732.UpdateService `inject:"#bpm-update-service"`
}


type pComBpmUpgrade struct {
	instance *command0xf0f741.BpmUpgrade
	 markup0x23084a.Component `class:"cli-handler"`
	Service service0xa5f732.UpgradeService `inject:"#bpm-upgrade-service"`
}


type pComBpm struct {
	instance *command0xf0f741.Bpm
	 markup0x23084a.Component `class:"cli-handler"`
}

