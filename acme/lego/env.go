// This file based on https://github.com/go-acme/lego/blob/v4.19.2/platform/config/env/env.go
// This file is licensed under the MIT License (MIT)

package lego

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Env struct {
	data map[string]string
}

func NewEnv(data map[string]string) Env {
	return Env{data: data}
}

// Get environment variables.
func (env Env) Get(names ...string) (map[string]string, error) {
	values := map[string]string{}

	var missingEnvVars []string
	for _, envVar := range names {
		value := env.GetOrFile(envVar)
		if value == "" {
			missingEnvVars = append(missingEnvVars, envVar)
		}
		values[envVar] = value
	}

	if len(missingEnvVars) > 0 {
		return nil, fmt.Errorf("some credentials information are missing: %s", strings.Join(missingEnvVars, ","))
	}

	return values, nil
}

// GetWithFallback Get environment variable values.
// The first name in each group is use as key in the result map.
//
// case 1:
//
//	// LEGO_ONE="ONE"
//	// LEGO_TWO="TWO"
//	env.GetWithFallback([]string{"LEGO_ONE", "LEGO_TWO"})
//	// => "LEGO_ONE" = "ONE"
//
// case 2:
//
//	// LEGO_ONE=""
//	// LEGO_TWO="TWO"
//	env.GetWithFallback([]string{"LEGO_ONE", "LEGO_TWO"})
//	// => "LEGO_ONE" = "TWO"
//
// case 3:
//
//	// LEGO_ONE=""
//	// LEGO_TWO=""
//	env.GetWithFallback([]string{"LEGO_ONE", "LEGO_TWO"})
//	// => error
func (env Env) GetWithFallback(groups ...[]string) (map[string]string, error) {
	values := map[string]string{}

	var missingEnvVars []string
	for _, names := range groups {
		if len(names) == 0 {
			return nil, errors.New("undefined environment variable names")
		}

		value, envVar := env.getOneWithFallback(names[0], names[1:]...)
		if value == "" {
			missingEnvVars = append(missingEnvVars, envVar)
			continue
		}
		values[envVar] = value
	}

	if len(missingEnvVars) > 0 {
		return nil, fmt.Errorf("some credentials information are missing: %s", strings.Join(missingEnvVars, ","))
	}

	return values, nil
}

func GetOneWithFallback[T any](env Env, main string, defaultValue T, fn func(string) (T, error), names ...string) T {
	v, _ := env.getOneWithFallback(main, names...)

	value, err := fn(v)
	if err != nil {
		return defaultValue
	}

	return value
}

func (env Env) getOneWithFallback(main string, names ...string) (string, string) {
	value := env.GetOrFile(main)
	if value != "" {
		return value, main
	}

	for _, name := range names {
		value := env.GetOrFile(name)
		if value != "" {
			return value, main
		}
	}

	return "", main
}

// GetOrDefaultString returns the given environment variable value as a string.
// Returns the default if the env var cannot be found.
func (env Env) GetOrDefaultString(envVar string, defaultValue string) string {
	return getOrDefault(env, envVar, defaultValue, ParseString)
}

// GetOrDefaultBool returns the given environment variable value as a boolean.
// Returns the default if the env var cannot be coopered to a boolean, or is not found.
func (env Env) GetOrDefaultBool(envVar string, defaultValue bool) bool {
	return getOrDefault(env, envVar, defaultValue, strconv.ParseBool)
}

// GetOrDefaultInt returns the given environment variable value as an integer.
// Returns the default if the env var cannot be coopered to an int, or is not found.
func (env Env) GetOrDefaultInt(envVar string, defaultValue int) int {
	return getOrDefault(env, envVar, defaultValue, strconv.Atoi)
}

// GetOrDefaultSecond returns the given environment variable value as a time.Duration (second).
// Returns the default if the env var cannot be coopered to an int, or is not found.
func (env Env) GetOrDefaultSecond(envVar string, defaultValue time.Duration) time.Duration {
	return getOrDefault(env, envVar, defaultValue, ParseSecond)
}

func getOrDefault[T any](env Env, envVar string, defaultValue T, fn func(string) (T, error)) T {
	v, err := fn(env.GetOrFile(envVar))
	if err != nil {
		return defaultValue
	}

	return v
}

// GetOrFile Attempts to resolve 'key' as an environment variable.
// Returns the value of the environment variable if it exists.
// The original logic in gen is as follows:
// Failing that, it will check to see if '<key>_FILE' exists.
// If so, it will attempt to read from the referenced file to populate a value.
func (env Env) GetOrFile(envVar string) string {
	return env.data[envVar]
}

// ParseSecond parses env var value (string) to a second (time.Duration).
func ParseSecond(s string) (time.Duration, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	if v < 0 {
		return 0, fmt.Errorf("unsupported value: %d", v)
	}

	return time.Duration(v) * time.Second, nil
}

// ParseString parses env var value (string) to a string but throws an error when the string is empty.
func ParseString(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string")
	}

	return s, nil
}
