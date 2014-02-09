// Package delaunay implements Delaunay Triangulation.
package delaunay

// Node is a data point in the list of objects to determine a Delaunay
// Triangulation. X and Y are Cartesian coordinates in a plane. At
// this time, Z is another float. TODO: Change struct so that Z can be
// any type of object -- int, float, string, list, etc. Perhaps a
// pointer?
type Node struct {
	X, Y float32
	Z float32
}

// Edge is the line between the X and Y components of two of the Nodes
// that make up a Delaunay Triangle.
type Edge struct {
	n1, n2 Node
}

// Triangle is an object that consists of three Nodes that make up a
// Delaunay Triangle. It also has an index I which is an index into a
// acyclic graph of all the Delaunay Triangles which make up a
// triangulation.
type Triangle struct {
	// I int
	n1, n2, n3 Node
}

// NewNode creates a Node object and returns its address
// 
func NewNode(a, b, c float32) *Node {
	n := new(Node)
	n.X, n.Y, n.Z = a, b, c
	return n
}

// NewEdge creates an Edge object and returns its address
// NOTE: This is the most succinct declaration
func NewEdge(n1, n2 Node) *Edge {
	return &Edge{n1, n2}
}

// NewTriangle creates a Triangle object and returns its address
// TODO: Implement this so it can take either three Nodes or three Edges
// NOTE: This is more succinct than NewNode
func NewTriangle(n1, n2, n3 Node) *Triangle {
	t := Triangle{n1, n2, n3}
	return &t
}
