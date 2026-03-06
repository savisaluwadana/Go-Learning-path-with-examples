package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"go-zero-to-hero/projects/phase-01-unit-converter/converter"
)

func main() {
	kind := flag.String("kind", "", "conversion kind: temperature|distance|weight")
	from := flag.String("from", "", "source unit")
	to := flag.String("to", "", "target unit")
	value := flag.Float64("value", 0, "numeric value to convert")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -kind <kind> -from <unit> -to <unit> -value <number>\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "\nExamples:")
		fmt.Fprintln(flag.CommandLine.Output(), "  go run ./projects/phase-01-unit-converter -kind temperature -from c -to f -value 37")
		fmt.Fprintln(flag.CommandLine.Output(), "  go run ./projects/phase-01-unit-converter -kind distance -from km -to mi -value 5")
		fmt.Fprintln(flag.CommandLine.Output(), "  go run ./projects/phase-01-unit-converter -kind weight -from kg -to lb -value 2")
	}

	flag.Parse()

	if *kind == "" || *from == "" || *to == "" {
		flag.Usage()
		os.Exit(2)
	}

	result, err := converter.Convert(*kind, *from, *to, *value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%.4f %s = %.4f %s\n", *value, strings.ToUpper(*from), result, strings.ToUpper(*to))
}
