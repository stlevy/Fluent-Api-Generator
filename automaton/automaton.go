package automaton

// Automaton is a finite state automaton, representing a Fluent API.
// Example:
// {
//	"Vertices":{
//		"v1": {
//			"Edges":{
//				"v2": "AddOne",
//				"v3": "AddTwo"
//			}
//		},
//		"v2": {
//			"Edges":{
//				"v3": "AddOne"
//			}
//		}
//	}
// }
type Automaton struct {
	Vertices map[string]Vertex
}

// Vertex is a node in the finite state automaton.
// Each vertex will be translated into a type in the final fluent API.
type Vertex struct {
	// Edges maps between the destination vertex and the edge data.
	Edges map[string]Edge
}

// Edge is a connection between two vertices in the finite state automaton, the actual string represents the required method name.
// Each Edge will be translated to a method for the translated type of the source vertex.
// This method's return type will be the translated type of the destination vertex.
type Edge string
