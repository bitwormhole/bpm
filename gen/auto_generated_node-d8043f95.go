// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	service0xa5f732 "github.com/bitwormhole/bpm/service"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	fs0x8698bb "github.com/bitwormhole/starter/io/fs"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComDeployServiceImpl struct {
	instance *service0xa5f732.DeployServiceImpl
	 markup0x23084a.Component `id:"bpm-deploy-service"`
	PM service0xa5f732.PackageManager `inject:"#bpm-package-manager"`
	Env service0xa5f732.EnvService `inject:"#bpm-env-service"`
}


type pComFetchServiceImpl struct {
	instance *service0xa5f732.FetchServiceImpl
	 markup0x23084a.Component `id:"bpm-fetch-service" class:"bpm-service"`
	PM service0xa5f732.PackageManager `inject:"#bpm-package-manager"`
	Env service0xa5f732.EnvService `inject:"#bpm-env-service"`
	Remote service0xa5f732.HTTPGetService `inject:"#bpm-remote-service"`
}


type pComInstallServiceImpl struct {
	instance *service0xa5f732.InstallServiceImpl
	 markup0x23084a.Component `id:"bpm-install-service" class:"bpm-service"`
	Fetch service0xa5f732.FetchService `inject:"#bpm-fetch-service"`
	Deploy service0xa5f732.DeployService `inject:"#bpm-deploy-service"`
}


type pComRunServiceImpl struct {
	instance *service0xa5f732.RunServiceImpl
	 markup0x23084a.Component `id:"bpm-run-service" class:"bpm-service"`
}


type pComUpdateServiceImpl struct {
	instance *service0xa5f732.UpdateServiceImpl
	 markup0x23084a.Component `id:"bpm-update-service" class:"bpm-service"`
	Remote service0xa5f732.HTTPGetService `inject:"#bpm-remote-service"`
	PM service0xa5f732.PackageManager `inject:"#bpm-package-manager"`
}


type pComUpgradeServiceImpl struct {
	instance *service0xa5f732.UpgradeServiceImpl
	 markup0x23084a.Component `id:"bpm-upgrade-service" class:"bpm-service"`
}


type pComEnvServiceImpl struct {
	instance *service0xa5f732.EnvServiceImpl
	 markup0x23084a.Component `id:"bpm-env-service" initMethod:"Init"`
	Context application0x67f6c5.Context `inject:"context"`
	home fs0x8698bb.Path ``
}


type pComPackageManagerImpl struct {
	instance *service0xa5f732.PackageManagerImpl
	 markup0x23084a.Component `id:"bpm-package-manager"`
	Env service0xa5f732.EnvService `inject:"#bpm-env-service"`
}


type pComHTTPGetServiceImpl struct {
	instance *service0xa5f732.HTTPGetServiceImpl
	 markup0x23084a.Component `id:"bpm-remote-service"`
}

