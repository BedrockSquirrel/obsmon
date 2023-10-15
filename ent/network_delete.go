// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bedrocksquirrel/obsmon/ent/network"
	"github.com/bedrocksquirrel/obsmon/ent/predicate"
)

// NetworkDelete is the builder for deleting a Network entity.
type NetworkDelete struct {
	config
	hooks    []Hook
	mutation *NetworkMutation
}

// Where appends a list predicates to the NetworkDelete builder.
func (nd *NetworkDelete) Where(ps ...predicate.Network) *NetworkDelete {
	nd.mutation.Where(ps...)
	return nd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nd *NetworkDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nd.sqlExec, nd.mutation, nd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NetworkDelete) ExecX(ctx context.Context) int {
	n, err := nd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nd *NetworkDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(network.Table, sqlgraph.NewFieldSpec(network.FieldID, field.TypeInt))
	if ps := nd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nd.mutation.done = true
	return affected, err
}

// NetworkDeleteOne is the builder for deleting a single Network entity.
type NetworkDeleteOne struct {
	nd *NetworkDelete
}

// Where appends a list predicates to the NetworkDelete builder.
func (ndo *NetworkDeleteOne) Where(ps ...predicate.Network) *NetworkDeleteOne {
	ndo.nd.mutation.Where(ps...)
	return ndo
}

// Exec executes the deletion query.
func (ndo *NetworkDeleteOne) Exec(ctx context.Context) error {
	n, err := ndo.nd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{network.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NetworkDeleteOne) ExecX(ctx context.Context) {
	if err := ndo.Exec(ctx); err != nil {
		panic(err)
	}
}
