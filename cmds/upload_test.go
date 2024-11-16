package cmds

import "testing"

func Test_getFileSizeGB(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "Successful",
			args:    args{filePath: "x.jpg"},
			want:    0.00002570357173681259,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFileSizeGB(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFileSizeGB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getFileSizeGB() = %v, want %v", got, tt.want)
			}
		})
	}
}
