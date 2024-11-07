// Code from https://github.com/go-acme/lego/tree/v4.19.2/providers/dns/gcloud
// License: MIT

package gcloud

import (
	"acme-manager/acme/lego"
	"cloud.google.com/go/compute/metadata"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-acme/lego/v4/providers/dns/gcloud"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/dns/v1"
)

// NewDNSProviderCredentials uses the supplied credentials
// to return a DNSProvider instance configured for Google Cloud DNS.
func NewDNSProviderCredentials(project string, env lego.Env) (*gcloud.DNSProvider, error) {
	if project == "" {
		return nil, errors.New("googlecloud: project name missing")
	}

	client, err := google.DefaultClient(context.Background(), dns.NdevClouddnsReadwriteScope)
	if err != nil {
		return nil, fmt.Errorf("googlecloud: unable to get Google Cloud client: %w", err)
	}

	config := NewConfig(env)
	config.Project = project
	config.HTTPClient = client

	return gcloud.NewDNSProviderConfig(config)
}

// NewDNSProviderServiceAccountKey uses the supplied service account JSON
// to return a DNSProvider instance configured for Google Cloud DNS.
func NewDNSProviderServiceAccountKey(saKey []byte, env lego.Env) (*gcloud.DNSProvider, error) {
	if len(saKey) == 0 {
		return nil, errors.New("googlecloud: Service Account is missing")
	}

	// If GCE_PROJECT is non-empty it overrides the project in the service
	// account file.
	project := env.GetOrDefaultString(EnvProject, "")
	if project == "" {
		// read project id from service account file
		var datJSON struct {
			ProjectID string `json:"project_id"`
		}
		err := json.Unmarshal(saKey, &datJSON)
		if err != nil || datJSON.ProjectID == "" {
			return nil, errors.New("googlecloud: project ID not found in Google Cloud Service Account file")
		}
		project = datJSON.ProjectID
	}

	conf, err := google.JWTConfigFromJSON(saKey, dns.NdevClouddnsReadwriteScope)
	if err != nil {
		return nil, fmt.Errorf("googlecloud: unable to acquire config: %w", err)
	}
	client := conf.Client(context.Background())

	config := gcloud.NewDefaultConfig()
	config.Project = project
	config.HTTPClient = client

	return gcloud.NewDNSProviderConfig(config)
}

func autodetectProjectID(ctx context.Context) string {
	if pid, err := metadata.ProjectIDWithContext(ctx); err == nil {
		return pid
	}

	return ""
}
