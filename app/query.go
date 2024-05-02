package app

import (
	"encoding/json"
	"fmt"
	"math"
	"slices"
)

type Query struct {
	Paths    *Path     `json:"paths,omitempty"`
	Cheapest *Cheapest `json:"cheapest,omitempty"`
}

type Path struct {
	Start string     `json:"start"`
	End   string     `json:"end"`
	Paths [][]string `json:"paths,omitempty"`
}

type Cheapest struct {
	Start string      `json:"start"`
	End   string      `json:"end"`
	Path  interface{} `json:"path,omitempty"`
}

type QuerySet struct {
	Queries []Query `json:"queries"`
}

func (c *Cheapest) FindCheapest(g *Graph) error {
	path, err := findCheapest(c.Start, c.End, g)
	if err != nil {
		logger.Errorf(err.Error())
	}
	if len(path) == 0 {
		c.Path = false
	} else {
		c.Path = path
	}
	return nil
}

func findCheapest(start, end string, g *Graph) ([]string, error) {
	topoSortedNodes, err := topologicalSort(g)
	if err != nil {
		return nil, err
	}

	// Initialize distances to all nodes as infinite except the start node
	dist := make(map[string]float64)
	for _, node := range g.Nodes {
		dist[node.Id] = math.Inf(1)
	}
	dist[start] = 0

	// Path reconstruction map
	prev := make(map[string]string)

	// Process nodes in topological order
	for _, node := range topoSortedNodes {
		if dist[node] < math.Inf(1) {
			for _, edge := range g.Edges {
				if edge.From == node {
					if dist[node]+edge.Cost < dist[edge.To] {
						dist[edge.To] = dist[node] + edge.Cost
						prev[edge.To] = node
					}
				}
			}
		}
	}

	// Reconstruct the path by backtracking from the end node
	if dist[end] == math.Inf(1) {
		logger.Infof("no path from %s to %s", start, end)
		return nil, nil
	}

	path := []string{}
	for at := end; at != ""; at = prev[at] {
		path = append([]string{at}, path...)
		if at == start {
			break
		}
	}
	return path, nil
}

func (p *Path) findAllPaths(g *Graph) error {

	paths := [][]string{}
	stack := [][]string{{p.Start}}

	for len(stack) > 0 {
		path := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		lastNode := path[len(path)-1]

		if lastNode == p.End {
			paths = append(paths, append([]string(nil), path...))
			continue
		}

		// Push all adjacent nodes not yet visited in the current path
		for _, edge := range g.Edges {
			if edge.From == lastNode && !slices.Contains(path, edge.To) {
				newPath := append([]string(nil), path...)
				newPath = append(newPath, edge.To)
				stack = append(stack, newPath)
			}
		}
	}
	p.Paths = paths
	return nil
}

func topologicalSort(g *Graph) ([]string, error) {
	inDegree := make(map[string]int)
	for _, node := range g.Nodes {
		inDegree[node.Id] = 0
	}
	for _, edge := range g.Edges {
		inDegree[edge.To]++
	}

	zeroInDegreeQueue := []string{}
	for node, degree := range inDegree {
		if degree == 0 {
			zeroInDegreeQueue = append(zeroInDegreeQueue, node)
		}
	}

	var topoOrder []string
	for len(zeroInDegreeQueue) > 0 {
		node := zeroInDegreeQueue[0]
		zeroInDegreeQueue = zeroInDegreeQueue[1:]
		topoOrder = append(topoOrder, node)

		for _, edge := range g.Edges {
			if edge.From == node {
				inDegree[edge.To]--
				if inDegree[edge.To] == 0 {
					zeroInDegreeQueue = append(zeroInDegreeQueue, edge.To)
				}
			}
		}
	}

	if len(topoOrder) != len(g.Nodes) {
		return nil, fmt.Errorf("graph has at least one cycle")
	}

	return topoOrder, nil
}

func readQueryFromJsonStr(str string) (*QuerySet, error) {
	var qs QuerySet
	err := json.Unmarshal([]byte(str), &qs)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
		return nil, err
	}
	return &qs, nil
}

func Process(jsonStr string) (string, error) {

	qs, err := readQueryFromJsonStr(jsonStr)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	err = qs.process()
	if err != nil {
		logger.Error(err)
		return "", err
	}

	jsonRes, err := json.Marshal(qs)
	if err != nil {
		return "", err
	}
	return string(jsonRes), nil
}

func (qs *QuerySet) process() error {
	graph, err := findLatestGraph()
	if err != nil {
		return err
	}
	for _, q := range qs.Queries {
		if q.Cheapest != nil {
			err := q.Cheapest.FindCheapest(graph)
			if err != nil {
				logger.Error(err)
				continue
			}
		}
		if q.Paths != nil {
			err := q.Paths.findAllPaths(graph)
			if err != nil {
				logger.Error(err)
				continue
			}
		}
	}
	return nil
}
