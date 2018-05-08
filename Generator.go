package main

import (
	"flag"
	"fmt"
	"os"
	"log"
)

func main() {
	serviceCommand := flag.NewFlagSet("service", flag.ExitOnError)

	serviceTextPtr := serviceCommand.String("name", "", "Generated Controller Name")
	serviceContextPtr := serviceCommand.String("context", "", "Used Context")
	servicePostrgresPtr := serviceCommand.Bool("postgres", false, "Use Postgres format [Default: Mssql]")

	if len(os.Args) < 2 {
		fmt.Println("service subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "service":
		serviceCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if serviceCommand.Parsed() {
		// Required Flags
		if *serviceTextPtr == "" {
			serviceCommand.PrintDefaults()
			os.Exit(1)
		}

		err := GenerateService(*serviceTextPtr, serviceContextPtr, servicePostrgresPtr)

		if err != nil {
			log.Fatal(err)
		}

	}
}
