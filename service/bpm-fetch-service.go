package service

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/bitwormhole/bpm/data/entity"
	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/bpm/tools"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

// FetchService 下载bpm包到本地
type FetchService interface {
	Fetch(ctx context.Context, in *vo.Fetch, out *vo.Fetch) error

	FetchPackage(ctx context.Context, pack *entity.AvailablePackageInfo) error
	FetchPackages(ctx context.Context, packs []*entity.AvailablePackageInfo) error
	FetchByNames(ctx context.Context, names []string) error
}

////////////////////////////////////////////////////////////////////////////////

// FetchServiceImpl 实现 FetchService
type FetchServiceImpl struct {
	markup.Component `id:"bpm-fetch-service" class:"bpm-service"`

	PM     PackageManager `inject:"#bpm-package-manager"`
	Env    EnvService     `inject:"#bpm-env-service"`
	Remote HTTPGetService `inject:"#bpm-remote-service"`
}

func (inst *FetchServiceImpl) _Impl() FetchService {
	return inst
}

// Fetch ...
func (inst *FetchServiceImpl) Fetch(ctx context.Context, in *vo.Fetch, out *vo.Fetch) error {
	return inst.FetchByNames(ctx, in.PackageNames)
}

// FetchPackage ...
func (inst *FetchServiceImpl) FetchPackage(ctx context.Context, pack *entity.AvailablePackageInfo) error {

	console, err := cli.GetConsole(ctx)
	if err != nil {
		return err
	}

	task := fetchServiceTask{}
	task.context = ctx
	task.parent = inst
	task.console = console
	task.pack = pack

	return task.run()
}

// FetchPackages ...
func (inst *FetchServiceImpl) FetchPackages(ctx context.Context, packs []*entity.AvailablePackageInfo) error {
	for _, item := range packs {
		err := inst.FetchPackage(ctx, item)
		if err != nil {
			return err
		}
	}
	return nil
}

// FetchByNames ...
func (inst *FetchServiceImpl) FetchByNames(ctx context.Context, names []string) error {
	packs, err := inst.PM.SelectAvailablePackages(names)
	if err != nil {
		return err
	}
	return inst.FetchPackages(ctx, packs)
}

////////////////////////////////////////////////////////////////////////////////

type fetchServiceTask struct {
	context context.Context
	console cli.Console
	parent  *FetchServiceImpl
	pack    *entity.AvailablePackageInfo
}

func (inst *fetchServiceTask) run() error {
	return inst.fetchPack(inst.pack)
}

func (inst *fetchServiceTask) fetchPack(pack *entity.AvailablePackageInfo) error {

	parent := inst.parent
	url := pack.URL
	wantSum := pack.SHA256
	bpmFiles := parent.Env.GetLocalBpmFiles(&pack.BasePackageInfo)
	bpmFile := bpmFiles.BPM

	// check cached
	if inst.existsFile(bpmFile, wantSum) {
		inst.console.WriteString("cached package:" + pack.Name + "@" + pack.Version + "\n")
		return nil
	}

	// do fetch
	inst.console.WriteString("fetch " + url + "\n")
	err := parent.Remote.LoadFile(url, bpmFile)
	if err != nil {
		return err
	}

	// check sha256sum
	if !inst.existsFile(bpmFile, wantSum) {
		if bpmFile.Exists() {
			bpmFile.Delete()
		}
		return errors.New("bad sha-256 checksum, want: " + wantSum)
	}

	// check size
	wantSize := pack.Size
	haveSize := bpmFile.Size()
	if wantSize != haveSize {
		wantSizeStr := strconv.FormatInt(wantSize, 10)
		return errors.New("bad package size, want: " + wantSizeStr)
	}

	// save meta
	err = inst.savePackageMeta(pack, bpmFiles)
	if err != nil {
		return err
	}

	// ok, done
	return nil
}

func (inst *fetchServiceTask) savePackageMeta(pack *entity.AvailablePackageInfo, bpmFiles *LocalBpmFiles) error {
	jsonFile := bpmFiles.MetaJSON
	data, err := json.Marshal(pack)
	if err != nil {
		return err
	}
	return jsonFile.GetIO().WriteBinary(data, nil, false)
}

func (inst *fetchServiceTask) existsFile(file fs.Path, wantSum string) bool {
	if !file.Exists() {
		return false
	}
	haveSum, err := tools.ComputeSHA256sum(file)
	if err != nil {
		return false
	}
	return haveSum == wantSum
}
