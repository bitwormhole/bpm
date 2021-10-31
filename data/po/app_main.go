package po

import "github.com/bitwormhole/bpm/data/entity"

// AppMain 是一个应用（可运行的包）的主配置
type AppMain struct {
	Main    entity.MainHead
	Scripts []*entity.MainScript
}
