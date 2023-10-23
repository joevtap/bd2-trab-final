package save

import (
	"log"

	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/model"
)

func SaveEvent(ch chan []model.Event) {
	for response := range ch {
		tx := Q.Begin()

		m := tx.Event

		for _, record := range response {
			if err := m.Create(&record); err != nil {
				tx.Rollback()
				panic(err)
			}
		}

		log.Printf("Done saving %v events. Next batch...\n", len(response))

		tx.Commit()
	}
}
