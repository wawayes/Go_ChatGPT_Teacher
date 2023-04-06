package main

import (
	"fmt"
	"time"
)

func demo002() {
	format := "2001-01-01"
	startDate, _ := time.Parse(format, "2022-01-01")
	endDate, _ := time.Parse(format, "2023-12-31")
	days := diff(startDate, endDate)
	fmt.Println(days)
}

func diff(startDate, endDate time.Time) int {
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, endDate.Location())
	diff := endDate.Sub(startDate)
	return int(diff.Hours() / 24)
}
