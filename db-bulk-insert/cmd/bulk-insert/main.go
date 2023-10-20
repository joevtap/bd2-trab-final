package main

import (
	"fmt"
	"os"

	"github.com/joevtap/bd2-trab-final/db-bulk-insert/internal/insert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/model"
	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/query"
)

const (
	BASE_URL = "https://spcultura.prefeitura.sp.gov.br/api"
)

func main() {

	resources := []insert.Resource{
		{
			Name: "event",
			Call: insert.Call{
				BaseUrl: BASE_URL,
				Route:   "/event/find",
			},
		},
		{
			Name: "agent",
			Call: insert.Call{
				BaseUrl: BASE_URL,
				Route:   "/agent/find",
			},
		},
		{
			Name: "space",
			Call: insert.Call{
				BaseUrl: BASE_URL,
				Route:   "/space/find",
			},
		},
		{
			Name: "project",
			Call: insert.Call{
				BaseUrl: BASE_URL,
				Route:   "/project/find",
			},
		},
	}

	wch := make(chan struct{})

	for i := range resources {
		go func(i int, resources []insert.Resource) {
			err := resources[i].SetCount()

			if err != nil {
				panic(err)
			}

			wch <- struct{}{}
		}(i, resources)

		<-wch
	}

	for _, resource := range resources {
		fmt.Printf("%+v\n", resource)
	}

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

	err := query.Use(gormdb).Event.Create(&model.Event{
		Name:                "Teste",
		ShortDescription:    "asasd",
		ClassificacaoEtaria: 18,
	})

	if err != nil {
		panic(err)
	}

	events, err := query.Use(gormdb).Event.Find()

	if err != nil {
		panic(err)
	}

	for _, e := range events {
		fmt.Printf("%+v\n", e)
	}
}
