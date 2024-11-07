package schema

import (
	"acme-manager/ent/schema/directive"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// AcmeServer holds the schema definition for the AcmeServer entity.
type AcmeServer struct {
	ent.Schema
}

// Annotations of the AcmeServer.
func (AcmeServer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "acme_server"},
		entgql.RelayConnection(),
	}
}

// Fields of the AcmeServer.
func (AcmeServer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			Annotations(entgql.Directives(directive.Model("AcmeServer"))),
		field.String("name").MaxLen(50).Comment("Name of the ACME server"),
		field.String("description").MaxLen(255).Optional().Nillable().
			Comment("Description of the ACME server"),
		field.String("url").MaxLen(255).Comment("URL of the ACME server"),
		field.Bool("built_in").Default(false).Comment("Is this a built-in ACME server"),
		field.Bool("external_account_required").Default(false).
			Comment("Does the ACME server require an External Account Binding"),
		field.Time("created_at").Immutable().Default(func() time.Time { return time.Now() }).
			Comment("Time the ACME server was created").
			Annotations(entgql.OrderField("CREATED_AT")),
		field.UUID("created_by", uuid.UUID{}).
			Comment("User that created the ACME server").
			Annotations(
				entgql.Type("ID"),
				entgql.Directives(directive.Model("User")),
			),
		field.Time("updated_at").Optional().Nillable().Comment("Time the ACME server was updated"),
		field.UUID("updated_by", uuid.UUID{}).Optional().Nillable().
			Comment("User that updated the ACME server").
			Annotations(
				entgql.Type("ID"),
				entgql.Directives(directive.Model("User")),
			),
	}
}

// Indexes of the AcmeServer.
func (AcmeServer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("url").Unique(),
		index.Fields("created_at"),
	}
}

// Edges of the AcmeServer.
func (AcmeServer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("acme_accounts", AcmeAccount.Type).
			Annotations(entgql.Skip()),
	}
}
