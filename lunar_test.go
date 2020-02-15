package calendar

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewLunarDate(t *testing.T) {
	t.SkipNow()
	testCases := []struct {
		No    int
		year  int
		month int
		date  int
		leap  DateType
		Err   error
	}{
		{1, 1899, 1, 1, NORMALMONTH, errors.New("Illegal year")},
		{2, 2101, 1, 1, NORMALMONTH, errors.New("Illegal year")},
		{3, 1990, 0, 1, NORMALMONTH, errors.New("Illegal month")},
		{4, 1990, 13, 1, NORMALMONTH, errors.New("Illegal month")},
		{5, 1990, 1, 0, NORMALMONTH, errors.New("Illegal date")},
		{6, 1990, 1, 31, NORMALMONTH, errors.New("Illegal date")},
		{7, 1990, 1, 30, LUNARMONTH, errors.New(fmt.Sprintf("Month: %d, is not lunar month in %d", 1, 1990))},
		{8, 1991, 1, 30, LUNARMONTH, errors.New(fmt.Sprintf("Month: %d, is not lunar month in %d", 1, 1991))},
		{9, 1900, 8, 30, LUNARMONTH, errors.New(fmt.Sprintf("Date: %d is not exists in %d年闰%d月", 30, 1900, 8))},
		{10, 1900, 8, 30, NORMALMONTH, errors.New(fmt.Sprintf("Date: %d is not exists in %d年%d月", 30, 1900, 8))},
		{11, 1900, 8, 29, NORMALMONTH, nil},
		{12, 1941, 6, 30, NORMALMONTH, errors.New(fmt.Sprintf("Date: %d is not exists in %d年%d月", 30, 1941, 6))},
		{13, 1941, 6, 30, LUNARMONTH, nil},
		{14, 1994, 7, 21, NORMALMONTH, nil},
		{15, 2000, 1, 11, NORMALMONTH, nil},
	}

	for _, c := range testCases {
		_, err := NewLunarDate(c.year, c.month, c.date, c.leap)
		if err == nil || c.Err == nil {
			if err == nil && c.Err == nil {
				continue
			}
			t.Errorf("No: %d failed. Should get %v, but got %v", c.No, c.Err, err)
		} else {
			if err.Error() != c.Err.Error() {
				t.Errorf("No: %d failed. Shourld get %v, but got %v", c.No, c.Err, err)
			}
		}
	}
}

func TestLunar2Solar(t *testing.T) {
	t.SkipNow()
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
		date := c.lDate.Lunar2Solar()
		if date.Year != c.sDate.Year || date.Month != c.sDate.Month || date.Date != c.sDate.Date {
			t.Errorf("Case %d failed. Should get %v, but got %v", c.No, c.sDate, date)
		}
	}
}
