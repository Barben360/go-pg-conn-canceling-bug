package go_pg_canceling_bug

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Foo struct {
	Id       int `pg:",pk"`
	Comments string
}

func (f *Foo) Insert(ctx context.Context, db *pg.DB) error {
	return db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := tx.Model(f).Insert()
		return err
	})
}

func (f *Foo) Update(ctx context.Context, db *pg.DB) error {
	return db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := tx.Model(f).WherePK().Update()
		return err
	})
}

func (f *Foo) Delete(ctx context.Context, db *pg.DB) error {
	return db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := tx.Model((*Foo)(nil)).Where("id = ?", f.Id).Delete()
		return err
	})
}

func Init(poolSize int) (*pg.DB, error) {
	// Initializing DB client
	db := pg.Connect(&pg.Options{
		Network:               "",
		Addr:                  "localhost:5454",
		Dialer:                nil,
		OnConnect:             nil,
		User:                  "postgres-user",
		Password:              "postgres-password",
		Database:              "db",
		ApplicationName:       "",
		MaxRetries:            5,
		RetryStatementTimeout: true,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		PoolSize:              poolSize,
		MinIdleConns:          0,
		MaxConnAge:            0,
		PoolTimeout:           0,
		IdleTimeout:           0,
		IdleCheckFrequency:    0,
	})
	// Creating Foo table if not exists
	err := db.Model(&Foo{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
