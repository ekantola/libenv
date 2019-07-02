package libenv

import (
	"fmt"
)

const envErrorHeader = "environmental variable error"

// Environment offers the methods to handle environmental variables
type Environment struct {
	variables map[string]string
}

// New returns a new instance of libenv.Environment based on the operating system's environment
func New() (env *Environment) {
	variables, _ := ParseOsEnvironment(true)

	return NewFromMap(variables)
}

// NewFromMap returns a new instance of libenv.Environment based on the environmental variables given as argument
func NewFromMap(envVars map[string]string) (env *Environment) {
	env = &Environment{envVars}

	return
}

// Variables returns an EnvMap that holds the environment as string-to-string-tuples
func (env *Environment) Variables() map[string]string {
	return env.variables
}

// ObligatoryEnvVars takes as its argument a variadic amount of environmental variable keys. If all matching variables
// can be found in the environment, it returns them. Otherwise it returns an error message that holds info
// about which environmental variables were missing.
func (env *Environment) ObligatoryEnvVars(obligatoryEnvVarKeys ...string) (envVars []string, err error) {
	missingEnvVars := []string{}

	for _, envVarKey := range obligatoryEnvVarKeys {
		if envVar, exists := env.variables[envVarKey]; !exists {
			missingEnvVars = append(missingEnvVars, envVarKey)
		} else {
			envVars = append(envVars, envVar)
		}
	}

	if len(missingEnvVars) != 0 {
		err = fmt.Errorf("%s, the following environmental variables were not set: %v", envErrorHeader, missingEnvVars)
	}

	return
}

// GetOrDefault returns either the environmental variable associated with the envVarKey, or the defaultValue
// if the environmental variable doesn't exist
func (env *Environment) GetOrDefault(envVarKey string, defaultValue string) string {
	if envVar, exists := env.variables[envVarKey]; exists {
		return envVar
	}

	return defaultValue
}

// Get returns the environmental variable corresponding to the key or an empty string if it doesn't exist
func (env *Environment) Get(envVarKey string) string {
	return env.variables[envVarKey]
}

// Set adds an environmental variable, overwriting existing if any
func (env *Environment) Set(envVarKey string, envVar string) {
	env.variables[envVarKey] = envVar
}

// Remove removes the environmental variable corresponsing to the key, if it exists
func (env *Environment) Remove(envVarKey string) {
	delete(env.variables, envVarKey)
}
