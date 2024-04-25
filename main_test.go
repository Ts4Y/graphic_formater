package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFormatTransactions(t *testing.T) {
	r := require.New(t)

	tests := []struct {
		name     string
		expected []Transaction
		args     struct {
			interval string
			tr       []Transaction
		}
	}{
		{
			name: "TestHourly",
			args: struct {
				interval string
				tr       []Transaction
			}{
				interval: "hour",
				tr: []Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
					{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
			},
			expected: []Transaction{

				{Value: 4321, Timestamp: time.Unix(1615888800, 0)},
				{Value: 4567, Timestamp: time.Unix(1615870800, 0)},
				{Value: 4456, Timestamp: time.Unix(1616025600, 0)},
				{Value: 4231, Timestamp: time.Unix(1616022000, 0)},
				{Value: 5212, Timestamp: time.Unix(1616018400, 0)},
			},
		},
		{
			name: "TestDaily",
			args: struct {
				interval string
				tr       []Transaction
			}{
				interval: "day",
				tr: []Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
					{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
			},
			expected: []Transaction{
				{Value: 4456, Timestamp: time.Unix(1616025600, 0)},
				{Value: 4231, Timestamp: time.Unix(1615939200, 0)},
				{Value: 4321, Timestamp: time.Unix(1615852800, 0)},
			},
		},
		{
			name: "TestWeekly",
			args: struct {
				interval string
				tr       []Transaction
			}{
				interval: "week",
				tr: []Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
					{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
			},
			expected: []Transaction{
				{Value: 4456, Timestamp: time.Unix(1615766400, 0)},
			},
		},
		{
			name: "TestMonthly",
			args: struct {
				interval string
				tr       []Transaction
			}{
				interval: "month",
				tr: []Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
					{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
			},
			expected: []Transaction{
				{Value: 4456, Timestamp: time.Unix(1614643200, 0)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatTransactions(tt.args.tr, tt.args.interval)
			r.Equal(tt.expected, got)
		})
	}

}
