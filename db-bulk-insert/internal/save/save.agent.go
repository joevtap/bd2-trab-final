package save

import (
	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/model"
)

func SaveAgent(ch chan []model.Agent) {
	for response := range ch {
		tx := Q.Begin()

		m := tx.Agent

		for _, record := range response {
			if err := m.Create(&record); err != nil {
				tx.Rollback()
				panic(err)
			}
		}

		tx.Commit()
	}
}
