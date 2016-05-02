package dummy

import (
	"testing"
)

// Test that the DummyFunc accepts anything implementing the Dummy interface
func TestDummyFunc(t *testing.T) {
	set_1 := (IntDummies{IntDummy(1), IntDummy(2), IntDummy(3)})
	set_2 := (FloatDummies{FloatDummy(3.12)})
	set_3 := (PointDummies{PointDummy{1, 2, 3}, PointDummy{3, 2, 1}})
	set_4 := (PersonDummies{PersonDummy{"Malloy", 7}})
	set_5 := StructDummies{string("jarjar"), []bool{true, false, true, false, false, false}}
	set_6 := &PointerDummies{7, 2}

	if DummyCount(set_1) != 3 {
		t.Errorf("set_1 %v should have been 3", DummyCount(set_1))
	}

	if DummyCount(set_2) != 1 {
		t.Errorf("set_2 %v should have been 1", DummyCount(set_2))
	}

	if DummyCount(set_3) != 2 {
		t.Errorf("set_3 %v should have been 2", DummyCount(set_3))
	}

	if DummyCount(set_4) != 1 {
		t.Errorf("set_4 %v should have been 1", DummyCount(set_4))
	}

	if DummyCount(set_5) != 6 {
		t.Errorf("set_5 %v should have been 6", DummyCount(set_5))
	}

	if DummyCount(set_6) != 15 {
		t.Errorf("set_6 %v should have been 6", DummyCount(set_6))
	}
}
