package strings

import (
	"testing"
)

func TestRemoveHyphen(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "正常系",args: args{s: "99999999-99999999999-9999999999"},want: "99999999999999999999999999999"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveHyphen(tt.args.s); got != tt.want {
				t.Errorf("RemoveHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}
