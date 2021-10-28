package service

import (
	"strings"

	"github.com/bitwormhole/bpm/data/entity"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

// LocalBpmFiles 包含缓存在本地的BPM包相关文件
type LocalBpmFiles struct {
	Dir       fs.Path
	BPM       fs.Path
	MetaJSON  fs.Path
	Manifest  fs.Path
	Signature fs.Path
}

// EnvService 提供环境上下文
type EnvService interface {
	GetBitwormholeHome() fs.Path

	// "etc/bpm/sources"
	GetSourcesFile() fs.Path

	// "etc/bpm/available"
	GetAvailableFile() fs.Path

	// "etc/bpm/installed"
	GetInstalledFile() fs.Path

	// "files"
	GetFilesFolder() fs.Path

	// "var/bpm/packages/pkgName/version/packages-version.bpm"
	GetLocalBpmFiles(p *entity.BasePackageInfo) *LocalBpmFiles
}

////////////////////////////////////////////////////////////////////////////////

// EnvServiceImpl 实现 EnvService
type EnvServiceImpl struct {
	markup.Component `id:"bpm-env-service" initMethod:"Init"`

	Context application.Context `inject:"context"`
	home    fs.Path
}

func (inst *EnvServiceImpl) _Impl() EnvService {
	return inst
}

// Init 初始化服务
func (inst *EnvServiceImpl) Init() error {
	home, err := inst.Context.GetEnvironment().GetEnv("BITWORMHOLE_HOME")
	if err != nil {
		return err
	}
	inst.home = fs.Default().GetPath(home)
	return nil
}

// GetBitwormholeHome ...
func (inst *EnvServiceImpl) GetBitwormholeHome() fs.Path {
	return inst.home
}

// GetSourcesFile ...
func (inst *EnvServiceImpl) GetSourcesFile() fs.Path {
	return inst.GetBitwormholeHome().GetChild("etc/bpm/sources")
}

// GetAvailableFile ...
func (inst *EnvServiceImpl) GetAvailableFile() fs.Path {
	return inst.GetBitwormholeHome().GetChild("etc/bpm/available")
}

// GetInstalledFile ...
func (inst *EnvServiceImpl) GetInstalledFile() fs.Path {
	return inst.GetBitwormholeHome().GetChild("etc/bpm/installed")
}

// GetFilesFolder ...
func (inst *EnvServiceImpl) GetFilesFolder() fs.Path {
	return inst.GetBitwormholeHome().GetChild("files")
}

// GetLocalBpmFiles ...
func (inst *EnvServiceImpl) GetLocalBpmFiles(p *entity.BasePackageInfo) *LocalBpmFiles {

	name := p.Name
	version := p.Version
	sum := p.SHA256

	path := strings.Builder{}
	path.WriteString("var/cache/bpm/packages/")
	path.WriteString(name)
	path.WriteString("/")
	path.WriteString(version)
	path.WriteString("/")
	path.WriteString(sum)
	dir := inst.GetBitwormholeHome().GetChild(path.String())

	path.Reset()
	path.WriteString(name)
	path.WriteString("-")
	path.WriteString(version)
	path.WriteString(".bpm")
	bpmFileName := path.String()

	result := &LocalBpmFiles{}
	result.Dir = dir
	result.BPM = dir.GetChild(bpmFileName)
	result.MetaJSON = dir.GetChild(bpmFileName + ".json")
	result.Manifest = dir.GetChild(bpmFileName + ".manifest")
	result.Signature = dir.GetChild(bpmFileName + ".signature")
	return result
}
