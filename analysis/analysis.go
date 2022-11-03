package analysis

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Measurements struct {
	Iterations int
	Input      []int
}

type Model struct {
	Name           string
	Measurements   []Measurements
	CurrentMeasure Measurements
}

type FullModel struct {
	Static  *Model
	StaticS *Model
	StaticD *Model
	Linked  *Model
	LinkedS *Model
	LinkedF *Model
}

// NewFullAnalysis creates a new full analysis
func NewFullAnalysis() *FullModel {
	return &FullModel{
		Static:  NewModel(),
		StaticS: NewModel(),
		StaticD: NewModel(),
		Linked:  NewModel(),
		LinkedS: NewModel(),
		LinkedF: NewModel(),
	}
}

// NewModel creates a new analysis model
func NewModel() *Model {
	return &Model{
		Measurements: []Measurements{},
	}
}

// NewMeasure creates a new measure
func (m *Model) NewMeasure(x int, y int) {
	m.CurrentMeasure = Measurements{
		Iterations: 0,
		Input:      []int{x, y},
	}
}

// EndMeasure ends the current measure
func (m *Model) EndMeasure() {
	m.Measurements = append(m.Measurements, m.CurrentMeasure)
}

// Points returns the points of the model
func (m *Model) Points() plotter.XYs {
	pts := make(plotter.XYs, len(m.Measurements))
	for i, m := range m.Measurements {
		// x is the input: x,y
		pts[i].X = float64(m.Input[0])
		pts[i].Y = float64(m.Iterations)
	}
	return pts
}

// Plot a graph comparing the input and the number of iterations
func (m *FullModel) Plot() {
	// use gonum/plot
	p := plot.New()
	p.Title.Text = "Analysis"
	p.X.Label.Text = "Input"
	p.Y.Label.Text = "Iterations"

	err := plotutil.AddLinePoints(p,
		"Static", m.Static.Points(),
		"StaticS", m.StaticS.Points(),
		"StaticD", m.StaticD.Points(),
		"Linked", m.Linked.Points(),
		"LinkedS", m.LinkedS.Points(),
		"LinkedF", m.LinkedF.Points(),
	)
	// At Axis X, show the input
	p.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{
		{Value: 100, Label: "100/" + fmt.Sprint(m.Static.Measurements[0].Input[1])},
		{Value: 200, Label: "200/" + fmt.Sprint(m.Static.Measurements[1].Input[1])},
		{Value: 300, Label: "300/" + fmt.Sprint(m.Static.Measurements[2].Input[1])},
		{Value: 400, Label: "400/" + fmt.Sprint(m.Static.Measurements[3].Input[1])},
		{Value: 500, Label: "500/" + fmt.Sprint(m.Static.Measurements[4].Input[1])},
	})

	// Show Y Axis in decimal notation
	p.Y.Tick.Marker = plot.DefaultTicks{}

	if err != nil {
		panic(err)
	}

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "clean.png"); err != nil {
		panic(err)
	}
}

var FullAnalysis = NewFullAnalysis()
var Analysis = NewModel()
