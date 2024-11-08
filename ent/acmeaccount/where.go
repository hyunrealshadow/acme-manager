// Code generated by ent, DO NOT EDIT.

package acmeaccount

import (
	"acme-manager/ent/predicate"
	"acme-manager/ent/schema/enum"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLTE(FieldID, id))
}

// AcmeServerID applies equality check predicate on the "acme_server_id" field. It's identical to AcmeServerIDEQ.
func AcmeServerID(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldAcmeServerID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldDescription, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldEmail, v))
}

// PrivateKey applies equality check predicate on the "private_key" field. It's identical to PrivateKeyEQ.
func PrivateKey(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldPrivateKey, v))
}

// KeyFingerprint applies equality check predicate on the "key_fingerprint" field. It's identical to KeyFingerprintEQ.
func KeyFingerprint(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldKeyFingerprint, v))
}

// EabKeyID applies equality check predicate on the "eab_key_id" field. It's identical to EabKeyIDEQ.
func EabKeyID(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldEabKeyID, v))
}

// EabHmacKey applies equality check predicate on the "eab_hmac_key" field. It's identical to EabHmacKeyEQ.
func EabHmacKey(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldEabHmacKey, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldUpdatedBy, v))
}

// AcmeServerIDEQ applies the EQ predicate on the "acme_server_id" field.
func AcmeServerIDEQ(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldAcmeServerID, v))
}

// AcmeServerIDNEQ applies the NEQ predicate on the "acme_server_id" field.
func AcmeServerIDNEQ(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNEQ(FieldAcmeServerID, v))
}

// AcmeServerIDIn applies the In predicate on the "acme_server_id" field.
func AcmeServerIDIn(vs ...uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIn(FieldAcmeServerID, vs...))
}

// AcmeServerIDNotIn applies the NotIn predicate on the "acme_server_id" field.
func AcmeServerIDNotIn(vs ...uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotIn(FieldAcmeServerID, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContainsFold(FieldDescription, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContainsFold(FieldEmail, v))
}

// KeyTypeEQ applies the EQ predicate on the "key_type" field.
func KeyTypeEQ(v enum.KeyType) predicate.AcmeAccount {
	vc := v
	return predicate.AcmeAccount(sql.FieldEQ(FieldKeyType, vc))
}

// KeyTypeNEQ applies the NEQ predicate on the "key_type" field.
func KeyTypeNEQ(v enum.KeyType) predicate.AcmeAccount {
	vc := v
	return predicate.AcmeAccount(sql.FieldNEQ(FieldKeyType, vc))
}

// KeyTypeIn applies the In predicate on the "key_type" field.
func KeyTypeIn(vs ...enum.KeyType) predicate.AcmeAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AcmeAccount(sql.FieldIn(FieldKeyType, v...))
}

// KeyTypeNotIn applies the NotIn predicate on the "key_type" field.
func KeyTypeNotIn(vs ...enum.KeyType) predicate.AcmeAccount {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AcmeAccount(sql.FieldNotIn(FieldKeyType, v...))
}

// PrivateKeyEQ applies the EQ predicate on the "private_key" field.
func PrivateKeyEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldPrivateKey, v))
}

// PrivateKeyNEQ applies the NEQ predicate on the "private_key" field.
func PrivateKeyNEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNEQ(FieldPrivateKey, v))
}

// PrivateKeyIn applies the In predicate on the "private_key" field.
func PrivateKeyIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIn(FieldPrivateKey, vs...))
}

// PrivateKeyNotIn applies the NotIn predicate on the "private_key" field.
func PrivateKeyNotIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotIn(FieldPrivateKey, vs...))
}

// PrivateKeyGT applies the GT predicate on the "private_key" field.
func PrivateKeyGT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGT(FieldPrivateKey, v))
}

// PrivateKeyGTE applies the GTE predicate on the "private_key" field.
func PrivateKeyGTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGTE(FieldPrivateKey, v))
}

// PrivateKeyLT applies the LT predicate on the "private_key" field.
func PrivateKeyLT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLT(FieldPrivateKey, v))
}

// PrivateKeyLTE applies the LTE predicate on the "private_key" field.
func PrivateKeyLTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLTE(FieldPrivateKey, v))
}

// PrivateKeyContains applies the Contains predicate on the "private_key" field.
func PrivateKeyContains(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContains(FieldPrivateKey, v))
}

// PrivateKeyHasPrefix applies the HasPrefix predicate on the "private_key" field.
func PrivateKeyHasPrefix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasPrefix(FieldPrivateKey, v))
}

// PrivateKeyHasSuffix applies the HasSuffix predicate on the "private_key" field.
func PrivateKeyHasSuffix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasSuffix(FieldPrivateKey, v))
}

// PrivateKeyEqualFold applies the EqualFold predicate on the "private_key" field.
func PrivateKeyEqualFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEqualFold(FieldPrivateKey, v))
}

// PrivateKeyContainsFold applies the ContainsFold predicate on the "private_key" field.
func PrivateKeyContainsFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContainsFold(FieldPrivateKey, v))
}

// KeyFingerprintEQ applies the EQ predicate on the "key_fingerprint" field.
func KeyFingerprintEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldKeyFingerprint, v))
}

// KeyFingerprintNEQ applies the NEQ predicate on the "key_fingerprint" field.
func KeyFingerprintNEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNEQ(FieldKeyFingerprint, v))
}

// KeyFingerprintIn applies the In predicate on the "key_fingerprint" field.
func KeyFingerprintIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIn(FieldKeyFingerprint, vs...))
}

// KeyFingerprintNotIn applies the NotIn predicate on the "key_fingerprint" field.
func KeyFingerprintNotIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotIn(FieldKeyFingerprint, vs...))
}

// KeyFingerprintGT applies the GT predicate on the "key_fingerprint" field.
func KeyFingerprintGT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGT(FieldKeyFingerprint, v))
}

// KeyFingerprintGTE applies the GTE predicate on the "key_fingerprint" field.
func KeyFingerprintGTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGTE(FieldKeyFingerprint, v))
}

// KeyFingerprintLT applies the LT predicate on the "key_fingerprint" field.
func KeyFingerprintLT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLT(FieldKeyFingerprint, v))
}

// KeyFingerprintLTE applies the LTE predicate on the "key_fingerprint" field.
func KeyFingerprintLTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLTE(FieldKeyFingerprint, v))
}

// KeyFingerprintContains applies the Contains predicate on the "key_fingerprint" field.
func KeyFingerprintContains(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContains(FieldKeyFingerprint, v))
}

// KeyFingerprintHasPrefix applies the HasPrefix predicate on the "key_fingerprint" field.
func KeyFingerprintHasPrefix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasPrefix(FieldKeyFingerprint, v))
}

// KeyFingerprintHasSuffix applies the HasSuffix predicate on the "key_fingerprint" field.
func KeyFingerprintHasSuffix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasSuffix(FieldKeyFingerprint, v))
}

// KeyFingerprintEqualFold applies the EqualFold predicate on the "key_fingerprint" field.
func KeyFingerprintEqualFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEqualFold(FieldKeyFingerprint, v))
}

// KeyFingerprintContainsFold applies the ContainsFold predicate on the "key_fingerprint" field.
func KeyFingerprintContainsFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContainsFold(FieldKeyFingerprint, v))
}

// EabKeyIDEQ applies the EQ predicate on the "eab_key_id" field.
func EabKeyIDEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldEabKeyID, v))
}

// EabKeyIDNEQ applies the NEQ predicate on the "eab_key_id" field.
func EabKeyIDNEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNEQ(FieldEabKeyID, v))
}

// EabKeyIDIn applies the In predicate on the "eab_key_id" field.
func EabKeyIDIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIn(FieldEabKeyID, vs...))
}

// EabKeyIDNotIn applies the NotIn predicate on the "eab_key_id" field.
func EabKeyIDNotIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotIn(FieldEabKeyID, vs...))
}

// EabKeyIDGT applies the GT predicate on the "eab_key_id" field.
func EabKeyIDGT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGT(FieldEabKeyID, v))
}

// EabKeyIDGTE applies the GTE predicate on the "eab_key_id" field.
func EabKeyIDGTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGTE(FieldEabKeyID, v))
}

// EabKeyIDLT applies the LT predicate on the "eab_key_id" field.
func EabKeyIDLT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLT(FieldEabKeyID, v))
}

// EabKeyIDLTE applies the LTE predicate on the "eab_key_id" field.
func EabKeyIDLTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLTE(FieldEabKeyID, v))
}

// EabKeyIDContains applies the Contains predicate on the "eab_key_id" field.
func EabKeyIDContains(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContains(FieldEabKeyID, v))
}

// EabKeyIDHasPrefix applies the HasPrefix predicate on the "eab_key_id" field.
func EabKeyIDHasPrefix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasPrefix(FieldEabKeyID, v))
}

// EabKeyIDHasSuffix applies the HasSuffix predicate on the "eab_key_id" field.
func EabKeyIDHasSuffix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasSuffix(FieldEabKeyID, v))
}

// EabKeyIDIsNil applies the IsNil predicate on the "eab_key_id" field.
func EabKeyIDIsNil() predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIsNull(FieldEabKeyID))
}

// EabKeyIDNotNil applies the NotNil predicate on the "eab_key_id" field.
func EabKeyIDNotNil() predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotNull(FieldEabKeyID))
}

// EabKeyIDEqualFold applies the EqualFold predicate on the "eab_key_id" field.
func EabKeyIDEqualFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEqualFold(FieldEabKeyID, v))
}

// EabKeyIDContainsFold applies the ContainsFold predicate on the "eab_key_id" field.
func EabKeyIDContainsFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContainsFold(FieldEabKeyID, v))
}

// EabHmacKeyEQ applies the EQ predicate on the "eab_hmac_key" field.
func EabHmacKeyEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldEabHmacKey, v))
}

// EabHmacKeyNEQ applies the NEQ predicate on the "eab_hmac_key" field.
func EabHmacKeyNEQ(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNEQ(FieldEabHmacKey, v))
}

// EabHmacKeyIn applies the In predicate on the "eab_hmac_key" field.
func EabHmacKeyIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIn(FieldEabHmacKey, vs...))
}

// EabHmacKeyNotIn applies the NotIn predicate on the "eab_hmac_key" field.
func EabHmacKeyNotIn(vs ...string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotIn(FieldEabHmacKey, vs...))
}

// EabHmacKeyGT applies the GT predicate on the "eab_hmac_key" field.
func EabHmacKeyGT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGT(FieldEabHmacKey, v))
}

// EabHmacKeyGTE applies the GTE predicate on the "eab_hmac_key" field.
func EabHmacKeyGTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGTE(FieldEabHmacKey, v))
}

// EabHmacKeyLT applies the LT predicate on the "eab_hmac_key" field.
func EabHmacKeyLT(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLT(FieldEabHmacKey, v))
}

// EabHmacKeyLTE applies the LTE predicate on the "eab_hmac_key" field.
func EabHmacKeyLTE(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLTE(FieldEabHmacKey, v))
}

// EabHmacKeyContains applies the Contains predicate on the "eab_hmac_key" field.
func EabHmacKeyContains(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContains(FieldEabHmacKey, v))
}

// EabHmacKeyHasPrefix applies the HasPrefix predicate on the "eab_hmac_key" field.
func EabHmacKeyHasPrefix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasPrefix(FieldEabHmacKey, v))
}

// EabHmacKeyHasSuffix applies the HasSuffix predicate on the "eab_hmac_key" field.
func EabHmacKeyHasSuffix(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldHasSuffix(FieldEabHmacKey, v))
}

// EabHmacKeyIsNil applies the IsNil predicate on the "eab_hmac_key" field.
func EabHmacKeyIsNil() predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIsNull(FieldEabHmacKey))
}

// EabHmacKeyNotNil applies the NotNil predicate on the "eab_hmac_key" field.
func EabHmacKeyNotNil() predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotNull(FieldEabHmacKey))
}

// EabHmacKeyEqualFold applies the EqualFold predicate on the "eab_hmac_key" field.
func EabHmacKeyEqualFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEqualFold(FieldEabHmacKey, v))
}

// EabHmacKeyContainsFold applies the ContainsFold predicate on the "eab_hmac_key" field.
func EabHmacKeyContainsFold(v string) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldContainsFold(FieldEabHmacKey, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotNull(FieldUpdatedAt))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v uuid.UUID) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.FieldNotNull(FieldUpdatedBy))
}

// HasCertificates applies the HasEdge predicate on the "certificates" edge.
func HasCertificates() predicate.AcmeAccount {
	return predicate.AcmeAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CertificatesTable, CertificatesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCertificatesWith applies the HasEdge predicate on the "certificates" edge with a given conditions (other predicates).
func HasCertificatesWith(preds ...predicate.Certificate) predicate.AcmeAccount {
	return predicate.AcmeAccount(func(s *sql.Selector) {
		step := newCertificatesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAcmeServer applies the HasEdge predicate on the "acme_server" edge.
func HasAcmeServer() predicate.AcmeAccount {
	return predicate.AcmeAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AcmeServerTable, AcmeServerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAcmeServerWith applies the HasEdge predicate on the "acme_server" edge with a given conditions (other predicates).
func HasAcmeServerWith(preds ...predicate.AcmeServer) predicate.AcmeAccount {
	return predicate.AcmeAccount(func(s *sql.Selector) {
		step := newAcmeServerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AcmeAccount) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AcmeAccount) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AcmeAccount) predicate.AcmeAccount {
	return predicate.AcmeAccount(sql.NotPredicates(p))
}
