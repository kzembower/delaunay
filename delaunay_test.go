// Testing suite for the delaunay go package.
package delaunay

import "testing"

func TestNewNode(t *testing.T) {
	const x, y, z float32 = 3.0, 4.0, 5.0
	n := NewNode(x, y, z)
	if n.X != x {t.Errorf("NewNode(%v) = %v, want %v", x, n.X, x)}
	if n.Y != y {t.Errorf("NewNode(%v) = %v, want %v", y, n.Y, y)}
	if n.Z != z {t.Errorf("NewNode(%v) = %v, want %v", z, n.Z, z)}
}
