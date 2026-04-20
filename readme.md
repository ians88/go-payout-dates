# Go Payout dates CLI tool

## Description

This CLI tool generates a csv of payout dates per month for the remainder of the year.

The first columns is the month name.

The second column contains the date of the last day of that month, unless that would be in the weekend.

The third column contains the date for the bonus payout, which is the 15th of next month. Unless that's in the weekend, in which case it should be on the Wednesday of the next week.

## Running the program

The linux/macos/windows folders each contain a binary to generate the csv.

Windows

`payout-dates.exe`

MacOS
`./payout-dates`

Linux
`./payout-dates`

Running this should generate a `payout-dates.csv`
