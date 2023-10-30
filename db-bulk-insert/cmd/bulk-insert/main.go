package main

import (
	"log"
	"sync"
	"time"

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
					"@select": "id,name,shortDescription,createTimestamp,updateTimestamp,owner,project,terms,classificacaoEtaria",
					"@order":  "id ASC",
				},
				Limit:    5000,
				CurrPage: 1,
			},
		},
		"agent": {
			Name: "agent",
			Call: insert.Call{
				BaseUrl: BASE_URL,
				Route:   "/agent/find",
				Params: insert.Params{
					"@select": "id,name,shortDescription,createTimestamp,updateTimestamp,parent,terms,children,spaces,events,projects",
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
					"@select": "id,location,name,shortDescription,createTimestamp,updateTimestamp,eventOccurrences,horario,telefonePublico,emailPublico,children,terms,parent,owner",
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
					"@select": "id,name,shortDescription,parent,owner,events,children,createTimestamp,updateTimestamp,registrationFrom,registrationTo",
					"@order":  "id ASC",
				},
				Limit:    5000,
				CurrPage: 1,
			},
		},
		"event_occurrence": {
			Name: "event_occurrence",
			Call: insert.Call{
				BaseUrl: BASE_URL,
				Route:   "/eventOccurrence/find",
				Params: insert.Params{
					"@select": "id,rule,frequency,separation,eventId,spaceId",
					"@order":  "id ASC",
				},
				Limit:    5000,
				CurrPage: 1,
			},
		},
	}

	wg := &sync.WaitGroup{}

	var sleep time.Duration = 1000

	wg.Add(1)
	go insert.SaveResource[*model.Event](resources["event"], sleep, wg)

	wg.Add(1)
	go insert.SaveResource[*model.Agent](resources["agent"], sleep, wg)

	wg.Add(1)
	go insert.SaveResource[*model.Space](resources["space"], sleep, wg)

	wg.Add(1)
	go insert.SaveResource[*model.Project](resources["project"], sleep, wg)

	wg.Add(1)
	go insert.SaveResource[*model.EventOccurrence](resources["event_occurrence"], sleep, wg)

	wg.Wait()

	invalidReferences := 0

	log.Println("Finished inserting resources...")

	log.Println("Removing invalid references on agents...")
	result := save.DB.Exec("DELETE FROM agents WHERE parent NOT IN (SELECT id FROM agents);")
	log.Println("Rows affected:", result.RowsAffected)
	invalidReferences += int(result.RowsAffected)
	save.DB.Exec("ALTER TABLE agents ADD CONSTRAINT fk_agents_parent FOREIGN KEY (parent) REFERENCES agents (id);")

	log.Println("Removing invalid references on events...")
	result = save.DB.Exec("DELETE FROM events WHERE owner NOT IN (SELECT id FROM agents);")
	log.Println("Rows affected:", result.RowsAffected)
	invalidReferences += int(result.RowsAffected)
	result = save.DB.Exec("DELETE FROM events WHERE project NOT IN (SELECT id FROM projects);")
	log.Println("Rows affected:", result.RowsAffected)
	invalidReferences += int(result.RowsAffected)
	save.DB.Exec("ALTER TABLE events ADD CONSTRAINT fk_events_owner FOREIGN KEY (owner) REFERENCES agents (id);")
	save.DB.Exec("ALTER TABLE events ADD CONSTRAINT fk_events_project FOREIGN KEY (project) REFERENCES projects (id);")

	log.Println("Removing invalid references on spaces...")
	result = save.DB.Exec("DELETE FROM spaces WHERE parent NOT IN (SELECT id FROM spaces);")
	log.Println("Rows affected:", result.RowsAffected)
	invalidReferences += int(result.RowsAffected)
	result = save.DB.Exec("DELETE FROM spaces WHERE owner NOT IN (SELECT id FROM agents);")
	log.Println("Rows affected:", result.RowsAffected)
	invalidReferences += int(result.RowsAffected)
	save.DB.Exec("ALTER TABLE spaces ADD CONSTRAINT fk_spaces_parent FOREIGN KEY (parent) REFERENCES spaces (id);")
	save.DB.Exec("ALTER TABLE spaces ADD CONSTRAINT fk_spaces_owner FOREIGN KEY (owner) REFERENCES agents (id);")

	log.Println("Removing invalid references on projects...")
	result = save.DB.Exec("DELETE FROM projects WHERE parent NOT IN (SELECT id FROM projects);")
	log.Println("Rows affected:", result.RowsAffected)
	invalidReferences += int(result.RowsAffected)
	result = save.DB.Exec("DELETE FROM projects WHERE owner NOT IN (SELECT id FROM agents);")
	log.Println("Rows affected:", result.RowsAffected)
	invalidReferences += int(result.RowsAffected)
	save.DB.Exec("ALTER TABLE projects ADD CONSTRAINT fk_projects_parent FOREIGN KEY (parent) REFERENCES projects (id);")
	save.DB.Exec("ALTER TABLE projects ADD CONSTRAINT fk_projects_owner FOREIGN KEY (owner) REFERENCES agents (id);")

	log.Println("Removing invalid references on event occurrences...")
	result = save.DB.Exec("DELETE FROM event_occurrences WHERE event NOT IN (SELECT id FROM events);")
	log.Println("Rows affected:", result.RowsAffected)
	invalidReferences += int(result.RowsAffected)
	result = save.DB.Exec("DELETE FROM event_occurrences WHERE space NOT IN (SELECT id FROM spaces);")
	log.Println("Rows affected:", result.RowsAffected)
	invalidReferences += int(result.RowsAffected)
	save.DB.Exec("ALTER TABLE event_occurrences ADD CONSTRAINT fk_event_occurrences_event FOREIGN KEY (event) REFERENCES events (id);")
	save.DB.Exec("ALTER TABLE event_occurrences ADD CONSTRAINT fk_event_occurrences_space FOREIGN KEY (space) REFERENCES spaces (id);")

	log.Println("Done removing invalid references. Total:", invalidReferences)
	log.Println("Done.")
}
