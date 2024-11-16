package db

import (
	"path/filepath"
	"sync"
)

func (c *Clients) Download(fileName string) {
	var wg sync.WaitGroup

	// wait for all Clients to complete
	wg.Add(len(c.Clients))

	for _, client := range c.Clients {

		go func() {
			defer wg.Done()
			fileExtension := filepath.Ext(fileName)[1:]
			client.download(fileName, fileExtension)
		}()

	}

	wg.Wait()

}
