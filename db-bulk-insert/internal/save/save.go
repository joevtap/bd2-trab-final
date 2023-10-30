package save

import (
	"log"

	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/model"
)

func Save[T model.Model](ch *chan []T) {
	for response := range *ch {
		if result := DB.Create(response); result.Error != nil {
			panic(result.Error)
		}

		log.Printf("Done saving %v %v. Next batch...\n", len(response), response[0].TableName())
	}
}
