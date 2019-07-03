package libenv

import (
	"fmt"
	"testing"
)

func TestParseEntriesWithValidEntriesOnly(t *testing.T) {
	entryTests := []struct {
		key   string
		value string
	}{
		{"first", "one"},
		{"second", "two"},
		{"third", "https://three.com?bananas=\"sure!\""},
		{"fourth", ""},
	}

	entries := []string{
		fmt.Sprintf("%s=%s", entryTests[0].key, entryTests[0].value),
		fmt.Sprintf(" %s = %s   ", entryTests[1].key, entryTests[1].value),
		fmt.Sprintf("%s=%s", entryTests[2].key, entryTests[2].value),
		fmt.Sprintf("%s=%s", entryTests[3].key, entryTests[3].value),
	}

	envVars, err := parseEntries(entries, true)

	if err != nil {
		t.Errorf("expected no errors but got: %v", err)
	}

	for _, test := range entryTests {
		t.Run(test.key, func(t *testing.T) {
			if actualValue := envVars[test.key]; actualValue != test.value {
				t.Errorf("expected var with key %s to be \"%s\" but was \"%s\"", test.key, test.value, actualValue)
			}
		})
	}
}

func TestParseEntriesWithInvalidEntryAndSuppressedErrors(t *testing.T) {
	entryTests := []struct {
		key   string
		value string
	}{
		{"first", "one"},
		{"fourth", ""},
	}

	entries := []string{
		fmt.Sprintf("%s=%s", entryTests[0].key, entryTests[0].value),
		"  = two   ",
		"third",
		fmt.Sprintf("%s=%s", entryTests[1].key, entryTests[1].value),
	}

	envVars, err := parseEntries(entries, true)

	if err != nil {
		t.Errorf("expected no errors but got: %v", err)
	}

	if n := len(envVars); n != 2 {
		t.Errorf("expected the amount of environmental variables to be %d but was %d", 2, n)
	}

	for _, test := range entryTests {
		t.Run(test.key, func(t *testing.T) {
			if actualValue := envVars[test.key]; actualValue != test.value {
				t.Errorf("expected var with key %s to be \"%s\" but was \"%s\"", test.key, test.value, actualValue)
			}
		})
	}
}

func TestParseEntriesWithEntryWithoutValueAndUnsuppressedErrors(t *testing.T) {
	var key, expectedValue string
	entries := []string{
		"first=one",
		"third",
		"fourth=",
	}

	envVars, err := parseEntries(entries, false)

	if err == nil {
		t.Errorf("expected an error but got none")
	}

	if n := len(envVars); n != 1 {
		t.Errorf("expected the amount of environmental variables to be %d but was %d", 1, n)
	}

	key, expectedValue = "first", "one"
	if actualValue := envVars[key]; actualValue != expectedValue {
		t.Errorf("expected var with key %s to be \"%s\" but was \"%s\"", key, expectedValue, actualValue)
	}
}

func TestParseEntriesWithEntryWithoutKeyAndUnsuppressedErrors(t *testing.T) {
	var key, expectedValue string
	entries := []string{
		"first=one",
		"=second",
		"fourth=",
	}

	envVars, err := parseEntries(entries, false)

	if err == nil {
		t.Errorf("expected an error but got none")
	}

	if n := len(envVars); n != 1 {
		t.Errorf("expected the amount of environmental variables to be %d but was %d", 1, n)
	}

	key, expectedValue = "first", "one"
	if actualValue := envVars[key]; actualValue != expectedValue {
		t.Errorf("expected var with key %s to be \"%s\" but was \"%s\"", key, expectedValue, actualValue)
	}
}
