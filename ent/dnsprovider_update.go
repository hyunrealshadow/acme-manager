// Code generated by ent, DO NOT EDIT.

package ent

import (
	"acme-manager/acme/lego"
	"acme-manager/ent/certificate"
	"acme-manager/ent/dnsprovider"
	"acme-manager/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DnsProviderUpdate is the builder for updating DnsProvider entities.
type DnsProviderUpdate struct {
	config
	hooks    []Hook
	mutation *DnsProviderMutation
}

// Where appends a list predicates to the DnsProviderUpdate builder.
func (dpu *DnsProviderUpdate) Where(ps ...predicate.DnsProvider) *DnsProviderUpdate {
	dpu.mutation.Where(ps...)
	return dpu
}

// SetName sets the "name" field.
func (dpu *DnsProviderUpdate) SetName(s string) *DnsProviderUpdate {
	dpu.mutation.SetName(s)
	return dpu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dpu *DnsProviderUpdate) SetNillableName(s *string) *DnsProviderUpdate {
	if s != nil {
		dpu.SetName(*s)
	}
	return dpu
}

// SetDescription sets the "description" field.
func (dpu *DnsProviderUpdate) SetDescription(s string) *DnsProviderUpdate {
	dpu.mutation.SetDescription(s)
	return dpu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dpu *DnsProviderUpdate) SetNillableDescription(s *string) *DnsProviderUpdate {
	if s != nil {
		dpu.SetDescription(*s)
	}
	return dpu
}

// ClearDescription clears the value of the "description" field.
func (dpu *DnsProviderUpdate) ClearDescription() *DnsProviderUpdate {
	dpu.mutation.ClearDescription()
	return dpu
}

// SetType sets the "type" field.
func (dpu *DnsProviderUpdate) SetType(s string) *DnsProviderUpdate {
	dpu.mutation.SetType(s)
	return dpu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (dpu *DnsProviderUpdate) SetNillableType(s *string) *DnsProviderUpdate {
	if s != nil {
		dpu.SetType(*s)
	}
	return dpu
}

// SetConfig sets the "config" field.
func (dpu *DnsProviderUpdate) SetConfig(lpc *lego.DnsProviderConfig) *DnsProviderUpdate {
	dpu.mutation.SetConfig(lpc)
	return dpu
}

// SetCreatedBy sets the "created_by" field.
func (dpu *DnsProviderUpdate) SetCreatedBy(u uuid.UUID) *DnsProviderUpdate {
	dpu.mutation.SetCreatedBy(u)
	return dpu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (dpu *DnsProviderUpdate) SetNillableCreatedBy(u *uuid.UUID) *DnsProviderUpdate {
	if u != nil {
		dpu.SetCreatedBy(*u)
	}
	return dpu
}

// SetUpdatedAt sets the "updated_at" field.
func (dpu *DnsProviderUpdate) SetUpdatedAt(t time.Time) *DnsProviderUpdate {
	dpu.mutation.SetUpdatedAt(t)
	return dpu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dpu *DnsProviderUpdate) SetNillableUpdatedAt(t *time.Time) *DnsProviderUpdate {
	if t != nil {
		dpu.SetUpdatedAt(*t)
	}
	return dpu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (dpu *DnsProviderUpdate) ClearUpdatedAt() *DnsProviderUpdate {
	dpu.mutation.ClearUpdatedAt()
	return dpu
}

// SetUpdatedBy sets the "updated_by" field.
func (dpu *DnsProviderUpdate) SetUpdatedBy(u uuid.UUID) *DnsProviderUpdate {
	dpu.mutation.SetUpdatedBy(u)
	return dpu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dpu *DnsProviderUpdate) SetNillableUpdatedBy(u *uuid.UUID) *DnsProviderUpdate {
	if u != nil {
		dpu.SetUpdatedBy(*u)
	}
	return dpu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (dpu *DnsProviderUpdate) ClearUpdatedBy() *DnsProviderUpdate {
	dpu.mutation.ClearUpdatedBy()
	return dpu
}

// AddCertificateIDs adds the "certificates" edge to the Certificate entity by IDs.
func (dpu *DnsProviderUpdate) AddCertificateIDs(ids ...uuid.UUID) *DnsProviderUpdate {
	dpu.mutation.AddCertificateIDs(ids...)
	return dpu
}

// AddCertificates adds the "certificates" edges to the Certificate entity.
func (dpu *DnsProviderUpdate) AddCertificates(c ...*Certificate) *DnsProviderUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dpu.AddCertificateIDs(ids...)
}

// Mutation returns the DnsProviderMutation object of the builder.
func (dpu *DnsProviderUpdate) Mutation() *DnsProviderMutation {
	return dpu.mutation
}

// ClearCertificates clears all "certificates" edges to the Certificate entity.
func (dpu *DnsProviderUpdate) ClearCertificates() *DnsProviderUpdate {
	dpu.mutation.ClearCertificates()
	return dpu
}

// RemoveCertificateIDs removes the "certificates" edge to Certificate entities by IDs.
func (dpu *DnsProviderUpdate) RemoveCertificateIDs(ids ...uuid.UUID) *DnsProviderUpdate {
	dpu.mutation.RemoveCertificateIDs(ids...)
	return dpu
}

// RemoveCertificates removes "certificates" edges to Certificate entities.
func (dpu *DnsProviderUpdate) RemoveCertificates(c ...*Certificate) *DnsProviderUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dpu.RemoveCertificateIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dpu *DnsProviderUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, dpu.sqlSave, dpu.mutation, dpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dpu *DnsProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := dpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dpu *DnsProviderUpdate) Exec(ctx context.Context) error {
	_, err := dpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpu *DnsProviderUpdate) ExecX(ctx context.Context) {
	if err := dpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dpu *DnsProviderUpdate) check() error {
	if v, ok := dpu.mutation.Name(); ok {
		if err := dnsprovider.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "DnsProvider.name": %w`, err)}
		}
	}
	if v, ok := dpu.mutation.Description(); ok {
		if err := dnsprovider.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "DnsProvider.description": %w`, err)}
		}
	}
	if v, ok := dpu.mutation.GetType(); ok {
		if err := dnsprovider.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "DnsProvider.type": %w`, err)}
		}
	}
	return nil
}

func (dpu *DnsProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := dpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(dnsprovider.Table, dnsprovider.Columns, sqlgraph.NewFieldSpec(dnsprovider.FieldID, field.TypeUUID))
	if ps := dpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dpu.mutation.Name(); ok {
		_spec.SetField(dnsprovider.FieldName, field.TypeString, value)
	}
	if value, ok := dpu.mutation.Description(); ok {
		_spec.SetField(dnsprovider.FieldDescription, field.TypeString, value)
	}
	if dpu.mutation.DescriptionCleared() {
		_spec.ClearField(dnsprovider.FieldDescription, field.TypeString)
	}
	if value, ok := dpu.mutation.GetType(); ok {
		_spec.SetField(dnsprovider.FieldType, field.TypeString, value)
	}
	if value, ok := dpu.mutation.Config(); ok {
		_spec.SetField(dnsprovider.FieldConfig, field.TypeJSON, value)
	}
	if value, ok := dpu.mutation.CreatedBy(); ok {
		_spec.SetField(dnsprovider.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := dpu.mutation.UpdatedAt(); ok {
		_spec.SetField(dnsprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if dpu.mutation.UpdatedAtCleared() {
		_spec.ClearField(dnsprovider.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := dpu.mutation.UpdatedBy(); ok {
		_spec.SetField(dnsprovider.FieldUpdatedBy, field.TypeUUID, value)
	}
	if dpu.mutation.UpdatedByCleared() {
		_spec.ClearField(dnsprovider.FieldUpdatedBy, field.TypeUUID)
	}
	if dpu.mutation.CertificatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dnsprovider.CertificatesTable,
			Columns: []string{dnsprovider.CertificatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certificate.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dpu.mutation.RemovedCertificatesIDs(); len(nodes) > 0 && !dpu.mutation.CertificatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dnsprovider.CertificatesTable,
			Columns: []string{dnsprovider.CertificatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certificate.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dpu.mutation.CertificatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dnsprovider.CertificatesTable,
			Columns: []string{dnsprovider.CertificatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certificate.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dnsprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dpu.mutation.done = true
	return n, nil
}

// DnsProviderUpdateOne is the builder for updating a single DnsProvider entity.
type DnsProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DnsProviderMutation
}

// SetName sets the "name" field.
func (dpuo *DnsProviderUpdateOne) SetName(s string) *DnsProviderUpdateOne {
	dpuo.mutation.SetName(s)
	return dpuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dpuo *DnsProviderUpdateOne) SetNillableName(s *string) *DnsProviderUpdateOne {
	if s != nil {
		dpuo.SetName(*s)
	}
	return dpuo
}

// SetDescription sets the "description" field.
func (dpuo *DnsProviderUpdateOne) SetDescription(s string) *DnsProviderUpdateOne {
	dpuo.mutation.SetDescription(s)
	return dpuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dpuo *DnsProviderUpdateOne) SetNillableDescription(s *string) *DnsProviderUpdateOne {
	if s != nil {
		dpuo.SetDescription(*s)
	}
	return dpuo
}

// ClearDescription clears the value of the "description" field.
func (dpuo *DnsProviderUpdateOne) ClearDescription() *DnsProviderUpdateOne {
	dpuo.mutation.ClearDescription()
	return dpuo
}

// SetType sets the "type" field.
func (dpuo *DnsProviderUpdateOne) SetType(s string) *DnsProviderUpdateOne {
	dpuo.mutation.SetType(s)
	return dpuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (dpuo *DnsProviderUpdateOne) SetNillableType(s *string) *DnsProviderUpdateOne {
	if s != nil {
		dpuo.SetType(*s)
	}
	return dpuo
}

// SetConfig sets the "config" field.
func (dpuo *DnsProviderUpdateOne) SetConfig(lpc *lego.DnsProviderConfig) *DnsProviderUpdateOne {
	dpuo.mutation.SetConfig(lpc)
	return dpuo
}

// SetCreatedBy sets the "created_by" field.
func (dpuo *DnsProviderUpdateOne) SetCreatedBy(u uuid.UUID) *DnsProviderUpdateOne {
	dpuo.mutation.SetCreatedBy(u)
	return dpuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (dpuo *DnsProviderUpdateOne) SetNillableCreatedBy(u *uuid.UUID) *DnsProviderUpdateOne {
	if u != nil {
		dpuo.SetCreatedBy(*u)
	}
	return dpuo
}

// SetUpdatedAt sets the "updated_at" field.
func (dpuo *DnsProviderUpdateOne) SetUpdatedAt(t time.Time) *DnsProviderUpdateOne {
	dpuo.mutation.SetUpdatedAt(t)
	return dpuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dpuo *DnsProviderUpdateOne) SetNillableUpdatedAt(t *time.Time) *DnsProviderUpdateOne {
	if t != nil {
		dpuo.SetUpdatedAt(*t)
	}
	return dpuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (dpuo *DnsProviderUpdateOne) ClearUpdatedAt() *DnsProviderUpdateOne {
	dpuo.mutation.ClearUpdatedAt()
	return dpuo
}

// SetUpdatedBy sets the "updated_by" field.
func (dpuo *DnsProviderUpdateOne) SetUpdatedBy(u uuid.UUID) *DnsProviderUpdateOne {
	dpuo.mutation.SetUpdatedBy(u)
	return dpuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dpuo *DnsProviderUpdateOne) SetNillableUpdatedBy(u *uuid.UUID) *DnsProviderUpdateOne {
	if u != nil {
		dpuo.SetUpdatedBy(*u)
	}
	return dpuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (dpuo *DnsProviderUpdateOne) ClearUpdatedBy() *DnsProviderUpdateOne {
	dpuo.mutation.ClearUpdatedBy()
	return dpuo
}

// AddCertificateIDs adds the "certificates" edge to the Certificate entity by IDs.
func (dpuo *DnsProviderUpdateOne) AddCertificateIDs(ids ...uuid.UUID) *DnsProviderUpdateOne {
	dpuo.mutation.AddCertificateIDs(ids...)
	return dpuo
}

// AddCertificates adds the "certificates" edges to the Certificate entity.
func (dpuo *DnsProviderUpdateOne) AddCertificates(c ...*Certificate) *DnsProviderUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dpuo.AddCertificateIDs(ids...)
}

// Mutation returns the DnsProviderMutation object of the builder.
func (dpuo *DnsProviderUpdateOne) Mutation() *DnsProviderMutation {
	return dpuo.mutation
}

// ClearCertificates clears all "certificates" edges to the Certificate entity.
func (dpuo *DnsProviderUpdateOne) ClearCertificates() *DnsProviderUpdateOne {
	dpuo.mutation.ClearCertificates()
	return dpuo
}

// RemoveCertificateIDs removes the "certificates" edge to Certificate entities by IDs.
func (dpuo *DnsProviderUpdateOne) RemoveCertificateIDs(ids ...uuid.UUID) *DnsProviderUpdateOne {
	dpuo.mutation.RemoveCertificateIDs(ids...)
	return dpuo
}

// RemoveCertificates removes "certificates" edges to Certificate entities.
func (dpuo *DnsProviderUpdateOne) RemoveCertificates(c ...*Certificate) *DnsProviderUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dpuo.RemoveCertificateIDs(ids...)
}

// Where appends a list predicates to the DnsProviderUpdate builder.
func (dpuo *DnsProviderUpdateOne) Where(ps ...predicate.DnsProvider) *DnsProviderUpdateOne {
	dpuo.mutation.Where(ps...)
	return dpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dpuo *DnsProviderUpdateOne) Select(field string, fields ...string) *DnsProviderUpdateOne {
	dpuo.fields = append([]string{field}, fields...)
	return dpuo
}

// Save executes the query and returns the updated DnsProvider entity.
func (dpuo *DnsProviderUpdateOne) Save(ctx context.Context) (*DnsProvider, error) {
	return withHooks(ctx, dpuo.sqlSave, dpuo.mutation, dpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dpuo *DnsProviderUpdateOne) SaveX(ctx context.Context) *DnsProvider {
	node, err := dpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dpuo *DnsProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := dpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpuo *DnsProviderUpdateOne) ExecX(ctx context.Context) {
	if err := dpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dpuo *DnsProviderUpdateOne) check() error {
	if v, ok := dpuo.mutation.Name(); ok {
		if err := dnsprovider.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "DnsProvider.name": %w`, err)}
		}
	}
	if v, ok := dpuo.mutation.Description(); ok {
		if err := dnsprovider.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "DnsProvider.description": %w`, err)}
		}
	}
	if v, ok := dpuo.mutation.GetType(); ok {
		if err := dnsprovider.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "DnsProvider.type": %w`, err)}
		}
	}
	return nil
}

func (dpuo *DnsProviderUpdateOne) sqlSave(ctx context.Context) (_node *DnsProvider, err error) {
	if err := dpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(dnsprovider.Table, dnsprovider.Columns, sqlgraph.NewFieldSpec(dnsprovider.FieldID, field.TypeUUID))
	id, ok := dpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DnsProvider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dnsprovider.FieldID)
		for _, f := range fields {
			if !dnsprovider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dnsprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dpuo.mutation.Name(); ok {
		_spec.SetField(dnsprovider.FieldName, field.TypeString, value)
	}
	if value, ok := dpuo.mutation.Description(); ok {
		_spec.SetField(dnsprovider.FieldDescription, field.TypeString, value)
	}
	if dpuo.mutation.DescriptionCleared() {
		_spec.ClearField(dnsprovider.FieldDescription, field.TypeString)
	}
	if value, ok := dpuo.mutation.GetType(); ok {
		_spec.SetField(dnsprovider.FieldType, field.TypeString, value)
	}
	if value, ok := dpuo.mutation.Config(); ok {
		_spec.SetField(dnsprovider.FieldConfig, field.TypeJSON, value)
	}
	if value, ok := dpuo.mutation.CreatedBy(); ok {
		_spec.SetField(dnsprovider.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := dpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(dnsprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if dpuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(dnsprovider.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := dpuo.mutation.UpdatedBy(); ok {
		_spec.SetField(dnsprovider.FieldUpdatedBy, field.TypeUUID, value)
	}
	if dpuo.mutation.UpdatedByCleared() {
		_spec.ClearField(dnsprovider.FieldUpdatedBy, field.TypeUUID)
	}
	if dpuo.mutation.CertificatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dnsprovider.CertificatesTable,
			Columns: []string{dnsprovider.CertificatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certificate.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dpuo.mutation.RemovedCertificatesIDs(); len(nodes) > 0 && !dpuo.mutation.CertificatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dnsprovider.CertificatesTable,
			Columns: []string{dnsprovider.CertificatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certificate.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dpuo.mutation.CertificatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dnsprovider.CertificatesTable,
			Columns: []string{dnsprovider.CertificatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certificate.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DnsProvider{config: dpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dnsprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dpuo.mutation.done = true
	return _node, nil
}