package core

import (
	"go-boilerplate/app/core/base"
	"go-boilerplate/app/core/client"
	"go.uber.org/fx"
)

type Modules struct {
	fx.In

	Base       Base
	Client     Client
	Repository Repository
}

type Base struct {
	fx.In

	Parameter base.Parameter
}

var BaseModule = fx.Module(
	"base",
	fx.Provide(
		base.NewGetParameter,
	),
)

type Client struct {
	fx.In

	PurchaseClient client.PurchaseClient
}

var ClientModule = fx.Module(
	"client",
	fx.Provide(
		client.NewPurchaseClient,
	),
)

type Repository struct {
	fx.In

	//Repo               repository.Repository
}

var RepositoryModule = fx.Module(
	"repository",
	fx.Provide(
	//repository.NewRepository,
	),
)
