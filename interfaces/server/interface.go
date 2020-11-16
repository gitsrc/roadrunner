package server

import (
	"context"
	"os/exec"

	"github.com/spiral/roadrunner/v2"
)

type Env map[string]string

// WorkerFactory creates workers for the application.
type WorkerFactory interface {
	CmdFactory(env Env) (func() *exec.Cmd, error)
	NewWorker(ctx context.Context, env Env) (roadrunner.WorkerBase, error)
	NewWorkerPool(ctx context.Context, opt roadrunner.PoolConfig, env Env) (roadrunner.Pool, error)
}
