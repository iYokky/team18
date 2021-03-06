// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team18/app/ent/counterstaff"
	"github.com/team18/app/ent/predicate"
)

// CounterStaffDelete is the builder for deleting a CounterStaff entity.
type CounterStaffDelete struct {
	config
	hooks      []Hook
	mutation   *CounterStaffMutation
	predicates []predicate.CounterStaff
}

// Where adds a new predicate to the delete builder.
func (csd *CounterStaffDelete) Where(ps ...predicate.CounterStaff) *CounterStaffDelete {
	csd.predicates = append(csd.predicates, ps...)
	return csd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (csd *CounterStaffDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(csd.hooks) == 0 {
		affected, err = csd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CounterStaffMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csd.mutation = mutation
			affected, err = csd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(csd.hooks) - 1; i >= 0; i-- {
			mut = csd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (csd *CounterStaffDelete) ExecX(ctx context.Context) int {
	n, err := csd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (csd *CounterStaffDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: counterstaff.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: counterstaff.FieldID,
			},
		},
	}
	if ps := csd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, csd.driver, _spec)
}

// CounterStaffDeleteOne is the builder for deleting a single CounterStaff entity.
type CounterStaffDeleteOne struct {
	csd *CounterStaffDelete
}

// Exec executes the deletion query.
func (csdo *CounterStaffDeleteOne) Exec(ctx context.Context) error {
	n, err := csdo.csd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{counterstaff.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (csdo *CounterStaffDeleteOne) ExecX(ctx context.Context) {
	csdo.csd.ExecX(ctx)
}
