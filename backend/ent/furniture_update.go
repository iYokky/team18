// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team18/app/ent/furniture"
	"github.com/team18/app/ent/furnituredetail"
	"github.com/team18/app/ent/predicate"
)

// FurnitureUpdate is the builder for updating Furniture entities.
type FurnitureUpdate struct {
	config
	hooks      []Hook
	mutation   *FurnitureMutation
	predicates []predicate.Furniture
}

// Where adds a new predicate for the builder.
func (fu *FurnitureUpdate) Where(ps ...predicate.Furniture) *FurnitureUpdate {
	fu.predicates = append(fu.predicates, ps...)
	return fu
}

// SetFurnitureName sets the furniture_name field.
func (fu *FurnitureUpdate) SetFurnitureName(s string) *FurnitureUpdate {
	fu.mutation.SetFurnitureName(s)
	return fu
}

// AddDetailIDs adds the details edge to FurnitureDetail by ids.
func (fu *FurnitureUpdate) AddDetailIDs(ids ...int) *FurnitureUpdate {
	fu.mutation.AddDetailIDs(ids...)
	return fu
}

// AddDetails adds the details edges to FurnitureDetail.
func (fu *FurnitureUpdate) AddDetails(f ...*FurnitureDetail) *FurnitureUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.AddDetailIDs(ids...)
}

// Mutation returns the FurnitureMutation object of the builder.
func (fu *FurnitureUpdate) Mutation() *FurnitureMutation {
	return fu.mutation
}

// RemoveDetailIDs removes the details edge to FurnitureDetail by ids.
func (fu *FurnitureUpdate) RemoveDetailIDs(ids ...int) *FurnitureUpdate {
	fu.mutation.RemoveDetailIDs(ids...)
	return fu
}

// RemoveDetails removes details edges to FurnitureDetail.
func (fu *FurnitureUpdate) RemoveDetails(f ...*FurnitureDetail) *FurnitureUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.RemoveDetailIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (fu *FurnitureUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := fu.mutation.FurnitureName(); ok {
		if err := furniture.FurnitureNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "furniture_name", err: fmt.Errorf("ent: validator failed for field \"furniture_name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FurnitureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FurnitureUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FurnitureUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FurnitureUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FurnitureUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   furniture.Table,
			Columns: furniture.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: furniture.FieldID,
			},
		},
	}
	if ps := fu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.FurnitureName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: furniture.FieldFurnitureName,
		})
	}
	if nodes := fu.mutation.RemovedDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   furniture.DetailsTable,
			Columns: []string{furniture.DetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: furnituredetail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.DetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   furniture.DetailsTable,
			Columns: []string{furniture.DetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: furnituredetail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{furniture.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FurnitureUpdateOne is the builder for updating a single Furniture entity.
type FurnitureUpdateOne struct {
	config
	hooks    []Hook
	mutation *FurnitureMutation
}

// SetFurnitureName sets the furniture_name field.
func (fuo *FurnitureUpdateOne) SetFurnitureName(s string) *FurnitureUpdateOne {
	fuo.mutation.SetFurnitureName(s)
	return fuo
}

// AddDetailIDs adds the details edge to FurnitureDetail by ids.
func (fuo *FurnitureUpdateOne) AddDetailIDs(ids ...int) *FurnitureUpdateOne {
	fuo.mutation.AddDetailIDs(ids...)
	return fuo
}

// AddDetails adds the details edges to FurnitureDetail.
func (fuo *FurnitureUpdateOne) AddDetails(f ...*FurnitureDetail) *FurnitureUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.AddDetailIDs(ids...)
}

// Mutation returns the FurnitureMutation object of the builder.
func (fuo *FurnitureUpdateOne) Mutation() *FurnitureMutation {
	return fuo.mutation
}

// RemoveDetailIDs removes the details edge to FurnitureDetail by ids.
func (fuo *FurnitureUpdateOne) RemoveDetailIDs(ids ...int) *FurnitureUpdateOne {
	fuo.mutation.RemoveDetailIDs(ids...)
	return fuo
}

// RemoveDetails removes details edges to FurnitureDetail.
func (fuo *FurnitureUpdateOne) RemoveDetails(f ...*FurnitureDetail) *FurnitureUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.RemoveDetailIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (fuo *FurnitureUpdateOne) Save(ctx context.Context) (*Furniture, error) {
	if v, ok := fuo.mutation.FurnitureName(); ok {
		if err := furniture.FurnitureNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "furniture_name", err: fmt.Errorf("ent: validator failed for field \"furniture_name\": %w", err)}
		}
	}

	var (
		err  error
		node *Furniture
	)
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FurnitureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FurnitureUpdateOne) SaveX(ctx context.Context) *Furniture {
	f, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return f
}

// Exec executes the query on the entity.
func (fuo *FurnitureUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FurnitureUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FurnitureUpdateOne) sqlSave(ctx context.Context) (f *Furniture, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   furniture.Table,
			Columns: furniture.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: furniture.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Furniture.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := fuo.mutation.FurnitureName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: furniture.FieldFurnitureName,
		})
	}
	if nodes := fuo.mutation.RemovedDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   furniture.DetailsTable,
			Columns: []string{furniture.DetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: furnituredetail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.DetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   furniture.DetailsTable,
			Columns: []string{furniture.DetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: furnituredetail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	f = &Furniture{config: fuo.config}
	_spec.Assign = f.assignValues
	_spec.ScanValues = f.scanValues()
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{furniture.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return f, nil
}
