package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	sbc := widgets.NewStackedBarChart()
	sbc.Title = "Year's Scores"
	sbc.Labels = []string{"1901", "1902", "1903", "1904"}

	sbc.Data = make([][]float64, 4)
	sbc.Data[0] = []float64{500, 500}
	sbc.Data[1] = []float64{250, 750}
	sbc.Data[2] = []float64{10, 990}
	sbc.Data[3] = []float64{0, 1000}
	sbc.SetRect(5, 5, 100, 30)
	sbc.BarWidth = 5

	ui.Render(sbc)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
