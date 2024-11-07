package internal

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/iancoleman/strcase"
	"strings"
)

const fileEnvSuffix = "_FILE"

var defaultAdditionalFields = mapset.NewSet("ttl", "propagationTimeout", "pollingInterval", "httpTimeout")

func processDnsProviderModel(provider *DnsProvider, m model) {
	provider.DisplayName = m.Name
	provider.Description = m.Description

	providerEnvNamespace := ""
	providerEnvNamespaceConstant := provider.EnvConstants["envNamespace"]
	altProviderEnvNamespace := ""
	altProviderEnvNamespaceConstant := provider.EnvConstants["altEnvNamespace"]
	if namespaceStr, ok := providerEnvNamespaceConstant.(string); ok {
		providerEnvNamespace = namespaceStr[1 : len(namespaceStr)-1]
	}
	if altNamespaceStr, ok := altProviderEnvNamespaceConstant.(string); ok {
		altProviderEnvNamespace = altNamespaceStr[1 : len(altNamespaceStr)-1]
	}
	envConstantsMap := make(map[string]string)
	credentials := make(map[string]string)
	additional := make(map[string]string)
	var credentialsFields []string
	credentialsFieldsSet := mapset.NewSet[string]()
	var additionalFields []string
	additionalFieldsSet := mapset.NewSet[string]()
	// Iterate over the credentials fields
	// if the key has a _FILE suffix, remove it
	// convert the key to camel case
	// add the key to the credentials map
	// envConstantsMap maps the original key to the new key
	for k, v := range m.Configuration.Credentials {
		key := k
		if strings.HasSuffix(key, fileEnvSuffix) {
			key = strings.TrimSuffix(key, fileEnvSuffix)
		}
		filedName := strcase.ToLowerCamel(strings.Replace(key, providerEnvNamespace, "", -1))
		if altProviderEnvNamespace != "" {
			filedName = strcase.ToLowerCamel(strings.Replace(k, altProviderEnvNamespace, "", -1))
		}
		envConstantsMap[k] = filedName
		credentials[filedName] = v
		credentialsFields = append(credentialsFields, filedName)
		credentialsFieldsSet.Add(filedName)
		if providerEnvNamespace == "" {
			strKey := "\"" + k + "\""
			strFiledName := "\"" + filedName + "\""
			provider.NewDefaultConfigFunc = strings.Replace(provider.NewDefaultConfigFunc, strKey, strFiledName, -1)
			provider.NewDNSProviderFunc = strings.Replace(provider.NewDNSProviderFunc, strKey, strFiledName, -1)
		}
	}
	// Iterate over the additional fields
	// convert the key to camel case
	// add the key to the additional map
	// envConstantsMap maps the original key to the new key
	for k, v := range m.Configuration.Additional {
		filedName := strcase.ToLowerCamel(strings.Replace(k, providerEnvNamespace, "", -1))
		if altProviderEnvNamespace != "" {
			filedName = strcase.ToLowerCamel(strings.Replace(k, altProviderEnvNamespace, "", -1))
		}
		envConstantsMap[k] = filedName
		additional[filedName] = v
		additionalFields = append(additionalFields, filedName)
		additionalFieldsSet.Add(filedName)
		if providerEnvNamespace == "" {
			strKey := "\"" + k + "\""
			strFiledName := "\"" + filedName + "\""
			provider.NewDefaultConfigFunc = strings.Replace(provider.NewDefaultConfigFunc, strKey, strFiledName, -1)
			provider.NewDNSProviderFunc = strings.Replace(provider.NewDNSProviderFunc, strKey, strFiledName, -1)
		}
	}

	// Iterate over the env constants
	// if the value is a string, and it is in the envConstantsMap
	// replace the value with the new key
	// if the key not in credentialsFieldsSet and additionalFieldsSet
	// and not in the defaultAdditionalFields
	// add the key to the additionalFields
	// because the toml configuration might not have the key
	var missingFields []string
	if provider.EnvConstants != nil {
		envConstants := make(map[string]any)
		for k, v := range provider.EnvConstants {
			if vStr, ok := v.(string); ok {
				vContent := vStr[1 : len(vStr)-1]
				if newKey, ok := envConstantsMap[vContent]; ok {
					envConstants[k] = "\"" + newKey + "\""
				} else if vContent != providerEnvNamespace && strings.HasPrefix(vContent, providerEnvNamespace) {
					fieldName := strcase.ToLowerCamel(strings.Replace(vContent, providerEnvNamespace, "", -1))
					envConstants[k] = "\"" + fieldName + "\""

					if defaultAdditionalFields.Contains(fieldName) {
						continue
					}
					if !credentialsFieldsSet.Contains(fieldName) && !additionalFieldsSet.Contains(fieldName) {
						missingFields = append(missingFields, fieldName)
					}
				} else {
					envConstants[k] = v
				}
			} else {
				envConstants[k] = v
			}
		}
		additionalFields = append(missingFields, additionalFields...)

		if providerEnvNamespace != "" {
			delete(envConstants, "envNamespace")
		}
		if altProviderEnvNamespace != "" {
			delete(envConstants, "altEnvNamespace")
		}
		provider.EnvConstants = envConstants
	}
	provider.CredentialsFields = credentialsFields
	provider.AdditionalFields = additionalFields
	jsonSchema, err := generateJsonSchema(m.Name, provider, credentials, additional)
	if err != nil {
		logger.Fatalf("could not generate JSON schema: %v", err)
	}
	provider.JsonSchema = jsonSchema
}
