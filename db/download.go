package db

import "path/filepath"

func (c *Clients) Download(fileName string) {
	// var wg sync.WaitGroup

	// wait for all clients to complete
	// wg.Add(len(c.clients))

	// i := 0
	// n := len(c.clients)

	// for i < n {

	// go func(index int) {
	// 		defer wg.Done()
	// 		fileExtension := filepath.Ext(fileName)
	// 		c.clients[i].download(fileName, fileExtension)

	// 	}(i)

	// 	i++

	// }

	// wg.Wait()

	c.clients[0].download(fileName, filepath.Ext(fileName)[1:])

}
