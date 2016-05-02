package dummy

// Dummy should contain a base entity
type Dummy interface {
	Value() int
}

// Dummies should represent a collection type of base Dummy items
type Dummies interface {
	Len() int
}

// IntDummy should inherit some behavior from the base Dummy
type IntDummy int
type FloatDummy float64
type PointDummy struct {
	x, y, z float64
}
type PersonDummy struct {
	name string
	age  int
}

// IntDummies should inherit some attributes form the base Dummies
type IntDummies []IntDummy
type FloatDummies []FloatDummy
type PointDummies []PointDummy
type PersonDummies []PersonDummy
type StructDummies struct {
	name    string
	content []bool
}
type PointerDummies struct {
	a, b int
}

func (d IntDummies) Len() int      { return len(d) }
func (d FloatDummies) Len() int    { return len(d) }
func (d PointDummies) Len() int    { return len(d) }
func (d PersonDummies) Len() int   { return len(d) }
func (d StructDummies) Len() int   { return len(d.content) }
func (d *PointerDummies) Len() int { return (d.a * d.b) + 1 }

func DummyCount(l Dummies) int { return l.Len() }
