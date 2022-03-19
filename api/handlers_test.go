package api

import (
	"reflect"
	"testing"
)

func TestListOfDatesUntilEndOfYear(t *testing.T) {
	data := Date(2022, 3, 10)

	list := listOfDatesUntilEndOfYear(data.Year(), int(data.Month()), data.Day())

	var listOfDays = []YearSuccessResponse{
		{"10-03-2022"},
		{"08-04-2022"},
		{"10-05-2022"},
		{"10-06-2022"},
		{"08-07-2022"},
		{"10-08-2022"},
		{"09-09-2022"},
		{"10-10-2022"},
		{"10-11-2022"},
		{"09-12-2022"},
	}
	if !reflect.DeepEqual(list, listOfDays) {
		t.Errorf("want: %v , got: %v", listOfDays, list)
	}
}

func TestListOfDatesUntilEndOfYearFinalOfYear(t *testing.T) {
	data := Date(2022, 12, 31)

	list := listOfDatesUntilEndOfYear(data.Year(), int(data.Month()), data.Day())

	var listOfDays = []YearSuccessResponse{
		{"30-12-2022"},
	}
	if !reflect.DeepEqual(list, listOfDays) {
		t.Errorf("want: %v , got: %v", listOfDays, list)
	}
}

func TestDaysUntilSalary(t *testing.T) {
	mockCurrentData := Date(2022, 3, 13)

	mockSuccess := DaySuccessResponse{29, "11-04-2022"}
	days := daysUntilSalary(10, mockCurrentData)
	if !reflect.DeepEqual(mockSuccess, days) {
		t.Errorf("want: %v , got: %v", mockSuccess, days)
	}
}
