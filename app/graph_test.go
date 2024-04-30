package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraph_parse(t *testing.T) {
	data := []byte("<graph> <id>g0</id><name>The Graph Name</name> <nodes><node> <id>a</id><name>A name</name> </node><node><id>e</id><name>E name</name> </node></nodes> <edges><node> <id>e1</id><from>a</from> <to>e</to> <cost>42</cost></node> ... <node><id>e5</id> <from>a</from> <to>a</to> <cost>0.42</cost></node> </edges></graph>")

	g := &Graph{}
	err := g.parse(data)
	assert.Nil(t, err)
}
