package db

import (
	"path/filepath"
	"sync"
)

func (c *Clients) Download(fileName string) {
	var wg sync.WaitGroup

	// wait for all clients to complete
	wg.Add(len(c.clients))

	for _, client := range c.clients {

		go func() {
			defer wg.Done()
			fileExtension := filepath.Ext(fileName)[1:]
			client.download(fileName, fileExtension)
		}()

	}

	wg.Wait()

}
