// Code generated by ent, DO NOT EDIT.

package webhook

import (
	"snack/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Webhook {
	return predicate.Webhook(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Webhook {
	return predicate.Webhook(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Webhook {
	return predicate.Webhook(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Webhook {
	return predicate.Webhook(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Webhook {
	return predicate.Webhook(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Webhook {
	return predicate.Webhook(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Webhook {
	return predicate.Webhook(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Webhook {
	return predicate.Webhook(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Webhook {
	return predicate.Webhook(sql.FieldLTE(FieldID, id))
}

// AppName applies equality check predicate on the "app_name" field. It's identical to AppNameEQ.
func AppName(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldEQ(FieldAppName, v))
}

// WebhookURL applies equality check predicate on the "webhook_url" field. It's identical to WebhookURLEQ.
func WebhookURL(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldEQ(FieldWebhookURL, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldEQ(FieldUpdatedAt, v))
}

// AppNameEQ applies the EQ predicate on the "app_name" field.
func AppNameEQ(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldEQ(FieldAppName, v))
}

// AppNameNEQ applies the NEQ predicate on the "app_name" field.
func AppNameNEQ(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldNEQ(FieldAppName, v))
}

// AppNameIn applies the In predicate on the "app_name" field.
func AppNameIn(vs ...string) predicate.Webhook {
	return predicate.Webhook(sql.FieldIn(FieldAppName, vs...))
}

// AppNameNotIn applies the NotIn predicate on the "app_name" field.
func AppNameNotIn(vs ...string) predicate.Webhook {
	return predicate.Webhook(sql.FieldNotIn(FieldAppName, vs...))
}

// AppNameGT applies the GT predicate on the "app_name" field.
func AppNameGT(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldGT(FieldAppName, v))
}

// AppNameGTE applies the GTE predicate on the "app_name" field.
func AppNameGTE(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldGTE(FieldAppName, v))
}

// AppNameLT applies the LT predicate on the "app_name" field.
func AppNameLT(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldLT(FieldAppName, v))
}

// AppNameLTE applies the LTE predicate on the "app_name" field.
func AppNameLTE(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldLTE(FieldAppName, v))
}

// AppNameContains applies the Contains predicate on the "app_name" field.
func AppNameContains(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldContains(FieldAppName, v))
}

// AppNameHasPrefix applies the HasPrefix predicate on the "app_name" field.
func AppNameHasPrefix(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldHasPrefix(FieldAppName, v))
}

// AppNameHasSuffix applies the HasSuffix predicate on the "app_name" field.
func AppNameHasSuffix(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldHasSuffix(FieldAppName, v))
}

// AppNameEqualFold applies the EqualFold predicate on the "app_name" field.
func AppNameEqualFold(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldEqualFold(FieldAppName, v))
}

// AppNameContainsFold applies the ContainsFold predicate on the "app_name" field.
func AppNameContainsFold(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldContainsFold(FieldAppName, v))
}

// WebhookURLEQ applies the EQ predicate on the "webhook_url" field.
func WebhookURLEQ(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldEQ(FieldWebhookURL, v))
}

// WebhookURLNEQ applies the NEQ predicate on the "webhook_url" field.
func WebhookURLNEQ(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldNEQ(FieldWebhookURL, v))
}

// WebhookURLIn applies the In predicate on the "webhook_url" field.
func WebhookURLIn(vs ...string) predicate.Webhook {
	return predicate.Webhook(sql.FieldIn(FieldWebhookURL, vs...))
}

// WebhookURLNotIn applies the NotIn predicate on the "webhook_url" field.
func WebhookURLNotIn(vs ...string) predicate.Webhook {
	return predicate.Webhook(sql.FieldNotIn(FieldWebhookURL, vs...))
}

// WebhookURLGT applies the GT predicate on the "webhook_url" field.
func WebhookURLGT(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldGT(FieldWebhookURL, v))
}

// WebhookURLGTE applies the GTE predicate on the "webhook_url" field.
func WebhookURLGTE(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldGTE(FieldWebhookURL, v))
}

// WebhookURLLT applies the LT predicate on the "webhook_url" field.
func WebhookURLLT(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldLT(FieldWebhookURL, v))
}

// WebhookURLLTE applies the LTE predicate on the "webhook_url" field.
func WebhookURLLTE(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldLTE(FieldWebhookURL, v))
}

// WebhookURLContains applies the Contains predicate on the "webhook_url" field.
func WebhookURLContains(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldContains(FieldWebhookURL, v))
}

// WebhookURLHasPrefix applies the HasPrefix predicate on the "webhook_url" field.
func WebhookURLHasPrefix(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldHasPrefix(FieldWebhookURL, v))
}

// WebhookURLHasSuffix applies the HasSuffix predicate on the "webhook_url" field.
func WebhookURLHasSuffix(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldHasSuffix(FieldWebhookURL, v))
}

// WebhookURLEqualFold applies the EqualFold predicate on the "webhook_url" field.
func WebhookURLEqualFold(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldEqualFold(FieldWebhookURL, v))
}

// WebhookURLContainsFold applies the ContainsFold predicate on the "webhook_url" field.
func WebhookURLContainsFold(v string) predicate.Webhook {
	return predicate.Webhook(sql.FieldContainsFold(FieldWebhookURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Webhook {
	return predicate.Webhook(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Webhook) predicate.Webhook {
	return predicate.Webhook(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Webhook) predicate.Webhook {
	return predicate.Webhook(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Webhook) predicate.Webhook {
	return predicate.Webhook(func(s *sql.Selector) {
		p(s.Not())
	})
}
