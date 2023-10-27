package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/model"
	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/query"
	"github.com/joevtap/bd2-trab-final/db-bulk-insert/internal/insert"
	"github.com/joevtap/bd2-trab-final/db-bulk-insert/internal/save"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	BASE_URL = "https://spcultura.prefeitura.sp.gov.br/api"
)

var (
	dsn = fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		os.Getenv("DB_HOST"),
		"postgres",
		"postgres",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	gormdb, _ = gorm.Open(postgres.Open(dsn))

	Q = query.Use(gormdb)
)

func main() {
	eventResource := insert.Resource{
		Name: "event",
		Call: insert.Call{
			BaseUrl: BASE_URL,
			Route:   "/event/find",
			Params: insert.Params{
				"@select": "id,_type,name,shortDescription,longDescription,rules,createTimestamp.{date},status,occurrences,owner,project",
				"@order":  "id ASC",
			},
			Limit:    10000,
			CurrPage: 1,
		},
	}

	projectResource := insert.Resource{
		Name: "project",
		Call: insert.Call{
			BaseUrl: BASE_URL,
			Route:   "/project/find",
			Params: insert.Params{
				"@select": "id,_type,name,shortDescription,longDescription,updateTimestamp.{date},registrationFrom.{date},registrationTo.{date},createTimestamp.{date},status,parent,_children,owner,_events",
				"@order":  "id ASC",
			},
			Limit:    10000,
			CurrPage: 1,
		},
	}

	client := &insert.Client{}
	client.SetBaseUrl(eventResource.BaseUrl)

	eventResource.SetCount()

	res, err := client.Get(eventResource.Route, eventResource.Params, insert.Params{
		"@limit": strconv.Itoa(eventResource.Limit),
		"@page":  strconv.Itoa(1),
	})

	if err != nil {
		panic(err)
	}

	result := []model.Event{}

	if err := json.Unmarshal([]byte(res), &result); err != nil {
		panic(err)
	}

	for i := range result {
		result[i].Format()
	}

	eventChan := make(chan []model.Event, 1)

	eventChan <- result
	close(eventChan)

	save.SaveEvent(eventChan)

	res, err = client.Get(projectResource.Route, projectResource.Params, insert.Params{
		"@limit": strconv.Itoa(projectResource.Limit),
		"@page":  strconv.Itoa(1),
	})

	if err != nil {
		panic(err)
	}

	resultP := []model.Project{}

	if err := json.Unmarshal([]byte(res), &resultP); err != nil {
		panic(err)
	}

	for i := range resultP {
		resultP[i].Format()
	}

	projectChan := make(chan []model.Project, 1)

	projectChan <- resultP
	close(projectChan)

	save.SaveProject(projectChan)

	// wg.Add(1)
	// go insert.SaveResource[model.Agent](resources["agent"], 1000, save.SaveAgent, wg)

	// wg.Add(1)
	// go insert.SaveResource[model.Space](resources["space"], 1000, save.SaveSpace, wg)

	// wg.Add(1)
	// go insert.SaveResource[model.Project](resources["project"], 1000, save.SaveProject, wg)

	// gormdb.Exec("ALTER TABLE events ADD CONSTRAINT fk_events_project FOREIGN KEY (project) REFERENCES projects (id) ON DELETE SET NULL;")

}
