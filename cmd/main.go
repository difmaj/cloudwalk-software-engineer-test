package main

import (
	"fmt"

	"github.com/difmaj/cloudwalk-software-engineer-test/internal/parser"
	"github.com/difmaj/cloudwalk-software-engineer-test/internal/report"
)

func main() {
	logData, err := parser.ParseLog("data/quake.log")
	if err != nil {
		panic(err)
	}

	response, err := report.Generate(logData)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
