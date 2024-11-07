// Code from https://github.com/go-acme/lego/tree/v4.17.4/providers/dns/dnshomede
// License: MIT

package dnshomede

import (
	"fmt"
	"strings"
)

func parseCredentials(raw string) (map[string]string, error) {
	credentials := make(map[string]string)

	credStrings := strings.Split(strings.TrimSuffix(raw, ","), ",")
	for _, credPair := range credStrings {
		data := strings.Split(credPair, ":")
		if len(data) != 2 {
			return nil, fmt.Errorf("invalid credential pair: %q", credPair)
		}

		credentials[strings.TrimSpace(data[0])] = strings.TrimSpace(data[1])
	}

	return credentials, nil
}
