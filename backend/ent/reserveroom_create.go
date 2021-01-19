// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team18/app/ent/checkin"
	"github.com/team18/app/ent/customer"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/promotion"
	"github.com/team18/app/ent/reserveroom"
	"github.com/team18/app/ent/statusreserve"
)

// ReserveRoomCreate is the builder for creating a ReserveRoom entity.
type ReserveRoomCreate struct {
	config
	mutation *ReserveRoomMutation
	hooks    []Hook
}

// SetReserveDate sets the reserve_date field.
func (rrc *ReserveRoomCreate) SetReserveDate(t time.Time) *ReserveRoomCreate {
	rrc.mutation.SetReserveDate(t)
	return rrc
}

// SetProvince sets the province field.
func (rrc *ReserveRoomCreate) SetProvince(s string) *ReserveRoomCreate {
	rrc.mutation.SetProvince(s)
	return rrc
}

// SetAmount sets the amount field.
func (rrc *ReserveRoomCreate) SetAmount(i int) *ReserveRoomCreate {
	rrc.mutation.SetAmount(i)
	return rrc
}

// SetPhoneNumber sets the phone_number field.
func (rrc *ReserveRoomCreate) SetPhoneNumber(s string) *ReserveRoomCreate {
	rrc.mutation.SetPhoneNumber(s)
	return rrc
}

// SetNetPrice sets the net_price field.
func (rrc *ReserveRoomCreate) SetNetPrice(f float64) *ReserveRoomCreate {
	rrc.mutation.SetNetPrice(f)
	return rrc
}

// SetCustomerID sets the customer edge to Customer by id.
func (rrc *ReserveRoomCreate) SetCustomerID(id int) *ReserveRoomCreate {
	rrc.mutation.SetCustomerID(id)
	return rrc
}

// SetNillableCustomerID sets the customer edge to Customer by id if the given value is not nil.
func (rrc *ReserveRoomCreate) SetNillableCustomerID(id *int) *ReserveRoomCreate {
	if id != nil {
		rrc = rrc.SetCustomerID(*id)
	}
	return rrc
}

// SetCustomer sets the customer edge to Customer.
func (rrc *ReserveRoomCreate) SetCustomer(c *Customer) *ReserveRoomCreate {
	return rrc.SetCustomerID(c.ID)
}

// SetPromotionID sets the promotion edge to Promotion by id.
func (rrc *ReserveRoomCreate) SetPromotionID(id int) *ReserveRoomCreate {
	rrc.mutation.SetPromotionID(id)
	return rrc
}

// SetNillablePromotionID sets the promotion edge to Promotion by id if the given value is not nil.
func (rrc *ReserveRoomCreate) SetNillablePromotionID(id *int) *ReserveRoomCreate {
	if id != nil {
		rrc = rrc.SetPromotionID(*id)
	}
	return rrc
}

// SetPromotion sets the promotion edge to Promotion.
func (rrc *ReserveRoomCreate) SetPromotion(p *Promotion) *ReserveRoomCreate {
	return rrc.SetPromotionID(p.ID)
}

// SetRoomID sets the room edge to DataRoom by id.
func (rrc *ReserveRoomCreate) SetRoomID(id int) *ReserveRoomCreate {
	rrc.mutation.SetRoomID(id)
	return rrc
}

// SetNillableRoomID sets the room edge to DataRoom by id if the given value is not nil.
func (rrc *ReserveRoomCreate) SetNillableRoomID(id *int) *ReserveRoomCreate {
	if id != nil {
		rrc = rrc.SetRoomID(*id)
	}
	return rrc
}

// SetRoom sets the room edge to DataRoom.
func (rrc *ReserveRoomCreate) SetRoom(d *DataRoom) *ReserveRoomCreate {
	return rrc.SetRoomID(d.ID)
}

// SetStatusID sets the status edge to StatusReserve by id.
func (rrc *ReserveRoomCreate) SetStatusID(id int) *ReserveRoomCreate {
	rrc.mutation.SetStatusID(id)
	return rrc
}

// SetNillableStatusID sets the status edge to StatusReserve by id if the given value is not nil.
func (rrc *ReserveRoomCreate) SetNillableStatusID(id *int) *ReserveRoomCreate {
	if id != nil {
		rrc = rrc.SetStatusID(*id)
	}
	return rrc
}

// SetStatus sets the status edge to StatusReserve.
func (rrc *ReserveRoomCreate) SetStatus(s *StatusReserve) *ReserveRoomCreate {
	return rrc.SetStatusID(s.ID)
}

// AddCheckinIDs adds the checkins edge to CheckIn by ids.
func (rrc *ReserveRoomCreate) AddCheckinIDs(ids ...int) *ReserveRoomCreate {
	rrc.mutation.AddCheckinIDs(ids...)
	return rrc
}

// AddCheckins adds the checkins edges to CheckIn.
func (rrc *ReserveRoomCreate) AddCheckins(c ...*CheckIn) *ReserveRoomCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return rrc.AddCheckinIDs(ids...)
}

// Mutation returns the ReserveRoomMutation object of the builder.
func (rrc *ReserveRoomCreate) Mutation() *ReserveRoomMutation {
	return rrc.mutation
}

// Save creates the ReserveRoom in the database.
func (rrc *ReserveRoomCreate) Save(ctx context.Context) (*ReserveRoom, error) {
	if _, ok := rrc.mutation.ReserveDate(); !ok {
		return nil, &ValidationError{Name: "reserve_date", err: errors.New("ent: missing required field \"reserve_date\"")}
	}
	if _, ok := rrc.mutation.Province(); !ok {
		return nil, &ValidationError{Name: "province", err: errors.New("ent: missing required field \"province\"")}
	}
	if v, ok := rrc.mutation.Province(); ok {
		if err := reserveroom.ProvinceValidator(v); err != nil {
			return nil, &ValidationError{Name: "province", err: fmt.Errorf("ent: validator failed for field \"province\": %w", err)}
		}
	}
	if _, ok := rrc.mutation.Amount(); !ok {
		return nil, &ValidationError{Name: "amount", err: errors.New("ent: missing required field \"amount\"")}
	}
	if v, ok := rrc.mutation.Amount(); ok {
		if err := reserveroom.AmountValidator(v); err != nil {
			return nil, &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}
	if _, ok := rrc.mutation.PhoneNumber(); !ok {
		return nil, &ValidationError{Name: "phone_number", err: errors.New("ent: missing required field \"phone_number\"")}
	}
	if v, ok := rrc.mutation.PhoneNumber(); ok {
		if err := reserveroom.PhoneNumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "phone_number", err: fmt.Errorf("ent: validator failed for field \"phone_number\": %w", err)}
		}
	}
	if _, ok := rrc.mutation.NetPrice(); !ok {
		return nil, &ValidationError{Name: "net_price", err: errors.New("ent: missing required field \"net_price\"")}
	}
	if v, ok := rrc.mutation.NetPrice(); ok {
		if err := reserveroom.NetPriceValidator(v); err != nil {
			return nil, &ValidationError{Name: "net_price", err: fmt.Errorf("ent: validator failed for field \"net_price\": %w", err)}
		}
	}
	var (
		err  error
		node *ReserveRoom
	)
	if len(rrc.hooks) == 0 {
		node, err = rrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReserveRoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rrc.mutation = mutation
			node, err = rrc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rrc.hooks) - 1; i >= 0; i-- {
			mut = rrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rrc *ReserveRoomCreate) SaveX(ctx context.Context) *ReserveRoom {
	v, err := rrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rrc *ReserveRoomCreate) sqlSave(ctx context.Context) (*ReserveRoom, error) {
	rr, _spec := rrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rrc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	rr.ID = int(id)
	return rr, nil
}

func (rrc *ReserveRoomCreate) createSpec() (*ReserveRoom, *sqlgraph.CreateSpec) {
	var (
		rr    = &ReserveRoom{config: rrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: reserveroom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: reserveroom.FieldID,
			},
		}
	)
	if value, ok := rrc.mutation.ReserveDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: reserveroom.FieldReserveDate,
		})
		rr.ReserveDate = value
	}
	if value, ok := rrc.mutation.Province(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reserveroom.FieldProvince,
		})
		rr.Province = value
	}
	if value, ok := rrc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reserveroom.FieldAmount,
		})
		rr.Amount = value
	}
	if value, ok := rrc.mutation.PhoneNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reserveroom.FieldPhoneNumber,
		})
		rr.PhoneNumber = value
	}
	if value, ok := rrc.mutation.NetPrice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: reserveroom.FieldNetPrice,
		})
		rr.NetPrice = value
	}
	if nodes := rrc.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reserveroom.CustomerTable,
			Columns: []string{reserveroom.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rrc.mutation.PromotionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reserveroom.PromotionTable,
			Columns: []string{reserveroom.PromotionColumn},
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
	if nodes := rrc.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reserveroom.RoomTable,
			Columns: []string{reserveroom.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dataroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rrc.mutation.StatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reserveroom.StatusTable,
			Columns: []string{reserveroom.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statusreserve.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rrc.mutation.CheckinsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   reserveroom.CheckinsTable,
			Columns: []string{reserveroom.CheckinsColumn},
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
	return rr, _spec
}
