package internal

import (
	"encoding/json"
	"github.com/emirpasic/gods/maps/linkedhashmap"
	"strconv"
)

func generateJsonSchema(name string, p *DnsProvider, credentials, additional map[string]string) (string, error) {
	schema := linkedhashmap.New()
	schema.Put("$schema", "http://json-schema.org/draft-07/schema#")
	schema.Put("title", name+" Configuration")
	schema.Put("type", "object")
	properties := linkedhashmap.New()
	credentialsMap := linkedhashmap.New()
	for _, field := range p.CredentialsFields {
		fieldMap := linkedhashmap.New()
		fieldMap.Put("type", "string")
		if description, ok := credentials[field]; ok {
			fieldMap.Put("description", description)
		}
		credentialsMap.Put(field, fieldMap)
	}
	additionalMap := linkedhashmap.New()
	if len(p.AdditionalFields) > 0 {
		for _, field := range p.AdditionalFields {
			fieldMap := linkedhashmap.New()
			fieldMap.Put("type", "string")
			if description, ok := additional[field]; ok {
				fieldMap.Put("description", description)
			}
			additionalMap.Put(field, fieldMap)
		}
	}
	properties.Put("credentials", credentialsMap)
	if additionalMap.Size() > 0 {
		properties.Put("additional", additionalMap)
	}
	schema.Put("properties", properties)

	jsonData, err := json.Marshal(schema)
	if err != nil {
		return "", err
	}
	return strconv.Quote(string(jsonData)), nil
}
