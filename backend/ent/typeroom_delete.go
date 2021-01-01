// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team18/app/ent/predicate"
	"github.com/team18/app/ent/typeroom"
)

// TypeRoomDelete is the builder for deleting a TypeRoom entity.
type TypeRoomDelete struct {
	config
	hooks      []Hook
	mutation   *TypeRoomMutation
	predicates []predicate.TypeRoom
}

// Where adds a new predicate to the delete builder.
func (trd *TypeRoomDelete) Where(ps ...predicate.TypeRoom) *TypeRoomDelete {
	trd.predicates = append(trd.predicates, ps...)
	return trd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (trd *TypeRoomDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(trd.hooks) == 0 {
		affected, err = trd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TypeRoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			trd.mutation = mutation
			affected, err = trd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(trd.hooks) - 1; i >= 0; i-- {
			mut = trd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, trd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (trd *TypeRoomDelete) ExecX(ctx context.Context) int {
	n, err := trd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (trd *TypeRoomDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: typeroom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: typeroom.FieldID,
			},
		},
	}
	if ps := trd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, trd.driver, _spec)
}

// TypeRoomDeleteOne is the builder for deleting a single TypeRoom entity.
type TypeRoomDeleteOne struct {
	trd *TypeRoomDelete
}

// Exec executes the deletion query.
func (trdo *TypeRoomDeleteOne) Exec(ctx context.Context) error {
	n, err := trdo.trd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{typeroom.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (trdo *TypeRoomDeleteOne) ExecX(ctx context.Context) {
	trdo.trd.ExecX(ctx)
}