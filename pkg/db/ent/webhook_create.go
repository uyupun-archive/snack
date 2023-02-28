// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"snack/pkg/db/ent/webhook"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WebhookCreate is the builder for creating a Webhook entity.
type WebhookCreate struct {
	config
	mutation *WebhookMutation
	hooks    []Hook
}

// SetAppName sets the "app_name" field.
func (wc *WebhookCreate) SetAppName(s string) *WebhookCreate {
	wc.mutation.SetAppName(s)
	return wc
}

// SetWebhookURL sets the "webhook_url" field.
func (wc *WebhookCreate) SetWebhookURL(s string) *WebhookCreate {
	wc.mutation.SetWebhookURL(s)
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WebhookCreate) SetCreatedAt(t time.Time) *WebhookCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableCreatedAt(t *time.Time) *WebhookCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetUpdatedAt sets the "updated_at" field.
func (wc *WebhookCreate) SetUpdatedAt(t time.Time) *WebhookCreate {
	wc.mutation.SetUpdatedAt(t)
	return wc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wc *WebhookCreate) SetNillableUpdatedAt(t *time.Time) *WebhookCreate {
	if t != nil {
		wc.SetUpdatedAt(*t)
	}
	return wc
}

// Mutation returns the WebhookMutation object of the builder.
func (wc *WebhookCreate) Mutation() *WebhookMutation {
	return wc.mutation
}

// Save creates the Webhook in the database.
func (wc *WebhookCreate) Save(ctx context.Context) (*Webhook, error) {
	wc.defaults()
	return withHooks[*Webhook, WebhookMutation](ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WebhookCreate) SaveX(ctx context.Context) *Webhook {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WebhookCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WebhookCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WebhookCreate) defaults() {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := webhook.DefaultCreatedAt
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		v := webhook.DefaultUpdatedAt
		wc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WebhookCreate) check() error {
	if _, ok := wc.mutation.AppName(); !ok {
		return &ValidationError{Name: "app_name", err: errors.New(`ent: missing required field "Webhook.app_name"`)}
	}
	if _, ok := wc.mutation.WebhookURL(); !ok {
		return &ValidationError{Name: "webhook_url", err: errors.New(`ent: missing required field "Webhook.webhook_url"`)}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Webhook.created_at"`)}
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Webhook.updated_at"`)}
	}
	return nil
}

func (wc *WebhookCreate) sqlSave(ctx context.Context) (*Webhook, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WebhookCreate) createSpec() (*Webhook, *sqlgraph.CreateSpec) {
	var (
		_node = &Webhook{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(webhook.Table, sqlgraph.NewFieldSpec(webhook.FieldID, field.TypeInt))
	)
	if value, ok := wc.mutation.AppName(); ok {
		_spec.SetField(webhook.FieldAppName, field.TypeString, value)
		_node.AppName = value
	}
	if value, ok := wc.mutation.WebhookURL(); ok {
		_spec.SetField(webhook.FieldWebhookURL, field.TypeString, value)
		_node.WebhookURL = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.SetField(webhook.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.UpdatedAt(); ok {
		_spec.SetField(webhook.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// WebhookCreateBulk is the builder for creating many Webhook entities in bulk.
type WebhookCreateBulk struct {
	config
	builders []*WebhookCreate
}

// Save creates the Webhook entities in the database.
func (wcb *WebhookCreateBulk) Save(ctx context.Context) ([]*Webhook, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Webhook, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WebhookMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WebhookCreateBulk) SaveX(ctx context.Context) []*Webhook {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WebhookCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WebhookCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}