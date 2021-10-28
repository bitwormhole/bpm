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

type UpdateService interface {
	Update(ctx context.Context, in *vo.Update, out *vo.Update) error
}

////////////////////////////////////////////////////////////////////////////////

type UpdateServiceImpl struct {
	markup.Component `id:"bpm-update-service" class:"bpm-service"`

	Remote HTTPGetService `inject:"#bpm-remote-service"`
	PM     PackageManager `inject:"#bpm-package-manager"`
}

func (inst *UpdateServiceImpl) _Impl() UpdateService {
	return inst
}

func (inst *UpdateServiceImpl) Update(ctx context.Context, in *vo.Update, out *vo.Update) error {

	// load source-list
	srclist, err := inst.PM.LoadPackageSourceList()
	if err != nil {
		return err
	}

	// fetch sources
	sources := srclist.Sources
	all := make(map[string]*entity.AvailablePackageInfo)
	for _, item := range sources {
		err := inst.fetchSource(ctx, item, all)
		if err != nil {
			return err
		}
	}

	// save packs
	packs := inst.makePackListFromTable(all)
	err = inst.PM.SaveAvailablePackages(packs)
	if err != nil {
		return err
	}
	return inst.PM.SaveAvailablePackages(packs)
}

func (inst *UpdateServiceImpl) fetchSourceByURL(url string) (*po.AvailablePackages, error) {

	text, err := inst.Remote.LoadText(url)
	if err != nil {
		return nil, err
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return nil, err
	}

	result := &po.AvailablePackages{}
	err = convert.LoadAvailablePackages(result, props)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (inst *UpdateServiceImpl) fetchSource(ctx context.Context, src *entity.PackSource, dst map[string]*entity.AvailablePackageInfo) error {

	console, err := cli.GetConsole(ctx)
	if err != nil {
		return err
	}

	url := src.URL
	console.WriteString("fetch " + url + " ... ")
	body, err := inst.fetchSourceByURL(url)
	if err != nil {
		console.WriteString("[fail]\n")
		return err
	}
	console.WriteString("[ok]\n")

	packs := body.Packages
	for _, pack := range packs {
		dst[pack.SHA256] = pack
	}
	return nil
}

func (inst *UpdateServiceImpl) makePackListFromTable(table map[string]*entity.AvailablePackageInfo) *po.AvailablePackages {
	body := &po.AvailablePackages{}
	list := make([]*entity.AvailablePackageInfo, 0)
	for _, item := range table {
		list = append(list, item)
	}
	body.Packages = list
	return body
}
