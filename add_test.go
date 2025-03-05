package test

import (
	"testing"
)

func Test_Add(t *testing.T) {
	tests := []struct {
		a       int
		b       int
		want    int
		wantErr bool
	}{
		{2, 3, 5, false}, {1, 2, 4, true},
	}

	for _, tt := range tests {
		t.Run("Testing Add", func(t *testing.T) {
			got, err := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d,%d)=%d,want %d", tt.a, tt.b, got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Add(%d,%d) err=%v, wantErr=%v", tt.a, tt.b, err, tt.wantErr)
			}

		})
	}
}
