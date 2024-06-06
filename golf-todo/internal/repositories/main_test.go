package repositories

import (
	"testing"

	"github.com/mariusfa/golf/database"
)

var dbConfig database.DbConfig

func TestMain(m *testing.M) {
	dbConfig = database.DbConfig{
		User:        "test",
		Password:    "test",
		Name:        "test",
		AppUser:     "app_user",
		AppPassword: "app_password",
	}
	database.SinglePostgresTestMain(m, &dbConfig, "../../migrations")
}
