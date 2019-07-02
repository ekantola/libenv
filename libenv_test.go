package libenv

import (
	"testing"
)

var mockEnvironment = map[string]string{
	"first": "one",
	"second": "two",
}

func TestNew(t *testing.T) {
	environment := New(mockEnvironment)

	if n := len(environment.Variables()); n != 2 {
		t.Errorf("expected %d variables but found %d", 2, n)
	}
}

func TestGetWhenVariableExists(t *testing.T) {
	environment := New(mockEnvironment)
	key, expected := "first", "one"

	if envVar := environment.Get(key); envVar != expected {
		t.Errorf("expected %s but got %s", expected, envVar)
	}
}

func TestGetWhenVariableDoesNotExist(t *testing.T) {
	environment := New(mockEnvironment)
	key, expected := "third", ""

	if envVar := environment.Get(key); envVar != expected {
		t.Errorf("expected %s but got %s", expected, envVar)
	}
}

func TestGetOrDefaultWhenVariableExists(t *testing.T) {
	environment := New(mockEnvironment)
	key, expected := "second", "two"

	if envVar := environment.GetOrDefault(key, "some"); envVar != expected {
		t.Errorf("expected %s but got %s", expected, envVar)
	}
}

func TestGetOrDefaultWhenVariableDoesNotExist(t *testing.T) {
	environment := New(mockEnvironment)
	key, expected := "third", "three"

	if envVar := environment.GetOrDefault(key, "three"); envVar != expected {
		t.Errorf("expected %s but got %s", expected, envVar)
	}
}

func TestObligatoryEnvVarsWhenAllVariablesExist(t *testing.T) {
	environment := New(mockEnvironment)
	envVars, err := environment.ObligatoryEnvVars("first", "second")

	if err != nil {
		t.Errorf("wasn't expecting an error but got one: %v", err)
	}

	if firstVar := envVars[0]; firstVar != "one" {
		t.Errorf("expected %s but got %s", "one", firstVar)
	}

	if secondVar := envVars[1]; secondVar != "two" {
		t.Errorf("expected %s but got %s", "two", secondVar)
	}
}

func TestObligatoryEnvVarsWhenSomeVariablesDoNotExist(t *testing.T) {
	environment := New(mockEnvironment)
	_, err := environment.ObligatoryEnvVars("first", "second", "third", "fourth")
	expectedErrorMessage := "environmental variable error, the following environmental variables were not set: [third fourth]"

	if err.Error() != expectedErrorMessage {
		t.Errorf("was expecting an error message \"%s\" but got: %v", expectedErrorMessage, err)
	}
}

func TestSetWorksCorrectly(t *testing.T) {
	environment := New(mockEnvironment)

	if nonExisting := environment.Get("third"); nonExisting != "" {
		t.Errorf("expected variable with key %s to be empty but got %s", "third", nonExisting)
	}

	environment.Set("third", "three")

	if existing := environment.Get("third"); existing != "three" {
		t.Errorf("expected %s but got %s", "three", existing)
	}
}

func TestRemoveWorksCorrectlyForAnExistingKey(t *testing.T) {
	environment := New(mockEnvironment)

	if existing := environment.Get("first"); existing != "one" {
		t.Errorf("expected variable with key %s to be %s but got %s", "first", "one", existing)
	}

	environment.Remove("first")

	if nonExisting := environment.Get("first"); nonExisting != "" {
		t.Errorf("expected variable with key %s to be empty but got %s", "first", nonExisting)
	}
}
