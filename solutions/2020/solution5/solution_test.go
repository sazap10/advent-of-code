package solution5

import (
	"reflect"
	"testing"
)

func Test_findRow(t *testing.T) {
	tests := []struct {
		args string
		want int
	}{
		{args: "FBFBBFF", want: 44},
		{args: "BFFFBBF", want: 70},
		{args: "FFFBBBF", want: 14},
		{args: "BBFFBBF", want: 102},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := findRow(tt.args); got != tt.want {
				t.Errorf("findRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findColumn(t *testing.T) {
	tests := []struct {
		args string
		want int
	}{
		{args: "RLR", want: 5},
		{args: "RRR", want: 7},
		{args: "LRL", want: 2},
		{args: "RLL", want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := findColumn(tt.args); got != tt.want {
				t.Errorf("findColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardingPassToSeatLocation(t *testing.T) {
	tests := []struct {
		boardingPass string
		want         SeatLocation
		wantErr      bool
	}{
		{boardingPass: "FBFBBFFRLR", want: SeatLocation{Row: 44, Column: 5}},
		{boardingPass: "BFFFBBFRRR", want: SeatLocation{Row: 70, Column: 7}},
		{boardingPass: "FFFBBBFRRR", want: SeatLocation{Row: 14, Column: 7}},
		{boardingPass: "BBFFBBFRLL", want: SeatLocation{Row: 102, Column: 4}},
	}
	for _, tt := range tests {
		t.Run(tt.boardingPass, func(t *testing.T) {
			got, err := boardingPassToSeatLocation(tt.boardingPass)
			if (err != nil) != tt.wantErr {
				t.Errorf("boardingPassToSeatLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boardingPassToSeatLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
