# plan
Command line tool for creating plans.

[![Go Report Card](https://goreportcard.com/badge/github.com/dherbst/plan?style=flat-square)](https://goreportcard.com/report/github.com/dherbst/plan)

## Commands

### Until - `plan until 2020-01-20`
Prints out the number of days until the specified target date.  The format for the target date is YYYY-MM-DD.

If you run this on 2020-12-28:

    $ plan until 2021-01-20
    Days: 22
    WorkingDays: 15
    Holidays: 2
