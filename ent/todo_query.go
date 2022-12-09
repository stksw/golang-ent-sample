// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"ent-sample/ent/predicate"
	"ent-sample/ent/todo"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoQuery is the builder for querying Todo entities.
type TodoQuery struct {
	config
	limit             *int
	offset            *int
	unique            *bool
	order             []OrderFunc
	fields            []string
	predicates        []predicate.Todo
	withChildren      *TodoQuery
	withParent        *TodoQuery
	withFKs           bool
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*Todo) error
	withNamedChildren map[string]*TodoQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TodoQuery builder.
func (tq *TodoQuery) Where(ps ...predicate.Todo) *TodoQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TodoQuery) Limit(limit int) *TodoQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TodoQuery) Offset(offset int) *TodoQuery {
	tq.offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TodoQuery) Unique(unique bool) *TodoQuery {
	tq.unique = &unique
	return tq
}

// Order adds an order step to the query.
func (tq *TodoQuery) Order(o ...OrderFunc) *TodoQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryChildren chains the current query on the "children" edge.
func (tq *TodoQuery) QueryChildren() *TodoQuery {
	query := &TodoQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, selector),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, todo.ChildrenTable, todo.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (tq *TodoQuery) QueryParent() *TodoQuery {
	query := &TodoQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, selector),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, todo.ParentTable, todo.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Todo entity from the query.
// Returns a *NotFoundError when no Todo was found.
func (tq *TodoQuery) First(ctx context.Context) (*Todo, error) {
	nodes, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{todo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TodoQuery) FirstX(ctx context.Context) *Todo {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Todo ID from the query.
// Returns a *NotFoundError when no Todo ID was found.
func (tq *TodoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{todo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TodoQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Todo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Todo entity is found.
// Returns a *NotFoundError when no Todo entities are found.
func (tq *TodoQuery) Only(ctx context.Context) (*Todo, error) {
	nodes, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{todo.Label}
	default:
		return nil, &NotSingularError{todo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TodoQuery) OnlyX(ctx context.Context) *Todo {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Todo ID in the query.
// Returns a *NotSingularError when more than one Todo ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TodoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{todo.Label}
	default:
		err = &NotSingularError{todo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TodoQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Todos.
func (tq *TodoQuery) All(ctx context.Context) ([]*Todo, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TodoQuery) AllX(ctx context.Context) []*Todo {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Todo IDs.
func (tq *TodoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tq.Select(todo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TodoQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TodoQuery) Count(ctx context.Context) (int, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TodoQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TodoQuery) Exist(ctx context.Context) (bool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TodoQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TodoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TodoQuery) Clone() *TodoQuery {
	if tq == nil {
		return nil
	}
	return &TodoQuery{
		config:       tq.config,
		limit:        tq.limit,
		offset:       tq.offset,
		order:        append([]OrderFunc{}, tq.order...),
		predicates:   append([]predicate.Todo{}, tq.predicates...),
		withChildren: tq.withChildren.Clone(),
		withParent:   tq.withParent.Clone(),
		// clone intermediate query.
		sql:    tq.sql.Clone(),
		path:   tq.path,
		unique: tq.unique,
	}
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TodoQuery) WithChildren(opts ...func(*TodoQuery)) *TodoQuery {
	query := &TodoQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withChildren = query
	return tq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TodoQuery) WithParent(opts ...func(*TodoQuery)) *TodoQuery {
	query := &TodoQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withParent = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Todo.Query().
//		GroupBy(todo.FieldText).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TodoQuery) GroupBy(field string, fields ...string) *TodoGroupBy {
	grbuild := &TodoGroupBy{config: tq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(ctx), nil
	}
	grbuild.label = todo.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//	}
//
//	client.Todo.Query().
//		Select(todo.FieldText).
//		Scan(ctx, &v)
func (tq *TodoQuery) Select(fields ...string) *TodoSelect {
	tq.fields = append(tq.fields, fields...)
	selbuild := &TodoSelect{TodoQuery: tq}
	selbuild.label = todo.Label
	selbuild.flds, selbuild.scan = &tq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a TodoSelect configured with the given aggregations.
func (tq *TodoQuery) Aggregate(fns ...AggregateFunc) *TodoSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TodoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tq.fields {
		if !todo.ValidColumn(f) {
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

func (tq *TodoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Todo, error) {
	var (
		nodes       = []*Todo{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withChildren != nil,
			tq.withParent != nil,
		}
	)
	if tq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, todo.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Todo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Todo{config: tq.config}
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
	if query := tq.withChildren; query != nil {
		if err := tq.loadChildren(ctx, query, nodes,
			func(n *Todo) { n.Edges.Children = []*Todo{} },
			func(n *Todo, e *Todo) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withParent; query != nil {
		if err := tq.loadParent(ctx, query, nodes, nil,
			func(n *Todo, e *Todo) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range tq.withNamedChildren {
		if err := tq.loadChildren(ctx, query, nodes,
			func(n *Todo) { n.appendNamedChildren(name) },
			func(n *Todo, e *Todo) { n.appendNamedChildren(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range tq.loadTotal {
		if err := tq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TodoQuery) loadChildren(ctx context.Context, query *TodoQuery, nodes []*Todo, init func(*Todo), assign func(*Todo, *Todo)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Todo)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.InValues(todo.ChildrenColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.todo_parent
		if fk == nil {
			return fmt.Errorf(`foreign-key "todo_parent" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "todo_parent" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *TodoQuery) loadParent(ctx context.Context, query *TodoQuery, nodes []*Todo, init func(*Todo), assign func(*Todo, *Todo)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Todo)
	for i := range nodes {
		if nodes[i].todo_parent == nil {
			continue
		}
		fk := *nodes[i].todo_parent
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(todo.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "todo_parent" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *TodoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.fields
	if len(tq.fields) > 0 {
		_spec.Unique = tq.unique != nil && *tq.unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TodoQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (tq *TodoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   todo.Table,
			Columns: todo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: todo.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todo.FieldID)
		for i := range fields {
			if fields[i] != todo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
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

func (tq *TodoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(todo.Table)
	columns := tq.fields
	if len(columns) == 0 {
		columns = todo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.unique != nil && *tq.unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedChildren tells the query-builder to eager-load the nodes that are connected to the "children"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tq *TodoQuery) WithNamedChildren(name string, opts ...func(*TodoQuery)) *TodoQuery {
	query := &TodoQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	if tq.withNamedChildren == nil {
		tq.withNamedChildren = make(map[string]*TodoQuery)
	}
	tq.withNamedChildren[name] = query
	return tq
}

// TodoGroupBy is the group-by builder for Todo entities.
type TodoGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TodoGroupBy) Aggregate(fns ...AggregateFunc) *TodoGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tgb *TodoGroupBy) Scan(ctx context.Context, v any) error {
	query, err := tgb.path(ctx)
	if err != nil {
		return err
	}
	tgb.sql = query
	return tgb.sqlScan(ctx, v)
}

func (tgb *TodoGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range tgb.fields {
		if !todo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TodoGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql.Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
		for _, f := range tgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tgb.fields...)...)
}

// TodoSelect is the builder for selecting fields of Todo entities.
type TodoSelect struct {
	*TodoQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TodoSelect) Aggregate(fns ...AggregateFunc) *TodoSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TodoSelect) Scan(ctx context.Context, v any) error {
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	ts.sql = ts.TodoQuery.sqlQuery(ctx)
	return ts.sqlScan(ctx, v)
}

func (ts *TodoSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(ts.sql))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ts.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ts.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ts.sql.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
