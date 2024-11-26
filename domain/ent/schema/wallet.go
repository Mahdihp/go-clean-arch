package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Wallet holds the schema definition for the Wallet entity.
type Wallet struct {
	ent.Schema
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("id"),
		field.String("name").MaxLen(50),
		field.Bool("is_active").Default(true),
		field.Bool("deleted").Default(false),

		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp(6)", // Override Postgres.
			}).Default(time.Now).Immutable(),

		field.Time("updated_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp(6)", // Override Postgres.
			}).Default(time.Now()),
	}
}

func (Wallet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).Ref("wallets").Unique(),
	}
}
