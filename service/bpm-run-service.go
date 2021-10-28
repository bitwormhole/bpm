package service

import (
	"context"

	"github.com/bitwormhole/bpm/data/vo"
	"github.com/bitwormhole/starter/markup"
)

type RunService interface {
	Run(ctx context.Context, in *vo.Run, out *vo.Run) error
}

type RunServiceImpl struct {
	markup.Component `id:"bpm-run-service" class:"bpm-service"`
}

func (inst *RunServiceImpl) _Impl() RunService {
	return inst
}

func (inst *RunServiceImpl) Run(ctx context.Context, in *vo.Run, out *vo.Run) error {

	return nil
}
