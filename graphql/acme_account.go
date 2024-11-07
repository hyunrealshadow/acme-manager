package graphql

import (
	"acme-manager/acme"
	"acme-manager/config"
	"acme-manager/ent"
	"acme-manager/ent/schema/enum"
	"acme-manager/graphql/model"
	"acme-manager/secret"
	"acme-manager/util"
	"context"
	"github.com/cockroachdb/errors"
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/registration"
	"time"
)

type registrationInfo struct {
	Email          string
	Registration   registration.Resource
	PrivateKey     string
	KeyFingerprint string
}

type registerAcmeAccountInput struct {
	KeyType                 enum.KeyType
	Email                   string
	CAUrl                   string
	ExternalAccountRequired bool
	EabKeyID                *string
	EabHmacKey              *string
}

func registerAcmeAccount(input registerAcmeAccountInput) (*registrationInfo, error) {
	legoKeyType, err := input.KeyType.LegoCertCryptoKeyType()
	if err != nil {
		return nil, err
	}
	privateKey, err := certcrypto.GeneratePrivateKey(legoKeyType)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	keyFingerprint, err := util.PrivateKeyFingerprint(privateKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	user := acme.NewUser(input.Email, privateKey)
	clientConfig := acme.ClientConfig{
		CAUrl:                   input.CAUrl,
		User:                    user,
		ExternalAccountRequired: input.ExternalAccountRequired,
		EabKeyID:                input.EabKeyID,
		EabHmacKey:              input.EabHmacKey,
	}

	client, err := acme.NewClient(clientConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resource, err := client.Register()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	pkcs8PrivateKey, err := util.PrivateKeyToPkcs8Pem(user.GetPrivateKey())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	encryptedPrivateKey, err := secret.Get().Encrypt(pkcs8PrivateKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &registrationInfo{
		Email:          user.GetEmail(),
		Registration:   *resource,
		PrivateKey:     encryptedPrivateKey,
		KeyFingerprint: keyFingerprint,
	}, nil
}

func (r *mutationResolver) updateAcmeAccount(ctx context.Context, input model.UpdateAcmeAccountInput,
	exist *ent.AcmeAccount) (*ent.AcmeAccount, error) {

	update := exist.Update().
		SetName(input.Name).
		SetUpdatedAt(time.Now()).
		SetUpdatedBy(config.SystemUserID)
	if input.Description == nil {
		update.ClearDescription()
	} else {
		update.SetDescription(*input.Description)
	}
	return update.Save(ctx)
}

func sensitiveAcmeAccount(acmeAccount *ent.AcmeAccount) {
	if acmeAccount.EabKeyID != nil {
		sensitive := util.MakeSensitive(*acmeAccount.EabKeyID)
		acmeAccount.EabKeyID = &sensitive
	}
	if acmeAccount.EabHmacKey != nil {
		sensitive := util.MakeSensitive(*acmeAccount.EabHmacKey)
		acmeAccount.EabHmacKey = &sensitive
	}
}
