// Code generated by ent, DO NOT EDIT.

package ent

import (
	"acme-manager/ent/acmeserver"
	"acme-manager/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AcmeServerDelete is the builder for deleting a AcmeServer entity.
type AcmeServerDelete struct {
	config
	hooks    []Hook
	mutation *AcmeServerMutation
}

// Where appends a list predicates to the AcmeServerDelete builder.
func (asd *AcmeServerDelete) Where(ps ...predicate.AcmeServer) *AcmeServerDelete {
	asd.mutation.Where(ps...)
	return asd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (asd *AcmeServerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, asd.sqlExec, asd.mutation, asd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (asd *AcmeServerDelete) ExecX(ctx context.Context) int {
	n, err := asd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (asd *AcmeServerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(acmeserver.Table, sqlgraph.NewFieldSpec(acmeserver.FieldID, field.TypeUUID))
	if ps := asd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, asd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	asd.mutation.done = true
	return affected, err
}

// AcmeServerDeleteOne is the builder for deleting a single AcmeServer entity.
type AcmeServerDeleteOne struct {
	asd *AcmeServerDelete
}

// Where appends a list predicates to the AcmeServerDelete builder.
func (asdo *AcmeServerDeleteOne) Where(ps ...predicate.AcmeServer) *AcmeServerDeleteOne {
	asdo.asd.mutation.Where(ps...)
	return asdo
}

// Exec executes the deletion query.
func (asdo *AcmeServerDeleteOne) Exec(ctx context.Context) error {
	n, err := asdo.asd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{acmeserver.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (asdo *AcmeServerDeleteOne) ExecX(ctx context.Context) {
	if err := asdo.Exec(ctx); err != nil {
		panic(err)
	}
}