package utils

import "time"

func GetLastMonthRange() (string, string) {
	now := time.Now()
	per1 := time.Date(now.Year(), now.Month(), 0, 0, 0, 0, 0, time.UTC)
	stDate := per1.Format("2006-01-02")
	endDate := now.Format("2006-01-02")
	return stDate, endDate
}

func GetWeekRange() (string, string) {
	now := time.Now()
	per1 := now.AddDate(0, 0, -7)
	stDate := per1.Format("2006-01-02")
	endDate := now.Format("2006-01-02")
	return stDate, endDate
}

func GetTwoDayRange() (string, string) {
	now := time.Now()
	per1 := now.AddDate(0, 0, -1)
	stDate := per1.Format("2006-01-02")
	endDate := now.Format("2006-01-02")
	return stDate, endDate
}

type Month time.Month

func (m Month) String() (month string) {
	switch m {
	case Month(time.January):
		month = "Январь"
	case Month(time.February):
		month = "Февраль"
	case Month(time.March):
		month = "Март"
	case Month(time.April):
		month = "Апрель"
	case Month(time.May):
		month = "Май"
	case Month(time.June):
		month = "Июнь"
	case Month(time.July):
		month = "Июль"
	case Month(time.August):
		month = "Август"
	case Month(time.September):
		month = "Сентябрь"
	case Month(time.October):
		month = "Октябрь"
	case Month(time.November):
		month = "Ноябрь"
	case Month(time.December):
		month = "Декабрь"
	}
	return
}

func (m Month) StringInflected() (month string) {
	switch m {
	case Month(time.January):
		month = "Января"
	case Month(time.February):
		month = "Февраля"
	case Month(time.March):
		month = "Марта"
	case Month(time.April):
		month = "Апреля"
	case Month(time.May):
		month = "Мая"
	case Month(time.June):
		month = "Июня"
	case Month(time.July):
		month = "Июля"
	case Month(time.August):
		month = "Августа"
	case Month(time.September):
		month = "Сентября"
	case Month(time.October):
		month = "Октября"
	case Month(time.November):
		month = "Ноября"
	case Month(time.December):
		month = "Декабря"
	}
	return
}
