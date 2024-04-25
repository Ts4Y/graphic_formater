package main

import (
	"fmt"
	"sort"
	"time"
)

type Transaction struct {
	Value     int
	Timestamp time.Time
}

func FormatTransactions(transactions []Transaction, interval string) []Transaction {
	lastTransaction := make(map[string]Transaction)

	for i, transaction := range transactions {
		switch interval {
		case "hour":
			transaction.Timestamp = transaction.Timestamp.Truncate(time.Hour)
		case "day":
			transaction.Timestamp = transaction.Timestamp.Truncate(24 * time.Hour)
		case "week":
			transaction.Timestamp = transaction.Timestamp.Truncate(24 * 7 * time.Hour)
		case "month":
			transaction.Timestamp = transaction.Timestamp.Truncate(time.Hour * 24 * 30)
		default:
			panic("unsupported interval")
		}

		transactions[i] = transaction

		ex, ok := lastTransaction[transaction.Timestamp.String()]
		if !ok || transaction.Timestamp.After(ex.Timestamp) {
			lastTransaction[transaction.Timestamp.String()] = transaction
		}
	}

	var result []Transaction
	for _, transaction := range lastTransaction {
		result = append(result, transaction)
	}

	switch interval {
	case "hour":
		sort.Slice(result, func(i, j int) bool {
			return result[i].Timestamp.Hour() > result[j].Timestamp.Hour()
		})
	case "day":
		sort.Slice(result, func(i, j int) bool {
			return result[i].Timestamp.Day() > result[j].Timestamp.Day()
		})
	case "week":
		sort.Slice(result, func(i, j int) bool {
			return beginningOfWeek(result[j].Timestamp).Before(beginningOfWeek(result[i].Timestamp))
		})
	case "month":
		sort.Slice(result, func(i, j int) bool {
			return result[i].Timestamp.Month() > result[j].Timestamp.Month()
		})
	}

	return result
}

func main() {

	transactions := []Transaction{
		{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
		{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
		{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
		{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
		{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
	}

	result := FormatTransactions(transactions, "month")

	for _, transaction := range result {
		fmt.Printf("{Value: %d, Timestamp: %d, Date: %v}\n", transaction.Value, transaction.Timestamp.Truncate(time.Hour*24).Unix(), transaction.Timestamp)
	}
}

func beginningOfWeek(t time.Time) time.Time {
	_, offset := t.Zone()
	return t.AddDate(0, 0, -int(t.Weekday())).Truncate(24 * time.Hour).Add(time.Duration(offset) * time.Second)
}
