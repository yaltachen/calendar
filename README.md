# calendar
solar date to lunar date, lunar date to solar date
# Usage
```go
lunarDate, err := NewLunarDate(1900,1,1,calendar.NORMALMONTH)
if err != nil {
    panic(err)
}
solarDate = lunarDate.Lunar2Solar()

solarDate, err := NewSolarDate(1900,1,31)
if err != nil {
    panic(err)
}
lunarDate = solarDate.Solar2Lunar()
```