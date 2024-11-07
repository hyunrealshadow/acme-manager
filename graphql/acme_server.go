package graphql

import (
	"acme-manager/config"
	"acme-manager/ent"
	"acme-manager/graphql/model"
	"context"
	"github.com/cockroachdb/errors"
	"github.com/go-acme/lego/v4/acme/api"
	"net/http"
	"time"
)

func loadAcmeServerInfo(url string) (*api.Core, error) {
	core, err := api.New(http.DefaultClient, config.LegoUserAgent, url, "", nil)
	if err != nil {
		return nil, errors.New("could not load ACME directory information")
	}
	return core, nil
}

func (r *mutationResolver) updateAcmeServer(ctx context.Context, input model.UpdateAcmeServerInput, exist *ent.AcmeServer) (*ent.AcmeServer, error) {
	if exist.BuiltIn {
		return nil, errors.Errorf("cannot update built-in ACME server")
	}
	update := exist.Update().
		SetName(input.Name).
		SetUpdatedAt(time.Now()).
		SetUpdatedBy(config.SystemUserID)
	if input.Description != nil {
		update.SetDescription(*input.Description)
	} else {
		update.ClearDescription()
	}
	return update.Save(ctx)
}
