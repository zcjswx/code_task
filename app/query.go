package app

import "slices"

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
	Start string `json:"start"`
	End   string `json:"end"`
	Path  string `json:"path,omitempty"`
}

type QuerySet struct {
	Queries []Query `json:"queries"`
}

func (c *Cheapest) findCheapest(g *Graph) error {
	return nil
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
