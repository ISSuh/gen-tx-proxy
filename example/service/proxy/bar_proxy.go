// source: /home/issuh/workspace/my_project/gen-tx-proxy/example/service/bar.go

package proxy

import (
	"database/sql"

	"context"

	entity "github.com/ISSuh/simple-gen-proxy/example/entity"

	service "github.com/ISSuh/simple-gen-proxy/example/service"
)

type BarProxyTransaction func() *sql.DB

type BarProxy struct {
	target service.Bar
	db     BarProxyTransaction
}

func NewBarProxy(target service.Bar, db BarProxyTransaction) *BarProxy {
	return &BarProxy{
		target: target,
		db:     db,
	}
}

func (p *BarProxy) CreateC(c context.Context, id int) (*entity.Foo, error) {

	db := p.db()
	tx, txErr := db.Begin()
	if txErr != nil {
		panic(txErr)
	}

	r0, err := p.target.CreateC(c, id)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			panic(txErr)
		}
	} else {
		if txErr := tx.Commit(); txErr != nil {
			panic(txErr)
		}
	}

	return r0, err

}

func (p *BarProxy) CreateD(c context.Context, id int) (*entity.Bar, error) {

	db := p.db()
	tx, txErr := db.Begin()
	if txErr != nil {
		panic(txErr)
	}

	r0, err := p.target.CreateD(c, id)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			panic(txErr)
		}
	} else {
		if txErr := tx.Commit(); txErr != nil {
			panic(txErr)
		}
	}

	return r0, err

}

func (p *BarProxy) CBarz(c context.Context, id int) (*entity.Foo, *entity.Bar, error) {

	return p.target.CBarz(c, id)

}

func (p *BarProxy) BFoos(c context.Context, a *entity.Foo, b *entity.Bar) error {

	db := p.db()
	tx, txErr := db.Begin()
	if txErr != nil {
		panic(txErr)
	}

	err := p.target.BFoos(c, a, b)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			panic(txErr)
		}
	} else {
		if txErr := tx.Commit(); txErr != nil {
			panic(txErr)
		}
	}

	return err

}

func (p *BarProxy) QFoosBarz(c context.Context) {

	p.target.QFoosBarz(c)

}
