package insert

import (
	"log"
	"sync"
	"time"
)

func SaveResource[RType any](resource Resource, sleep time.Duration, fn func(chan []RType), wg *sync.WaitGroup) {
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

	fn(responsesChan)
	wg.Done()

	log.Printf("Done saving %q in %v\n", resource.Name, time.Since(now))
}
