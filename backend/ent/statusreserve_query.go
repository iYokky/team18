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
	"github.com/team18/app/ent/predicate"
	"github.com/team18/app/ent/reserveroom"
	"github.com/team18/app/ent/statusreserve"
)

// StatusReserveQuery is the builder for querying StatusReserve entities.
type StatusReserveQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.StatusReserve
	// eager-loading edges.
	withReserves *ReserveRoomQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (srq *StatusReserveQuery) Where(ps ...predicate.StatusReserve) *StatusReserveQuery {
	srq.predicates = append(srq.predicates, ps...)
	return srq
}

// Limit adds a limit step to the query.
func (srq *StatusReserveQuery) Limit(limit int) *StatusReserveQuery {
	srq.limit = &limit
	return srq
}

// Offset adds an offset step to the query.
func (srq *StatusReserveQuery) Offset(offset int) *StatusReserveQuery {
	srq.offset = &offset
	return srq
}

// Order adds an order step to the query.
func (srq *StatusReserveQuery) Order(o ...OrderFunc) *StatusReserveQuery {
	srq.order = append(srq.order, o...)
	return srq
}

// QueryReserves chains the current query on the reserves edge.
func (srq *StatusReserveQuery) QueryReserves() *ReserveRoomQuery {
	query := &ReserveRoomQuery{config: srq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(statusreserve.Table, statusreserve.FieldID, srq.sqlQuery()),
			sqlgraph.To(reserveroom.Table, reserveroom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, statusreserve.ReservesTable, statusreserve.ReservesColumn),
		)
		fromU = sqlgraph.SetNeighbors(srq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StatusReserve entity in the query. Returns *NotFoundError when no statusreserve was found.
func (srq *StatusReserveQuery) First(ctx context.Context) (*StatusReserve, error) {
	srs, err := srq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(srs) == 0 {
		return nil, &NotFoundError{statusreserve.Label}
	}
	return srs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (srq *StatusReserveQuery) FirstX(ctx context.Context) *StatusReserve {
	sr, err := srq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return sr
}

// FirstID returns the first StatusReserve id in the query. Returns *NotFoundError when no id was found.
func (srq *StatusReserveQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = srq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{statusreserve.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (srq *StatusReserveQuery) FirstXID(ctx context.Context) int {
	id, err := srq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only StatusReserve entity in the query, returns an error if not exactly one entity was returned.
func (srq *StatusReserveQuery) Only(ctx context.Context) (*StatusReserve, error) {
	srs, err := srq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(srs) {
	case 1:
		return srs[0], nil
	case 0:
		return nil, &NotFoundError{statusreserve.Label}
	default:
		return nil, &NotSingularError{statusreserve.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (srq *StatusReserveQuery) OnlyX(ctx context.Context) *StatusReserve {
	sr, err := srq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return sr
}

// OnlyID returns the only StatusReserve id in the query, returns an error if not exactly one id was returned.
func (srq *StatusReserveQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = srq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{statusreserve.Label}
	default:
		err = &NotSingularError{statusreserve.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (srq *StatusReserveQuery) OnlyIDX(ctx context.Context) int {
	id, err := srq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StatusReserves.
func (srq *StatusReserveQuery) All(ctx context.Context) ([]*StatusReserve, error) {
	if err := srq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return srq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (srq *StatusReserveQuery) AllX(ctx context.Context) []*StatusReserve {
	srs, err := srq.All(ctx)
	if err != nil {
		panic(err)
	}
	return srs
}

// IDs executes the query and returns a list of StatusReserve ids.
func (srq *StatusReserveQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := srq.Select(statusreserve.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (srq *StatusReserveQuery) IDsX(ctx context.Context) []int {
	ids, err := srq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (srq *StatusReserveQuery) Count(ctx context.Context) (int, error) {
	if err := srq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return srq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (srq *StatusReserveQuery) CountX(ctx context.Context) int {
	count, err := srq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (srq *StatusReserveQuery) Exist(ctx context.Context) (bool, error) {
	if err := srq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return srq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (srq *StatusReserveQuery) ExistX(ctx context.Context) bool {
	exist, err := srq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (srq *StatusReserveQuery) Clone() *StatusReserveQuery {
	return &StatusReserveQuery{
		config:     srq.config,
		limit:      srq.limit,
		offset:     srq.offset,
		order:      append([]OrderFunc{}, srq.order...),
		unique:     append([]string{}, srq.unique...),
		predicates: append([]predicate.StatusReserve{}, srq.predicates...),
		// clone intermediate query.
		sql:  srq.sql.Clone(),
		path: srq.path,
	}
}

//  WithReserves tells the query-builder to eager-loads the nodes that are connected to
// the "reserves" edge. The optional arguments used to configure the query builder of the edge.
func (srq *StatusReserveQuery) WithReserves(opts ...func(*ReserveRoomQuery)) *StatusReserveQuery {
	query := &ReserveRoomQuery{config: srq.config}
	for _, opt := range opts {
		opt(query)
	}
	srq.withReserves = query
	return srq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		StatusName string `json:"status_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StatusReserve.Query().
//		GroupBy(statusreserve.FieldStatusName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (srq *StatusReserveQuery) GroupBy(field string, fields ...string) *StatusReserveGroupBy {
	group := &StatusReserveGroupBy{config: srq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return srq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		StatusName string `json:"status_name,omitempty"`
//	}
//
//	client.StatusReserve.Query().
//		Select(statusreserve.FieldStatusName).
//		Scan(ctx, &v)
//
func (srq *StatusReserveQuery) Select(field string, fields ...string) *StatusReserveSelect {
	selector := &StatusReserveSelect{config: srq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return srq.sqlQuery(), nil
	}
	return selector
}

func (srq *StatusReserveQuery) prepareQuery(ctx context.Context) error {
	if srq.path != nil {
		prev, err := srq.path(ctx)
		if err != nil {
			return err
		}
		srq.sql = prev
	}
	return nil
}

func (srq *StatusReserveQuery) sqlAll(ctx context.Context) ([]*StatusReserve, error) {
	var (
		nodes       = []*StatusReserve{}
		_spec       = srq.querySpec()
		loadedTypes = [1]bool{
			srq.withReserves != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &StatusReserve{config: srq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
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
	if err := sqlgraph.QueryNodes(ctx, srq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := srq.withReserves; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*StatusReserve)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.ReserveRoom(func(s *sql.Selector) {
			s.Where(sql.InValues(statusreserve.ReservesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.status_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "status_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "status_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Reserves = append(node.Edges.Reserves, n)
		}
	}

	return nodes, nil
}

func (srq *StatusReserveQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := srq.querySpec()
	return sqlgraph.CountNodes(ctx, srq.driver, _spec)
}

func (srq *StatusReserveQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := srq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (srq *StatusReserveQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statusreserve.Table,
			Columns: statusreserve.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statusreserve.FieldID,
			},
		},
		From:   srq.sql,
		Unique: true,
	}
	if ps := srq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := srq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := srq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := srq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (srq *StatusReserveQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(srq.driver.Dialect())
	t1 := builder.Table(statusreserve.Table)
	selector := builder.Select(t1.Columns(statusreserve.Columns...)...).From(t1)
	if srq.sql != nil {
		selector = srq.sql
		selector.Select(selector.Columns(statusreserve.Columns...)...)
	}
	for _, p := range srq.predicates {
		p(selector)
	}
	for _, p := range srq.order {
		p(selector)
	}
	if offset := srq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := srq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StatusReserveGroupBy is the builder for group-by StatusReserve entities.
type StatusReserveGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (srgb *StatusReserveGroupBy) Aggregate(fns ...AggregateFunc) *StatusReserveGroupBy {
	srgb.fns = append(srgb.fns, fns...)
	return srgb
}

// Scan applies the group-by query and scan the result into the given value.
func (srgb *StatusReserveGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := srgb.path(ctx)
	if err != nil {
		return err
	}
	srgb.sql = query
	return srgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (srgb *StatusReserveGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := srgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (srgb *StatusReserveGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(srgb.fields) > 1 {
		return nil, errors.New("ent: StatusReserveGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := srgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (srgb *StatusReserveGroupBy) StringsX(ctx context.Context) []string {
	v, err := srgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (srgb *StatusReserveGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = srgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusreserve.Label}
	default:
		err = fmt.Errorf("ent: StatusReserveGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (srgb *StatusReserveGroupBy) StringX(ctx context.Context) string {
	v, err := srgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (srgb *StatusReserveGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(srgb.fields) > 1 {
		return nil, errors.New("ent: StatusReserveGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := srgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (srgb *StatusReserveGroupBy) IntsX(ctx context.Context) []int {
	v, err := srgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (srgb *StatusReserveGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = srgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusreserve.Label}
	default:
		err = fmt.Errorf("ent: StatusReserveGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (srgb *StatusReserveGroupBy) IntX(ctx context.Context) int {
	v, err := srgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (srgb *StatusReserveGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(srgb.fields) > 1 {
		return nil, errors.New("ent: StatusReserveGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := srgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (srgb *StatusReserveGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := srgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (srgb *StatusReserveGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = srgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusreserve.Label}
	default:
		err = fmt.Errorf("ent: StatusReserveGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (srgb *StatusReserveGroupBy) Float64X(ctx context.Context) float64 {
	v, err := srgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (srgb *StatusReserveGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(srgb.fields) > 1 {
		return nil, errors.New("ent: StatusReserveGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := srgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (srgb *StatusReserveGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := srgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (srgb *StatusReserveGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = srgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusreserve.Label}
	default:
		err = fmt.Errorf("ent: StatusReserveGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (srgb *StatusReserveGroupBy) BoolX(ctx context.Context) bool {
	v, err := srgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (srgb *StatusReserveGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := srgb.sqlQuery().Query()
	if err := srgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (srgb *StatusReserveGroupBy) sqlQuery() *sql.Selector {
	selector := srgb.sql
	columns := make([]string, 0, len(srgb.fields)+len(srgb.fns))
	columns = append(columns, srgb.fields...)
	for _, fn := range srgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(srgb.fields...)
}

// StatusReserveSelect is the builder for select fields of StatusReserve entities.
type StatusReserveSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (srs *StatusReserveSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := srs.path(ctx)
	if err != nil {
		return err
	}
	srs.sql = query
	return srs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (srs *StatusReserveSelect) ScanX(ctx context.Context, v interface{}) {
	if err := srs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (srs *StatusReserveSelect) Strings(ctx context.Context) ([]string, error) {
	if len(srs.fields) > 1 {
		return nil, errors.New("ent: StatusReserveSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := srs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (srs *StatusReserveSelect) StringsX(ctx context.Context) []string {
	v, err := srs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (srs *StatusReserveSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = srs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusreserve.Label}
	default:
		err = fmt.Errorf("ent: StatusReserveSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (srs *StatusReserveSelect) StringX(ctx context.Context) string {
	v, err := srs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (srs *StatusReserveSelect) Ints(ctx context.Context) ([]int, error) {
	if len(srs.fields) > 1 {
		return nil, errors.New("ent: StatusReserveSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := srs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (srs *StatusReserveSelect) IntsX(ctx context.Context) []int {
	v, err := srs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (srs *StatusReserveSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = srs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusreserve.Label}
	default:
		err = fmt.Errorf("ent: StatusReserveSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (srs *StatusReserveSelect) IntX(ctx context.Context) int {
	v, err := srs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (srs *StatusReserveSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(srs.fields) > 1 {
		return nil, errors.New("ent: StatusReserveSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := srs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (srs *StatusReserveSelect) Float64sX(ctx context.Context) []float64 {
	v, err := srs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (srs *StatusReserveSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = srs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusreserve.Label}
	default:
		err = fmt.Errorf("ent: StatusReserveSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (srs *StatusReserveSelect) Float64X(ctx context.Context) float64 {
	v, err := srs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (srs *StatusReserveSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(srs.fields) > 1 {
		return nil, errors.New("ent: StatusReserveSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := srs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (srs *StatusReserveSelect) BoolsX(ctx context.Context) []bool {
	v, err := srs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (srs *StatusReserveSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = srs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusreserve.Label}
	default:
		err = fmt.Errorf("ent: StatusReserveSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (srs *StatusReserveSelect) BoolX(ctx context.Context) bool {
	v, err := srs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (srs *StatusReserveSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := srs.sqlQuery().Query()
	if err := srs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (srs *StatusReserveSelect) sqlQuery() sql.Querier {
	selector := srs.sql
	selector.Select(selector.Columns(srs.fields...)...)
	return selector
}