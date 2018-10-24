package calendar

import (
	"errors"
	"fmt"
	"time"
)

type DateType bool

const (
	LUNARMONTH  DateType = true
	NORMALMONTH DateType = false
)

type lunarDate struct {
	year  int
	month int
	date  int
	leap  DateType
}

// return a lunar date
func NewLunarDate(year, month, date int, leap DateType) (*lunarDate, error) {
	var (
		err error
	)
	if err = vaildateLunarDate(year, month, date, leap); err != nil {
		return nil, err
	}
	return &lunarDate{year: year, month: month, date: date, leap: leap}, nil
}

// lunar date to solar date
func (date lunarDate) Lunar2Solar() *solarDate {
	days := date.CalDaysInterval(lunarDate{1900, 1, 1, NORMALMONTH})
	d := time.Date(1900, 1, 31, 1, 0, 0, 0, time.Local).Add(time.Hour * 24 * time.Duration(days))
	return &solarDate{d.Year(), int(d.Month()), d.Day()}
}

// cal days count between target and date
func (date lunarDate) CalDaysInterval(target lunarDate) int {
	var (
		before lunarDate
		after  lunarDate
	)
	if date.isAfter(target) {
		before = target
		after = date
	} else {
		before = date
		after = target
	}
	return calInterval(after, before)
}

// if date is after target return true
func (date lunarDate) isAfter(target lunarDate) bool {
	if date.year > target.year {
		return true
	}
	if date.year == target.year && date.month > target.month {
		return true
	}
	if date.year == target.year && date.month == target.month && date.leap && !target.leap {
		return true
	}
	if date.year == target.year && date.month == target.month && date.leap == target.leap && date.date > target.date {
		return true
	}
	return false
}

// check date is a legal lunar date
func vaildateLunarDate(year, month, date int, leap DateType) error {

	// year between 1900 to 2100
	if year < 1900 || year > 2100 {
		return errors.New("Illegal year")
	}

	// month between 1 to 12
	if month < 1 || month > 12 {
		return errors.New("Illegal month")
	}

	// date between 29 to 30
	if date < 1 || date > 30 {
		return errors.New("Illegal date")
	}

	// check lunar month in specific year
	if leap == LUNARMONTH && month != getLunarMonth(year) {
		return errors.New(fmt.Sprintf("Month: %d, is not lunar month in %d", month, year))
	}

	// check date in specific year and specific month
	if !isContains(year, month, date, leap) {
		if leap {
			return errors.New(fmt.Sprintf("Date: %d is not exists in %d年闰%d月", date, year, month))
		} else {
			return errors.New(fmt.Sprintf("Date: %d is not exists in %d年%d月", date, year, month))
		}
	}
	return nil
}

// get lunar month in specific year
// 0 means this year doesn't contain lunar month
func getLunarMonth(year int) int {
	return lunarInfo[year-ORIGINYEAR] & 0x0000f
}

// get months info in year
func getLunarYearMonths(year int) []int {
	var (
		month      []int
		tmpMonth   []int
		addedMonth = 29
	)
	// get lunar month
	lunarMonth := getLunarMonth(year)

	// get lunar month days
	year = lunarInfo[year-ORIGINYEAR]
	bigOrSmall := year & 0xf0000 >> 16

	// big month 30 days
	if bigOrSmall == 1 {
		addedMonth = 30
	}

	// get normal month info
	monthData := fmt.Sprintf("%b", year&0x0fff0>>4)

	for len(monthData) != 12 {
		monthData = "0" + monthData
	}

	for _, v := range monthData {
		if v == '1' {
			tmpMonth = append(tmpMonth, 30)
		} else {
			tmpMonth = append(tmpMonth, 29)
		}
	}

	month = tmpMonth

	// handle contains lunar month
	if lunarMonth != 0 {
		month = append([]int{}, tmpMonth[:lunarMonth]...)
		month = append(month, addedMonth)
		month = append(month, tmpMonth[lunarMonth:]...)
	}

	return month
}

// check is date is legal.
func isContains(year, month, date int, leap DateType) bool {

	lunarMonth := getLunarMonth(year)
	months := getLunarYearMonths(year)

	// normal year
	if lunarMonth == 0 && date > months[month-1] {
		return false
	}

	// lunar year,before lunar month
	if lunarMonth != 0 && month <= lunarMonth && date > months[month-1] && leap != LUNARMONTH {
		return false
	}

	// lunar year,lunar month and after lunar month
	if lunarMonth != 0 && month >= lunarMonth && date > months[month] {
		return false
	}

	return true
}

// date1 is over date2
func calInterval(date1, date2 lunarDate) int {
	var (
		i     int
		j     int
		k     int
		count int
	)

	m := getLunarMonth(date1.year)
	if (m != 0 && date1.month > m) || date1.leap == LUNARMONTH {
		date1.month++
	}

	m = getLunarMonth(date2.year)
	if (m != 0 && date2.month > m) || date2.leap == LUNARMONTH {
		date2.month++
	}

	// 开始累加
	for i = date2.year; i <= date1.year; i++ {
		months := getLunarYearMonths(i)
		if i == date2.year {
			// 起始年份
			for j = date2.month; j <= len(months); j++ {

				if j == date2.month {
					// 起始月份
					// 从当日开始
					for k = date2.date; k <= months[j-1]; k++ {
						if i == date1.year && j == date1.month && k == date1.date {
							return count
						}
						count++
					}
				} else {
					// 非起始分yue
					// 从1号开始
					for k = 1; k <= months[j-1]; k++ {
						if i == date1.year && j == date1.month && k == date1.date {
							return count
						}
						count++
					}
				}
			}
		} else {
			// 非起始年份，继续循环，从1月1日开始
			for j = 1; j <= len(months); j++ {
				for k = 1; k <= months[j-1]; k++ {
					if i == date1.year && j == date1.month && k == date1.date {
						return count
					}

					count++
				}
			}
		}
	}
	return -1
}
