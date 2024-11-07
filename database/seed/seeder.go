package seed

import (
	"acme-manager/ent"
	"context"
)

type Seeder interface {
	Seed() error
}

func Seeders(ctx context.Context, client ent.Client) []Seeder {
	return []Seeder{
		AcmeServerSeeder{ctx: ctx, client: client},
	}
}
