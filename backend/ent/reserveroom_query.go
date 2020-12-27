// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team18/app/ent/checkin"
	"github.com/team18/app/ent/customer"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/predicate"
	"github.com/team18/app/ent/promotion"
	"github.com/team18/app/ent/reserveroom"
)

// ReserveRoomQuery is the builder for querying ReserveRoom entities.
type ReserveRoomQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.ReserveRoom
	// eager-loading edges.
	withCustomer  *CustomerQuery
	withPromotion *PromotionQuery
	withRoom      *DataRoomQuery
	withCheckins  *CheckInQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (rrq *ReserveRoomQuery) Where(ps ...predicate.ReserveRoom) *ReserveRoomQuery {
	rrq.predicates = append(rrq.predicates, ps...)
	return rrq
}

// Limit adds a limit step to the query.
func (rrq *ReserveRoomQuery) Limit(limit int) *ReserveRoomQuery {
	rrq.limit = &limit
	return rrq
}

// Offset adds an offset step to the query.
func (rrq *ReserveRoomQuery) Offset(offset int) *ReserveRoomQuery {
	rrq.offset = &offset
	return rrq
}

// Order adds an order step to the query.
func (rrq *ReserveRoomQuery) Order(o ...OrderFunc) *ReserveRoomQuery {
	rrq.order = append(rrq.order, o...)
	return rrq
}

// QueryCustomer chains the current query on the customer edge.
func (rrq *ReserveRoomQuery) QueryCustomer() *CustomerQuery {
	query := &CustomerQuery{config: rrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reserveroom.Table, reserveroom.FieldID, rrq.sqlQuery()),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, reserveroom.CustomerTable, reserveroom.CustomerColumn),
		)
		fromU = sqlgraph.SetNeighbors(rrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPromotion chains the current query on the promotion edge.
func (rrq *ReserveRoomQuery) QueryPromotion() *PromotionQuery {
	query := &PromotionQuery{config: rrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reserveroom.Table, reserveroom.FieldID, rrq.sqlQuery()),
			sqlgraph.To(promotion.Table, promotion.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, reserveroom.PromotionTable, reserveroom.PromotionColumn),
		)
		fromU = sqlgraph.SetNeighbors(rrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoom chains the current query on the room edge.
func (rrq *ReserveRoomQuery) QueryRoom() *DataRoomQuery {
	query := &DataRoomQuery{config: rrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reserveroom.Table, reserveroom.FieldID, rrq.sqlQuery()),
			sqlgraph.To(dataroom.Table, dataroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, reserveroom.RoomTable, reserveroom.RoomColumn),
		)
		fromU = sqlgraph.SetNeighbors(rrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCheckins chains the current query on the checkins edge.
func (rrq *ReserveRoomQuery) QueryCheckins() *CheckInQuery {
	query := &CheckInQuery{config: rrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reserveroom.Table, reserveroom.FieldID, rrq.sqlQuery()),
			sqlgraph.To(checkin.Table, checkin.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, reserveroom.CheckinsTable, reserveroom.CheckinsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ReserveRoom entity in the query. Returns *NotFoundError when no reserveroom was found.
func (rrq *ReserveRoomQuery) First(ctx context.Context) (*ReserveRoom, error) {
	rrs, err := rrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(rrs) == 0 {
		return nil, &NotFoundError{reserveroom.Label}
	}
	return rrs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rrq *ReserveRoomQuery) FirstX(ctx context.Context) *ReserveRoom {
	rr, err := rrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return rr
}

// FirstID returns the first ReserveRoom id in the query. Returns *NotFoundError when no id was found.
func (rrq *ReserveRoomQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{reserveroom.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (rrq *ReserveRoomQuery) FirstXID(ctx context.Context) int {
	id, err := rrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only ReserveRoom entity in the query, returns an error if not exactly one entity was returned.
func (rrq *ReserveRoomQuery) Only(ctx context.Context) (*ReserveRoom, error) {
	rrs, err := rrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(rrs) {
	case 1:
		return rrs[0], nil
	case 0:
		return nil, &NotFoundError{reserveroom.Label}
	default:
		return nil, &NotSingularError{reserveroom.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rrq *ReserveRoomQuery) OnlyX(ctx context.Context) *ReserveRoom {
	rr, err := rrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return rr
}

// OnlyID returns the only ReserveRoom id in the query, returns an error if not exactly one id was returned.
func (rrq *ReserveRoomQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{reserveroom.Label}
	default:
		err = &NotSingularError{reserveroom.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rrq *ReserveRoomQuery) OnlyIDX(ctx context.Context) int {
	id, err := rrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ReserveRooms.
func (rrq *ReserveRoomQuery) All(ctx context.Context) ([]*ReserveRoom, error) {
	if err := rrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rrq *ReserveRoomQuery) AllX(ctx context.Context) []*ReserveRoom {
	rrs, err := rrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return rrs
}

// IDs executes the query and returns a list of ReserveRoom ids.
func (rrq *ReserveRoomQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rrq.Select(reserveroom.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rrq *ReserveRoomQuery) IDsX(ctx context.Context) []int {
	ids, err := rrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rrq *ReserveRoomQuery) Count(ctx context.Context) (int, error) {
	if err := rrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rrq *ReserveRoomQuery) CountX(ctx context.Context) int {
	count, err := rrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rrq *ReserveRoomQuery) Exist(ctx context.Context) (bool, error) {
	if err := rrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rrq *ReserveRoomQuery) ExistX(ctx context.Context) bool {
	exist, err := rrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rrq *ReserveRoomQuery) Clone() *ReserveRoomQuery {
	return &ReserveRoomQuery{
		config:     rrq.config,
		limit:      rrq.limit,
		offset:     rrq.offset,
		order:      append([]OrderFunc{}, rrq.order...),
		unique:     append([]string{}, rrq.unique...),
		predicates: append([]predicate.ReserveRoom{}, rrq.predicates...),
		// clone intermediate query.
		sql:  rrq.sql.Clone(),
		path: rrq.path,
	}
}

//  WithCustomer tells the query-builder to eager-loads the nodes that are connected to
// the "customer" edge. The optional arguments used to configure the query builder of the edge.
func (rrq *ReserveRoomQuery) WithCustomer(opts ...func(*CustomerQuery)) *ReserveRoomQuery {
	query := &CustomerQuery{config: rrq.config}
	for _, opt := range opts {
		opt(query)
	}
	rrq.withCustomer = query
	return rrq
}

//  WithPromotion tells the query-builder to eager-loads the nodes that are connected to
// the "promotion" edge. The optional arguments used to configure the query builder of the edge.
func (rrq *ReserveRoomQuery) WithPromotion(opts ...func(*PromotionQuery)) *ReserveRoomQuery {
	query := &PromotionQuery{config: rrq.config}
	for _, opt := range opts {
		opt(query)
	}
	rrq.withPromotion = query
	return rrq
}

//  WithRoom tells the query-builder to eager-loads the nodes that are connected to
// the "room" edge. The optional arguments used to configure the query builder of the edge.
func (rrq *ReserveRoomQuery) WithRoom(opts ...func(*DataRoomQuery)) *ReserveRoomQuery {
	query := &DataRoomQuery{config: rrq.config}
	for _, opt := range opts {
		opt(query)
	}
	rrq.withRoom = query
	return rrq
}

//  WithCheckins tells the query-builder to eager-loads the nodes that are connected to
// the "checkins" edge. The optional arguments used to configure the query builder of the edge.
func (rrq *ReserveRoomQuery) WithCheckins(opts ...func(*CheckInQuery)) *ReserveRoomQuery {
	query := &CheckInQuery{config: rrq.config}
	for _, opt := range opts {
		opt(query)
	}
	rrq.withCheckins = query
	return rrq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ReserveDate time.Time `json:"reserve_date,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ReserveRoom.Query().
//		GroupBy(reserveroom.FieldReserveDate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rrq *ReserveRoomQuery) GroupBy(field string, fields ...string) *ReserveRoomGroupBy {
	group := &ReserveRoomGroupBy{config: rrq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rrq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		ReserveDate time.Time `json:"reserve_date,omitempty"`
//	}
//
//	client.ReserveRoom.Query().
//		Select(reserveroom.FieldReserveDate).
//		Scan(ctx, &v)
//
func (rrq *ReserveRoomQuery) Select(field string, fields ...string) *ReserveRoomSelect {
	selector := &ReserveRoomSelect{config: rrq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rrq.sqlQuery(), nil
	}
	return selector
}

func (rrq *ReserveRoomQuery) prepareQuery(ctx context.Context) error {
	if rrq.path != nil {
		prev, err := rrq.path(ctx)
		if err != nil {
			return err
		}
		rrq.sql = prev
	}
	return nil
}

func (rrq *ReserveRoomQuery) sqlAll(ctx context.Context) ([]*ReserveRoom, error) {
	var (
		nodes       = []*ReserveRoom{}
		withFKs     = rrq.withFKs
		_spec       = rrq.querySpec()
		loadedTypes = [4]bool{
			rrq.withCustomer != nil,
			rrq.withPromotion != nil,
			rrq.withRoom != nil,
			rrq.withCheckins != nil,
		}
	)
	if rrq.withCustomer != nil || rrq.withPromotion != nil || rrq.withRoom != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, reserveroom.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &ReserveRoom{config: rrq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, rrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rrq.withCustomer; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ReserveRoom)
		for i := range nodes {
			if fk := nodes[i].customer_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(customer.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "customer_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Customer = n
			}
		}
	}

	if query := rrq.withPromotion; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ReserveRoom)
		for i := range nodes {
			if fk := nodes[i].promotion_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(promotion.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "promotion_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Promotion = n
			}
		}
	}

	if query := rrq.withRoom; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ReserveRoom)
		for i := range nodes {
			if fk := nodes[i].room_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(dataroom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "room_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Room = n
			}
		}
	}

	if query := rrq.withCheckins; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ReserveRoom)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.CheckIn(func(s *sql.Selector) {
			s.Where(sql.InValues(reserveroom.CheckinsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.reserves_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "reserves_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "reserves_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Checkins = append(node.Edges.Checkins, n)
		}
	}

	return nodes, nil
}

func (rrq *ReserveRoomQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rrq.querySpec()
	return sqlgraph.CountNodes(ctx, rrq.driver, _spec)
}

func (rrq *ReserveRoomQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rrq *ReserveRoomQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reserveroom.Table,
			Columns: reserveroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: reserveroom.FieldID,
			},
		},
		From:   rrq.sql,
		Unique: true,
	}
	if ps := rrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rrq *ReserveRoomQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(rrq.driver.Dialect())
	t1 := builder.Table(reserveroom.Table)
	selector := builder.Select(t1.Columns(reserveroom.Columns...)...).From(t1)
	if rrq.sql != nil {
		selector = rrq.sql
		selector.Select(selector.Columns(reserveroom.Columns...)...)
	}
	for _, p := range rrq.predicates {
		p(selector)
	}
	for _, p := range rrq.order {
		p(selector)
	}
	if offset := rrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReserveRoomGroupBy is the builder for group-by ReserveRoom entities.
type ReserveRoomGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rrgb *ReserveRoomGroupBy) Aggregate(fns ...AggregateFunc) *ReserveRoomGroupBy {
	rrgb.fns = append(rrgb.fns, fns...)
	return rrgb
}

// Scan applies the group-by query and scan the result into the given value.
func (rrgb *ReserveRoomGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rrgb.path(ctx)
	if err != nil {
		return err
	}
	rrgb.sql = query
	return rrgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rrgb *ReserveRoomGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rrgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (rrgb *ReserveRoomGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rrgb.fields) > 1 {
		return nil, errors.New("ent: ReserveRoomGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rrgb *ReserveRoomGroupBy) StringsX(ctx context.Context) []string {
	v, err := rrgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (rrgb *ReserveRoomGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rrgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reserveroom.Label}
	default:
		err = fmt.Errorf("ent: ReserveRoomGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rrgb *ReserveRoomGroupBy) StringX(ctx context.Context) string {
	v, err := rrgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (rrgb *ReserveRoomGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rrgb.fields) > 1 {
		return nil, errors.New("ent: ReserveRoomGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rrgb *ReserveRoomGroupBy) IntsX(ctx context.Context) []int {
	v, err := rrgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (rrgb *ReserveRoomGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rrgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reserveroom.Label}
	default:
		err = fmt.Errorf("ent: ReserveRoomGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rrgb *ReserveRoomGroupBy) IntX(ctx context.Context) int {
	v, err := rrgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (rrgb *ReserveRoomGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rrgb.fields) > 1 {
		return nil, errors.New("ent: ReserveRoomGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rrgb *ReserveRoomGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rrgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (rrgb *ReserveRoomGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rrgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reserveroom.Label}
	default:
		err = fmt.Errorf("ent: ReserveRoomGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rrgb *ReserveRoomGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rrgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (rrgb *ReserveRoomGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rrgb.fields) > 1 {
		return nil, errors.New("ent: ReserveRoomGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rrgb *ReserveRoomGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rrgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (rrgb *ReserveRoomGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rrgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reserveroom.Label}
	default:
		err = fmt.Errorf("ent: ReserveRoomGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rrgb *ReserveRoomGroupBy) BoolX(ctx context.Context) bool {
	v, err := rrgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rrgb *ReserveRoomGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rrgb.sqlQuery().Query()
	if err := rrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rrgb *ReserveRoomGroupBy) sqlQuery() *sql.Selector {
	selector := rrgb.sql
	columns := make([]string, 0, len(rrgb.fields)+len(rrgb.fns))
	columns = append(columns, rrgb.fields...)
	for _, fn := range rrgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(rrgb.fields...)
}

// ReserveRoomSelect is the builder for select fields of ReserveRoom entities.
type ReserveRoomSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (rrs *ReserveRoomSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := rrs.path(ctx)
	if err != nil {
		return err
	}
	rrs.sql = query
	return rrs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rrs *ReserveRoomSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rrs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (rrs *ReserveRoomSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rrs.fields) > 1 {
		return nil, errors.New("ent: ReserveRoomSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rrs *ReserveRoomSelect) StringsX(ctx context.Context) []string {
	v, err := rrs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (rrs *ReserveRoomSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rrs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reserveroom.Label}
	default:
		err = fmt.Errorf("ent: ReserveRoomSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rrs *ReserveRoomSelect) StringX(ctx context.Context) string {
	v, err := rrs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (rrs *ReserveRoomSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rrs.fields) > 1 {
		return nil, errors.New("ent: ReserveRoomSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rrs *ReserveRoomSelect) IntsX(ctx context.Context) []int {
	v, err := rrs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (rrs *ReserveRoomSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rrs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reserveroom.Label}
	default:
		err = fmt.Errorf("ent: ReserveRoomSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rrs *ReserveRoomSelect) IntX(ctx context.Context) int {
	v, err := rrs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (rrs *ReserveRoomSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rrs.fields) > 1 {
		return nil, errors.New("ent: ReserveRoomSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rrs *ReserveRoomSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rrs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (rrs *ReserveRoomSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rrs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reserveroom.Label}
	default:
		err = fmt.Errorf("ent: ReserveRoomSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rrs *ReserveRoomSelect) Float64X(ctx context.Context) float64 {
	v, err := rrs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (rrs *ReserveRoomSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rrs.fields) > 1 {
		return nil, errors.New("ent: ReserveRoomSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rrs *ReserveRoomSelect) BoolsX(ctx context.Context) []bool {
	v, err := rrs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (rrs *ReserveRoomSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rrs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reserveroom.Label}
	default:
		err = fmt.Errorf("ent: ReserveRoomSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rrs *ReserveRoomSelect) BoolX(ctx context.Context) bool {
	v, err := rrs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rrs *ReserveRoomSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rrs.sqlQuery().Query()
	if err := rrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rrs *ReserveRoomSelect) sqlQuery() sql.Querier {
	selector := rrs.sql
	selector.Select(selector.Columns(rrs.fields...)...)
	return selector
}
