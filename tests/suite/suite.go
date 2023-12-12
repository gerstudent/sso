package suite

import (
	"context"
	ssov1 "github.com/gerstudent/protos/gen/go/sso"
	"github.com/gerstudent/sso/internal/config"
	"testing"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient ssov1.AuthClient
}

const configPath

func (s *Suite) New() {
	t.Helper()
	t.Parallel

	cfg := config.MustLoadPath(configPath)
	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)
	t.Cleanup() {
		t.Helper()
		cancelCtx()
	}
}
