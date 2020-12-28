package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/dherbst/plan"
)

var funcMap map[string]func(context.Context)

func init() {
	funcMap = map[string]func(context.Context){
		"until":   Until,
		"version": Version,
	}
}

// Version prints the version and exits.
func Version(ctx context.Context) {
	fmt.Printf("Version: %v\n", plan.Version)
}

// Usage prints how to invoke `plan` from the command line.
func Usage(ctx context.Context) {
	fmt.Printf(`
Usage:

plan until <2006-01-01>   ; number of days until specified date.
plan version              ; print out version information
`)

}

// Until prints the days until a specified date in format 2006-01-02.
func Until(ctx context.Context) {
	duration, err := plan.TimeUntil(flag.Arg(1))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	days := int(duration.Hours() / 24)
	fmt.Printf("%v days\n", days)
}

func main() {
	flag.Parse()

	ctx := context.Background()

	command := flag.Arg(0)

	f := funcMap[command]
	if f == nil {
		fmt.Println("Unknown command")
		Usage(ctx)
		return
	}
	f(ctx)
}
