package plesk

import (
	"acme-manager/util"
	"github.com/go-acme/lego/v4/providers/dns/plesk"
)

func setBaseUrl(config *plesk.Config, baseUrl string) error {
	return util.SetStructPtrUnExportedStrField(config, "baseUrl", baseUrl)
}
