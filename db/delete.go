package db

import (
	"path/filepath"
	"sync"
)

func (c *Clients) Delete(fileName string) error {

	var wg sync.WaitGroup

	// wait for all clients to complete
	wg.Add(len(c.clients))

	i := 0
	n := len(c.clients)

	for i < n {

		go func(index int) {
			defer wg.Done()
			c.clients[0].delete(fileName, filepath.Ext(fileName)[1:])
		}(i)

		i++

	}

	wg.Wait()

	return nil
}
