package schema

import (
	"acme-manager/ent/schema/directive"
	"acme-manager/ent/schema/enum"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// Certificate holds the schema definition for the Certificate entity.
type Certificate struct {
	ent.Schema
}

// Annotations of the Certificate.
func (Certificate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "certificate"},
		entgql.RelayConnection(),
	}
}

// Fields of the Certificate.
func (Certificate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			Annotations(entgql.Directives(directive.Model("Certificate"))),
		field.UUID("acme_account_id", uuid.UUID{}).
			Comment("ACME account ID").
			Annotations(
				entgql.Skip(),
				entgql.Directives(directive.Model("AcmeAccount")),
			),
		field.UUID("dns_provider_id", uuid.UUID{}).
			Comment("DNS provider ID").
			Annotations(
				entgql.Skip(),
				entgql.Directives(directive.Model("DnsProvider")),
			),
		field.String("common_name").MaxLen(255).
			Comment("Common name of the certificate"),
		field.Strings("subject_alternative_name").Optional().
			Comment("Subject alternative name of the certificate").
			SchemaType(map[string]string{dialect.Postgres: "varchar[]"}),
		// The following fields may not be compatible with public ACME servers
		field.String("organization").MaxLen(50).Optional().Nillable().
			Comment("Organization name of the certificate"),
		field.String("organizational_unit").MaxLen(100).Optional().Nillable().
			Comment("Organizational unit name of the certificate"),
		field.String("country").MaxLen(2).Optional().Nillable().
			Comment("Country code of the certificate"),
		field.String("state").MaxLen(50).Optional().Nillable().
			Comment("State or province of the certificate"),
		field.String("locality").MaxLen(50).Optional().Nillable().
			Comment("Locality of the certificate"),
		field.String("street_address").MaxLen(255).Optional().Nillable().
			Comment("Street address of the certificate"),
		// End of fields that may not be compatible with public ACME servers
		field.Enum("key_type").
			GoType(enum.KeyType("")).
			Comment("Key type of the certificate").
			SchemaType(map[string]string{dialect.Postgres: "varchar"}).
			Annotations(entgql.Type("KeyType")),
		field.String("csr").Optional().Nillable().
			Comment("Certificate signing request of the certificate").
			Annotations(entgql.Skip()),
		field.String("private_key").Optional().Nillable().
			Comment("Private key of the certificate").
			Annotations(entgql.Skip()),
		field.Text("certificate").Optional().Nillable().
			Comment("Certificate of the certificate").
			Annotations(entgql.Skip()),
		field.Strings("certificate_chain").Optional().Comment("CertificateChain of the certificate").
			Annotations(entgql.Skip()),
		field.String("fingerprint").MaxLen(64).Optional().Nillable().
			Comment("Fingerprint of the certificate"),
		field.Enum("status").
			GoType(enum.Status("")).
			Comment("Status of the certificate").
			SchemaType(map[string]string{dialect.Postgres: "varchar"}).
			Annotations(entgql.Type("Status")),
		field.Time("issued_at").Optional().Nillable().Comment("Time the certificate was issued"),
		field.Time("expires_at").Optional().Nillable().Comment("Time the certificate expires"),
		field.Time("created_at").Immutable().Default(func() time.Time { return time.Now() }).
			Comment("Time the certificate was created").
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.UUID("created_by", uuid.UUID{}).
			Comment("User that created the certificate").
			Annotations(
				entgql.Type("ID"),
				entgql.Directives(directive.Model("User")),
			),
		field.Time("updated_at").Optional().Nillable().Comment("Time the certificate was updated"),
		field.UUID("updated_by", uuid.UUID{}).Optional().Nillable().
			Comment("User that updated the certificate").
			Annotations(
				entgql.Type("ID"),
				entgql.Directives(directive.Model("User")),
			),
	}
}

// Indexes of the Certificate.
func (Certificate) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("common_name"),
		index.Fields("common_name", "acme_account_id", "key_type").Unique(),
		index.Fields("created_at"),
	}
}

// Edges of the Certificate.
func (Certificate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("acme_account", AcmeAccount.Type).Ref("certificates").Unique().
			Field("acme_account_id").Required(),
		edge.From("dns_provider", DnsProvider.Type).Ref("certificates").Unique().
			Field("dns_provider_id").Required(),
	}
}
