package libenv

import (
	"reflect"
	"testing"
)

func TestStringMapIsCopiedProperly(t *testing.T) {
	original := map[string]string{
		"one": "first",
		"two": "second",
	}

	copy := CopyStringMap(original)

	if &copy == &original {
		t.Errorf("expected maps to be different instances but they weren't")
	}

	if !reflect.DeepEqual(original, copy) {
		t.Errorf("expected the original and copy to be deeply equal but they weren't")
	}
}
