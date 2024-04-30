package app

import (
	"encoding/xml"
	"errors"
	"strings"
)

type Edge struct {
	Id   string  `xml:"id" gorm:"column:edge_id"`
	From string  `xml:"from" gorm:"from"`
	To   string  `xml:"to" gorm:"to"`
	Cost float64 `xml:"cost" gorm:"cost"`
}

type Node struct {
	Id   string `xml:"id" gorm:"column:node_id"`
	Name string `xml:"name" gorm:"name"`
}

type Graph struct {
	XMLName xml.Name `xml:"graph" gorm:"-"`
	Id      string   `xml:"id" gorm:"column:graph_id"`
	Name    string   `xml:"name" gorm:"name"`
	Nodes   []*Node  `xml:"nodes>node" gorm:"-"`
	Edges   []*Edge  `xml:"edges>node" gorm:"-"`
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
