package calendar

import (
	"testing"
)

func TestSolar2Lunar(t *testing.T) {
	cases := []struct {
		No    int
		lDate LunarDate
		sDate SolarDate
	}{
		{1, LunarDate{2018, 9, 16, NORMALMONTH}, SolarDate{2018, 10, 24}},
		{2, LunarDate{1960, 6, 12, LUNARMONTH}, SolarDate{1960, 8, 4}},
		{3, LunarDate{1963, 9, 17, NORMALMONTH}, SolarDate{1963, 11, 2}},
		{4, LunarDate{1995, 11, 29, NORMALMONTH}, SolarDate{1996, 1, 19}},
		{5, LunarDate{1990, 5, 18, NORMALMONTH}, SolarDate{1990, 6, 10}},
		{6, LunarDate{1990, 5, 18, LUNARMONTH}, SolarDate{1990, 7, 10}},
	}

	for _, c := range cases {
		date := c.sDate.Solar2Lunar()
		if date.Year != c.lDate.Year || date.Month != c.lDate.Month || date.Date != c.lDate.Date || date.Leap != c.lDate.Leap {
			t.Errorf("Case %d failed. Should get %v, but got %v", c.No, c.lDate, date)
		}
	}
}
