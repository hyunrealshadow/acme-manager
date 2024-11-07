package schema

import (
	"acme-manager/ent/schema/directive"
	"acme-manager/ent/schema/enum"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/go-acme/lego/v4/registration"
	"github.com/google/uuid"
	"time"
)

// AcmeAccount holds the schema definition for the AcmeAccount entity.
type AcmeAccount struct {
	ent.Schema
}

// Annotations of the AcmeAccount.
func (AcmeAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "acme_account"},
		entgql.RelayConnection(),
	}
}

// Fields of the AcmeAccount.
func (AcmeAccount) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			Annotations(entgql.Directives(directive.Model("AcmeAccount"))),
		field.UUID("acme_server_id", uuid.UUID{}).
			Comment("ACME server ID").
			Annotations(
				entgql.Skip(),
				entgql.Directives(directive.Model("AcmeServer")),
			),
		field.String("name").MaxLen(50).
			Comment("Name of the ACME account"),
		field.String("description").MaxLen(255).Optional().Nillable().
			Comment("Description of the ACME account"),
		field.String("email").MaxLen(50).
			Comment("Email address associated with the ACME account"),
		field.Enum("key_type").GoType(enum.KeyType("")).
			Comment("Type of private key associated with the ACME account").
			SchemaType(map[string]string{"postgres": "varchar"}).
			Annotations(entgql.Type("KeyType")),
		field.String("private_key").
			Comment("Private key associated with the ACME account").
			Annotations(entgql.Skip()),
		field.String("key_fingerprint").MaxLen(64).
			Comment("Fingerprint of the private key associated with the ACME account"),
		field.JSON("registration", registration.Resource{}).
			Comment("Registration information associated with the ACME account").
			Annotations(entgql.Skip()),
		field.String("eab_key_id").Optional().Nillable().
			Comment("External Account Binding (EAB) key ID"),
		field.String("eab_hmac_key").Optional().Nillable().
			Comment("External Account Binding (EAB) HMAC key"),
		field.Time("created_at").Immutable().Default(func() time.Time { return time.Now() }).
			Comment("Time the ACME account was created").
			Annotations(entgql.OrderField("CREATED_AT")),
		field.UUID("created_by", uuid.UUID{}).
			Comment("User that created the ACME account").
			Annotations(
				entgql.Type("ID"),
				entgql.Directives(directive.Model("User")),
			),
		field.Time("updated_at").Optional().Nillable().Comment("Time the ACME account was updated"),
		field.UUID("updated_by", uuid.UUID{}).Optional().Nillable().
			Comment("User that updated the ACME account").
			Annotations(
				entgql.Type("ID"),
				entgql.Directives(directive.Model("User")),
			),
	}
}

// Indexes of the AcmeServer.
func (AcmeAccount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("email", "acme_server_id").Unique(),
		index.Fields("created_at"),
	}
}

// Edges of the AcmeAccount.
func (AcmeAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("certificates", Certificate.Type).
			Annotations(entgql.Skip()),
		edge.From("acme_server", AcmeServer.Type).Ref("acme_accounts").Unique().
			Comment("ACME server associated with the ACME account").
			Field("acme_server_id").Required(),
	}
}
