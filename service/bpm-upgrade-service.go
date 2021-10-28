package service

import (
	"context"

	"github.com/bitwormhole/bpm/data/entity"
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// UpgradeService 升级已安装的包
type UpgradeService interface {
	Upgrade(ctx context.Context, in *vo.Upgrade, out *vo.Upgrade) error
	UpgradePackage(ctx context.Context, pack *entity.InstalledPackageInfo) error
	UpgradePackages(ctx context.Context, packs []*entity.InstalledPackageInfo) error
	UpgradeByNames(ctx context.Context, names []string) error
}

// UpgradeServiceImpl 实现 UpgradeService
type UpgradeServiceImpl struct {
	markup.Component `id:"bpm-upgrade-service" class:"bpm-service"`

	PM     PackageManager `inject:"#bpm-package-manager"`
	Fetch  FetchService   `inject:"#bpm-fetch-service"`
	Deploy DeployService  `inject:"#bpm-deploy-service"`
}

func (inst *UpgradeServiceImpl) _Impl() UpgradeService {
	return inst
}

// Upgrade ...
func (inst *UpgradeServiceImpl) Upgrade(ctx context.Context, in *vo.Upgrade, out *vo.Upgrade) error {
	return inst.UpgradeByNames(ctx, in.PackageNames)
}

// UpgradePackage ...
func (inst *UpgradeServiceImpl) UpgradePackage(ctx context.Context, pack *entity.InstalledPackageInfo) error {

	// todo : fetch

	// todo : deploy

	return inst.Deploy.DeployPackage(ctx, nil)
}

// UpgradePackages ...
func (inst *UpgradeServiceImpl) UpgradePackages(ctx context.Context, packs []*entity.InstalledPackageInfo) error {
	for _, pack := range packs {
		err := inst.UpgradePackage(ctx, pack)
		if err != nil {
			vlog.Warn(err)
		}
	}
	return nil
}

// UpgradeByNames ...
func (inst *UpgradeServiceImpl) UpgradeByNames(ctx context.Context, names []string) error {
	packs, err := inst.PM.SelectInstalledPackages(names)
	if err != nil {
		return err
	}
	return inst.UpgradePackages(ctx, packs)
}
