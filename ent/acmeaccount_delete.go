// Code generated by ent, DO NOT EDIT.

package ent

import (
	"acme-manager/ent/acmeaccount"
	"acme-manager/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AcmeAccountDelete is the builder for deleting a AcmeAccount entity.
type AcmeAccountDelete struct {
	config
	hooks    []Hook
	mutation *AcmeAccountMutation
}

// Where appends a list predicates to the AcmeAccountDelete builder.
func (aad *AcmeAccountDelete) Where(ps ...predicate.AcmeAccount) *AcmeAccountDelete {
	aad.mutation.Where(ps...)
	return aad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aad *AcmeAccountDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, aad.sqlExec, aad.mutation, aad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (aad *AcmeAccountDelete) ExecX(ctx context.Context) int {
	n, err := aad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aad *AcmeAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(acmeaccount.Table, sqlgraph.NewFieldSpec(acmeaccount.FieldID, field.TypeUUID))
	if ps := aad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, aad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	aad.mutation.done = true
	return affected, err
}

// AcmeAccountDeleteOne is the builder for deleting a single AcmeAccount entity.
type AcmeAccountDeleteOne struct {
	aad *AcmeAccountDelete
}

// Where appends a list predicates to the AcmeAccountDelete builder.
func (aado *AcmeAccountDeleteOne) Where(ps ...predicate.AcmeAccount) *AcmeAccountDeleteOne {
	aado.aad.mutation.Where(ps...)
	return aado
}

// Exec executes the deletion query.
func (aado *AcmeAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := aado.aad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{acmeaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aado *AcmeAccountDeleteOne) ExecX(ctx context.Context) {
	if err := aado.Exec(ctx); err != nil {
		panic(err)
	}
}
