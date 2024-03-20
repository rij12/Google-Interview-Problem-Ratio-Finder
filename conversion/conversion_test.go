package conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() ConversionService {
	conersionService := NewConversionService()
	return conersionService
}

func setupWithPopulatedGraph() ConversionService {
	cs := NewConversionService()
	cs.Graph["millimeter"] = make(map[string]float64, 2)
	cs.Graph["millimeter"]["meter"] = 0.01
	cs.Graph["millimeter"]["centimeter"] = 0.1
	return cs
}

func TestLoadDataIntoGraph(t *testing.T) {

}

func TestAddNodeToGraph(t *testing.T) {
	cs := setup()

	type args struct {
		origin string
		dest   string
		raito  float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should Add Centimeters to Inches",
			args: args{"centimeter", "inch", 2.540},
		},
		{
			name: "Should Add Millimeters to Meters",
			args: args{"millimeter", "meters", 0.1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs.AddNodeToGraph(tt.args.origin, tt.args.dest, tt.args.raito)
			// t.Logf("Raito = [%f]", cs.Graph[tt.args.origin][tt.args.dest])
			assert.Equal(t, cs.Graph[tt.args.origin][tt.args.dest], tt.args.raito)
		})
	}

}

func TestPopulateGraphWithStringNodes(t *testing.T) {

	cs := setup()

	tests := []struct {
		name             string
		inputStringNodes []string
		origin           string
		dest             string
		raito            float64
	}{
		{
			name:             "Should Add millimeter to centimeter",
			inputStringNodes: []string{"millimeter, centimeter, 10"},
			origin:           "millimeter",
			dest:             "centimeter",
			raito:            10,
		},
		{
			name:             "Should Add foot to Meters",
			inputStringNodes: []string{"foot, meter, 0.3048"},
			origin:           "foot",
			dest:             "meter",
			raito:            0.3048,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs.PopulateGraphWithStringNodes(tt.inputStringNodes)
			// t.Logf("Raito = [%f]", cs.Graph[tt.args.origin][tt.args.dest])
			assert.Equal(t, cs.Graph[tt.origin][tt.dest], tt.raito)
		})
	}
}

func TestGetNeighbors(t *testing.T) {

	cs := setupWithPopulatedGraph()

	want := []string{"meter", "centimeter"}

	neighbors, err := cs.GetNeighbors("millimeter")

	if assert.Equal(t, want, neighbors) != true || err != nil {
		t.FailNow()
	}
}

// func TestBreadthFirstSearchGraph(t *testing.T) {

// 	cs := setupWithPopulatedGraph()

// 	tests := []struct {
// 		name   string
// 		origin string
// 		dest   string
// 		raito  float64
// 	}{
// 		{
// 			name:   "Raito to millimeters to meters",
// 			origin: "millimeters",
// 			dest:   "meters",
// 			raito:  0.001,
// 		},
// 		{
// 			name:   "",
// 			origin: "millimeters",
// 			dest:   "light year",
// 			raito:  math.Pow(1.057, -19),
// 		},
// 	}

// }
