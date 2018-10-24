package calendar

import (
	"testing"
)

func TestSolar2Lunar(t *testing.T) {
	cases := []struct {
		No    int
		lDate lunarDate
		sDate solarDate
	}{
		{1, lunarDate{2018, 9, 16, NORMALMONTH}, solarDate{2018, 10, 24}},
		{2, lunarDate{1960, 6, 12, LUNARMONTH}, solarDate{1960, 8, 4}},
		{3, lunarDate{1963, 9, 17, NORMALMONTH}, solarDate{1963, 11, 2}},
		{4, lunarDate{1995, 11, 29, NORMALMONTH}, solarDate{1996, 1, 19}},
		{5, lunarDate{1990, 5, 18, NORMALMONTH}, solarDate{1990, 6, 10}},
		{6, lunarDate{1990, 5, 18, LUNARMONTH}, solarDate{1990, 7, 10}},
	}

	for _, c := range cases {
		date := c.sDate.Solar2Lunar()
		if date.year != c.lDate.year || date.month != c.lDate.month || date.date != c.lDate.date || date.leap != c.lDate.leap {
			t.Errorf("Case %d failed. Should get %v, but got %v", c.No, c.lDate, date)
		}
	}
}
