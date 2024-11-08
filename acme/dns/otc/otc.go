// Code generated by acme-manager, DO NOT EDIT.
// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/otc
// License: MIT

package otc

import (
	"acme-manager/acme/lego"
	"acme-manager/logger"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/otc"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	EnvDomainName           = "domainName"
	EnvHTTPTimeout          = "httpTimeout"
	EnvIdentityEndpoint     = "identityEndpoint"
	EnvPassword             = "password"
	EnvPollingInterval      = "pollingInterval"
	EnvProjectName          = "projectName"
	EnvPropagationTimeout   = "propagationTimeout"
	EnvSequenceInterval     = "sequenceInterval"
	EnvTTL                  = "ttl"
	EnvUserName             = "userName"
	defaultIdentityEndpoint = "https://iam.eu-de.otc.t-systems.com:443/v3/auth/tokens"
	minTTL                  = 300
)

const JsonSchema = "{\"$schema\":\"http://json-schema.org/draft-07/schema#\",\"title\":\"Open Telekom Cloud Configuration\",\"type\":\"object\",\"properties\":{\"credentials\":{\"userName\":{\"type\":\"string\",\"description\":\"User name\"},\"password\":{\"type\":\"string\",\"description\":\"Password\"},\"projectName\":{\"type\":\"string\",\"description\":\"Project name\"},\"domainName\":{\"type\":\"string\",\"description\":\"Domain name\"},\"identityEndpoint\":{\"type\":\"string\",\"description\":\"Identity endpoint URL\"}},\"additional\":{\"propagationTimeout\":{\"type\":\"string\",\"description\":\"Maximum waiting time for DNS propagation\"},\"sequenceInterval\":{\"type\":\"string\",\"description\":\"Time between sequential requests\"},\"ttl\":{\"type\":\"string\",\"description\":\"The TTL of the TXT record used for the DNS challenge\"},\"httpTimeout\":{\"type\":\"string\",\"description\":\"API request timeout\"},\"pollingInterval\":{\"type\":\"string\",\"description\":\"Time between DNS propagation check\"}}}}"

var compiledJsonSchema *jsonschema.Schema

var credentialsFields = mapset.NewSet("userName", "password", "projectName", "domainName", "identityEndpoint")
var additionalFields = mapset.NewSet("propagationTimeout", "sequenceInterval", "ttl", "httpTimeout", "pollingInterval")

func init() {
	schema, err := jsonschema.UnmarshalJSON(strings.NewReader(JsonSchema))
	if err != nil {
		logger.Fatalf("Failed to unmarshal otc JSON schema: %v", err)
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", schema); err != nil {
		logger.Fatalf("Failed to add otc schema resource to compiler: %v", err)
	}
	compiledJsonSchema, err = compiler.Compile("schema.json")
	if err != nil {
		logger.Fatalf("Failed to compile otc schema: %v", err)
	}
}

func ParseConfig(conf *lego.DnsProviderConfig) (*lego.Env, error) {
	data := make(map[string]string)
	if conf.Credentials != nil {
		for key, value := range conf.Credentials {
			if credentialsFields.Contains(key) {
				data[key] = value
			} else {
				delete(conf.Credentials, key)
			}
		}
	}
	if conf.Additional != nil {
		for key, value := range conf.Additional {
			if additionalFields.Contains(key) {
				data[key] = value
			} else {
				delete(conf.Additional, key)
			}
		}
	}

	anyData := make(map[string]any)
	for key, value := range data {
		anyData[key] = value
	}
	if err := compiledJsonSchema.Validate(anyData); err != nil {
		return nil, fmt.Errorf("otc: configuration validation failed: %v", err)
	}

	env := lego.NewEnv(data)
	return &env, nil
}

func NewConfig(env lego.Env) *otc.Config {
	return &otc.Config{
		TTL:                env.GetOrDefaultInt(EnvTTL, minTTL),
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, dns01.DefaultPropagationTimeout),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, dns01.DefaultPollingInterval),
		IdentityEndpoint:   env.GetOrDefaultString(EnvIdentityEndpoint, defaultIdentityEndpoint),
		SequenceInterval:   env.GetOrDefaultSecond(EnvSequenceInterval, dns01.DefaultPropagationTimeout),
		HTTPClient: &http.Client{
			Timeout: env.GetOrDefaultSecond(EnvHTTPTimeout, 10*time.Second),
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,

				// Workaround for keep alive bug in otc api
				DisableKeepAlives: true,
			},
		},
	}
}

func NewDNSProvider(env lego.Env) (*otc.DNSProvider, error) {
	values, err := env.Get(EnvDomainName, EnvUserName, EnvPassword, EnvProjectName)
	if err != nil {
		return nil, fmt.Errorf("otc: %w", err)
	}

	config := NewConfig(env)
	config.DomainName = values[EnvDomainName]
	config.UserName = values[EnvUserName]
	config.Password = values[EnvPassword]
	config.ProjectName = values[EnvProjectName]

	return otc.NewDNSProviderConfig(config)
}
