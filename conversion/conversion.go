package conversion

import (
	"errors"
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/gammazero/deque"
)

type node struct {
	unit  string
	raito float64
}

type ConversionService struct {
	// map of a map [inches][centermeters] = float64
	Graph map[string]map[string]float64
}

func NewConversionService() ConversionService {
	g := make(map[string]map[string]float64)
	return ConversionService{g}
}

func (conversionService ConversionService) PopulateGraphWithStringNodes(nodes []string) {

	for _, node := range nodes {
		origin, destination, raito := CreateNodeFromString(node)
		conversionService.AddNodeToGraph(origin, destination, raito)
	}

}

func (conversionService ConversionService) AddNodeToGraph(origin string, destination string, raito float64) {

	if conversionService.Graph[origin] == nil {
		conversionService.Graph[origin] = make(map[string]float64)
	}
	conversionService.Graph[origin][destination] = raito

	if conversionService.Graph[destination] == nil {
		conversionService.Graph[destination] = make(map[string]float64)
	}
	conversionService.Graph[destination][origin] = 1 / raito

}

func (conversionService ConversionService) GetNeighbors(node string) ([]string, error) {
	// Meters -> ["inch", "cm", "foot"]
	if conversionService.Graph[node] == nil {
		return nil, fmt.Errorf(fmt.Sprintf("Unable to find node with name: [%s]", node))
	}
	return ConvertMaptokeysToList(conversionService.Graph[node]), nil
}

func (conversionService ConversionService) GetNodes() []string {
	return ConvertMapOfMapstokeysToList(conversionService.Graph)
}

func (conversionService ConversionService) BreadthFirstSearchGraph(start string, end string) (float64, error) {

	var nodesToVist deque.Deque[node]
	nodesToVist.PushBack(node{unit: start, raito: 1.0})
	visted := hashset.New()

	for nodesToVist.Len() != 0 {
		currentNode := nodesToVist.PopBack()

		if currentNode.unit == end {
			return currentNode.raito, nil
		}
		// Ensure we don't view the same unit twice
		visted.Add(currentNode.unit)

		// We've not found the path to the unit so get the current nodes neighbours
		nodes, err := conversionService.GetNeighbors(currentNode.unit)

		if err != nil {
			return 0, errors.New("node doesn't exsit")
		}

		for _, n := range nodes {
			// Loop through neighbours and if we've not seen the unit before add it to the queue.
			if !visted.Contains(n) {
				nodesToVist.PushBack(node{unit: n,
					raito: currentNode.raito * conversionService.getEdgeWeightBetweenToNodes(currentNode.unit, n)})

			}
		}
		return 0, fmt.Errorf("unable to find path between units: [%s] and [%s]", start, end)
	}

	// Nodes to visit: These are the neighbours of the current node

	// then you need to save the node that we

	return 0, fmt.Errorf("couldn't find path to unit")

}
func (conversionService ConversionService) convertUnits(origin string, destination string) {
	return
}

func (convconversionService ConversionService) getEdgeWeightBetweenToNodes(start string, end string) float64 {
	return convconversionService.Graph[start][end]
}
