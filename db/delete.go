package db

import "path/filepath"

func (c *Clients) Delete(fileName string) error {

	c.clients[0].delete(fileName, filepath.Ext(fileName)[1:])

	return nil
}
