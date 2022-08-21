package database

import (
	"context"
	"time"

	pg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/ssibrahimbas/go-uuid/config"
	"github.com/ssibrahimbas/go-uuid/entity"
)

type Postgres struct {
	cnf *config.Config
	Db  *pg.DB
}

func NewPostgres(cnf *config.Config) *Postgres {
	return &Postgres{
		cnf: cnf,
		Db:  nil,
	}
}

func (p *Postgres) Connect() {
	time.Sleep(10 * time.Second)
	p.Db = pg.Connect(&pg.Options{
		User:     p.cnf.Pg.User,
		Password: p.cnf.Pg.Password,
		Addr:     p.cnf.Pg.Host + ":" + p.cnf.Pg.Port,
		Database: p.cnf.Pg.Database,
	})
	p.Db.Ping(context.Background())
	err := p.createSchema()
	if err != nil {
		panic(err)
	}
}

func (p *Postgres) createSchema() error {
	models := []interface{}{
		(*entity.User)(nil),
	}
	for _, model := range models {
		err := p.Db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
