package seed

import (
	"acme-manager/config"
	"acme-manager/ent"
	"acme-manager/ent/acmeserver"
	"context"
	"github.com/cockroachdb/errors"
	"github.com/go-acme/lego/v4/lego"
	"time"
)

type AcmeServerSeeder struct {
	ctx    context.Context
	client ent.Client
}

func (s AcmeServerSeeder) createAcmeServer(name, description, url string) error {
	exist, err := s.client.AcmeServer.Query().Where(acmeserver.URLEQ(url)).Exist(s.ctx)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}
	_, err = s.client.AcmeServer.Create().
		SetName(name).
		SetDescription(description).
		SetURL(url).
		SetBuiltIn(true).
		SetExternalAccountRequired(false).
		SetCreatedAt(time.Now()).
		SetCreatedBy(config.SystemUserID).
		Save(s.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s AcmeServerSeeder) Seed() error {
	err := s.createAcmeServer("Let's Encrypt", "Let's Encrypt Production Server", lego.LEDirectoryProduction)
	if err != nil {
		return errors.Wrap(err, "create let's encrypt production server")
	}
	err = s.createAcmeServer("Let's Encrypt Staging", "Let's Encrypt Staging Server", lego.LEDirectoryStaging)
	if err != nil {
		return errors.Wrap(err, "create let's encrypt staging server")
	}
	return nil
}
