package times_test

import (
	"testing"
	"time"

	"github.com/noi/go-snippets/times"
)

func TestAddDays(t *testing.T) {
	t.Parallel()

	type args struct {
		t    time.Time
		days int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "sample",
			args: args{
				t:    str2jst(t, "2000-01-31 12:34:56.7"),
				days: 1,
			},
			want: str2jst(t, "2000-02-01 12:34:56.7"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assertTimeEquals(t, tt.want, times.AddDays(tt.args.t, tt.args.days))
		})
	}
}

func assertTimeEquals(t *testing.T, left, right time.Time) {
	if !left.Equal(right) || left.Location() != right.Location() {
		t.Errorf("not equals: left=%q, right=%q", left, right)
	}
}

func str2jst(t *testing.T, str string) time.Time {
	jst, err := time.ParseInLocation("2006-01-02 15:04:05.999999999", str, times.JST)
	if err != nil {
		t.Fatal(err)
	}
	return jst
}
