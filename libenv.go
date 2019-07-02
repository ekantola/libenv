package libenv

import (
	"fmt"
)

const envErrorHeader = "environmental variable error"

// Service offers the methods to handle environmental variables
type Service struct {
	variables map[string]string
}

// New returns a new instance of libenv.Service based on the environmental variables given as argument
func New(envVars map[string]string) (service *Service) {
	service = &Service{envVars}

	return
}

// Variables returns an EnvMap that holds the environment as string-to-string-tuples
func (service *Service) Variables() map[string]string {
	return service.variables
}

// ObligatoryEnvVars takes as its argument a variadic amount of environmental variable keys. If all matching variables
// can be found in the environment, it returns them. Otherwise it returns an error message that holds info
// about which environmental variables were missing.
func (service *Service) ObligatoryEnvVars(obligatoryEnvVarKeys ...string) (envVars []string, err error) {
	missingEnvVars := []string{}

	for _, envVarKey := range obligatoryEnvVarKeys {
		if envVar, exists := service.variables[envVarKey]; !exists {
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
func (service *Service) GetOrDefault(envVarKey string, defaultValue string) string {
	if envVar, exists := service.variables[envVarKey]; exists {
		return envVar
	}

	return defaultValue
}

// Get returns the environmental variable corresponding to the key or an empty string if it doesn't exist
func (service *Service) Get(envVarKey string) string {
	return service.variables[envVarKey]
}

// Set adds an environmental variable, overwriting existing if any
func (service *Service) Set(envVarKey string, envVar string) {
	service.variables[envVarKey] = envVar
}

// Remove removes the environmental variable corresponsing to the key, if it exists
func (service *Service) Remove(envVarKey string) {
	delete(service.variables, envVarKey)
}
