package main

import (
	"reflect"
	"testing"
)

func Test_findMostFrequentRanks(t *testing.T) {
	type args struct {
		cards []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test",
			args: args{
				cards: []string{"5s", "4c", "5d", "5c", "7s"},
			},
			want: []string{"5s", "5d", "5c"},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMostFrequentRanks(tt.args.cards); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMostFrequentRanks() = %v, want %v", got, tt.want)
			}
		})
	}
}
