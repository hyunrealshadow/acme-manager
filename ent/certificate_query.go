// Code generated by ent, DO NOT EDIT.

package ent

import (
	"acme-manager/ent/acmeaccount"
	"acme-manager/ent/certificate"
	"acme-manager/ent/dnsprovider"
	"acme-manager/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CertificateQuery is the builder for querying Certificate entities.
type CertificateQuery struct {
	config
	ctx             *QueryContext
	order           []certificate.OrderOption
	inters          []Interceptor
	predicates      []predicate.Certificate
	withAcmeAccount *AcmeAccountQuery
	withDNSProvider *DnsProviderQuery
	modifiers       []func(*sql.Selector)
	loadTotal       []func(context.Context, []*Certificate) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CertificateQuery builder.
func (cq *CertificateQuery) Where(ps ...predicate.Certificate) *CertificateQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CertificateQuery) Limit(limit int) *CertificateQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CertificateQuery) Offset(offset int) *CertificateQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CertificateQuery) Unique(unique bool) *CertificateQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CertificateQuery) Order(o ...certificate.OrderOption) *CertificateQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryAcmeAccount chains the current query on the "acme_account" edge.
func (cq *CertificateQuery) QueryAcmeAccount() *AcmeAccountQuery {
	query := (&AcmeAccountClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(certificate.Table, certificate.FieldID, selector),
			sqlgraph.To(acmeaccount.Table, acmeaccount.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, certificate.AcmeAccountTable, certificate.AcmeAccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDNSProvider chains the current query on the "dns_provider" edge.
func (cq *CertificateQuery) QueryDNSProvider() *DnsProviderQuery {
	query := (&DnsProviderClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(certificate.Table, certificate.FieldID, selector),
			sqlgraph.To(dnsprovider.Table, dnsprovider.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, certificate.DNSProviderTable, certificate.DNSProviderColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Certificate entity from the query.
// Returns a *NotFoundError when no Certificate was found.
func (cq *CertificateQuery) First(ctx context.Context) (*Certificate, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{certificate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CertificateQuery) FirstX(ctx context.Context) *Certificate {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Certificate ID from the query.
// Returns a *NotFoundError when no Certificate ID was found.
func (cq *CertificateQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{certificate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CertificateQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Certificate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Certificate entity is found.
// Returns a *NotFoundError when no Certificate entities are found.
func (cq *CertificateQuery) Only(ctx context.Context) (*Certificate, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{certificate.Label}
	default:
		return nil, &NotSingularError{certificate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CertificateQuery) OnlyX(ctx context.Context) *Certificate {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Certificate ID in the query.
// Returns a *NotSingularError when more than one Certificate ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CertificateQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{certificate.Label}
	default:
		err = &NotSingularError{certificate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CertificateQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Certificates.
func (cq *CertificateQuery) All(ctx context.Context) ([]*Certificate, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Certificate, *CertificateQuery]()
	return withInterceptors[[]*Certificate](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CertificateQuery) AllX(ctx context.Context) []*Certificate {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Certificate IDs.
func (cq *CertificateQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryIDs)
	if err = cq.Select(certificate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CertificateQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CertificateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CertificateQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CertificateQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CertificateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CertificateQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CertificateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CertificateQuery) Clone() *CertificateQuery {
	if cq == nil {
		return nil
	}
	return &CertificateQuery{
		config:          cq.config,
		ctx:             cq.ctx.Clone(),
		order:           append([]certificate.OrderOption{}, cq.order...),
		inters:          append([]Interceptor{}, cq.inters...),
		predicates:      append([]predicate.Certificate{}, cq.predicates...),
		withAcmeAccount: cq.withAcmeAccount.Clone(),
		withDNSProvider: cq.withDNSProvider.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithAcmeAccount tells the query-builder to eager-load the nodes that are connected to
// the "acme_account" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CertificateQuery) WithAcmeAccount(opts ...func(*AcmeAccountQuery)) *CertificateQuery {
	query := (&AcmeAccountClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withAcmeAccount = query
	return cq
}

// WithDNSProvider tells the query-builder to eager-load the nodes that are connected to
// the "dns_provider" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CertificateQuery) WithDNSProvider(opts ...func(*DnsProviderQuery)) *CertificateQuery {
	query := (&DnsProviderClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withDNSProvider = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AcmeAccountID uuid.UUID `json:"acme_account_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Certificate.Query().
//		GroupBy(certificate.FieldAcmeAccountID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CertificateQuery) GroupBy(field string, fields ...string) *CertificateGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CertificateGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = certificate.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AcmeAccountID uuid.UUID `json:"acme_account_id,omitempty"`
//	}
//
//	client.Certificate.Query().
//		Select(certificate.FieldAcmeAccountID).
//		Scan(ctx, &v)
func (cq *CertificateQuery) Select(fields ...string) *CertificateSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CertificateSelect{CertificateQuery: cq}
	sbuild.label = certificate.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CertificateSelect configured with the given aggregations.
func (cq *CertificateQuery) Aggregate(fns ...AggregateFunc) *CertificateSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CertificateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !certificate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CertificateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Certificate, error) {
	var (
		nodes       = []*Certificate{}
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withAcmeAccount != nil,
			cq.withDNSProvider != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Certificate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Certificate{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withAcmeAccount; query != nil {
		if err := cq.loadAcmeAccount(ctx, query, nodes, nil,
			func(n *Certificate, e *AcmeAccount) { n.Edges.AcmeAccount = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withDNSProvider; query != nil {
		if err := cq.loadDNSProvider(ctx, query, nodes, nil,
			func(n *Certificate, e *DnsProvider) { n.Edges.DNSProvider = e }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CertificateQuery) loadAcmeAccount(ctx context.Context, query *AcmeAccountQuery, nodes []*Certificate, init func(*Certificate), assign func(*Certificate, *AcmeAccount)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Certificate)
	for i := range nodes {
		fk := nodes[i].AcmeAccountID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(acmeaccount.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "acme_account_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CertificateQuery) loadDNSProvider(ctx context.Context, query *DnsProviderQuery, nodes []*Certificate, init func(*Certificate), assign func(*Certificate, *DnsProvider)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Certificate)
	for i := range nodes {
		fk := nodes[i].DNSProviderID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(dnsprovider.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "dns_provider_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cq *CertificateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CertificateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(certificate.Table, certificate.Columns, sqlgraph.NewFieldSpec(certificate.FieldID, field.TypeUUID))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, certificate.FieldID)
		for i := range fields {
			if fields[i] != certificate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if cq.withAcmeAccount != nil {
			_spec.Node.AddColumnOnce(certificate.FieldAcmeAccountID)
		}
		if cq.withDNSProvider != nil {
			_spec.Node.AddColumnOnce(certificate.FieldDNSProviderID)
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CertificateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(certificate.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = certificate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CertificateGroupBy is the group-by builder for Certificate entities.
type CertificateGroupBy struct {
	selector
	build *CertificateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CertificateGroupBy) Aggregate(fns ...AggregateFunc) *CertificateGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CertificateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CertificateQuery, *CertificateGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CertificateGroupBy) sqlScan(ctx context.Context, root *CertificateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CertificateSelect is the builder for selecting fields of Certificate entities.
type CertificateSelect struct {
	*CertificateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CertificateSelect) Aggregate(fns ...AggregateFunc) *CertificateSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CertificateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CertificateQuery, *CertificateSelect](ctx, cs.CertificateQuery, cs, cs.inters, v)
}

func (cs *CertificateSelect) sqlScan(ctx context.Context, root *CertificateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}