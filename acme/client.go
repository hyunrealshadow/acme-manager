package acme

import (
	"acme-manager/config"
	"github.com/cockroachdb/errors"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

type ClientConfig struct {
	CAUrl                   string
	User                    User
	ExternalAccountRequired bool
	EabKeyID                *string
	EabHmacKey              *string
}

type Client struct {
	config ClientConfig
	lego   *lego.Client
}

func NewClient(clientConfig ClientConfig) (*Client, error) {
	legoConfig := lego.NewConfig(&clientConfig.User)
	legoConfig.CADirURL = clientConfig.CAUrl
	legoConfig.UserAgent = config.LegoUserAgent
	if clientConfig.ExternalAccountRequired {
		if clientConfig.EabKeyID == nil || clientConfig.EabHmacKey == nil {
			return nil, errors.New("the EabKeyID and the EabHmacKey are required for External Account Binding")
		}
	}

	client, err := lego.NewClient(legoConfig)
	if err != nil {
		return nil, err
	}
	return &Client{
		config: clientConfig,
		lego:   client,
	}, nil
}

func (c *Client) Register() (*registration.Resource, error) {
	if c.config.ExternalAccountRequired {
		return c.lego.Registration.RegisterWithExternalAccountBinding(
			registration.RegisterEABOptions{
				TermsOfServiceAgreed: true,
				Kid:                  *c.config.EabKeyID,
				HmacEncoded:          *c.config.EabHmacKey,
			})
	} else {
		return c.lego.Registration.Register(registration.RegisterOptions{
			TermsOfServiceAgreed: true,
		})
	}
}
