package times_test

import (
	"testing"
	"time"

	"github.com/noi/go-snippets/times"
)

func TestBeginDayOfMonth(t *testing.T) {
	t.Parallel()

	assertTimeEquals(t, str2jst(t, "2023-04-01 12:34:56.7"), times.BeginDayOfMonth(str2jst(t, "2023-04-05 12:34:56.7")))
}

func TestEndDayOfMonth(t *testing.T) {
	t.Parallel()

	assertTimeEquals(t, str2jst(t, "2023-04-30 12:34:56.7"), times.EndDayOfMonth(str2jst(t, "2023-04-05 12:34:56.7")))
}

func TestTruncateToDay(t *testing.T) {
	t.Parallel()

	assertTimeEquals(t, str2jst(t, "2023-04-05 00:00:00"), times.TruncateToDay(str2jst(t, "2023-04-05 12:34:56.7")))
}

func TestAddYears(t *testing.T) {
	t.Parallel()

	type args struct {
		t     time.Time
		years int
		mode  times.AddMode
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "with NormalizeExcessDays (normalized)",
			args: args{
				t:     str2jst(t, "2004-02-29 12:34:56.7"),
				years: 3,
				mode:  times.NormalizeExcessDays,
			},
			want: str2jst(t, "2007-03-01 12:34:56.7"),
		},
		{
			name: "with NormalizeExcessDays (negative & normalized)",
			args: args{
				t:     str2jst(t, "2004-02-29 12:34:56.7"),
				years: -3,
				mode:  times.NormalizeExcessDays,
			},
			want: str2jst(t, "2001-03-01 12:34:56.7"),
		},
		{
			name: "with NormalizeExcessDays (not normalized)",
			args: args{
				t:     str2jst(t, "2001-02-27 12:34:56.7"),
				years: 3,
				mode:  times.NormalizeExcessDays,
			},
			want: str2jst(t, "2004-02-27 12:34:56.7"),
		},
		{
			name: "with NormalizeExcessDays (negative & not normalized)",
			args: args{
				t:     str2jst(t, "2004-02-27 12:34:56.7"),
				years: -3,
				mode:  times.NormalizeExcessDays,
			},
			want: str2jst(t, "2001-02-27 12:34:56.7"),
		},
		{
			name: "with TruncateExcessDays (truncated)",
			args: args{
				t:     str2jst(t, "2004-02-29 12:34:56.7"),
				years: 3,
				mode:  times.TruncateExcessDays,
			},
			want: str2jst(t, "2007-02-28 12:34:56.7"),
		},
		{
			name: "with TruncateExcessDays (negative & truncated)",
			args: args{
				t:     str2jst(t, "2004-02-29 12:34:56.7"),
				years: -3,
				mode:  times.TruncateExcessDays,
			},
			want: str2jst(t, "2001-02-28 12:34:56.7"),
		},
		{
			name: "with TruncateExcessDays (not truncated & not preserved)",
			args: args{
				t:     str2jst(t, "2001-02-28 12:34:56.7"),
				years: 3,
				mode:  times.TruncateExcessDays,
			},
			want: str2jst(t, "2004-02-28 12:34:56.7"),
		},
		{
			name: "with TruncateExcessDays (negative & not truncated & not preserved)",
			args: args{
				t:     str2jst(t, "2007-02-28 12:34:56.7"),
				years: -3,
				mode:  times.TruncateExcessDays,
			},
			want: str2jst(t, "2004-02-28 12:34:56.7"),
		},
		{
			name: "with PreserveEndDayOfMonth (truncated)",
			args: args{
				t:     str2jst(t, "2004-02-29 12:34:56.7"),
				years: 3,
				mode:  times.PreserveEndDayOfMonth,
			},
			want: str2jst(t, "2007-02-28 12:34:56.7"),
		},
		{
			name: "with PreserveEndDayOfMonth (negative & truncated)",
			args: args{
				t:     str2jst(t, "2004-02-29 12:34:56.7"),
				years: -3,
				mode:  times.PreserveEndDayOfMonth,
			},
			want: str2jst(t, "2001-02-28 12:34:56.7"),
		},
		{
			name: "with PreserveEndDayOfMonth (preserved)",
			args: args{
				t:     str2jst(t, "2001-02-28 12:34:56.7"),
				years: 3,
				mode:  times.PreserveEndDayOfMonth,
			},
			want: str2jst(t, "2004-02-29 12:34:56.7"),
		},
		{
			name: "with PreserveEndDayOfMonth (negative & preserved)",
			args: args{
				t:     str2jst(t, "2007-02-28 12:34:56.7"),
				years: -3,
				mode:  times.PreserveEndDayOfMonth,
			},
			want: str2jst(t, "2004-02-29 12:34:56.7"),
		},
		{
			name: "with PreserveEndDayOfMonth (not truncated & not preserved)",
			args: args{
				t:     str2jst(t, "2001-02-27 12:34:56.7"),
				years: 3,
				mode:  times.PreserveEndDayOfMonth,
			},
			want: str2jst(t, "2004-02-27 12:34:56.7"),
		},
		{
			name: "with PreserveEndDayOfMonth (negative & not truncated & not preserved)",
			args: args{
				t:     str2jst(t, "2004-02-27 12:34:56.7"),
				years: -3,
				mode:  times.PreserveEndDayOfMonth,
			},
			want: str2jst(t, "2001-02-27 12:34:56.7"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assertTimeEquals(t, tt.want, times.AddYears(tt.args.t, tt.args.years, tt.args.mode))
		})
	}
}

func TestAddMonths(t *testing.T) {
	t.Parallel()

	type args struct {
		t      time.Time
		months int
		mode   times.AddMode
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "with NormalizeExcessDays (normalized & year changed)",
			args: args{
				t:      str2jst(t, "2004-02-29 12:34:56.7"),
				months: 12,
				mode:   times.NormalizeExcessDays,
			},
			want: str2jst(t, "2005-03-01 12:34:56.7"),
		},
		{
			name: "with NormalizeExcessDays (negative & normalized & year changed)",
			args: args{
				t:      str2jst(t, "2004-02-29 12:34:56.7"),
				months: -12,
				mode:   times.NormalizeExcessDays,
			},
			want: str2jst(t, "2003-03-01 12:34:56.7"),
		},
		{
			name: "with NormalizeExcessDays (not normalized)",
			args: args{
				t:      str2jst(t, "2004-02-27 12:34:56.7"),
				months: 3,
				mode:   times.NormalizeExcessDays,
			},
			want: str2jst(t, "2004-05-27 12:34:56.7"),
		},
		{
			name: "with NormalizeExcessDays (negative & not normalized)",
			args: args{
				t:      str2jst(t, "2004-05-27 12:34:56.7"),
				months: -3,
				mode:   times.NormalizeExcessDays,
			},
			want: str2jst(t, "2004-02-27 12:34:56.7"),
		},
		{
			name: "with TruncateExcessDays (truncated & year changed)",
			args: args{
				t:      str2jst(t, "2004-02-29 12:34:56.7"),
				months: 12,
				mode:   times.TruncateExcessDays,
			},
			want: str2jst(t, "2005-02-28 12:34:56.7"),
		},
		{
			name: "with TruncateExcessDays (negative & truncated & year changed)",
			args: args{
				t:      str2jst(t, "2004-02-29 12:34:56.7"),
				months: -12,
				mode:   times.TruncateExcessDays,
			},
			want: str2jst(t, "2003-02-28 12:34:56.7"),
		},
		{
			name: "with TruncateExcessDays (not truncated & not preserved)",
			args: args{
				t:      str2jst(t, "2003-02-28 12:34:56.7"),
				months: 3,
				mode:   times.TruncateExcessDays,
			},
			want: str2jst(t, "2003-05-28 12:34:56.7"),
		},
		{
			name: "with TruncateExcessDays (negative & not truncated & not preserved)",
			args: args{
				t:      str2jst(t, "2003-02-28 12:34:56.7"),
				months: -1,
				mode:   times.TruncateExcessDays,
			},
			want: str2jst(t, "2003-01-28 12:34:56.7"),
		},
		{
			name: "with PreserveEndDayOfMonth (truncated & year changed)",
			args: args{
				t:      str2jst(t, "2004-02-29 12:34:56.7"),
				months: 12,
				mode:   times.PreserveEndDayOfMonth,
			},
			want: str2jst(t, "2005-02-28 12:34:56.7"),
		},
		{
			name: "with PreserveEndDayOfMonth (negative & truncated & year changed)",
			args: args{
				t:      str2jst(t, "2004-02-29 12:34:56.7"),
				months: -12,
				mode:   times.PreserveEndDayOfMonth,
			},
			want: str2jst(t, "2003-02-28 12:34:56.7"),
		},
		{
			name: "with PreserveEndDayOfMonth (preserved & year changed)",
			args: args{
				t:      str2jst(t, "2003-02-28 12:34:56.7"),
				months: 12,
				mode:   times.PreserveEndDayOfMonth,
			},
			want: str2jst(t, "2004-02-29 12:34:56.7"),
		},
		{
			name: "with PreserveEndDayOfMonth (negative & preserved & year changed)",
			args: args{
				t:      str2jst(t, "2005-02-28 12:34:56.7"),
				months: -12,
				mode:   times.PreserveEndDayOfMonth,
			},
			want: str2jst(t, "2004-02-29 12:34:56.7"),
		},
		{
			name: "with PreserveEndDayOfMonth (not truncated & not preserved)",
			args: args{
				t:      str2jst(t, "2003-02-27 12:34:56.7"),
				months: 3,
				mode:   times.PreserveEndDayOfMonth,
			},
			want: str2jst(t, "2003-05-27 12:34:56.7"),
		},
		{
			name: "with PreserveEndDayOfMonth (negative & not truncated & not preserved)",
			args: args{
				t:      str2jst(t, "2003-02-27 12:34:56.7"),
				months: -1,
				mode:   times.PreserveEndDayOfMonth,
			},
			want: str2jst(t, "2003-01-27 12:34:56.7"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assertTimeEquals(t, tt.want, times.AddMonths(tt.args.t, tt.args.months, tt.args.mode))
		})
	}
}

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
			name: "day changed",
			args: args{
				t:    str2jst(t, "2000-01-01 12:34:56.7"),
				days: 3,
			},
			want: str2jst(t, "2000-01-04 12:34:56.7"),
		},
		{
			name: "day changed (negative)",
			args: args{
				t:    str2jst(t, "2000-01-04 12:34:56.7"),
				days: -3,
			},
			want: str2jst(t, "2000-01-01 12:34:56.7"),
		},
		{
			name: "month changed",
			args: args{
				t:    str2jst(t, "2000-01-30 12:34:56.7"),
				days: 3,
			},
			want: str2jst(t, "2000-02-02 12:34:56.7"),
		},
		{
			name: "month changed (negative)",
			args: args{
				t:    str2jst(t, "2000-02-02 12:34:56.7"),
				days: -3,
			},
			want: str2jst(t, "2000-01-30 12:34:56.7"),
		},
		{
			name: "year changed",
			args: args{
				t:    str2jst(t, "2000-12-30 12:34:56.7"),
				days: 3,
			},
			want: str2jst(t, "2001-01-02 12:34:56.7"),
		},
		{
			name: "year changed (negative)",
			args: args{
				t:    str2jst(t, "2001-01-02 12:34:56.7"),
				days: -3,
			},
			want: str2jst(t, "2000-12-30 12:34:56.7"),
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
		t.Fatalf("str2jst failed: %+v", err)
	}
	return jst
}
