package service

import "context"

// AssemblyLineService ...
type AssemblyLineService interface {
	Assembly(ctx context.Context, targetNames []string) error
}
