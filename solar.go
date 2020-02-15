package calendar

import (
	"errors"
	"fmt"
	"time"
)

type SolarDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Date  int `json:"date"`
}

func NewSolarDate(year, month, date int) (*SolarDate, error) {
	var (
		err error
	)
	if err = vaildateSolarDate(year, month, date); err != nil {
		return nil, err
	}
	return &SolarDate{year, month, date}, err
}

func (date SolarDate) CalDaysInterval(target SolarDate) int {
	var (
		before SolarDate
		after  SolarDate
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

func (date SolarDate) Solar2Lunar() *LunarDate {
	daysDiff := calSolarInterval(date, SolarDate{1900, 1, 31})

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
							return &LunarDate{i + 1900, j, k, NORMALMONTH}
						}

						// lunar year
						if j <= lunarInfo[i]&0x0000f {
							// before lunar month
							return &LunarDate{i + 1900, j, k, NORMALMONTH}
						} else if j == lunarInfo[i]&0x0000f+1 {
							// is lunar month
							return &LunarDate{i + 1900, j - 1, k, LUNARMONTH}
						} else {
							// after lunar month
							return &LunarDate{i + 1900, j - 1, k, NORMALMONTH}
						}
					}
				}
			}
		}
	}
	return nil
}

func (date SolarDate) isAfter(target SolarDate) bool {
	if date.Year > target.Year {
		return true
	} else if date.Year == target.Year && date.Month > target.Month {
		return true
	} else if date.Year == target.Year && date.Month > target.Month && date.Date > target.Date {
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

	if !isSolarDateExits(year, month, date) {
		return errors.New("Illegal date")
	}
	return nil
}

// VaildateSolarDate VaildateSolarDate
func VaildateSolarDate(date SolarDate) error {
	return vaildateSolarDate(date.Year, date.Month, date.Date)
}

func isSolarDateExits(year, month, date int) bool {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("month is 1 3 5 7 8 10 12")
		if date > 31 {
			fmt.Println("date > 31")
			return false
		}
	case 4, 6, 9, 11:
		fmt.Println("month is 4 6 9 11")
		if date > 30 {
			fmt.Println("date > 30")
			return false
		}
	}

	if isNormalYear(year) && month == 2 && date > 28 {
		fmt.Println("year is normal and month = 2 but date > 28")
		return false
	}

	if !isNormalYear(year) && month == 2 && date > 29 {
		fmt.Println("year is not normal and month = 2 but date > 29")
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

func calSolarInterval(date1, date2 SolarDate) int {
	now := time.Now()
	return int(
		time.Date(date1.Year, time.Month(date1.Month), date1.Date, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local).Sub(
			time.Date(date2.Year, time.Month(date2.Month), date2.Date, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local)).Hours() / 24)

}
