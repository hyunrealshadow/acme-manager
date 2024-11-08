// Code generated by ent, DO NOT EDIT.

package acmeaccount

import (
	"acme-manager/ent/schema/enum"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the acmeaccount type in the database.
	Label = "acme_account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAcmeServerID holds the string denoting the acme_server_id field in the database.
	FieldAcmeServerID = "acme_server_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldKeyType holds the string denoting the key_type field in the database.
	FieldKeyType = "key_type"
	// FieldPrivateKey holds the string denoting the private_key field in the database.
	FieldPrivateKey = "private_key"
	// FieldKeyFingerprint holds the string denoting the key_fingerprint field in the database.
	FieldKeyFingerprint = "key_fingerprint"
	// FieldRegistration holds the string denoting the registration field in the database.
	FieldRegistration = "registration"
	// FieldEabKeyID holds the string denoting the eab_key_id field in the database.
	FieldEabKeyID = "eab_key_id"
	// FieldEabHmacKey holds the string denoting the eab_hmac_key field in the database.
	FieldEabHmacKey = "eab_hmac_key"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// EdgeCertificates holds the string denoting the certificates edge name in mutations.
	EdgeCertificates = "certificates"
	// EdgeAcmeServer holds the string denoting the acme_server edge name in mutations.
	EdgeAcmeServer = "acme_server"
	// Table holds the table name of the acmeaccount in the database.
	Table = "acme_account"
	// CertificatesTable is the table that holds the certificates relation/edge.
	CertificatesTable = "certificate"
	// CertificatesInverseTable is the table name for the Certificate entity.
	// It exists in this package in order to avoid circular dependency with the "certificate" package.
	CertificatesInverseTable = "certificate"
	// CertificatesColumn is the table column denoting the certificates relation/edge.
	CertificatesColumn = "acme_account_id"
	// AcmeServerTable is the table that holds the acme_server relation/edge.
	AcmeServerTable = "acme_account"
	// AcmeServerInverseTable is the table name for the AcmeServer entity.
	// It exists in this package in order to avoid circular dependency with the "acmeserver" package.
	AcmeServerInverseTable = "acme_server"
	// AcmeServerColumn is the table column denoting the acme_server relation/edge.
	AcmeServerColumn = "acme_server_id"
)

// Columns holds all SQL columns for acmeaccount fields.
var Columns = []string{
	FieldID,
	FieldAcmeServerID,
	FieldName,
	FieldDescription,
	FieldEmail,
	FieldKeyType,
	FieldPrivateKey,
	FieldKeyFingerprint,
	FieldRegistration,
	FieldEabKeyID,
	FieldEabHmacKey,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// KeyFingerprintValidator is a validator for the "key_fingerprint" field. It is called by the builders before save.
	KeyFingerprintValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// KeyTypeValidator is a validator for the "key_type" field enum values. It is called by the builders before save.
func KeyTypeValidator(kt enum.KeyType) error {
	switch kt {
	case "RSA2048", "RSA3072", "RSA4096", "RSA8192", "EC256", "EC384":
		return nil
	default:
		return fmt.Errorf("acmeaccount: invalid enum value for key_type field: %q", kt)
	}
}

// OrderOption defines the ordering options for the AcmeAccount queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAcmeServerID orders the results by the acme_server_id field.
func ByAcmeServerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAcmeServerID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByKeyType orders the results by the key_type field.
func ByKeyType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKeyType, opts...).ToFunc()
}

// ByPrivateKey orders the results by the private_key field.
func ByPrivateKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrivateKey, opts...).ToFunc()
}

// ByKeyFingerprint orders the results by the key_fingerprint field.
func ByKeyFingerprint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKeyFingerprint, opts...).ToFunc()
}

// ByEabKeyID orders the results by the eab_key_id field.
func ByEabKeyID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEabKeyID, opts...).ToFunc()
}

// ByEabHmacKey orders the results by the eab_hmac_key field.
func ByEabHmacKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEabHmacKey, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByCertificatesCount orders the results by certificates count.
func ByCertificatesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCertificatesStep(), opts...)
	}
}

// ByCertificates orders the results by certificates terms.
func ByCertificates(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCertificatesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAcmeServerField orders the results by acme_server field.
func ByAcmeServerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAcmeServerStep(), sql.OrderByField(field, opts...))
	}
}
func newCertificatesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CertificatesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CertificatesTable, CertificatesColumn),
	)
}
func newAcmeServerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AcmeServerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AcmeServerTable, AcmeServerColumn),
	)
}

var (
	// enum.KeyType must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enum.KeyType)(nil)
	// enum.KeyType must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enum.KeyType)(nil)
)
