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

	Fetch  FetchService  `inject:"#bpm-fetch-service"`
	Deploy DeployService `inject:"#bpm-deploy-service"`
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
	err := inst.Fetch.FetchPackage(ctx, pack)
	if err != nil {
		return err
	}
	return inst.Deploy.DeployPackage(ctx, pack)
}

// InstallPackages ...
func (inst *InstallServiceImpl) InstallPackages(ctx context.Context, packs []*entity.AvailablePackageInfo) error {
	err := inst.Fetch.FetchPackages(ctx, packs)
	if err != nil {
		return err
	}
	return inst.Deploy.DeployPackages(ctx, packs)
}

// InstallByNames ...
func (inst *InstallServiceImpl) InstallByNames(ctx context.Context, names []string) error {
	err := inst.Fetch.FetchByNames(ctx, names)
	if err != nil {
		return err
	}
	return inst.Deploy.DeployByNames(ctx, names)
}
