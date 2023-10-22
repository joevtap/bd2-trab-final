package save

import (
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

		tx.Commit()
	}
}