package internal

import "strings"

var customFunctionReplacements = map[string]map[string]string{
	"gcloud": {
		"NewDNSProviderServiceAccountKey([]byte(saKey))": "NewDNSProviderServiceAccountKey([]byte(saKey), env)",
		"NewDNSProviderCredentials(project)":             "NewDNSProviderCredentials(project, env)",
	},
	"plesk": {
		"config.baseURL = values[EnvServerBaseURL]": `
err = setBaseUrl(config, values[EnvServerBaseURL])
if err != nil {
	return nil, fmt.Errorf("plesk: %w", err)
}
`,
	},
	"oraclecloud": {
		"newConfigProvider(": "newConfigProvider(env, ",
	},
	"ovh": {
		"ovh.DefaultTimeout": "DefaultTimeout",
		"&OAuth2Config":      "&ovh.OAuth2Config",
	},
	"selectel": {
		"selectel.DefaultSelectelBaseURL": "DefaultSelectelBaseURL",
	},
	"selectelv2": {
		"selectel.DefaultSelectelBaseURL": "DefaultSelectelBaseURL",
	},
	"vscale": {
		"selectel.DefaultVScaleBaseURL": "DefaultVScaleBaseURL",
	},
}

var customNewDNSProviderCode = map[string]string{
	"joker": `
func NewDNSProvider(env lego.Env) (challenge.ProviderTimeout, error) {
	config := NewConfig(env)
	return joker.NewDNSProviderConfig(config)
}`,
}

func customReplace(providerName, code string) string {
	if replacements, ok := customFunctionReplacements[providerName]; ok {
		for oldCode, newCode := range replacements {
			code = strings.Replace(code, oldCode, newCode, -1)
		}
	}
	return code

}

func generateNewDefaultConfigFunc(provider *DnsProvider) {
	code := provider.NewDefaultConfigFunc
	code = strings.Replace(code, "NewDefaultConfig(", "NewConfig(env lego.Env", -1)
	code = strings.Replace(code, "*Config", "*"+provider.Name+".Config", -1)
	code = strings.Replace(code, "&Config{", "&"+provider.Name+".Config{", -1)
	code = strings.Replace(code, "internal.", "", -1)
	if strings.Contains(code, "GetOneWithFallback") {
		code = strings.Replace(code, "env.ParseSecond", "lego.ParseSecond", -1)
		code = strings.Replace(code, "env.ParseString", "lego.ParseString", -1)
		code = strings.Replace(code, "env.GetOneWithFallback(", "lego.GetOneWithFallback(e, ", -1)
		code = strings.Replace(code, "env lego.Env", "e lego.Env", -1)
		code = strings.Replace(code, "env.", "e", -1)
	}

	code = customReplace(provider.Name, code)

	provider.NewDefaultConfigFunc = code
}

func generateNewDNSProviderFunc(provider *DnsProvider) {
	if funcCode, ok := customNewDNSProviderCode[provider.Name]; ok {
		provider.NewDNSProviderFunc = funcCode
		return
	}

	code := provider.NewDNSProviderFunc
	code = strings.Replace(code, "NewDNSProvider(", "NewDNSProvider(env lego.Env", -1)
	code = strings.Replace(code, "*DNSProvider", "*"+provider.Name+".DNSProvider", -1)
	code = strings.Replace(code, "NewDefaultConfig()", "NewConfig(env)", -1)
	code = strings.Replace(code, "NewDNSProviderConfig", provider.Name+".NewDNSProviderConfig", -1)
	code = strings.Replace(code, "internal.", "", -1)
	if strings.Contains(code, "GetOneWithFallback") {
		code = strings.Replace(code, "env.ParseSecond", "lego.ParseSecond", -1)
		code = strings.Replace(code, "env.ParseString", "lego.ParseString", -1)
		code = strings.Replace(code, "env.GetOneWithFallback(", "lego.GetOneWithFallback(e, ", -1)
		code = strings.Replace(code, "env lego.Env", "e lego.Env", -1)
		code = strings.Replace(code, "env.", "e.", -1)
		code = strings.Replace(code, "NewConfig(env)", "NewConfig(e)", -1)
	}

	code = customReplace(provider.Name, code)

	provider.NewDNSProviderFunc = code
}
