package db

func (c *Clients) Update(fPath string) error {
	c.Delete(fPath)
	c.Upload(fPath)
	return nil
}
