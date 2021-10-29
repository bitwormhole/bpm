package service

import (
	"crypto/sha256"
	"errors"
	"strings"

	"github.com/bitwormhole/bpm/data/convert"
	"github.com/bitwormhole/bpm/data/entity"
	"github.com/bitwormhole/bpm/data/po"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
)

// PackageManager 包管理器
type PackageManager interface {
	LoadInstalledPackages() (*po.InstalledPackages, error)
	LoadPackageSourceList() (*po.PackageSourceList, error)
	LoadAvailablePackages() (*po.AvailablePackages, error)

	SaveInstalledPackages(p *po.InstalledPackages) error
	SavePackageSourceList(p *po.PackageSourceList) error
	SaveAvailablePackages(p *po.AvailablePackages) error

	SelectAvailablePackages(namelist []string) ([]*entity.AvailablePackageInfo, error)
	SelectInstalledPackages(namelist []string) ([]*entity.InstalledPackageInfo, error)

	ComputeSHA256sum(file fs.Path) (string, error)
}

////////////////////////////////////////////////////////////////////////////////

// PackageManagerImpl ...
type PackageManagerImpl struct {
	markup.Component `id:"bpm-package-manager"`

	Env EnvService `inject:"#bpm-env-service"`
}

func (inst *PackageManagerImpl) _Impl() PackageManager {
	return inst
}

// LoadInstalledPackages ...
func (inst *PackageManagerImpl) LoadInstalledPackages() (*po.InstalledPackages, error) {
	file := inst.Env.GetInstalledFile()
	props, err := inst.loadProperties(file)
	if err != nil {
		return nil, err
	}
	o := &po.InstalledPackages{}
	err = convert.LoadInstalledPackages(o, props)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// LoadPackageSourceList ...
func (inst *PackageManagerImpl) LoadPackageSourceList() (*po.PackageSourceList, error) {
	file := inst.Env.GetSourcesFile()
	props, err := inst.loadProperties(file)
	if err != nil {
		return nil, err
	}
	o := &po.PackageSourceList{}
	err = convert.LoadPackageSourceList(o, props)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// LoadAvailablePackages ...
func (inst *PackageManagerImpl) LoadAvailablePackages() (*po.AvailablePackages, error) {
	file := inst.Env.GetAvailableFile()
	props, err := inst.loadProperties(file)
	if err != nil {
		return nil, err
	}
	o := &po.AvailablePackages{}
	err = convert.LoadAvailablePackages(o, props)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// SaveInstalledPackages ...
func (inst *PackageManagerImpl) SaveInstalledPackages(o *po.InstalledPackages) error {
	file := inst.Env.GetInstalledFile()
	props := collection.CreateProperties()
	err := convert.SaveInstalledPackages(o, props)
	if err != nil {
		return err
	}
	return inst.saveProperties(props, file)
}

// SavePackageSourceList ...
func (inst *PackageManagerImpl) SavePackageSourceList(o *po.PackageSourceList) error {
	file := inst.Env.GetSourcesFile()
	props := collection.CreateProperties()
	err := convert.SavePackageSourceList(o, props)
	if err != nil {
		return err
	}
	return inst.saveProperties(props, file)
}

// SaveAvailablePackages ...
func (inst *PackageManagerImpl) SaveAvailablePackages(o *po.AvailablePackages) error {
	file := inst.Env.GetAvailableFile()
	props := collection.CreateProperties()
	err := convert.SaveAvailablePackages(o, props)
	if err != nil {
		return err
	}
	return inst.saveProperties(props, file)
}

func (inst *PackageManagerImpl) saveProperties(p collection.Properties, file fs.Path) error {
	text := collection.FormatPropertiesWithSegment(p)
	return file.GetIO().WriteText(text, nil, true)
}

func (inst *PackageManagerImpl) loadProperties(file fs.Path) (collection.Properties, error) {
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	return collection.ParseProperties(text, nil)
}

// ComputeSHA256sum ...
func (inst *PackageManagerImpl) ComputeSHA256sum(file fs.Path) (string, error) {
	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	str := util.StringifyBytes(sum[:])
	return str, nil
}

// SelectAvailablePackages ...
func (inst *PackageManagerImpl) SelectAvailablePackages(namelist []string) ([]*entity.AvailablePackageInfo, error) {

	// 生成名单查找表
	todolist := make(map[string]bool)
	donelist := make(map[string]bool)
	for _, name := range namelist {
		todolist[name] = true
		donelist[name] = false
	}

	// 加载可用包列表
	packs, err := inst.LoadAvailablePackages()
	if err != nil {
		return nil, err
	}
	all := packs.Packages

	// 遍历筛选
	results := make([]*entity.AvailablePackageInfo, 0)
	for _, pack := range all {
		name := pack.Name
		if todolist[name] {
			results = append(results, pack)
			donelist[name] = true
		}
	}

	// 检查遗漏
	errCount := 0
	errMaker := strings.Builder{}
	errMaker.WriteString("cannot find these package(s):")
	for name, ok := range donelist {
		if !ok {
			errMaker.WriteString(name)
			errMaker.WriteString(",")
			errCount++
		}
	}
	if errCount > 0 {
		return nil, errors.New(errMaker.String())
	}

	return results, nil
}

// SelectInstalledPackages ...
func (inst *PackageManagerImpl) SelectInstalledPackages(namelist []string) ([]*entity.InstalledPackageInfo, error) {

	// load all
	src, err := inst.LoadInstalledPackages()
	if err != nil {
		return nil, err
	}

	// src >>> all
	all := make(map[string]*entity.InstalledPackageInfo) // map[name]pack
	for _, item := range src.Packages {
		all[item.Name] = item
	}

	// find in all
	dst := make([]*entity.InstalledPackageInfo, 0)
	for _, name := range namelist {
		pack := all[name]
		if pack != nil {
			dst = append(dst, pack)
		} else {
			return nil, errors.New("no installed package named:" + name)
		}
	}
	return dst, nil
}
