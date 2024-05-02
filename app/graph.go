package app

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
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
			} else {
				nodeIdMap[v.Id] = struct{}{}
			}
		}
		return nil
	}
	if err = validateNodesId(g); err != nil {
		return
	}

	validateEdge := func(g *Graph) error {
		for _, v := range g.Edges {

			if _, ok := nodeIdMap[v.From]; !ok {
				return errors.New(fmt.Sprintf("node %s not defined", v.Id))
			}
			if _, ok := nodeIdMap[v.To]; !ok {
				return errors.New(fmt.Sprintf("node %s not defined", v.Id))
			}

			if v.Cost < 0 {
				return errors.New(fmt.Sprintf("cost of %s must be non-negative", v.Id))
			}
		}
		return nil
	}
	if err = validateEdge(g); err != nil {
		return
	}

	return
}

func ProcessGraph(url string) error {
	filePath, err := DownloadFileToTmp(url)
	if err != nil {
		logger.Error(err)
		return err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		logger.Error(err)
		return err
	}

	g := &Graph{}
	err = g.Parse(data)

	if err != nil {
		logger.Error(err)
		return err
	}

	ge := &GraphEntity{Graph: g}
	err = ge.create()

	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
