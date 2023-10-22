package save

import (
	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/model"
)

func SaveSpace(ch chan []model.Space) {
	for response := range ch {
		tx := Q.Begin()

		m := tx.Space

		for _, record := range response {
			if err := m.Create(&record); err != nil {
				tx.Rollback()
				panic(err)
			}
		}

		tx.Commit()
	}
}
