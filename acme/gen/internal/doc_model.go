// Code from https://github.com/go-acme/lego/blob/v4.17.4/internal/dnsdocs/generator.go
// License: MIT

package internal

type model struct {
	Name          string        // Real name of the DNS provider
	Code          string        // DNS code
	Since         string        // First gen version
	URL           string        // DNS provider URL
	Description   string        // Provider summary
	Example       string        // CLI example
	Configuration configuration // Environment variables
	Links         links         // Links
	Additional    string        // Extra documentation
	GeneratedFrom string        // Source file
}

type configuration struct {
	Credentials map[string]string
	Additional  map[string]string
}

type links struct {
	API      string
	GoClient string
}
