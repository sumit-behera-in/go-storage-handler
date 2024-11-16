package db

import "testing"

func Test_mongoClient_upload(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name    string
		mc      *mongoClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.mc.upload(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("mongoClient.upload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mongoClient_download(t *testing.T) {
	type args struct {
		fileName string
		fileType string
	}
	tests := []struct {
		name string
		mc   *mongoClient
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mc.download(tt.args.fileName, tt.args.fileType)
		})
	}
}

func Test_mongoClient_delete(t *testing.T) {
	type args struct {
		fileName string
		fileType string
	}
	tests := []struct {
		name    string
		mc      *mongoClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.mc.delete(tt.args.fileName, tt.args.fileType); (err != nil) != tt.wantErr {
				t.Errorf("mongoClient.delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mongoClient_UpdateSpace(t *testing.T) {
	tests := []struct {
		name string
		mc   *mongoClient
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mc.UpdateSpace(); got != tt.want {
				t.Errorf("mongoClient.UpdateSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mongoClient_find(t *testing.T) {
	type args struct {
		fileName string
		fileType string
	}
	tests := []struct {
		name string
		mc   *mongoClient
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mc.find(tt.args.fileName, tt.args.fileType); got != tt.want {
				t.Errorf("mongoClient.find() = %v, want %v", got, tt.want)
			}
		})
	}
}
