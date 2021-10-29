package service

import (
	"context"

	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

// MakeService ...
type MakeService interface {
	Make(ctx context.Context, in *vo.Make, out *vo.Make) error
	MakePackage(ctx context.Context, pwd fs.Path) error
}

////////////////////////////////////////////////////////////////////////////////

// MakeServiceImpl ...
type MakeServiceImpl struct {
	markup.Component `id:"bpm-make-service" class:"bpm-service"`

	// Fetch  FetchService  `inject:"#bpm-fetch-service"`
	// Deploy DeployService `inject:"#bpm-deploy-service"`
}

func (inst *MakeServiceImpl) _Impl() MakeService {
	return inst
}

func (inst *MakeServiceImpl) Make(ctx context.Context, in *vo.Make, out *vo.Make) error {

	return nil
}

func (inst *MakeServiceImpl) MakePackage(ctx context.Context, pwd fs.Path) error {

	return nil
}
