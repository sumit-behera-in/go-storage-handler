package db

import "testing"

func TestData_isEmpty(t *testing.T) {
	tests := []struct {
		name string
		d    *Data
		want bool
	}{
		{
			name: "Test Successful",
			d:    &Data{FileName: "myfile", FileType: "exe", File: []byte("This is a sample mock data.")},
			want: false,
		},
		{
			name: "Test Unsuccessful",
			d:    &Data{FileName: "myfile", FileType: "exe", File: []byte{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.IsEmpty(); got != tt.want {
				t.Errorf("Data.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
