package rules

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_parseSingle(t *testing.T) {
	type args struct {
		rule string
	}
	tests := []struct {
		args args
		want Bag
	}{
		{
			args: args{"light red bags contain 1 bright white bag, 2 muted yellow bags."},
		},
		{
			args: args{"bright white bags contain 1 shiny gold bag."},
		},
		{
			args: args{"dotted black bags contain no other bags."},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ParseSingle(tt.args.rule); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got)
				t.Errorf("ParseSingle() = %v, want %v", got, tt.want)
			}
		})
	}
}
