package app

import (
	"encoding/xml"
)

type Edge struct {
	Id   string `xml:"id"`
	From string `xml:"from"`
	To   string `xml:"to"`
	Cost string `xml:"cost"`
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

func (g *Graph) parse(data []byte) error {
	if err := xml.Unmarshal(data, g); err != nil {
		logger.Errorf("parse failed with err: %s", err)
		return err
	}

	return nil
}
