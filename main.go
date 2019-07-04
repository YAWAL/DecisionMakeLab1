package main

import (
	"fmt"

	"github.com/YAWAL/DecisionMakeLab1/calculation"
	"github.com/YAWAL/DecisionMakeLab1/reader"
	"github.com/YAWAL/DecisionMakeLab1/reporter"
)

func main() {
	file := "input.csv"
	input, err := reader.ReadCSV(file)
	if err != nil {
		fmt.Errorf("err %s:", err.Error)
	}

	gurvic := calculation.Gurvic(input)
	fmt.Println("gurvic ", gurvic[0])
	fmt.Println("gurvic ", gurvic[1])
	fmt.Println("gurvic ", gurvic[2])

	wald := calculation.Wald(input)
	fmt.Println("wald ", wald[0])
	fmt.Println("wald ", wald[1])
	fmt.Println("wald ", wald[2])

	laplas := calculation.Laplas(input)
	fmt.Println("laplas ", laplas[0])
	fmt.Println("laplas ", laplas[1])
	fmt.Println("laplas ", laplas[2])

	bayerLaplas := calculation.BayerLaplas(input)
	fmt.Println("bayerLaplas ", bayerLaplas[0])
	fmt.Println("bayerLaplas ", bayerLaplas[1])
	fmt.Println("bayerLaplas ", bayerLaplas[2])

	err = reporter.Report(gurvic, wald, laplas, bayerLaplas, input)
	if err != nil {
		fmt.Errorf("err %s:", err.Error)
	}
}
