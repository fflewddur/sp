package libsp

import "testing"

func TestQTypeToString(t *testing.T) {
	qt := Embedded

	if qt.String() != "Embedded" {
		t.Errorf("QType.String() = %s; wanted 'Embedded'", qt.String())
	}
}

func TestUnknownQType(t *testing.T) {
	qt := newQTypeFromString("foo", "bar", "etc")

	if qt != Unknown {
		t.Errorf("QType = %s; wanted 'Unknown'", qt)
	}
}
