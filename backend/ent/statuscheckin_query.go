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
	"github.com/team18/app/ent/predicate"
	"github.com/team18/app/ent/statuscheckin"
)

// StatusCheckInQuery is the builder for querying StatusCheckIn entities.
type StatusCheckInQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.StatusCheckIn
	// eager-loading edges.
	withCheckins *CheckInQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (sciq *StatusCheckInQuery) Where(ps ...predicate.StatusCheckIn) *StatusCheckInQuery {
	sciq.predicates = append(sciq.predicates, ps...)
	return sciq
}

// Limit adds a limit step to the query.
func (sciq *StatusCheckInQuery) Limit(limit int) *StatusCheckInQuery {
	sciq.limit = &limit
	return sciq
}

// Offset adds an offset step to the query.
func (sciq *StatusCheckInQuery) Offset(offset int) *StatusCheckInQuery {
	sciq.offset = &offset
	return sciq
}

// Order adds an order step to the query.
func (sciq *StatusCheckInQuery) Order(o ...OrderFunc) *StatusCheckInQuery {
	sciq.order = append(sciq.order, o...)
	return sciq
}

// QueryCheckins chains the current query on the checkins edge.
func (sciq *StatusCheckInQuery) QueryCheckins() *CheckInQuery {
	query := &CheckInQuery{config: sciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(statuscheckin.Table, statuscheckin.FieldID, sciq.sqlQuery()),
			sqlgraph.To(checkin.Table, checkin.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, statuscheckin.CheckinsTable, statuscheckin.CheckinsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StatusCheckIn entity in the query. Returns *NotFoundError when no statuscheckin was found.
func (sciq *StatusCheckInQuery) First(ctx context.Context) (*StatusCheckIn, error) {
	scis, err := sciq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(scis) == 0 {
		return nil, &NotFoundError{statuscheckin.Label}
	}
	return scis[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sciq *StatusCheckInQuery) FirstX(ctx context.Context) *StatusCheckIn {
	sci, err := sciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return sci
}

// FirstID returns the first StatusCheckIn id in the query. Returns *NotFoundError when no id was found.
func (sciq *StatusCheckInQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sciq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{statuscheckin.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (sciq *StatusCheckInQuery) FirstXID(ctx context.Context) int {
	id, err := sciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only StatusCheckIn entity in the query, returns an error if not exactly one entity was returned.
func (sciq *StatusCheckInQuery) Only(ctx context.Context) (*StatusCheckIn, error) {
	scis, err := sciq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(scis) {
	case 1:
		return scis[0], nil
	case 0:
		return nil, &NotFoundError{statuscheckin.Label}
	default:
		return nil, &NotSingularError{statuscheckin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sciq *StatusCheckInQuery) OnlyX(ctx context.Context) *StatusCheckIn {
	sci, err := sciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return sci
}

// OnlyID returns the only StatusCheckIn id in the query, returns an error if not exactly one id was returned.
func (sciq *StatusCheckInQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sciq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{statuscheckin.Label}
	default:
		err = &NotSingularError{statuscheckin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sciq *StatusCheckInQuery) OnlyIDX(ctx context.Context) int {
	id, err := sciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StatusCheckIns.
func (sciq *StatusCheckInQuery) All(ctx context.Context) ([]*StatusCheckIn, error) {
	if err := sciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sciq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sciq *StatusCheckInQuery) AllX(ctx context.Context) []*StatusCheckIn {
	scis, err := sciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return scis
}

// IDs executes the query and returns a list of StatusCheckIn ids.
func (sciq *StatusCheckInQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sciq.Select(statuscheckin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sciq *StatusCheckInQuery) IDsX(ctx context.Context) []int {
	ids, err := sciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sciq *StatusCheckInQuery) Count(ctx context.Context) (int, error) {
	if err := sciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sciq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sciq *StatusCheckInQuery) CountX(ctx context.Context) int {
	count, err := sciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sciq *StatusCheckInQuery) Exist(ctx context.Context) (bool, error) {
	if err := sciq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sciq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sciq *StatusCheckInQuery) ExistX(ctx context.Context) bool {
	exist, err := sciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sciq *StatusCheckInQuery) Clone() *StatusCheckInQuery {
	return &StatusCheckInQuery{
		config:     sciq.config,
		limit:      sciq.limit,
		offset:     sciq.offset,
		order:      append([]OrderFunc{}, sciq.order...),
		unique:     append([]string{}, sciq.unique...),
		predicates: append([]predicate.StatusCheckIn{}, sciq.predicates...),
		// clone intermediate query.
		sql:  sciq.sql.Clone(),
		path: sciq.path,
	}
}

//  WithCheckins tells the query-builder to eager-loads the nodes that are connected to
// the "checkins" edge. The optional arguments used to configure the query builder of the edge.
func (sciq *StatusCheckInQuery) WithCheckins(opts ...func(*CheckInQuery)) *StatusCheckInQuery {
	query := &CheckInQuery{config: sciq.config}
	for _, opt := range opts {
		opt(query)
	}
	sciq.withCheckins = query
	return sciq
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
//	client.StatusCheckIn.Query().
//		GroupBy(statuscheckin.FieldStatusName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sciq *StatusCheckInQuery) GroupBy(field string, fields ...string) *StatusCheckInGroupBy {
	group := &StatusCheckInGroupBy{config: sciq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sciq.sqlQuery(), nil
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
//	client.StatusCheckIn.Query().
//		Select(statuscheckin.FieldStatusName).
//		Scan(ctx, &v)
//
func (sciq *StatusCheckInQuery) Select(field string, fields ...string) *StatusCheckInSelect {
	selector := &StatusCheckInSelect{config: sciq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sciq.sqlQuery(), nil
	}
	return selector
}

func (sciq *StatusCheckInQuery) prepareQuery(ctx context.Context) error {
	if sciq.path != nil {
		prev, err := sciq.path(ctx)
		if err != nil {
			return err
		}
		sciq.sql = prev
	}
	return nil
}

func (sciq *StatusCheckInQuery) sqlAll(ctx context.Context) ([]*StatusCheckIn, error) {
	var (
		nodes       = []*StatusCheckIn{}
		_spec       = sciq.querySpec()
		loadedTypes = [1]bool{
			sciq.withCheckins != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &StatusCheckIn{config: sciq.config}
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
	if err := sqlgraph.QueryNodes(ctx, sciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sciq.withCheckins; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*StatusCheckIn)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.CheckIn(func(s *sql.Selector) {
			s.Where(sql.InValues(statuscheckin.CheckinsColumn, fks...))
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
			node.Edges.Checkins = append(node.Edges.Checkins, n)
		}
	}

	return nodes, nil
}

func (sciq *StatusCheckInQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sciq.querySpec()
	return sqlgraph.CountNodes(ctx, sciq.driver, _spec)
}

func (sciq *StatusCheckInQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sciq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (sciq *StatusCheckInQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statuscheckin.Table,
			Columns: statuscheckin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statuscheckin.FieldID,
			},
		},
		From:   sciq.sql,
		Unique: true,
	}
	if ps := sciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sciq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sciq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sciq *StatusCheckInQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(sciq.driver.Dialect())
	t1 := builder.Table(statuscheckin.Table)
	selector := builder.Select(t1.Columns(statuscheckin.Columns...)...).From(t1)
	if sciq.sql != nil {
		selector = sciq.sql
		selector.Select(selector.Columns(statuscheckin.Columns...)...)
	}
	for _, p := range sciq.predicates {
		p(selector)
	}
	for _, p := range sciq.order {
		p(selector)
	}
	if offset := sciq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sciq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StatusCheckInGroupBy is the builder for group-by StatusCheckIn entities.
type StatusCheckInGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scigb *StatusCheckInGroupBy) Aggregate(fns ...AggregateFunc) *StatusCheckInGroupBy {
	scigb.fns = append(scigb.fns, fns...)
	return scigb
}

// Scan applies the group-by query and scan the result into the given value.
func (scigb *StatusCheckInGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := scigb.path(ctx)
	if err != nil {
		return err
	}
	scigb.sql = query
	return scigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (scigb *StatusCheckInGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := scigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (scigb *StatusCheckInGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(scigb.fields) > 1 {
		return nil, errors.New("ent: StatusCheckInGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := scigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (scigb *StatusCheckInGroupBy) StringsX(ctx context.Context) []string {
	v, err := scigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (scigb *StatusCheckInGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = scigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statuscheckin.Label}
	default:
		err = fmt.Errorf("ent: StatusCheckInGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (scigb *StatusCheckInGroupBy) StringX(ctx context.Context) string {
	v, err := scigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (scigb *StatusCheckInGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(scigb.fields) > 1 {
		return nil, errors.New("ent: StatusCheckInGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := scigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (scigb *StatusCheckInGroupBy) IntsX(ctx context.Context) []int {
	v, err := scigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (scigb *StatusCheckInGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = scigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statuscheckin.Label}
	default:
		err = fmt.Errorf("ent: StatusCheckInGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (scigb *StatusCheckInGroupBy) IntX(ctx context.Context) int {
	v, err := scigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (scigb *StatusCheckInGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(scigb.fields) > 1 {
		return nil, errors.New("ent: StatusCheckInGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := scigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (scigb *StatusCheckInGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := scigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (scigb *StatusCheckInGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = scigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statuscheckin.Label}
	default:
		err = fmt.Errorf("ent: StatusCheckInGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (scigb *StatusCheckInGroupBy) Float64X(ctx context.Context) float64 {
	v, err := scigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (scigb *StatusCheckInGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(scigb.fields) > 1 {
		return nil, errors.New("ent: StatusCheckInGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := scigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (scigb *StatusCheckInGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := scigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (scigb *StatusCheckInGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = scigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statuscheckin.Label}
	default:
		err = fmt.Errorf("ent: StatusCheckInGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (scigb *StatusCheckInGroupBy) BoolX(ctx context.Context) bool {
	v, err := scigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (scigb *StatusCheckInGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := scigb.sqlQuery().Query()
	if err := scigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (scigb *StatusCheckInGroupBy) sqlQuery() *sql.Selector {
	selector := scigb.sql
	columns := make([]string, 0, len(scigb.fields)+len(scigb.fns))
	columns = append(columns, scigb.fields...)
	for _, fn := range scigb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(scigb.fields...)
}

// StatusCheckInSelect is the builder for select fields of StatusCheckIn entities.
type StatusCheckInSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (scis *StatusCheckInSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := scis.path(ctx)
	if err != nil {
		return err
	}
	scis.sql = query
	return scis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (scis *StatusCheckInSelect) ScanX(ctx context.Context, v interface{}) {
	if err := scis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (scis *StatusCheckInSelect) Strings(ctx context.Context) ([]string, error) {
	if len(scis.fields) > 1 {
		return nil, errors.New("ent: StatusCheckInSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := scis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (scis *StatusCheckInSelect) StringsX(ctx context.Context) []string {
	v, err := scis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (scis *StatusCheckInSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = scis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statuscheckin.Label}
	default:
		err = fmt.Errorf("ent: StatusCheckInSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (scis *StatusCheckInSelect) StringX(ctx context.Context) string {
	v, err := scis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (scis *StatusCheckInSelect) Ints(ctx context.Context) ([]int, error) {
	if len(scis.fields) > 1 {
		return nil, errors.New("ent: StatusCheckInSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := scis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (scis *StatusCheckInSelect) IntsX(ctx context.Context) []int {
	v, err := scis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (scis *StatusCheckInSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = scis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statuscheckin.Label}
	default:
		err = fmt.Errorf("ent: StatusCheckInSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (scis *StatusCheckInSelect) IntX(ctx context.Context) int {
	v, err := scis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (scis *StatusCheckInSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(scis.fields) > 1 {
		return nil, errors.New("ent: StatusCheckInSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := scis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (scis *StatusCheckInSelect) Float64sX(ctx context.Context) []float64 {
	v, err := scis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (scis *StatusCheckInSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = scis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statuscheckin.Label}
	default:
		err = fmt.Errorf("ent: StatusCheckInSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (scis *StatusCheckInSelect) Float64X(ctx context.Context) float64 {
	v, err := scis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (scis *StatusCheckInSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(scis.fields) > 1 {
		return nil, errors.New("ent: StatusCheckInSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := scis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (scis *StatusCheckInSelect) BoolsX(ctx context.Context) []bool {
	v, err := scis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (scis *StatusCheckInSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = scis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statuscheckin.Label}
	default:
		err = fmt.Errorf("ent: StatusCheckInSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (scis *StatusCheckInSelect) BoolX(ctx context.Context) bool {
	v, err := scis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (scis *StatusCheckInSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := scis.sqlQuery().Query()
	if err := scis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (scis *StatusCheckInSelect) sqlQuery() sql.Querier {
	selector := scis.sql
	selector.Select(selector.Columns(scis.fields...)...)
	return selector
}