package db

import "testing"

func Test_postgresClient_createTable(t *testing.T) {
	type args struct {
		fileType string
	}
	tests := []struct {
		name    string
		pc      *postgresClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pc.createTable(tt.args.fileType); (err != nil) != tt.wantErr {
				t.Errorf("postgresClient.createTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_postgresClient_upload(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name    string
		pc      *postgresClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pc.upload(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("postgresClient.upload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_postgresClient_download(t *testing.T) {
	type args struct {
		fileName string
		fileType string
	}
	tests := []struct {
		name string
		pc   *postgresClient
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pc.download(tt.args.fileName, tt.args.fileType)
		})
	}
}

func Test_postgresClient_delete(t *testing.T) {
	type args struct {
		fileName string
		fileType string
	}
	tests := []struct {
		name    string
		pc      *postgresClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pc.delete(tt.args.fileName, tt.args.fileType); (err != nil) != tt.wantErr {
				t.Errorf("postgresClient.delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_postgresClient_updateSpace(t *testing.T) {
	tests := []struct {
		name string
		pc   *postgresClient
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.updateSpace(); got != tt.want {
				t.Errorf("postgresClient.updateSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresClient_find(t *testing.T) {
	type args struct {
		fileName string
		fileType string
	}
	tests := []struct {
		name string
		pc   *postgresClient
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.find(tt.args.fileName, tt.args.fileType); got != tt.want {
				t.Errorf("postgresClient.find() = %v, want %v", got, tt.want)
			}
		})
	}
}
