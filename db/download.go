package db

import (
	"context"
	"path/filepath"
	"sync"
)

func (c *Clients) Download(fileName string) Data {
	var wg sync.WaitGroup
	var data Data
	var mu sync.Mutex
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancellation happens in all cases

	wg.Add(len(c.Clients))

	for _, client := range c.Clients {
		go func(client Client) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				// Exit if context is canceled
				return
			default:
				fileExtension := filepath.Ext(fileName)[1:]
				retrieveData := client.download(fileName, fileExtension)
				if !retrieveData.IsEmpty() {
					// Protect shared data with a mutex
					mu.Lock()
					data = retrieveData
					mu.Unlock()
					// Cancel other goroutines
					cancel()
				}
			}
		}(client)
	}

	wg.Wait()

	return data
}
