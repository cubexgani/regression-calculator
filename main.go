package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cubexgani/regression-calculator/tui"
	"github.com/cubexgani/regression-calculator/utils"
)

// The closest I can get to a constant array apparently
var RegTypes = [2]string{"quadratic", "linear"}

func main() {
	// doRegression()
	p := tea.NewProgram(tui.DadModel{
		Choice: tui.NewChoiceModel(),
		XYIn:   tui.XYInModel{},
		Result: tui.ResultModel{},
	}, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

// For testing out regressions through raw CLI
func doRegression() {
	var regtype string
	fmt.Print("Regression type: ")
	_, err := fmt.Scan(&regtype)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Ditch this part assuming that the selected regression type is valid

	exists := false
	for i := range len(RegTypes) {
		if strings.ToLower(regtype) == RegTypes[i] {
			exists = true
		}
	}
	if !exists {
		fmt.Printf("Invalid regression type %s\n", regtype)
		return
	}

	n, x, y, err := scanXY()

	if err != nil {
		return
	}

	table, err := utils.InitTable(n, x, y, regtype)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, p, err := table.Solve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Solution curve:", p)
	}
}

func scanXY() (n int, x, y []float32, err error) {
	fmt.Print("Enter number of points: ")
	_, err = fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error while scanning n:", err)
		return
	}
	x = make([]float32, n)
	y = make([]float32, n)
	fmt.Println("Points:")
	for i := range n {
		_, err = fmt.Scan(&x[i], &y[i])
		if err != nil {
			fmt.Println("Error while scanning:", err)
			return
		}
	}
	return
}
