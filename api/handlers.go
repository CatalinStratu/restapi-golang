package api

import (
	"encoding/json"
	"math"
	"net/http"
	"time"
)

type ErrorResponse struct {
	Status  int    `json:"statusCode"`
	Message string `json:"message"`
}

type DaySuccessResponse struct {
	Days int    `json:"days"`
	Date string `json:"date"`
}

type YearSuccessResponse struct {
	Date string `json:"date"`
}

func HowMuch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		statusCode, day := requestDayValidation(r)
		if statusCode == http.StatusBadRequest {
			w.WriteHeader(http.StatusBadRequest)
			jsonResp, err := json.Marshal(ErrorResponse{http.StatusBadRequest, "Bad request"})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}
			w.Write(jsonResp)
			return
		}

		currentTime := time.Now()
		w.WriteHeader(http.StatusOK)
		jsonResp, err := json.Marshal(daysUntilSalary(day, currentTime))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		w.Write(jsonResp)
		return
	default:
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func ListHowMuch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		currentTime := time.Now()
		statusCode, day := requestDayValidation(r)
		if statusCode == http.StatusBadRequest {
			w.WriteHeader(http.StatusBadRequest)
			jsonResp, err := json.Marshal(ErrorResponse{http.StatusBadRequest, "Bad request"})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}
			w.Write(jsonResp)
			return
		}

		successResponses := listOfDatesUntilEndOfYear(currentTime.Year(), int(currentTime.Month()), day)
		jsonResp, err := json.Marshal(successResponses)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
		return
	default:
		w.Header().Add("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

//daysUntilSalary return the number of days remaining until the salary and the date when the salary will be received
func daysUntilSalary(day int, currentTime time.Time) DaySuccessResponse {
	month := currentTime.Month() + 1

	//the day the employee will receive the salary in this month
	salaryTime := Date(currentTime.Year(), month, day)

	//If the salary was received this month, we will transfer it to the next month
	if salaryTime.Day()-currentTime.Day() < 0 {
		salaryTime = Date(salaryTime.Year(), salaryTime.Month()+1, day)
	}

	//Difference in days between the current date and the date when the salary must be received
	days := diffInDays(salaryTime, currentTime)

	//at the current time we add days + 1
	salaryTime = currentTime.AddDate(0, 0, days+1)

	//check if the payday is on the weekend
	salaryTime = weekendValidation(salaryTime)

	nrDays := salaryTime.Sub(currentTime).Hours() / 24
	return DaySuccessResponse{int(math.Ceil(nrDays)), salaryTime.Format("02-01-2006")}
}

// listOfDatesUntilEndOfYear return list of salary data until the end of the year
func listOfDatesUntilEndOfYear(year, month, day int) []YearSuccessResponse {
	var successResponses []YearSuccessResponse
	for i := month; i <= 12; i++ {
		salaryTime := Date(year, time.Month(i), day)
		salaryTime = weekendValidation(salaryTime)
		item := YearSuccessResponse{Date: salaryTime.Format("02-01-2006")}
		successResponses = append(successResponses, item)
	}
	return successResponses
}
