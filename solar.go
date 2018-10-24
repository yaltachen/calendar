package calendar

import (
	"errors"
	"time"
)

type solarDate struct {
	year  int
	month int
	date  int
}

func NewSolarDate(year, month, date int) (*solarDate, error) {
	var (
		err error
	)
	if err = vaildateSolarDate(year, month, date); err != nil {
		return nil, err
	}
	return &solarDate{year, month, date}, err
}

func (date solarDate) CalDaysInterval(target solarDate) int {
	var (
		before solarDate
		after  solarDate
	)
	if date.isAfter(target) {
		after = date
		before = target
	} else {
		after = target
		before = date
	}

	return calSolarInterval(after, before)
}

func (date solarDate) Solar2Lunar() lunarDate {
	daysDiff := calSolarInterval(date, solarDate{1900, 1, 31})

	// i：year
	// j：month
	// k：date
	var i, j, k, sum = 0, 0, 0, -1

	// do add loop
	for sum != daysDiff {
		for i = 0; i < len(lunarInfo); i++ {
			yearInfo := getLunarYearMonths(i + ORIGINYEAR)

			for j = 1; j <= len(yearInfo); j++ {

				for k = 1; k <= yearInfo[j-1]; k++ {
					sum++
					// find
					if sum == daysDiff {
						// not lunar year
						if lunarInfo[i]&0x0000f == 0 {
							// return directly
							return lunarDate{i + 1900, j, k, NORMALMONTH}
						}

						// lunar year
						if j <= lunarInfo[i]&0x0000f {
							// before lunar month
							return lunarDate{i + 1900, j, k, NORMALMONTH}
						} else if j == lunarInfo[i]&0x0000f+1 {
							// is lunar month
							return lunarDate{i + 1900, j - 1, k, LUNARMONTH}
						} else {
							// after lunar month
							return lunarDate{i + 1900, j - 1, k, NORMALMONTH}
						}
					}
				}
			}
		}
	}
	return lunarDate{}
}

func (date solarDate) isAfter(target solarDate) bool {
	if date.year > target.year {
		return true
	} else if date.year == target.year && date.month > target.month {
		return true
	} else if date.year == target.year && date.month > target.month && date.date > target.date {
		return true
	}

	return false
}

func vaildateSolarDate(year, month, date int) error {
	// year between 1900 to 2100
	if year < 1900 || year > 2100 {
		return errors.New("Illegal year")
	}

	// month between 1 to 12
	if month < 1 || month > 12 {
		return errors.New("Illegal month")
	}

	// date between 29 to 30
	if date < 1 || date > 31 {
		return errors.New("Illegal date")
	}

	if isSolarDateExits(year, month, date) {
		return errors.New("Illegal date")
	}
	return nil
}

func isSolarDateExits(year, month, date int) bool {
	switch month {
	case 1:
	case 3:
	case 5:
	case 7:
	case 8:
	case 10:
	case 12:
		if date > 31 {
			return false
		}
	case 4:
	case 6:
	case 9:
	case 11:
		if date > 30 {
			return false
		}
	}

	if isNormalYear(year) && month == 2 && date > 28 {
		return false
	}

	if !isNormalYear(year) && month == 2 && date > 29 {
		return false
	}

	return true
}

// if year has 366 days return false
func isNormalYear(year int) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return false
	}
	return true
}

func calSolarInterval(date1, date2 solarDate) int {
	now := time.Now()
	return int(
		time.Date(date1.year, time.Month(date1.month), date1.date, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local).Sub(
			time.Date(date2.year, time.Month(date2.month), date2.date, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local)).Hours() / 24)

}
