package database

import (
	"acme-manager/config"
	"acme-manager/database/seed"
	"acme-manager/ent"
	"acme-manager/logger"
	"acme-manager/util"
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var Client *ent.Client

func Connect() {
	var cfg = config.Get()
	logger.Infof("Connecting to postgres: %s", util.MaskDSN(cfg.Database.DSN))
	db, err := sql.Open("pgx", cfg.Database.DSN)
	if err != nil {
		logger.Fatalf("Failed opening connection to postgres: %v", err)
	}
	driver := entsql.OpenDB(dialect.Postgres, db)
	sqlDriver := dialect.DebugWithContext(driver, func(ctx context.Context, i ...any) {
		logger.Debugf("Ent: %v", i...)
	})
	Client = ent.NewClient(ent.Driver(sqlDriver))
}

func Migration() {
	logger.Info("Running schema migration")
	ctx := context.Background()
	err := Client.Schema.Create(ctx)
	if err != nil {
		logger.Fatalf("Failed creating schema resources: %v", err)
	}
}

func Seed() {
	logger.Info("Running seeders")
	ctx := context.Background()
	seeders := seed.Seeders(ctx, *Client)
	for _, seeder := range seeders {
		err := seeder.Seed()
		if err != nil {
			logger.Fatalf("Failed seeding: %v", err)
		}
	}
}

func Close() {
	logger.Info("Closing connection to postgres")
	err := Client.Close()
	if err != nil {
		logger.Fatalf("Failed closing connection to postgres: %v", err)
	}
}
