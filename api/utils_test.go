package api

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRequestDayValidationInvalidNumber(t *testing.T) {
	r, _ := http.NewRequest("GET", "?pay_day=400", nil)
	statusCode, day := requestDayValidation(r)
	if statusCode != 400 && day != 0 {
		t.Errorf("Error invalide number")
	}
}

func TestRequestDayValidationInvalidSymbol(t *testing.T) {
	r, _ := http.NewRequest("GET", "?pay_day=asdzxc", nil)
	statusCode, day := requestDayValidation(r)
	if statusCode != 400 && day != 0 {
		t.Errorf("Error invalide number")
	}
}

func TestRequestDayValidationValidNumber(t *testing.T) {
	r, _ := http.NewRequest("GET", "?pay_day=10", nil)
	statusCode, day := requestDayValidation(r)
	if statusCode != 200 && day != 10 {
		t.Errorf("Error invalide number")
	}
}

func TestWeekendValidationSaturday(t *testing.T) {
	mockTime := Date(2022, 3, 19) // weekend Saturday
	date := Date(2022, 3, 18)
	salaryTime := weekendValidation(mockTime)
	if !reflect.DeepEqual(salaryTime, date) {
		t.Errorf("want: %v , got: %v", date, salaryTime)
	}
}

func TestWeekendValidationSunday(t *testing.T) {
	mockTime := Date(2022, 3, 20) // weekend Sunday
	date := Date(2022, 3, 18)
	salaryTime := weekendValidation(mockTime)
	if !reflect.DeepEqual(salaryTime, date) {
		t.Errorf("want: %v , got: %v", date, salaryTime)
	}
}

func TestWeekendValidationWeek(t *testing.T) {
	mockTime := Date(2022, 3, 17) // simple day
	date := Date(2022, 3, 17)
	salaryTime := weekendValidation(mockTime)
	if !reflect.DeepEqual(salaryTime, date) {
		t.Errorf("want: %v , got: %v", date, salaryTime)
	}
}

func TestDiffInDaysZeroDays(t *testing.T) {
	date1 := Date(2022, 3, 17)
	date2 := Date(2022, 3, 17)

	difference := diffInDays(date1, date2)
	if difference != 0 {
		t.Errorf("want: %v , got: %v", difference, 0)
	}
}

func TestDiffInDaysSomeDays(t *testing.T) {
	date1 := Date(2022, 3, 10)
	date2 := Date(2022, 3, 20)

	difference := diffInDays(date1, date2)
	if difference != 10 {
		t.Errorf("want: %v , got: %v", difference, 10)
	}
}
