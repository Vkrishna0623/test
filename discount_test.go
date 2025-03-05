package test

import "testing"

func TestDiscount(t *testing.T) {
	tests := []struct {
		wantErr  string
		price    float64
		discount float64
		want     float64
	}{{"0 Discount", 200, 0, 200}, {"have some discount", 200, 40, 160},
		{"Price can't be -ve", -200, 40, 160}, {"Discount can't be -ve", 200, -40, 160},
		{"discount can't be 100%", 200, 200, 0}}

	for _, tt := range tests {
		t.Run("Cal Discount", func(t *testing.T) {
			got, err := CalculateDiscount(tt.price, tt.discount)

			if got != tt.want {
				t.Errorf("CalculateDiscount(%v,%v)=%v, want=%v", tt.price, tt.discount, got, tt.want)
			}

			if err != nil && err.Error() != tt.wantErr {
				t.Errorf("CalculateDiscount(%v,%v)=%v, error %v wantErr %v ", tt.price, tt.discount, tt.want, err, tt.wantErr)
			} else if err == nil && tt.wantErr != "" {
				t.Errorf("CalculateDiscount(%v,%v)=%v, wantErr=%v", tt.price, tt.discount, tt.want, tt.wantErr)

			}
		})
	}
}
