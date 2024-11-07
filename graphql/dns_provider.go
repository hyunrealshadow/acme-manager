package graphql

import (
	"acme-manager/acme/dns"
	"acme-manager/acme/lego"
	"acme-manager/ent"
	"encoding/json"
	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
)

func prepareDnsProviderConfig(typ string, config string) (*lego.DnsProviderConfig, error) {
	conf := &lego.DnsProviderConfig{}
	err := json.Unmarshal([]byte(config), conf)
	if err != nil {
		return nil, errors.Wrap(err, "dns provider config unmarshal failed")
	}
	_, err = dns.NewDnsProvider(typ, conf)
	if err != nil {
		return nil, errors.Wrap(err, "dns provider create failed")
	}
	err = conf.Encrypt()
	if err != nil {
		return nil, errors.Wrap(err, "dns provider config encrypt failed")
	}
	return conf, nil
}

func sensitiveDnsProvider(dnsProvider *ent.DnsProvider) {
	err := dnsProvider.Config.Decrypt()
	if err != nil {
		logrus.Errorf("DNS provider decryption failed: %v", err)
	}
	dnsProvider.Config.Sensitive()
}
