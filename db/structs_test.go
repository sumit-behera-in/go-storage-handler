package db

import "testing"

func TestData_isEmpty(t *testing.T) {
	tests := []struct {
		name string
		d    *Data
		want bool
	}{
		{
			name: "Test Sucessful",
			d:    &Data{fileName: "myfile", FileType: "exe", File: []byte("This is a sample mock data.")},
			want: false,
		},
		{
			name: "Test Unsucessful",
			d:    &Data{fileName: "myfile", FileType: "exe", File: []byte{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.isEmpty(); got != tt.want {
				t.Errorf("Data.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
