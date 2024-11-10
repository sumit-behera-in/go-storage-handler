package db

import (
	"path/filepath"
	"sync"
)

func (c *Clients) Delete(fileName string) error {

	var wg sync.WaitGroup

	// wait for all clients to complete
	wg.Add(len(c.clients))

	for _, client := range c.clients {

		go func() {
			defer wg.Done()
			client.delete(fileName, filepath.Ext(fileName)[1:])
		}()

	}

	wg.Wait()

	return nil
}
