// Code generated by ent, DO NOT EDIT.

package ent

import (
	"snack/pkg/db/ent/webhook"
	"snack/pkg/db/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	webhookFields := schema.Webhook{}.Fields()
	_ = webhookFields
	// webhookDescCreatedAt is the schema descriptor for created_at field.
	webhookDescCreatedAt := webhookFields[2].Descriptor()
	// webhook.DefaultCreatedAt holds the default value on creation for the created_at field.
	webhook.DefaultCreatedAt = webhookDescCreatedAt.Default.(time.Time)
	// webhookDescUpdatedAt is the schema descriptor for updated_at field.
	webhookDescUpdatedAt := webhookFields[3].Descriptor()
	// webhook.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	webhook.DefaultUpdatedAt = webhookDescUpdatedAt.Default.(time.Time)
}
