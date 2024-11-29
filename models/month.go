package models

import (
	"strings"
	"time"
)

type Month struct {
	Name        string
	ShortedName string
	Time        time.Month
}
type Year struct {
	Name string
	Time time.Time
}

// Arrays
type Months []Month
type Years []Year

func (m Months) GetMonths() Months {
	return Months{
		{Name: "January", ShortedName: "Jan", Time: time.January},
		{Name: "February", ShortedName: "Feb", Time: time.February},
		{Name: "March", ShortedName: "Mar", Time: time.March},
		{Name: "April", ShortedName: "Apr", Time: time.April},
		{Name: "May", ShortedName: "May", Time: time.May},
		{Name: "June", ShortedName: "Jun", Time: time.June},
		{Name: "July", ShortedName: "Jul", Time: time.July},
		{Name: "August", ShortedName: "Aug", Time: time.August},
		{Name: "September", ShortedName: "Sep", Time: time.September},
		{Name: "October", ShortedName: "Oct", Time: time.October},
		{Name: "November", ShortedName: "Nov", Time: time.November},
		{Name: "December", ShortedName: "Dec", Time: time.December},
		{Name: "All Year", ShortedName: "All", Time: 0},
	}
}
func (months Months) GetMonthByName(monthName string) Month {
	for _, month := range months {

		if strings.ToLower(month.Name) == strings.ToLower(monthName) || strings.ToLower(month.ShortedName) == strings.ToLower(monthName) {
			return month
		}
	}
	return Month{}
}
func (y Year) GetYears() Years {
	return Years{
		{Name: "2024", Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
		{Name: "2023", Time: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
		{Name: "2022", Time: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
		{Name: "2021", Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		{Name: "All Years", Time: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)},
	}
}
