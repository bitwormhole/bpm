package service

import (
	"context"

	"github.com/bitwormhole/bpm/data/entity"
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/starter/markup"
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

	PM        PackageManager `inject:"#bpm-package-manager"`
	FetchSer  FetchService   `inject:"#bpm-fetch-service"`
	DeploySer DeployService  `inject:"#bpm-deploy-service"`
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
	list := []*entity.InstalledPackageInfo{pack}
	return inst.doUpgradeAll(ctx, list)
}

// UpgradePackages ...
func (inst *UpgradeServiceImpl) UpgradePackages(ctx context.Context, packs []*entity.InstalledPackageInfo) error {
	return inst.doUpgradeAll(ctx, packs)
}

// UpgradeByNames ...
func (inst *UpgradeServiceImpl) UpgradeByNames(ctx context.Context, names []string) error {
	packs, err := inst.PM.SelectInstalledPackages(names)
	if err != nil {
		return err
	}
	return inst.doUpgradeAll(ctx, packs)
}

func (inst *UpgradeServiceImpl) convertInstalledToAvailable(src []*entity.InstalledPackageInfo) []*entity.AvailablePackageInfo {
	namelist := make([]string, 0)
	for _, item1 := range src {
		namelist = append(namelist, item1.Name)
	}
	dst, err := inst.PM.SelectAvailablePackages(namelist)
	if err != nil {
		return []*entity.AvailablePackageInfo{}
	}
	return dst
}

func (inst *UpgradeServiceImpl) doUpgradeAll(ctx context.Context, packs []*entity.InstalledPackageInfo) error {
	all := inst.convertInstalledToAvailable(packs)
	err := inst.doFetchAll(ctx, all)
	if err != nil {
		return err
	}
	return inst.doDeployAll(ctx, all)
}

func (inst *UpgradeServiceImpl) doFetchAll(ctx context.Context, packs []*entity.AvailablePackageInfo) error {
	return inst.FetchSer.FetchPackages(ctx, packs)
}

func (inst *UpgradeServiceImpl) doDeployAll(ctx context.Context, packs []*entity.AvailablePackageInfo) error {
	return inst.DeploySer.DeployPackages(ctx, packs, inst)
}

func (inst *UpgradeServiceImpl) AcceptDeploy(prev *entity.InstalledPackageInfo, next *entity.AvailablePackageInfo) bool {
	if prev == nil || next == nil {
		return false // 如果还没有安装，那么就不能升级
	}
	if prev.Revision == next.Revision {
		return false // 如果已经是最新的，那么就不需要升级
	}
	return true
}
