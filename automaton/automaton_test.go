package automaton

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONUnmarshal(t *testing.T) {
	repr := `{
	"Vertices":{
		"v1": {
			"Edges":{
				"v2": "AddOne",
				"v3": "AddTwo"
			}
		},
		"v2": {
			"Edges":{
				"v3": "AddOne"
			}
		}
	}
}`
	var automaton Automaton
	err := json.Unmarshal([]byte(repr), &automaton)
	assert.Nil(t, err, err)

	assert.Equal(t, 2, len(automaton.Vertices))
	assert.Equal(t, Edge("AddOne"), automaton.Vertices["v1"].Edges["v2"])
	assert.Equal(t, Edge("AddTwo"), automaton.Vertices["v1"].Edges["v3"])
	assert.Equal(t, Edge("AddOne"), automaton.Vertices["v2"].Edges["v3"])
}
