// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team18/app/ent/checkin"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/fixroom"
	"github.com/team18/app/ent/furnituredetail"
	"github.com/team18/app/ent/promotion"
	"github.com/team18/app/ent/reserveroom"
	"github.com/team18/app/ent/statusroom"
	"github.com/team18/app/ent/typeroom"
)

// DataRoomCreate is the builder for creating a DataRoom entity.
type DataRoomCreate struct {
	config
	mutation *DataRoomMutation
	hooks    []Hook
}

// SetPrice sets the price field.
func (drc *DataRoomCreate) SetPrice(f float64) *DataRoomCreate {
	drc.mutation.SetPrice(f)
	return drc
}

// SetRoomnumber sets the roomnumber field.
func (drc *DataRoomCreate) SetRoomnumber(s string) *DataRoomCreate {
	drc.mutation.SetRoomnumber(s)
	return drc
}

// AddReserfIDs adds the reserves edge to ReserveRoom by ids.
func (drc *DataRoomCreate) AddReserfIDs(ids ...int) *DataRoomCreate {
	drc.mutation.AddReserfIDs(ids...)
	return drc
}

// AddReserves adds the reserves edges to ReserveRoom.
func (drc *DataRoomCreate) AddReserves(r ...*ReserveRoom) *DataRoomCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return drc.AddReserfIDs(ids...)
}

// AddFixIDs adds the fixs edge to FixRoom by ids.
func (drc *DataRoomCreate) AddFixIDs(ids ...int) *DataRoomCreate {
	drc.mutation.AddFixIDs(ids...)
	return drc
}

// AddFixs adds the fixs edges to FixRoom.
func (drc *DataRoomCreate) AddFixs(f ...*FixRoom) *DataRoomCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return drc.AddFixIDs(ids...)
}

// AddDetailIDs adds the details edge to FurnitureDetail by ids.
func (drc *DataRoomCreate) AddDetailIDs(ids ...int) *DataRoomCreate {
	drc.mutation.AddDetailIDs(ids...)
	return drc
}

// AddDetails adds the details edges to FurnitureDetail.
func (drc *DataRoomCreate) AddDetails(f ...*FurnitureDetail) *DataRoomCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return drc.AddDetailIDs(ids...)
}

// AddCheckinIDs adds the checkins edge to CheckIn by ids.
func (drc *DataRoomCreate) AddCheckinIDs(ids ...int) *DataRoomCreate {
	drc.mutation.AddCheckinIDs(ids...)
	return drc
}

// AddCheckins adds the checkins edges to CheckIn.
func (drc *DataRoomCreate) AddCheckins(c ...*CheckIn) *DataRoomCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return drc.AddCheckinIDs(ids...)
}

// SetPromotionID sets the promotion edge to Promotion by id.
func (drc *DataRoomCreate) SetPromotionID(id int) *DataRoomCreate {
	drc.mutation.SetPromotionID(id)
	return drc
}

// SetNillablePromotionID sets the promotion edge to Promotion by id if the given value is not nil.
func (drc *DataRoomCreate) SetNillablePromotionID(id *int) *DataRoomCreate {
	if id != nil {
		drc = drc.SetPromotionID(*id)
	}
	return drc
}

// SetPromotion sets the promotion edge to Promotion.
func (drc *DataRoomCreate) SetPromotion(p *Promotion) *DataRoomCreate {
	return drc.SetPromotionID(p.ID)
}

// SetStatusroomID sets the statusroom edge to StatusRoom by id.
func (drc *DataRoomCreate) SetStatusroomID(id int) *DataRoomCreate {
	drc.mutation.SetStatusroomID(id)
	return drc
}

// SetNillableStatusroomID sets the statusroom edge to StatusRoom by id if the given value is not nil.
func (drc *DataRoomCreate) SetNillableStatusroomID(id *int) *DataRoomCreate {
	if id != nil {
		drc = drc.SetStatusroomID(*id)
	}
	return drc
}

// SetStatusroom sets the statusroom edge to StatusRoom.
func (drc *DataRoomCreate) SetStatusroom(s *StatusRoom) *DataRoomCreate {
	return drc.SetStatusroomID(s.ID)
}

// SetTyperoomID sets the typeroom edge to TypeRoom by id.
func (drc *DataRoomCreate) SetTyperoomID(id int) *DataRoomCreate {
	drc.mutation.SetTyperoomID(id)
	return drc
}

// SetNillableTyperoomID sets the typeroom edge to TypeRoom by id if the given value is not nil.
func (drc *DataRoomCreate) SetNillableTyperoomID(id *int) *DataRoomCreate {
	if id != nil {
		drc = drc.SetTyperoomID(*id)
	}
	return drc
}

// SetTyperoom sets the typeroom edge to TypeRoom.
func (drc *DataRoomCreate) SetTyperoom(t *TypeRoom) *DataRoomCreate {
	return drc.SetTyperoomID(t.ID)
}

// Mutation returns the DataRoomMutation object of the builder.
func (drc *DataRoomCreate) Mutation() *DataRoomMutation {
	return drc.mutation
}

// Save creates the DataRoom in the database.
func (drc *DataRoomCreate) Save(ctx context.Context) (*DataRoom, error) {
	if _, ok := drc.mutation.Price(); !ok {
		return nil, &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	if v, ok := drc.mutation.Price(); ok {
		if err := dataroom.PriceValidator(v); err != nil {
			return nil, &ValidationError{Name: "price", err: fmt.Errorf("ent: validator failed for field \"price\": %w", err)}
		}
	}
	if _, ok := drc.mutation.Roomnumber(); !ok {
		return nil, &ValidationError{Name: "roomnumber", err: errors.New("ent: missing required field \"roomnumber\"")}
	}
	if v, ok := drc.mutation.Roomnumber(); ok {
		if err := dataroom.RoomnumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "roomnumber", err: fmt.Errorf("ent: validator failed for field \"roomnumber\": %w", err)}
		}
	}
	var (
		err  error
		node *DataRoom
	)
	if len(drc.hooks) == 0 {
		node, err = drc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DataRoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			drc.mutation = mutation
			node, err = drc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(drc.hooks) - 1; i >= 0; i-- {
			mut = drc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, drc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (drc *DataRoomCreate) SaveX(ctx context.Context) *DataRoom {
	v, err := drc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (drc *DataRoomCreate) sqlSave(ctx context.Context) (*DataRoom, error) {
	dr, _spec := drc.createSpec()
	if err := sqlgraph.CreateNode(ctx, drc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	dr.ID = int(id)
	return dr, nil
}

func (drc *DataRoomCreate) createSpec() (*DataRoom, *sqlgraph.CreateSpec) {
	var (
		dr    = &DataRoom{config: drc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dataroom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dataroom.FieldID,
			},
		}
	)
	if value, ok := drc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dataroom.FieldPrice,
		})
		dr.Price = value
	}
	if value, ok := drc.mutation.Roomnumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dataroom.FieldRoomnumber,
		})
		dr.Roomnumber = value
	}
	if nodes := drc.mutation.ReservesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dataroom.ReservesTable,
			Columns: []string{dataroom.ReservesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reserveroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := drc.mutation.FixsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dataroom.FixsTable,
			Columns: []string{dataroom.FixsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fixroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := drc.mutation.DetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dataroom.DetailsTable,
			Columns: []string{dataroom.DetailsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := drc.mutation.CheckinsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dataroom.CheckinsTable,
			Columns: []string{dataroom.CheckinsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: checkin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := drc.mutation.PromotionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dataroom.PromotionTable,
			Columns: []string{dataroom.PromotionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := drc.mutation.StatusroomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dataroom.StatusroomTable,
			Columns: []string{dataroom.StatusroomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statusroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := drc.mutation.TyperoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dataroom.TyperoomTable,
			Columns: []string{dataroom.TyperoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: typeroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return dr, _spec
}
