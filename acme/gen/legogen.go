//go:build ignore

package main

import (
	"acme-manager/acme/gen/internal"
	"acme-manager/logger"
	"bytes"
	_ "embed"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"os"
	"strings"
	"text/template"
)

const generateDir = "./acme/dns"

var unsupportedDnsProviders = mapset.NewSet(
	"acmedns",   // This package is missing a config struct, preventing auto-generation
	"edgedns",   // This package must use environment variables for configuration, which is currently not supported
	"exec",      // This package requires a command for configuration, which is not supported
	"hyperone",  // This package must use config file for configuration, which is currently not supported
	"designate", // This package must use environment variables for configuration, which is currently not supported
)

//go:embed internal/dns_provider.tmpl
var dnsProviderTemplate string

//go:embed internal/create.tmpl
var createTemplate string

func generateDnsProvider(provider *internal.DnsProvider) {
	internal.EnsurePathExists(generateDir)
	internal.EnsurePathExists(generateDir + "/" + provider.Name)

	funcMap := template.FuncMap{
		"join": func(array []string, sep string) string {
			strArray := make([]string, len(array))
			for i, v := range array {
				strArray[i] = "\"" + v + "\""
			}
			return strings.Join(strArray, sep)
		},
	}

	tmpl, err := template.New("dns_provider").Funcs(funcMap).Parse(dnsProviderTemplate)
	if err != nil {
		logger.Fatalf("could not parse template: %v", err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, provider)
	if err != nil {
		logger.Fatalf("could not execute template: %v", err)
	}
	filePath := fmt.Sprintf("%s/%s/%s.go", generateDir, provider.Name, provider.Name)
	source := buf.Bytes()
	formattedSource, err := internal.FormatSource(provider.Name, source)
	if err != nil {
		logger.Fatalf("could not format source: %v\n%s", err, source)
	}
	err = os.WriteFile(filePath, formattedSource, 0644)
	if err != nil {
		logger.Fatalf("could not write file %s: %v", filePath, err)
	}
}

func main() {
	dnsProviderPackages := internal.GetAllDnsProviderPackages()
	dnsProviders := make([]*internal.DnsProvider, 0)
	for _, pkg := range dnsProviderPackages {
		logger.Infof("Processing package %s", pkg)
		dnsProvider, err := internal.GetDnsProvider(pkg)
		if err != nil {
			logger.Warnf("could not get DNS provider info for %s", pkg)
			continue
		}
		if unsupportedDnsProviders.Contains(dnsProvider.Name) {
			continue
		}
		generateDnsProvider(dnsProvider)
		dnsProviders = append(dnsProviders, dnsProvider)
	}

	data := map[string]interface{}{
		"Providers": dnsProviders,
	}
	tmpl, err := template.New("create").Parse(createTemplate)
	if err != nil {
		logger.Fatalf("could not parse template: %v", err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		logger.Fatalf("could not execute template: %v", err)
	}
	filePath := fmt.Sprintf("%s/provider.go", generateDir)
	source := buf.Bytes()
	formattedSource, err := internal.FormatSource("create.go", source)
	if err != nil {
		logger.Fatalf("could not format source: %v\n%s", err, source)
	}
	err = os.WriteFile(filePath, formattedSource, 0644)
	if err != nil {
		logger.Fatalf("could not write file %s: %v", filePath, err)
	}

	logger.Infof("Generated %d DNS providers", len(dnsProviders))
}
