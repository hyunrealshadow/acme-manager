package enum

import (
	"github.com/cockroachdb/errors"
	"github.com/go-acme/lego/v4/certcrypto"
	"io"
)

type KeyType string

const (
	RSA2048 KeyType = "RSA2048"
	RSA3072 KeyType = "RSA3072"
	RSA4096 KeyType = "RSA4096"
	RSA8192 KeyType = "RSA8192"
	EC256   KeyType = "EC256"
	EC384   KeyType = "EC384"
)

// Values provides list valid values for Enum.
func (KeyType) Values() (kinds []string) {
	for _, s := range []KeyType{RSA2048, RSA3072, RSA4096, RSA8192, EC256, EC384} {
		kinds = append(kinds, string(s))
	}
	return
}

// UnmarshalGQL Make entgo generated code happy
func (k KeyType) UnmarshalGQL(v any) error {
	panic("implement me")
}

// MarshalGQL Make entgo generated code happy
func (k KeyType) MarshalGQL(w io.Writer) {
	panic("implement me")
}

func (k KeyType) LegoCertCryptoKeyType() (certcrypto.KeyType, error) {
	switch k {
	case RSA2048:
		return certcrypto.RSA2048, nil
	case RSA3072:
		return certcrypto.RSA3072, nil
	case RSA4096:
		return certcrypto.RSA4096, nil
	case RSA8192:
		return certcrypto.RSA8192, nil
	case EC256:
		return certcrypto.EC256, nil
	case EC384:
		return certcrypto.EC384, nil
	default:
		return "", errors.New("invalid key type")
	}
}
