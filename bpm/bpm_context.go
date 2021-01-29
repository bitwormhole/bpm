package bpm

import (
	"net/url"

	"github.com/bitwormhole/go-wormhole-core/io/fs"
)

// ManagerContext 表示管理器的上下文
type ManagerContext struct {
	BitwormholeHomeDir fs.Path
	LocalPackagesRepo  fs.Path
	RemotePackagesRepo *url.URL

	CommandHandlers map[string]CommandHandler
}

// PackageContext 表示一个软件包
type PackageContext struct {
	Manager            *ManagerContext
	Name               string // the package name
	Directory          fs.Path
	CurrentVersionFile fs.Path
}

// VersionContext 表示包的某个版本
type VersionContext struct {
	Package         *PackageContext
	Version         string
	Directory       fs.Path
	PackageJSONFile fs.Path
}

// CommandContext 表示一个命令上下文
type CommandContext struct {
	Command string // the name of command
	Handler CommandHandler

	Manager *ManagerContext
	Package *PackageContext
	Version *VersionContext
}
