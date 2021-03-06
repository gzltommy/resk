package gorpc

import (
	"github.com/gzl-tommy/infra"
	"github.com/gzl-tommy/infra/base"
)

type GoRpcApiStarter struct {
	infra.BaseStarter
}

func (g *GoRpcApiStarter) Init(ctx infra.StarterContext) {
	base.RpcRegister(new(EnvelopeRpc))
}