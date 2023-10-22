package save

import (
	"fmt"
	"os"

	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/query"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dsn = fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	gormdb, _ = gorm.Open(postgres.Open(dsn))

	Q = query.Use(gormdb)
)
