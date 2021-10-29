package service

import (
	"context"

	"github.com/bitwormhole/bpm/data/convert"
	"github.com/bitwormhole/bpm/data/entity"
	"github.com/bitwormhole/bpm/data/po"
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/markup"
)

// PackInfoService ...
type PackInfoService interface {
	DisplayPackInfo(ctx context.Context, in *vo.PackInfo, out *vo.PackInfo) error
	DisplayPackInfoByNames(ctx context.Context, names []string) error
}

////////////////////////////////////////////////////////////////////////////////

// PackInfoServiceImpl ...
type PackInfoServiceImpl struct {
	markup.Component `id:"bpm-pack-info-service" class:"bpm-service"`

	PM PackageManager `inject:"#bpm-package-manager"`
}

func (inst *PackInfoServiceImpl) _Impl() PackInfoService {
	return inst
}

func (inst *PackInfoServiceImpl) DisplayPackInfo(ctx context.Context, in *vo.PackInfo, out *vo.PackInfo) error {
	return inst.DisplayPackInfoByNames(ctx, in.PackageNames)
}

func (inst *PackInfoServiceImpl) DisplayPackInfoByNames(ctx context.Context, names []string) error {

	console, err := cli.GetConsole(ctx)
	if err != nil {
		return err
	}

	list1, err := inst.PM.SelectAvailablePackages(names)
	if err == nil {
		inst.displayAvailablePackInfo(console, list1)
	}

	list2, err := inst.PM.SelectInstalledPackages(names)
	if err == nil {
		inst.displayInstalledPackInfo(console, list2)
	}

	return nil
}

func (inst *PackInfoServiceImpl) displayInstalledPackInfo(console cli.Console, items []*entity.InstalledPackageInfo) {

	props := collection.CreateProperties()
	body := &po.InstalledPackages{}
	body.Packages = items
	convert.SaveInstalledPackages(body, props)
	text := collection.FormatProperties(props)
	console.WriteString(text + "\n")
}

func (inst *PackInfoServiceImpl) displayAvailablePackInfo(console cli.Console, items []*entity.AvailablePackageInfo) {

	props := collection.CreateProperties()
	body := &po.AvailablePackages{}
	body.Packages = items
	convert.SaveAvailablePackages(body, props)
	text := collection.FormatProperties(props)
	console.WriteString(text + "\n")
}
