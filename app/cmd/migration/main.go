package main

import (
	"github.com/TadayoshiOtsuka/go_test_sample/app/individual/migration"
)

func main() {
	migration.NewMigrateHandler().Migrate()
}
