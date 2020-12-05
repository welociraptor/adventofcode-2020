package boardingpass

import (
	"testing"
)

func TestSeatID(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{"BFFFBBFRRR"},
			567,
		},
		{
			args{"FFFBBBFRRR"},
			119,
		},
		{
			args{"BBFFBBFRLL"},
			820,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SeatID(tt.args.code); got != tt.want {
				t.Errorf("SeatID() = %v, want %v", got, tt.want)
			}
		})
	}
}
