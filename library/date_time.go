package library

import (
	"fmt"
	"time"
)

/*GetCurrentDate ...
@desc Get current date with format of YYYY-MM-DD HH:II:SS
*/
func GetCurrentDate() string {
	dt := time.Now()
	currentDate := fmt.Sprintf("%v-%v-%v %v:%v:%v", dt.Year(), monthMapping(dt.Month()), singleValueMapping(dt.Day()), singleValueMapping(dt.Hour()), singleValueMapping(dt.Minute()), singleValueMapping(dt.Second()))
	return currentDate
}

/*
	Convert single value to formatted string with 0 prefix
*/
func singleValueMapping(value int) string {
	if value < 10 {
		return fmt.Sprintf("0%v", value)
	}

	return fmt.Sprintf("%v", value)
}

/*
	Convert string month to string of number
*/
func monthMapping(month time.Month) string {
	var formattedMonth string

	switch month.String() {
	case "January":
		formattedMonth = "01"
	case "February":
		formattedMonth = "02"
	case "March":
		formattedMonth = "03"
	case "April":
		formattedMonth = "04"
	case "May":
		formattedMonth = "05"
	case "June":
		formattedMonth = "06"
	case "July":
		formattedMonth = "07"
	case "August":
		formattedMonth = "08"
	case "September":
		formattedMonth = "09"
	case "October":
		formattedMonth = "10"
	case "November":
		formattedMonth = "11"
	case "December":
		formattedMonth = "12"
	}

	return formattedMonth
}
