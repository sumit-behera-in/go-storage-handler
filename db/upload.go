package db

import (
	"fmt"
	"sync"
)

func (c *Clients) Upload(data Data, sizeOfTheData float64) error {

	// check if the data exists anywhere
	err := c.checkFileExists(data.FileName, data.FileType)
	if err != nil {
		return err
	}

	var client Client
	var space_available = false

	i := 0

	for i < len(c.DBCollection.Database) {
		c.DBCollection.Database[i].UsedSpaceGB = c.Clients[i].UpdateSpace()
		if space_available = c.isAvailspace(sizeOfTheData, i); space_available {
			break
		}
		i++
	}

	if !space_available {
		return fmt.Errorf("unable to upload as any of the databases cant hold this data")
	}

	client = c.Clients[i]

	err = client.upload(data)
	if err != nil {
		return err
	}

	println("Successfully uploaded to database index: %v", i)
	return nil

}

func (c *Clients) isAvailspace(data float64, index int) bool {
	return (c.DBCollection.Database[index].UsedSpaceGB + data) <= 0.8*c.DBCollection.Database[index].TotalSpaceGB
}

func (c *Clients) checkFileExists(fileName string, fileType string) error {
	var wg sync.WaitGroup
	var exists bool
	var mu sync.Mutex
	resultChan := make(chan int, 1) // Buffered channel to signal the first match

	for i, client := range c.Clients {
		wg.Add(1)
		go func(idx int, cl Client) {
			defer wg.Done()
			if cl.find(fileName, fileType) {
				mu.Lock()
				if !exists { // Update exists only once
					exists = true
					resultChan <- idx // Signal the index of the match
				}
				mu.Unlock()
			}
		}(i, client)
	}

	// Wait for all goroutines or an early termination
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Handle the result
	if idx, ok := <-resultChan; ok {
		return fmt.Errorf("the file already exists on database with index: %v", idx)
	}
	return nil
}
