package database

import (
	"technical-test-skyshi/app/database/seeder"
	"technical-test-skyshi/config"
	"testing"
)

func TestMigrateAndSeedDB(t *testing.T) {
	c := config.LoadConfigFromEnv()
	t.Run("Test Migrate DB", func(t *testing.T) {
		Migrate(NewDB(c))
	})
	t.Run("Test Seed DB", func(t *testing.T) {
		seeder.SeedDB(NewDB(c))
	})
}
