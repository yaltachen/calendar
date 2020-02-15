package service

import (
	"github.com/yaltachen/calendar"
)

// DateTransServer struct
type DateTransServer struct{}

// Date struct
type Date struct {
	Year  int
	Month int
	Date  int
	Leap  bool
}

// Solar2Lunar 阳历转阴历
func (DateTransServer) Solar2Lunar(date Date,
	result *calendar.LunarDate) error {
	solardate, err := calendar.NewSolarDate(date.Year, date.Month, date.Date)
	if err != nil {
		return err
	}
	*result = *solardate.Solar2Lunar()
	return nil
}

// Lunar2Solar 阴历转阳历
func (DateTransServer) Lunar2Solar(date Date,
	result *calendar.SolarDate) error {
	lunarDate, err := calendar.NewLunarDate(
		date.Year, date.Month, date.Date, calendar.DateType(date.Leap))
	if err != nil {
		return err
	}
	*result = *lunarDate.Lunar2Solar()
	return nil
}
