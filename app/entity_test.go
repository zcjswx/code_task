package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraphEntity_create(t *testing.T) {
	data := []byte("<graph> <id>g0</id><name>The Graph Name</name> <nodes><node> <id>a</id><name>A name</name> </node><node><id>e</id><name>E name</name> </node></nodes> <edges><node> <id>e1</id><from>a</from> <to>e</to> <cost>42</cost></node> ... <node><id>e5</id> <from>a</from> <to>a</to> <cost>0.42</cost></node> </edges></graph>")

	g := &Graph{}
	_ = g.Parse(data)

	ge := GraphEntity{Graph: g}
	err := ge.create()
	assert.NoError(t, err)
}

func TestGraphEntity_findLatest(t *testing.T) {
	latest, err := findLatestGraph()

	assert.NoError(t, err)
	assert.NotNil(t, latest)
	assert.NotNil(t, latest.Nodes)
	assert.NotNil(t, latest.Edges)
}
