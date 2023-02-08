package database

import (
	"technical-test-skyshi/app/database/seeder"
	"technical-test-skyshi/config"
	"testing"
)

func TestMigrateAndSeedDB(t *testing.T) {
	c := &config.Config{
		MySqlHost:     "localhost",
		MySqlPort:     "3306",
		MySqlUser:     "root",
		MySqlPassword: "password",
		MySqlDBName:   "technical_test_skyshi_database",
	}
	t.Run("Test Migrate DB", func(t *testing.T) {
		Migrate(NewDB(c))
	})
	t.Run("Test Seed DB", func(t *testing.T) {
		seeder.SeedDB(NewDB(c))
	})
}
