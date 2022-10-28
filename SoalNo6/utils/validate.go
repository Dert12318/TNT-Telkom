package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CheckDate(str ...string) bool {

	for _, s := range str {
		if len(s) != 10 {
			return false
		}

		layout := "2006-01-02"
		t, _ := time.Parse(layout, s)
		if t.String() == "0001-01-01 00:00:00 +0000 UTC" {
			return false
		}

		now := time.Now()
		currentYear, currentMonth, _ := t.Date()
		fmt.Println(currentYear, currentMonth)
		currentLocation := now.Location()

		firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

		strArr := strings.Split(lastOfMonth.String(), " ")
		tempDate := strings.Split(strArr[0], "-")
		lastDate, err := strconv.Atoi(tempDate[2])
		if err != nil {
			return false
		}

		dateArr := strings.Split(s, "-")
		tempDateMonth, err := strconv.Atoi(dateArr[1])
		if err != nil {
			return false
		}
		if tempDateMonth > 12 {
			return false
		}
		tempDateDay, err := strconv.Atoi(dateArr[2])
		if err != nil {
			return false
		}
		if tempDateDay > lastDate {
			return false
		}

		if s[4] != '-' || s[7] != '-' {
			return false
		}
		for i := 0; i < 10; i++ {
			if i == 4 || i == 7 {
				continue
			}

			if s[i] < '0' || s[i] > '9' {
				return false
			}
		}
	}
	return true

}
