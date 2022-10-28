package utils

import (
	"strings"

	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"
)

func LowerCamelCase(v string) string {
	s := strings.Split(v, "")
	s[0] = strings.ToLower(s[0])

	res := strings.Join(s, "")

	return res
}

func ConvertMonth(str string) string {
	var ss string
	sTemp := strings.Split(str, "-")
	switch sTemp[1] {
	case "01":
		ss = models.January
	case "02":
		ss = models.February
	case "03":
		ss = models.March
	case "04":
		ss = models.April
	case "05":
		ss = models.May
	case "06":
		ss = models.June
	case "07":
		ss = models.July
	case "08":
		ss = models.August
	case "09":
		ss = models.September
	case "10":
		ss = models.October
	case "11":
		ss = models.November
	case "12":
		ss = models.December
	}

	return ss
}
