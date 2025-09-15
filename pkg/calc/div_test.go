package calc

import (
	"testing"
)

func TestDivAndRemainder(t *testing.T) {
	tests := []struct {
		name          string
		num, denom    int
		wantQuotient  int
		wantRemainder int
		wantErr       bool
	}{
		{"10を2で割る", 10, 2, 5, 0, false},
		{"10を3で割る", 10, 3, 3, 1, false},
		{"10を0で割る", 10, 0, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuotient, gotRemainder, err := divAndRemainder(tt.num, tt.denom)
			if (err != nil) != tt.wantErr {
				t.Fatalf("divAndRemainder(%d, %d) error = %v, wantErr %v", tt.num, tt.denom, err, tt.wantErr)
			}
			if gotQuotient != tt.wantQuotient {
				t.Errorf("divAndRemainder(%d, %d) gotQuotient = %d, want %d", tt.num, tt.denom, gotQuotient, tt.wantQuotient)
			}
			if gotRemainder != tt.wantRemainder {
				t.Errorf("divAndRemainder(%d, %d) gotRemainder = %d, want %d", tt.num, tt.denom, gotRemainder, tt.wantRemainder)
			}
		})
	}
}
