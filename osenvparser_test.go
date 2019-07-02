package libenv

import (
	"testing"
)

func TestParseEntriesWithValidEntriesOnly(t *testing.T) {
	var key, expectedValue string
	entries := []string{
		"first=one",
		" second = two   ",
		"third=https://three.com?bananas=\"sure!\"",
		"fourth=",
	}

	envVars, err := parseEntries(entries, true)

	if err != nil {
		t.Errorf("expected no errors but got: %v", err)
	}

	key, expectedValue = "first", "one"
	if actualValue := envVars[key]; actualValue != expectedValue {
		t.Errorf("expected var with key %s to be \"%s\" but was \"%s\"", key, expectedValue, actualValue)
	}

	key, expectedValue = "second", "two"
	if actualValue := envVars[key]; actualValue != expectedValue {
		t.Errorf("expected var with key %s to be \"%s\" but was \"%s\"", key, expectedValue, actualValue)
	}

	key, expectedValue = "third", "https://three.com?bananas=\"sure!\""
	if actualValue := envVars[key]; actualValue != expectedValue {
		t.Errorf("expected var with key %s to be \"%s\" but was \"%s\"", key, expectedValue, actualValue)
	}

	key, expectedValue = "fourth", ""
	if actualValue := envVars[key]; actualValue != expectedValue {
		t.Errorf("expected var with key %s to be \"%s\" but was \"%s\"", key, expectedValue, actualValue)
	}
}

func TestParseEntriesWithInvalidEntryAndSuppressedErrors(t *testing.T) {
	var key, expectedValue string
	entries := []string{
		"first=one",
		"  = two   ",
		"third",
		"fourth=",
	}

	envVars, err := parseEntries(entries, true)

	if err != nil {
		t.Errorf("expected no errors but got: %v", err)
	}

	if n := len(envVars); n != 2 {
		t.Errorf("expected the amount of environmental variables to be %d but was %d", 2, n)
	}

	key, expectedValue = "first", "one"
	if actualValue := envVars[key]; actualValue != expectedValue {
		t.Errorf("expected var with key %s to be \"%s\" but was \"%s\"", key, expectedValue, actualValue)
	}

	key, expectedValue = "fourth", ""
	if actualValue := envVars[key]; actualValue != expectedValue {
		t.Errorf("expected var with key %s to be \"%s\" but was \"%s\"", key, expectedValue, actualValue)
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
