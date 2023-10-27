package main

import (
	"sync"

	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/model"
	"github.com/joevtap/bd2-trab-final/db-bulk-insert/internal/insert"
	"github.com/joevtap/bd2-trab-final/db-bulk-insert/internal/save"
)

const (
	BASE_URL = "https://spcultura.prefeitura.sp.gov.br/api"
)

func main() {
	resources := map[string]insert.Resource{
		"event": {
			Name: "event",
			Call: insert.Call{
				BaseUrl: BASE_URL,
				Route:   "/event/find",
				Params: insert.Params{
					"@select": "id,_type,name,shortDescription,longDescription,rules,createTimestamp.{date},status,occurrences,owner,project",
					"@order":  "id ASC",
				},
				Limit:    3000,
				CurrPage: 1,
			},
		},
		"agent": {
			Name: "agent",
			Call: insert.Call{
				BaseUrl: BASE_URL,
				Route:   "/agent/find",
				Params: insert.Params{
					"@select": "id,_type,name,shortDescription,longDescription,status,parent,_children,_spaces,_projects,_events",
					"@order":  "id ASC",
				},
				Limit:    5000,
				CurrPage: 1,
			},
		},
		"space": {
			Name: "space",
			Call: insert.Call{
				BaseUrl: BASE_URL,
				Route:   "/space/find",
				Params: insert.Params{
					"@select": "id,location,_geoLocation,name,public,shortDescription,longDescription,status,_type,eventOccurrences,parent,_children,owner",
					"@order":  "id ASC",
				},
				Limit:    5000,
				CurrPage: 1,
			},
		},
		"project": {
			Name: "project",
			Call: insert.Call{
				BaseUrl: BASE_URL,
				Route:   "/project/find",
				Params: insert.Params{
					"@select": "id,_type,name,shortDescription,longDescription,updateTimestamp,registrationFrom,registrationTo,createTimestamp,status,parent,_children,owner,_events",
					"@order":  "id ASC",
				},
				Limit:    5000,
				CurrPage: 1,
			},
		},
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go insert.SaveResource[model.Event](resources["event"], 1000, save.SaveEvent, wg)

	// wg.Add(1)
	// go insert.SaveResource[model.Agent](resources["agent"], 1000, save.SaveAgent, wg)

	// wg.Add(1)
	// go insert.SaveResource[model.Space](resources["space"], 1000, save.SaveSpace, wg)

	// wg.Add(1)
	// go insert.SaveResource[model.Project](resources["project"], 1000, save.SaveProject, wg)

	wg.Wait()
}
