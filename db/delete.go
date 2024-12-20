package db

import (
	"path/filepath"
	"sync"
)

func (c *Clients) Delete(fileName string) error {

	var wg sync.WaitGroup

	// wait for all Clients to complete
	wg.Add(len(c.Clients))

	for _, client := range c.Clients {

		go func() {
			defer wg.Done()
			client.delete(fileName, filepath.Ext(fileName)[1:])
		}()

	}

	wg.Wait()

	return nil
}
