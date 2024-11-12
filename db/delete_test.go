package db

import "testing"

func TestClients_Delete(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		c       *Clients
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Delete(tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("Clients.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
