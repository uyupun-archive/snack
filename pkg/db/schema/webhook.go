package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Webhook holds the schema definition for the Webhook entity.
type Webhook struct {
	ent.Schema
}

// Fields of the Webhook.
func (Webhook) Fields() []ent.Field {
	return []ent.Field{
		field.String("app_name").Unique(),
		field.String("webhook_url").Unique(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Webhook.
func (Webhook) Edges() []ent.Edge {
	return nil
}
