package config

import "github.com/google/uuid"

var (
	SystemUserID = uuid.NewSHA1(uuid.NameSpaceURL, []byte("urn:acme_manager:user:system"))
)
