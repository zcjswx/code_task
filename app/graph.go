package app

import (
	"encoding/xml"
	"errors"
	"strings"
)

type Edge struct {
	Id   string  `xml:"id"`
	From string  `xml:"from"`
	To   string  `xml:"to"`
	Cost float64 `xml:"cost"`
}

type Node struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
}

type Graph struct {
	XMLName xml.Name `xml:"graph"`
	Id      string   `xml:"id"`
	Name    string   `xml:"name"`
	Nodes   []Node   `xml:"nodes>node"`
	Edges   []Edge   `xml:"edges>node"`
}

func (g *Graph) Parse(data []byte) error {
	if err := xml.Unmarshal(data, g); err != nil {
		logger.Errorf("parse failed with err: %s", err)
		return err
	}

	return g.validate()
}

func (g *Graph) validate() (err error) {
	nodeIdMap := make(map[string]struct{})

	validateExist := func(g *Graph) error {
		if strings.TrimSpace(g.Id) == "" {
			return errors.New("there must be an <id> for the <graph>")
		}
		if strings.TrimSpace(g.Name) == "" {
			return errors.New("there must be an <name> for the <graph>")
		}
		return nil
	}

	if err = validateExist(g); err != nil {
		return
	}

	validateNodeExist := func(g *Graph) error {
		if len(g.Nodes) == 0 {
			return errors.New("there must be at least one <node> in the <nodes> group")
		}
		return nil
	}
	if err = validateNodeExist(g); err != nil {
		return
	}

	validateNodesId := func(g *Graph) error {

		for _, v := range g.Nodes {
			if _, ok := nodeIdMap[v.Id]; ok {
				return errors.New("all nodes must have different <id> tags")
			}
		}
		return nil
	}
	if err = validateNodesId(g); err != nil {
		return
	}

	validateEdge := func(g *Graph) error {
		for _, v := range g.Edges {

			if v.From == v.To {
				return errors.New("from and to must be different")
			}

			if _, ok := nodeIdMap[v.Id]; !ok {
				return errors.New("node must have been defined")
			}
		}
		return nil
	}
	if err = validateEdge(g); err != nil {
		return
	}

	return
}
