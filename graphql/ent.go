package graphql

import "acme-manager/ent"

func sensitiveNode(node any) {
	switch v := node.(type) {
	case *ent.AcmeAccount:
		sensitiveAcmeAccount(v)
	case *ent.DnsProvider:
		sensitiveDnsProvider(v)
	}
}
