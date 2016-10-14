package main

import (
	"fmt"

	"github.com/kr/pretty"
	"github.com/pedrommone/sentry-mttr-mtbf-calculator/log"
	"github.com/Sirupsen/logrus"

	_ "github.com/joho/godotenv/autoload"
)

var (
	activities	[]Activity
)

func main() {
	calculator := NewCalculator()
	calculator.Start()
}

func NewCalculator() *Calculator {
	calc := new(Calculator)
	calc.Log = log.NewLogrus()
	calc.Collector = Collect()

	return calc;
}

func (c *Calculator) Start() {
	fmt.Print(fmt.Sprintf("%# v", pretty.Formatter("Hello World")))
}
