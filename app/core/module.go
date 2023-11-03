package core

import (
	"go.uber.org/fx"
)

type Repository struct {
	fx.In

	// core repository
	//Repo repository.Repository
}

var RepositoryModule = fx.Module(
	"repository",
	fx.Provide(
	// core repository
	//repository.NewRepository,
	),
)
