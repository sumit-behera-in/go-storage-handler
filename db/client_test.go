package db

import (
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestNew(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    Clients
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClients_addConnect(t *testing.T) {
	type args struct {
		dbProvider string
		db         Database
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
			if err := tt.c.addConnect(tt.args.dbProvider, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Clients.addConnect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mongoClient_connect(t *testing.T) {
	type args struct {
		db Database
	}
	tests := []struct {
		name    string
		client  *mongoClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.client.connect(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("mongoClient.connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_postgresClient_connect(t *testing.T) {
	type args struct {
		db Database
	}
	tests := []struct {
		name    string
		client  *postgresClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.client.connect(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("postgresClient.connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_postgresClient_close(t *testing.T) {
	tests := []struct {
		name   string
		client *postgresClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.client.close()
		})
	}
}

func Test_mongoClient_close(t *testing.T) {
	tests := []struct {
		name   string
		client *mongoClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.client.close()
		})
	}
}

func TestClients_Close(t *testing.T) {
	tests := []struct {
		name string
		c    *Clients
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Close()
		})
	}
}
