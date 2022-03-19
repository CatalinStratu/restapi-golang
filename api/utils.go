package api

import (
	"net/http"
	"strconv"
	"time"
)

func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// requestDayValidation return request Validation
func requestDayValidation(r *http.Request) (int, int) {
	query, ok := r.URL.Query()["pay_day"]
	day, err := strconv.Atoi(query[0])

	if day > 31 || len(query[0]) < 1 || err != nil || !ok {
		return http.StatusBadRequest, 0
	}

	return http.StatusOK, day
}

// weekendValidation validates whether salaryTime is Saturday or Sunday and return Friday.
func weekendValidation(salaryTime time.Time) time.Time {
	if salaryTime.Weekday().String() == "Saturday" {
		salaryTime = salaryTime.AddDate(0, 0, -1)
		return salaryTime
	} else if salaryTime.Weekday().String() == "Sunday" {
		salaryTime = salaryTime.AddDate(0, 0, -2)
		return salaryTime
	}

	return salaryTime
}

// diffInDays return difference in days between two dates
func diffInDays(a, b time.Time) int {

	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year := y2 - y1
	month := int(M2 - M1)
	day := d2 - d1
	hour := h2 - h1
	min := m2 - m1
	sec := s2 - s1

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return day
}
