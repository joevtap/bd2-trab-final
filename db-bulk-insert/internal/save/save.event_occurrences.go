package save

import (
	"log"

	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/model"
)

func SaveEventOccurrences(ch chan []model.EventOccurrence) {
	for response := range ch {
		tx := Q.Begin()

		m := tx.EventOccurrence

		for _, record := range response {
			if err := m.Create(&record); err != nil {
				tx.Rollback()
				panic(err)
			}
		}

		log.Printf("Done saving %v event occurrences. Next batch...\n", len(response))

		tx.Commit()
	}
}
