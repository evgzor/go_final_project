package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func afterNow(date, now time.Time) bool {
	return date.After(now)
}

func lastDayOfMonth(t time.Time) int {
	return time.Date(
		t.Year(),
		t.Month()+1,
		0,
		0, 0, 0, 0,
		t.Location(),
	).Day()
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {

	date, err := time.Parse("20060102", dstart)

	if err != nil {
		return "", err
	}

	if repeat == "" {
		return "", fmt.Errorf("empty repeat parameter")
	}
	rules := strings.Split(repeat, " ")

	mainRule := []rune(rules[0])

	symbol := mainRule[0]

	year := 0
	day := 0
	switch symbol {
	case 'y':
		if len(rules) != 1 {
			return "", errors.New("invalid y format")
		}
		year = 1
	case 'd':
		if len(rules) != 2 {
			return "", errors.New("invalid d format")
		}
		day, err = strconv.Atoi(rules[1])
		if err != nil || day < 1 || day > 400 {
			return "", errors.New("Invalid 'd' format")
		}
	case 'w':
		if len(rules) != 2 {
			return "", errors.New("invalid w format")
		}

		var week [8]bool
		for _, rule := range strings.Split(rules[1], ",") {
			value, err := strconv.Atoi(rule)
			if err != nil || value < 1 || value > 7 {
				return "", errors.New("invalid weekday")
			}
			week[value] = true
		}

		for {
			date = date.AddDate(0, 0, 1)

			if !afterNow(date, now) {
				continue
			}

			weekday := int(date.Weekday())
			if weekday == 0 {
				weekday = 7
			}

			if week[weekday] {
				return date.Format("20060102"), nil
			}
		}
	case 'm':
		days := strings.Split(rules[1], ",")
		if len(rules) < 2 || len(rules) > 3 {
			return "", errors.New("Wrong format 'm'")
		}
		var day [32]bool
		var month [13]bool
		var isLastDay = false
		var isPreLast = false

		for _, element := range days {

			value, err := strconv.Atoi(element)
			if err != nil {
				return "", err
			}

			switch value {
			case -1:
				isLastDay = true
			case -2:
				isPreLast = true
			default:
				if value < 1 || value > 31 {
					return "", fmt.Errorf("wrong day value %d", value)
				}
				day[value] = true
			}
		}

		if len(rules) > 2 {
			monthes := strings.Split(rules[2], ",")

			for _, element := range monthes {
				value, err := strconv.Atoi(element)
				if err != nil {
					return "", err
				}
				if value < 1 || value > 12 {
					return "", fmt.Errorf("wrong month value %d", value)
				}
				month[value] = true
			}
		} else {
			for i := 1; i <= 12; i++ {
				month[i] = true
			}
		}

		limit := 0
		for {
			limit++
			if limit > 366*10 {
				return "", errors.New("no date found")
			}

			date = date.AddDate(0, 0, 1)

			if !afterNow(date, now) {
				continue
			}

			dayCheck := int(date.Day())
			monthCheck := int(date.Month())
			lastDayOfMonthCheck := lastDayOfMonth(date)

			resultDay := day[dayCheck] || (isLastDay && date.Day() == lastDayOfMonthCheck) || (isPreLast && date.Day() == lastDayOfMonthCheck-1)

			if resultDay && month[monthCheck] {
				return date.Format("20060102"), nil
			}
		}

	default:
		return "", errors.ErrUnsupported
	}
	for {
		date = date.AddDate(year, 0, day)
		if afterNow(date, now) {
			break
		}
	}

	resultDate := date.Format("20060102")

	return resultDate, nil
}

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	now := r.URL.Query().Get("now")
	var nowDate time.Time
	var err error
	if len(now) != 0 {
		nowDate, err = time.Parse("20060102", now)
		if err != nil {
			http.Error(w, "invalid 'now' date format, expected YYYYMMDD", http.StatusBadRequest)
			return
		}
	} else {
		nowDate = time.Now()
	}

	date := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	resp, err := NextDate(nowDate, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
