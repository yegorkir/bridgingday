package model

type Date struct {
	year  int
	month int
	day   int
}

func NewDate(year, month, day int) Date {
	return Date{year: year, month: month, day: day}
}

func NewDateFromYearDay(year, yearDay int) Date {
	d := Date{}
	d.SetYearDate(year, yearDay)
	return d
}

func (d *Date) SetYearDate(year, yearDay int) {
	d.year = year
	d.month = 1
	d.day = yearDay
	for d.day > daysInMonth(d.year, d.month) {
		d.day -= daysInMonth(d.year, d.month)
		d.month++
	}
}

func (d *Date) SetDate(year, month, day int) {
	d.year = year
	d.month = month
	d.day = day
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) YearDay() int {
	yearDay := d.day
	for m := 1; m < d.month; m++ {
		yearDay += daysInMonth(d.year, m)
	}
	return yearDay
}

func (d *Date) Date() (int, int, int) {
	return d.year, d.month, d.day
}

func (d *Date) WeekDay() int {
	// January 1, 1 AD is a Monday, use Zeller's Congruence to find day of the week
	year, month, day := d.year, d.month, d.day
	if month < 3 {
		month += 12
		year--
	}
	k := year % 100
	j := year / 100
	h := day + 13*(month+1)/5 + k + k/4 + j/4 + 5*j
	return h % 7
}

func (d *Date) GetPrevYear() *Date {
	return &Date{year: d.year - 1, month: d.month, day: d.day}
}

func (d *Date) GetNextYear() *Date {
	return &Date{year: d.year + 1, month: d.month, day: d.day}
}

func (d *Date) GetPrevMonth() *Date {
	year, month, day := d.year, d.month, d.day
	if month == 1 {
		year--
		month = 12
	} else {
		month--
	}
	if day > daysInMonth(year, month) {
		day = daysInMonth(year, month)
	}
	return &Date{year: year, month: month, day: day}
}

func (d *Date) GetNextMonth() *Date {
	year, month, day := d.year, d.month, d.day
	if month == 12 {
		year++
		month = 1
	} else {
		month++
	}
	if day > daysInMonth(year, month) {
		day = daysInMonth(year, month)
	}
	return &Date{year: year, month: month, day: day}
}

func (d *Date) GetPrevDay() *Date {
	year, month, day := d.year, d.month, d.day
	if day == 1 {
		if month == 1 {
			year--
			month = 12
		} else {
			month--
		}
		day = daysInMonth(year, month)
	} else {
		day--
	}
	return &Date{year: year, month: month, day: day}
}

func (d *Date) GetNextDay() *Date {
	year, month, day := d.year, d.month, d.day
	if day == daysInMonth(year, month) {
		day = 1
		if month == 12 {
			year++
			month = 1
		} else {
			month++
		}
	} else {
		day++
	}
	return &Date{year: year, month: month, day: day}
}

func (d *Date) More(date Date) bool {
	return (d.year > date.year) ||
		(d.year == date.year && d.month > date.month) ||
		(d.year == date.year && d.month == date.month && d.day > date.day)
}

func (d *Date) Less(date Date) bool {
	return (d.year < date.year) ||
		(d.year == date.year && d.month < date.month) ||
		(d.year == date.year && d.month == date.month && d.day < date.day)
}

func isLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}

func daysInMonth(year, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 0
	}
}
