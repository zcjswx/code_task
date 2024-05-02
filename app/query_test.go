package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPath_findAllPaths(t *testing.T) {
	node0 := &Node{Id: "0", Name: "0"}
	node1 := &Node{Id: "1", Name: "1"}
	node2 := &Node{Id: "2", Name: "2"}
	node3 := &Node{Id: "3", Name: "3"}
	node4 := &Node{Id: "4", Name: "4"}

	edgea := &Edge{"a", "3", "0", 0}
	edgeb := &Edge{"b", "3", "2", 0}
	edgec := &Edge{"c", "3", "1", 0}
	edged := &Edge{"d", "3", "4", 0}
	edgee := &Edge{"e", "2", "1", 0}
	edgef := &Edge{"f", "0", "1", 0}
	edgeg := &Edge{"g", "4", "1", 0}
	edgeh := &Edge{"h", "4", "2", 0}

	g := &Graph{Id: "graphid", Name: "graphname",
		Nodes: []*Node{node0, node1, node2, node3, node4},
		Edges: []*Edge{edgea, edgeb, edgec, edged, edgee, edgef, edgeg, edgeh},
	}

	p := Path{Start: "3", End: "1"}
	err := p.findAllPaths(g)
	assert.NoError(t, err)
	assert.Equal(t, len(p.Paths), 5)
}

func Test_findCheapest(t *testing.T) {
	node0 := &Node{Id: "0", Name: "0"}
	node1 := &Node{Id: "1", Name: "1"}
	node2 := &Node{Id: "2", Name: "2"}
	node3 := &Node{Id: "3", Name: "3"}
	node4 := &Node{Id: "4", Name: "4"}

	edgea := &Edge{"a", "3", "0", 5}
	edgeb := &Edge{"b", "3", "2", 6}
	edgec := &Edge{"c", "3", "1", 8.5}
	edged := &Edge{"d", "3", "4", 1.5}
	edgee := &Edge{"e", "2", "1", 3.5}
	edgef := &Edge{"f", "0", "1", 7}
	edgeg := &Edge{"g", "4", "1", 8}
	edgeh := &Edge{"h", "4", "2", 2.5}

	g := &Graph{Id: "graphid", Name: "graphname",
		Nodes: []*Node{node0, node1, node2, node3, node4},
		Edges: []*Edge{edgea, edgeb, edgec, edged, edgee, edgef, edgeg, edgeh},
	}

	ge := &GraphEntity{Graph: g}
	ge.create()

	path, err := findCheapest("3", "1", g)
	assert.NoError(t, err)
	assert.Equal(t, len(path), 4)
}

func TestProcess(t *testing.T) {
	str := `{
  "queries": [
    {
      "paths": {
        "start": "3",
        "end": "1"
      }
    },
    {
      "cheapest": {
        "start": "3",
        "end": "1"
      }
    },
    {
      "cheapest": {
        "start": "3",
        "end": "2"
      }
    },
{
      "cheapest": {
        "start": "1",
        "end": "0"
      }
    }
  ]
}`
	res, err := Process(str)
	assert.NoError(t, err)
	fmt.Println(res)
}
