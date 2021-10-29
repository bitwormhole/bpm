package service

import (
	"context"

	"github.com/bitwormhole/bpm/data/entity"
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/starter/markup"
)

// InstallService ...
type InstallService interface {
	Install(ctx context.Context, in *vo.Install, out *vo.Install) error

	InstallPackage(ctx context.Context, pack *entity.AvailablePackageInfo) error
	InstallPackages(ctx context.Context, packs []*entity.AvailablePackageInfo) error
	InstallByNames(ctx context.Context, names []string) error
}

////////////////////////////////////////////////////////////////////////////////

// InstallServiceImpl ...
type InstallServiceImpl struct {
	markup.Component `id:"bpm-install-service" class:"bpm-service"`

	PM        PackageManager `inject:"#bpm-package-manager"`
	FetchSer  FetchService   `inject:"#bpm-fetch-service"`
	DeploySer DeployService  `inject:"#bpm-deploy-service"`
}

func (inst *InstallServiceImpl) _Impl() InstallService {
	return inst
}

// Install ...
func (inst *InstallServiceImpl) Install(ctx context.Context, in *vo.Install, out *vo.Install) error {
	return inst.InstallByNames(ctx, in.PackageNames)
}

// InstallPackage ...
func (inst *InstallServiceImpl) InstallPackage(ctx context.Context, pack *entity.AvailablePackageInfo) error {
	list := []*entity.AvailablePackageInfo{pack}
	return inst.doInstallAll(ctx, list)
}

// InstallPackages ...
func (inst *InstallServiceImpl) InstallPackages(ctx context.Context, packs []*entity.AvailablePackageInfo) error {
	return inst.doInstallAll(ctx, packs)
}

// InstallByNames ...
func (inst *InstallServiceImpl) InstallByNames(ctx context.Context, names []string) error {
	packs, err := inst.PM.SelectAvailablePackages(names)
	if err != nil {
		return err
	}
	return inst.doInstallAll(ctx, packs)
}

func (inst *InstallServiceImpl) doInstallAll(ctx context.Context, packs []*entity.AvailablePackageInfo) error {
	err := inst.doFetchAll(ctx, packs)
	if err != nil {
		return err
	}
	return inst.doDeployAll(ctx, packs)
}

func (inst *InstallServiceImpl) doFetchAll(ctx context.Context, packs []*entity.AvailablePackageInfo) error {
	return inst.FetchSer.FetchPackages(ctx, packs)
}

func (inst *InstallServiceImpl) doDeployAll(ctx context.Context, packs []*entity.AvailablePackageInfo) error {
	return inst.DeploySer.DeployPackages(ctx, packs, inst)
}

func (inst *InstallServiceImpl) AcceptDeploy(prev *entity.InstalledPackageInfo, next *entity.AvailablePackageInfo) bool {
	if prev != nil {
		return false // 如果已经安装，那么就不能重复安装
	}
	return true
}
