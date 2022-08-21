package database

import (
	"github.com/ssibrahimbas/go-uuid/config"
)

type Databases struct {
	Postgres *Postgres

	cnf *config.Config
}

func New(cnf *config.Config) *Databases {
	pg := NewPostgres(cnf)
	return &Databases{
		cnf:      cnf,
		Postgres: pg,
	}
}

func (d *Databases) ConnectAll() {
	d.Postgres.Connect()
}
