package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type User struct{}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./gorm/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	gormdb, _ := gorm.Open(postgres.Open(dsn))

	g.UseDB(gormdb)

	g.ApplyBasic(
		g.GenerateModel("events"),
	)

	g.Execute()
}
