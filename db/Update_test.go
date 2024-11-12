package db

import "testing"

func TestClients_Update(t *testing.T) {
	type args struct {
		fPath string
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
			if err := tt.c.Update(tt.args.fPath); (err != nil) != tt.wantErr {
				t.Errorf("Clients.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
