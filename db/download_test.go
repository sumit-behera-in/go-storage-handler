package db

import "testing"

func TestClients_Download(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		c    *Clients
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Download(tt.args.fileName)
		})
	}
}
