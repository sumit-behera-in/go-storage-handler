package db

func (c *Clients) Update(data Data, sizeOfTheData float64) error {
	c.Delete(data.FileName)
	c.Upload(data, sizeOfTheData)
	return nil
}
