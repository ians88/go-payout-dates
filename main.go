package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func payoutDay(month time.Month, now time.Time) time.Time {
	// Get the first day of the next month, substract a day to get the last day of the current month
	preferredPayoutDate := time.Date(now.Year(), month+1, 1, 0, 0, 0, 0, now.Location()).AddDate(0, 0, -1)

	var actualPayoutDate time.Time

	switch preferredPayoutDate.Weekday() {
	case time.Saturday:
		actualPayoutDate = preferredPayoutDate.AddDate(0, 0, -1)
	case time.Sunday:
		actualPayoutDate = preferredPayoutDate.AddDate(0, 0, -2)
	default:
		actualPayoutDate = preferredPayoutDate
	}

	return actualPayoutDate
}

func bonusPayoutDay(month time.Month, now time.Time) time.Time {
	// Get the preferred payout date (15th of the month)
	preferredPayoutDate := time.Date(now.Year(), month, 15, 0, 0, 0, 0, now.Location())

	var actualPayoutDate time.Time

	switch preferredPayoutDate.Weekday() {
	case time.Saturday, time.Sunday:
		actualPayoutDate = nextWednesday(preferredPayoutDate)
	default:
		actualPayoutDate = preferredPayoutDate
	}

	return actualPayoutDate
}

func nextWednesday(from time.Time) time.Time {
	days := int(time.Wednesday - from.Weekday())
	if days <= 0 {
		days += 7
	}
	return from.AddDate(0, 0, days)
}

func main() {
	now := time.Now()
	currentMonth := now.Month()

	file, err := os.Create("payout-dates.csv")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	writer.Write([]string{"Month", "Payout day", "Bonus payout day"})

	for month := currentMonth; month <= time.December; month++ {
		if err := writer.Write([]string{
			month.String(),
			payoutDay(month, now).Format("2006-01-02"),
			bonusPayoutDay(month+1, now).Format("2006-01-02"),
		}); err != nil {
			log.Fatalf("error writing record: %v", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatalf("error flushing csv: %v", err)
	}

	fmt.Println("CSV file created successfully: payout-dates.csv")
}
