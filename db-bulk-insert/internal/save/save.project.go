package save

import (
	"log"

	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/model"
)

func SaveProject(ch chan []model.Project) {
	for response := range ch {
		tx := Q.Begin()

		m := tx.Project

		for _, record := range response {
			if err := m.Create(&record); err != nil {
				tx.Rollback()
				panic(err)
			}
		}

		log.Printf("Done saving %v projects. Next batch...\n", len(response))

		tx.Commit()
	}
}
