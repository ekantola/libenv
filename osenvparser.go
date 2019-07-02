package libenv

import (
	"fmt"
	"os"
	"strings"
)

// ParseOsEnvironment parses operating system's environmental variables and returns them as a map
func ParseOsEnvironment(suppressErrors bool) (envVars map[string]string, err error) {
	environment := os.Environ()

	envVars, err = parseEntries(environment, suppressErrors)

	return
}

func parseEntries(entries []string, suppressErrors bool) (envVars map[string]string, err error) {
	envVars = map[string]string{}

	for _, entry := range entries {
		envKey, envVar, parseErr := parseEnvEntry(entry)
		if parseErr != nil && !suppressErrors {
			err = parseErr
			return
		}

		if len(envKey) > 0 {
			envVars[envKey] = envVar
		}
	}

	return
}

func parseEnvEntry(entry string) (envKey string, envVar string, err error) {
	keyVarPair := strings.SplitN(entry, "=", 2)
	if k, v, valid := parsePair(keyVarPair); !valid {
		err = fmt.Errorf("encountered illegal entry while parsing: %s", keyVarPair)
	} else {
		envKey, envVar = k, v
	}

	return
}

func parsePair(keyVarPair []string) (envKey string, envVar string, valid bool) {
	if len(keyVarPair) != 2 {
		return "", "", false
	}

	envKey, envVar = strings.TrimSpace(keyVarPair[0]), strings.TrimSpace(keyVarPair[1])
	if len(envKey) < 1 {
		return "", "", false
	}

	valid = true
	return
}