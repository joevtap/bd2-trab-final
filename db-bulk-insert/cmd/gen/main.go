package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

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
		g.GenerateModel("events", gen.FieldJSONTagWithNS(nameStrategy)),
		g.GenerateModel("agents", gen.FieldJSONTagWithNS(nameStrategy)),
		g.GenerateModel("spaces", gen.FieldJSONTagWithNS(nameStrategy)),
		g.GenerateModel("projects", gen.FieldJSONTagWithNS(nameStrategy)),
	)

	g.Execute()
}

func nameStrategy(columnName string) string {
	strParts := strings.Split(columnName, "_")

	tagContent := ""

	for _, part := range strParts {
		tagContent += cases.Title(language.English).String(part)
	}

	tagContent = cases.Lower(language.English).String(tagContent[:1]) + tagContent[1:]

	return tagContent
}
