// Code generated by ent, DO NOT EDIT.

package ent

import (
	"DemoApp/ent/class"
	"DemoApp/ent/predicate"
	"DemoApp/ent/teacher"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeacherQuery is the builder for querying Teacher entities.
type TeacherQuery struct {
	config
	ctx         *QueryContext
	order       []teacher.OrderOption
	inters      []Interceptor
	predicates  []predicate.Teacher
	withClasses *ClassQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TeacherQuery builder.
func (tq *TeacherQuery) Where(ps ...predicate.Teacher) *TeacherQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TeacherQuery) Limit(limit int) *TeacherQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TeacherQuery) Offset(offset int) *TeacherQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TeacherQuery) Unique(unique bool) *TeacherQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TeacherQuery) Order(o ...teacher.OrderOption) *TeacherQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryClasses chains the current query on the "classes" edge.
func (tq *TeacherQuery) QueryClasses() *ClassQuery {
	query := (&ClassClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(teacher.Table, teacher.FieldID, selector),
			sqlgraph.To(class.Table, class.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, teacher.ClassesTable, teacher.ClassesColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Teacher entity from the query.
// Returns a *NotFoundError when no Teacher was found.
func (tq *TeacherQuery) First(ctx context.Context) (*Teacher, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{teacher.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TeacherQuery) FirstX(ctx context.Context) *Teacher {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Teacher ID from the query.
// Returns a *NotFoundError when no Teacher ID was found.
func (tq *TeacherQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{teacher.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TeacherQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Teacher entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Teacher entity is found.
// Returns a *NotFoundError when no Teacher entities are found.
func (tq *TeacherQuery) Only(ctx context.Context) (*Teacher, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{teacher.Label}
	default:
		return nil, &NotSingularError{teacher.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TeacherQuery) OnlyX(ctx context.Context) *Teacher {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Teacher ID in the query.
// Returns a *NotSingularError when more than one Teacher ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TeacherQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{teacher.Label}
	default:
		err = &NotSingularError{teacher.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TeacherQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Teachers.
func (tq *TeacherQuery) All(ctx context.Context) ([]*Teacher, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryAll)
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Teacher, *TeacherQuery]()
	return withInterceptors[[]*Teacher](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TeacherQuery) AllX(ctx context.Context) []*Teacher {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Teacher IDs.
func (tq *TeacherQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryIDs)
	if err = tq.Select(teacher.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TeacherQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TeacherQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryCount)
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TeacherQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TeacherQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TeacherQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryExist)
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TeacherQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TeacherQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TeacherQuery) Clone() *TeacherQuery {
	if tq == nil {
		return nil
	}
	return &TeacherQuery{
		config:      tq.config,
		ctx:         tq.ctx.Clone(),
		order:       append([]teacher.OrderOption{}, tq.order...),
		inters:      append([]Interceptor{}, tq.inters...),
		predicates:  append([]predicate.Teacher{}, tq.predicates...),
		withClasses: tq.withClasses.Clone(),
		// clone intermediate query.
		sql:       tq.sql.Clone(),
		path:      tq.path,
		modifiers: append([]func(*sql.Selector){}, tq.modifiers...),
	}
}

// WithClasses tells the query-builder to eager-load the nodes that are connected to
// the "classes" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TeacherQuery) WithClasses(opts ...func(*ClassQuery)) *TeacherQuery {
	query := (&ClassClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withClasses = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Teacher.Query().
//		GroupBy(teacher.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TeacherQuery) GroupBy(field string, fields ...string) *TeacherGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TeacherGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = teacher.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Teacher.Query().
//		Select(teacher.FieldName).
//		Scan(ctx, &v)
func (tq *TeacherQuery) Select(fields ...string) *TeacherSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TeacherSelect{TeacherQuery: tq}
	sbuild.label = teacher.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TeacherSelect configured with the given aggregations.
func (tq *TeacherQuery) Aggregate(fns ...AggregateFunc) *TeacherSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TeacherQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !teacher.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TeacherQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Teacher, error) {
	var (
		nodes       = []*Teacher{}
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withClasses != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Teacher).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Teacher{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withClasses; query != nil {
		if err := tq.loadClasses(ctx, query, nodes, nil,
			func(n *Teacher, e *Class) { n.Edges.Classes = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TeacherQuery) loadClasses(ctx context.Context, query *ClassQuery, nodes []*Teacher, init func(*Teacher), assign func(*Teacher, *Class)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Teacher)
	for i := range nodes {
		fk := nodes[i].ClassID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(class.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "class_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *TeacherQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TeacherQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(teacher.Table, teacher.Columns, sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teacher.FieldID)
		for i := range fields {
			if fields[i] != teacher.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tq.withClasses != nil {
			_spec.Node.AddColumnOnce(teacher.FieldClassID)
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TeacherQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(teacher.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = teacher.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range tq.modifiers {
		m(selector)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tq *TeacherQuery) Modify(modifiers ...func(s *sql.Selector)) *TeacherSelect {
	tq.modifiers = append(tq.modifiers, modifiers...)
	return tq.Select()
}

// TeacherGroupBy is the group-by builder for Teacher entities.
type TeacherGroupBy struct {
	selector
	build *TeacherQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TeacherGroupBy) Aggregate(fns ...AggregateFunc) *TeacherGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TeacherGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, ent.OpQueryGroupBy)
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TeacherQuery, *TeacherGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TeacherGroupBy) sqlScan(ctx context.Context, root *TeacherQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TeacherSelect is the builder for selecting fields of Teacher entities.
type TeacherSelect struct {
	*TeacherQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TeacherSelect) Aggregate(fns ...AggregateFunc) *TeacherSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TeacherSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, ent.OpQuerySelect)
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TeacherQuery, *TeacherSelect](ctx, ts.TeacherQuery, ts, ts.inters, v)
}

func (ts *TeacherSelect) sqlScan(ctx context.Context, root *TeacherQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ts *TeacherSelect) Modify(modifiers ...func(s *sql.Selector)) *TeacherSelect {
	ts.modifiers = append(ts.modifiers, modifiers...)
	return ts
}
