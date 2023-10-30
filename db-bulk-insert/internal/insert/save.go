package insert

import (
	"log"
	"sync"
	"time"

	"github.com/joevtap/bd2-trab-final/db-bulk-insert/gorm/model"
	"github.com/joevtap/bd2-trab-final/db-bulk-insert/internal/save"
)

func SaveResource[RType model.Model](resource Resource, sleep time.Duration, wg *sync.WaitGroup) {
	responsesChan := make(chan []RType)

	now := time.Now()

	log.Printf("Saving %q...\n", resource.Name)

	go func() {
		err := GetBatch[RType](resource, sleep, responsesChan)

		if err != nil {
			panic(err)
		}

		close(responsesChan)
	}()

	save.Save[RType](&responsesChan)
	wg.Done()

	log.Printf("Done saving %q in %v\n", resource.Name, time.Since(now))
}
