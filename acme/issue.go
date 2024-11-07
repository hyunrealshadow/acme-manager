package acme

import (
	"acme-manager/acme/dns"
	legoi "acme-manager/acme/lego"
	"acme-manager/config"
	"acme-manager/ent/schema/enum"
	"acme-manager/util"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"github.com/cockroachdb/errors"
	"github.com/go-acme/lego/v4/certcrypto"
	acmecert "github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"github.com/mholt/acmez/v2"
	"github.com/mholt/acmez/v2/acme"
	"net/http"
	"strings"
	"time"
)

type IssueCertificateRequest struct {
	AccountEmail           string
	AccountPrivateKey      string
	AccountRegistration    registration.Resource
	AcmeServerURL          string
	KeyType                enum.KeyType
	DnsProviderType        string
	DnsProviderConfig      legoi.DnsProviderConfig
	CommonName             string
	SubjectAlternativeName []string
	Organization           *string
	OrganizationalUnit     *string
	Country                *string
	State                  *string
	Locality               *string
	StreetAddress          *string
}

type IssueCertificateResponse struct {
	IssuedAt         time.Time
	ExpiresAt        time.Time
	Certificate      string
	Fingerprint      string
	CSR              string
	PrivateKey       string
	CertificateChain []string
}

func IssueCertificate(
	request IssueCertificateRequest,
) (*IssueCertificateResponse, error) {
	// Create a new ACME client
	keyType, err := request.KeyType.LegoCertCryptoKeyType()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	accountPrivateKey, err := util.Pkcs8PemToPrivateKey(request.AccountPrivateKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	user := NewRegisteredUser(request.AccountEmail, accountPrivateKey, request.AccountRegistration)
	legoConfig := lego.NewConfig(&user)
	legoConfig.CADirURL = request.AcmeServerURL
	legoConfig.Certificate.KeyType = keyType
	legoConfig.UserAgent = config.LegoUserAgent
	client, err := lego.NewClient(legoConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dns01, err := dns.NewDnsProvider(request.DnsProviderType, &request.DnsProviderConfig)
	if err != nil {
		return nil, errors.WithStack(err)

	}
	err = client.Challenge.SetDNS01Provider(dns01)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Issue the certificate
	csr, privateKey, err := createCertificateRequest(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	privateKeyPem, err := util.PrivateKeyToPkcs8Pem(privateKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	acmeRequest := acmecert.ObtainForCSRRequest{
		CSR: csr,
	}
	certificateResource, err := client.Certificate.ObtainForCSR(acmeRequest)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	certificate := string(certificateResource.Certificate)
	issuerCertificate := string(certificateResource.IssuerCertificate)
	x509Cert, err := util.ParseX509Certificate(certificate)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	certSha256 := sha256.Sum256(x509Cert.Raw)
	fingerprint := hex.EncodeToString(certSha256[:])
	csrString := string(certificateResource.CSR)
	certificateChain, err := getCertificateChain(request, accountPrivateKey, issuerCertificate, certificateResource.CertURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &IssueCertificateResponse{
		IssuedAt:         time.Now(),
		ExpiresAt:        x509Cert.NotAfter,
		Certificate:      certificate,
		Fingerprint:      fingerprint,
		CSR:              csrString,
		PrivateKey:       privateKeyPem,
		CertificateChain: certificateChain,
	}, nil
}

func createCertificateRequest(request IssueCertificateRequest) (*x509.CertificateRequest, crypto.PrivateKey, error) {
	subject := pkix.Name{
		CommonName: request.CommonName,
	}
	if request.Organization != nil {
		subject.Organization = []string{*request.Organization}
	}
	if request.OrganizationalUnit != nil {
		subject.OrganizationalUnit = []string{*request.OrganizationalUnit}
	}
	if request.Country != nil {
		subject.Country = []string{*request.Country}
	}
	if request.State != nil {
		subject.Province = []string{*request.State}
	}
	if request.Locality != nil {
		subject.Locality = []string{*request.Locality}
	}
	if request.StreetAddress != nil {
		subject.StreetAddress = []string{*request.StreetAddress}
	}
	signatureAlgorithm := getSignatureAlgorithm(request.KeyType)
	legoKeyType, err := request.KeyType.LegoCertCryptoKeyType()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	privateKey, err := certcrypto.GeneratePrivateKey(legoKeyType)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	csr := &x509.CertificateRequest{
		Subject:            subject,
		SignatureAlgorithm: signatureAlgorithm,
	}
	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, csr, privateKey)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	csr.Raw = csrBytes
	if request.SubjectAlternativeName != nil {
		csr.DNSNames = request.SubjectAlternativeName
	}
	return csr, privateKey, nil
}

func getSignatureAlgorithm(keyType enum.KeyType) x509.SignatureAlgorithm {
	switch keyType {
	case enum.RSA2048:
		return x509.SHA256WithRSA
	case enum.RSA4096:
		return x509.SHA384WithRSA
	case enum.RSA8192:
		return x509.SHA512WithRSA
	case enum.EC256:
		return x509.ECDSAWithSHA256
	case enum.EC384:
		return x509.ECDSAWithSHA384
	default:
		return x509.UnknownSignatureAlgorithm
	}
}

func getCertificateChain(request IssueCertificateRequest, accountPrivateKey crypto.Signer, issuerCertificate string, certificateUrl string) ([]string, error) {
	// We need to use standard acme client to get the certificate, because lego does not provide a way to get the certificate
	client := acmez.Client{
		Client: &acme.Client{
			Directory:  request.AcmeServerURL,
			HTTPClient: &http.Client{},
		},
	}
	reg := request.AccountRegistration
	account := acme.Account{
		Status:               acme.StatusValid,
		Contact:              reg.Body.Contact,
		Orders:               reg.Body.Orders,
		TermsOfServiceAgreed: reg.Body.TermsOfServiceAgreed,
		Location:             reg.URI,
		PrivateKey:           accountPrivateKey,
	}
	ctx := context.Background()
	certificates, err := client.GetCertificateChain(ctx, account, certificateUrl)
	var certificateChain []string
	for i := range certificates {
		currentCertificate := string(certificates[i].ChainPEM)
		if strings.Contains(currentCertificate, issuerCertificate) {
			certificateChain = append(certificateChain, strings.Split(currentCertificate, "\n\n")...)
		}
	}
	if len(certificateChain) != 0 {
		for i := range certificateChain {
			certificateChain[i] = strings.TrimSpace(certificateChain[i])
		}
	}
	return certificateChain, errors.WithStack(err)
}
