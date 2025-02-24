//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/weplanx/collector/v2/app"
	"github.com/weplanx/collector/v2/common"
)

func NewApp() (*app.App, error) {
	wire.Build(
		wire.Struct(new(common.Inject), "*"),
		LoadStaticValues,
		UseElastic,
		UseNats,
		UseJetStream,
		UseKeyValue,
		app.Initialize,
	)
	return &app.App{}, nil
}
