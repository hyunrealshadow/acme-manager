package schema

import (
	"acme-manager/acme/lego"
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

// DnsProvider holds the schema definition for the DnsProvider entity.
type DnsProvider struct {
	ent.Schema
}

// Annotations of the Certificate.
func (DnsProvider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "dns_provider"},
		entgql.RelayConnection(),
	}
}

// Fields of the DnsProvider.
func (DnsProvider) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			Annotations(entgql.Directives(directive.Model("DnsProvider"))),
		field.String("name").MaxLen(50).Comment("Name of the DNS provider"),
		field.String("description").MaxLen(255).Optional().Nillable().
			Comment("Description of the DNS provider"),
		field.String("type").MaxLen(20).Comment("Type of the DNS provider"),
		field.JSON("config", &lego.DnsProviderConfig{}).Comment("Configuration of the DNS provider").
			Annotations(entgql.Type("String")),
		field.Time("created_at").Immutable().Default(func() time.Time { return time.Now() }).
			Comment("Time the DNS provider was created").
			Annotations(entgql.OrderField("CREATED_AT")),
		field.UUID("created_by", uuid.UUID{}).
			Comment("User that created the ACME account").
			Annotations(
				entgql.Type("ID"),
				entgql.Directives(directive.Model("User")),
			),
		field.Time("updated_at").Optional().Nillable().Comment("Time the DNS provider was updated"),
		field.UUID("updated_by", uuid.UUID{}).Optional().Nillable().
			Comment("User that updated the ACME account").
			Annotations(
				entgql.Type("ID"),
				entgql.Directives(directive.Model("User")),
			),
	}
}

// Indexes of the AcmeServer.
func (DnsProvider) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("type"),
		index.Fields("created_at"),
		index.Fields("name", "type").Unique(),
	}
}

// Edges of the DnsProvider.
func (DnsProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("certificates", Certificate.Type).
			Annotations(entgql.Skip()),
	}
}
