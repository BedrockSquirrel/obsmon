// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/bedrocksquirrel/obsmon/ent/environment"
	"github.com/bedrocksquirrel/obsmon/ent/network"
	"github.com/bedrocksquirrel/obsmon/ent/node"
	"github.com/bedrocksquirrel/obsmon/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeEnvironment = "Environment"
	TypeNetwork     = "Network"
	TypeNode        = "Node"
)

// EnvironmentMutation represents an operation that mutates the Environment nodes in the graph.
type EnvironmentMutation struct {
	config
	op              Op
	typ             string
	id              *int
	name            *string
	clearedFields   map[string]struct{}
	networks        map[int]struct{}
	removednetworks map[int]struct{}
	clearednetworks bool
	done            bool
	oldValue        func(context.Context) (*Environment, error)
	predicates      []predicate.Environment
}

var _ ent.Mutation = (*EnvironmentMutation)(nil)

// environmentOption allows management of the mutation configuration using functional options.
type environmentOption func(*EnvironmentMutation)

// newEnvironmentMutation creates new mutation for the Environment entity.
func newEnvironmentMutation(c config, op Op, opts ...environmentOption) *EnvironmentMutation {
	m := &EnvironmentMutation{
		config:        c,
		op:            op,
		typ:           TypeEnvironment,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEnvironmentID sets the ID field of the mutation.
func withEnvironmentID(id int) environmentOption {
	return func(m *EnvironmentMutation) {
		var (
			err   error
			once  sync.Once
			value *Environment
		)
		m.oldValue = func(ctx context.Context) (*Environment, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Environment.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEnvironment sets the old Environment of the mutation.
func withEnvironment(node *Environment) environmentOption {
	return func(m *EnvironmentMutation) {
		m.oldValue = func(context.Context) (*Environment, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EnvironmentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EnvironmentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EnvironmentMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *EnvironmentMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Environment.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *EnvironmentMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *EnvironmentMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Environment entity.
// If the Environment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EnvironmentMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *EnvironmentMutation) ResetName() {
	m.name = nil
}

// AddNetworkIDs adds the "networks" edge to the Network entity by ids.
func (m *EnvironmentMutation) AddNetworkIDs(ids ...int) {
	if m.networks == nil {
		m.networks = make(map[int]struct{})
	}
	for i := range ids {
		m.networks[ids[i]] = struct{}{}
	}
}

// ClearNetworks clears the "networks" edge to the Network entity.
func (m *EnvironmentMutation) ClearNetworks() {
	m.clearednetworks = true
}

// NetworksCleared reports if the "networks" edge to the Network entity was cleared.
func (m *EnvironmentMutation) NetworksCleared() bool {
	return m.clearednetworks
}

// RemoveNetworkIDs removes the "networks" edge to the Network entity by IDs.
func (m *EnvironmentMutation) RemoveNetworkIDs(ids ...int) {
	if m.removednetworks == nil {
		m.removednetworks = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.networks, ids[i])
		m.removednetworks[ids[i]] = struct{}{}
	}
}

// RemovedNetworks returns the removed IDs of the "networks" edge to the Network entity.
func (m *EnvironmentMutation) RemovedNetworksIDs() (ids []int) {
	for id := range m.removednetworks {
		ids = append(ids, id)
	}
	return
}

// NetworksIDs returns the "networks" edge IDs in the mutation.
func (m *EnvironmentMutation) NetworksIDs() (ids []int) {
	for id := range m.networks {
		ids = append(ids, id)
	}
	return
}

// ResetNetworks resets all changes to the "networks" edge.
func (m *EnvironmentMutation) ResetNetworks() {
	m.networks = nil
	m.clearednetworks = false
	m.removednetworks = nil
}

// Where appends a list predicates to the EnvironmentMutation builder.
func (m *EnvironmentMutation) Where(ps ...predicate.Environment) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the EnvironmentMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *EnvironmentMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Environment, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *EnvironmentMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *EnvironmentMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Environment).
func (m *EnvironmentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EnvironmentMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, environment.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EnvironmentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case environment.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EnvironmentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case environment.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Environment field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EnvironmentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case environment.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Environment field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EnvironmentMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EnvironmentMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EnvironmentMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Environment numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EnvironmentMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EnvironmentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EnvironmentMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Environment nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EnvironmentMutation) ResetField(name string) error {
	switch name {
	case environment.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Environment field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EnvironmentMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.networks != nil {
		edges = append(edges, environment.EdgeNetworks)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EnvironmentMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case environment.EdgeNetworks:
		ids := make([]ent.Value, 0, len(m.networks))
		for id := range m.networks {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EnvironmentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removednetworks != nil {
		edges = append(edges, environment.EdgeNetworks)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EnvironmentMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case environment.EdgeNetworks:
		ids := make([]ent.Value, 0, len(m.removednetworks))
		for id := range m.removednetworks {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EnvironmentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearednetworks {
		edges = append(edges, environment.EdgeNetworks)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EnvironmentMutation) EdgeCleared(name string) bool {
	switch name {
	case environment.EdgeNetworks:
		return m.clearednetworks
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EnvironmentMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Environment unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EnvironmentMutation) ResetEdge(name string) error {
	switch name {
	case environment.EdgeNetworks:
		m.ResetNetworks()
		return nil
	}
	return fmt.Errorf("unknown Environment edge %s", name)
}

// NetworkMutation represents an operation that mutates the Network nodes in the graph.
type NetworkMutation struct {
	config
	op                 Op
	typ                string
	id                 *int
	deployTime         *time.Time
	githubBuildNum     *int
	addgithubBuildNum  *int
	clearedFields      map[string]struct{}
	environment        *int
	clearedenvironment bool
	nodes              map[int]struct{}
	removednodes       map[int]struct{}
	clearednodes       bool
	done               bool
	oldValue           func(context.Context) (*Network, error)
	predicates         []predicate.Network
}

var _ ent.Mutation = (*NetworkMutation)(nil)

// networkOption allows management of the mutation configuration using functional options.
type networkOption func(*NetworkMutation)

// newNetworkMutation creates new mutation for the Network entity.
func newNetworkMutation(c config, op Op, opts ...networkOption) *NetworkMutation {
	m := &NetworkMutation{
		config:        c,
		op:            op,
		typ:           TypeNetwork,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withNetworkID sets the ID field of the mutation.
func withNetworkID(id int) networkOption {
	return func(m *NetworkMutation) {
		var (
			err   error
			once  sync.Once
			value *Network
		)
		m.oldValue = func(ctx context.Context) (*Network, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Network.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withNetwork sets the old Network of the mutation.
func withNetwork(node *Network) networkOption {
	return func(m *NetworkMutation) {
		m.oldValue = func(context.Context) (*Network, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m NetworkMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m NetworkMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *NetworkMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *NetworkMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Network.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetDeployTime sets the "deployTime" field.
func (m *NetworkMutation) SetDeployTime(t time.Time) {
	m.deployTime = &t
}

// DeployTime returns the value of the "deployTime" field in the mutation.
func (m *NetworkMutation) DeployTime() (r time.Time, exists bool) {
	v := m.deployTime
	if v == nil {
		return
	}
	return *v, true
}

// OldDeployTime returns the old "deployTime" field's value of the Network entity.
// If the Network object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *NetworkMutation) OldDeployTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeployTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeployTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeployTime: %w", err)
	}
	return oldValue.DeployTime, nil
}

// ResetDeployTime resets all changes to the "deployTime" field.
func (m *NetworkMutation) ResetDeployTime() {
	m.deployTime = nil
}

// SetGithubBuildNum sets the "githubBuildNum" field.
func (m *NetworkMutation) SetGithubBuildNum(i int) {
	m.githubBuildNum = &i
	m.addgithubBuildNum = nil
}

// GithubBuildNum returns the value of the "githubBuildNum" field in the mutation.
func (m *NetworkMutation) GithubBuildNum() (r int, exists bool) {
	v := m.githubBuildNum
	if v == nil {
		return
	}
	return *v, true
}

// OldGithubBuildNum returns the old "githubBuildNum" field's value of the Network entity.
// If the Network object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *NetworkMutation) OldGithubBuildNum(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGithubBuildNum is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGithubBuildNum requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGithubBuildNum: %w", err)
	}
	return oldValue.GithubBuildNum, nil
}

// AddGithubBuildNum adds i to the "githubBuildNum" field.
func (m *NetworkMutation) AddGithubBuildNum(i int) {
	if m.addgithubBuildNum != nil {
		*m.addgithubBuildNum += i
	} else {
		m.addgithubBuildNum = &i
	}
}

// AddedGithubBuildNum returns the value that was added to the "githubBuildNum" field in this mutation.
func (m *NetworkMutation) AddedGithubBuildNum() (r int, exists bool) {
	v := m.addgithubBuildNum
	if v == nil {
		return
	}
	return *v, true
}

// ResetGithubBuildNum resets all changes to the "githubBuildNum" field.
func (m *NetworkMutation) ResetGithubBuildNum() {
	m.githubBuildNum = nil
	m.addgithubBuildNum = nil
}

// SetEnvironmentID sets the "environment" edge to the Environment entity by id.
func (m *NetworkMutation) SetEnvironmentID(id int) {
	m.environment = &id
}

// ClearEnvironment clears the "environment" edge to the Environment entity.
func (m *NetworkMutation) ClearEnvironment() {
	m.clearedenvironment = true
}

// EnvironmentCleared reports if the "environment" edge to the Environment entity was cleared.
func (m *NetworkMutation) EnvironmentCleared() bool {
	return m.clearedenvironment
}

// EnvironmentID returns the "environment" edge ID in the mutation.
func (m *NetworkMutation) EnvironmentID() (id int, exists bool) {
	if m.environment != nil {
		return *m.environment, true
	}
	return
}

// EnvironmentIDs returns the "environment" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// EnvironmentID instead. It exists only for internal usage by the builders.
func (m *NetworkMutation) EnvironmentIDs() (ids []int) {
	if id := m.environment; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetEnvironment resets all changes to the "environment" edge.
func (m *NetworkMutation) ResetEnvironment() {
	m.environment = nil
	m.clearedenvironment = false
}

// AddNodeIDs adds the "nodes" edge to the Node entity by ids.
func (m *NetworkMutation) AddNodeIDs(ids ...int) {
	if m.nodes == nil {
		m.nodes = make(map[int]struct{})
	}
	for i := range ids {
		m.nodes[ids[i]] = struct{}{}
	}
}

// ClearNodes clears the "nodes" edge to the Node entity.
func (m *NetworkMutation) ClearNodes() {
	m.clearednodes = true
}

// NodesCleared reports if the "nodes" edge to the Node entity was cleared.
func (m *NetworkMutation) NodesCleared() bool {
	return m.clearednodes
}

// RemoveNodeIDs removes the "nodes" edge to the Node entity by IDs.
func (m *NetworkMutation) RemoveNodeIDs(ids ...int) {
	if m.removednodes == nil {
		m.removednodes = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.nodes, ids[i])
		m.removednodes[ids[i]] = struct{}{}
	}
}

// RemovedNodes returns the removed IDs of the "nodes" edge to the Node entity.
func (m *NetworkMutation) RemovedNodesIDs() (ids []int) {
	for id := range m.removednodes {
		ids = append(ids, id)
	}
	return
}

// NodesIDs returns the "nodes" edge IDs in the mutation.
func (m *NetworkMutation) NodesIDs() (ids []int) {
	for id := range m.nodes {
		ids = append(ids, id)
	}
	return
}

// ResetNodes resets all changes to the "nodes" edge.
func (m *NetworkMutation) ResetNodes() {
	m.nodes = nil
	m.clearednodes = false
	m.removednodes = nil
}

// Where appends a list predicates to the NetworkMutation builder.
func (m *NetworkMutation) Where(ps ...predicate.Network) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the NetworkMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *NetworkMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Network, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *NetworkMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *NetworkMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Network).
func (m *NetworkMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *NetworkMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.deployTime != nil {
		fields = append(fields, network.FieldDeployTime)
	}
	if m.githubBuildNum != nil {
		fields = append(fields, network.FieldGithubBuildNum)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *NetworkMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case network.FieldDeployTime:
		return m.DeployTime()
	case network.FieldGithubBuildNum:
		return m.GithubBuildNum()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *NetworkMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case network.FieldDeployTime:
		return m.OldDeployTime(ctx)
	case network.FieldGithubBuildNum:
		return m.OldGithubBuildNum(ctx)
	}
	return nil, fmt.Errorf("unknown Network field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *NetworkMutation) SetField(name string, value ent.Value) error {
	switch name {
	case network.FieldDeployTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeployTime(v)
		return nil
	case network.FieldGithubBuildNum:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGithubBuildNum(v)
		return nil
	}
	return fmt.Errorf("unknown Network field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *NetworkMutation) AddedFields() []string {
	var fields []string
	if m.addgithubBuildNum != nil {
		fields = append(fields, network.FieldGithubBuildNum)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *NetworkMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case network.FieldGithubBuildNum:
		return m.AddedGithubBuildNum()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *NetworkMutation) AddField(name string, value ent.Value) error {
	switch name {
	case network.FieldGithubBuildNum:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddGithubBuildNum(v)
		return nil
	}
	return fmt.Errorf("unknown Network numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *NetworkMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *NetworkMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *NetworkMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Network nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *NetworkMutation) ResetField(name string) error {
	switch name {
	case network.FieldDeployTime:
		m.ResetDeployTime()
		return nil
	case network.FieldGithubBuildNum:
		m.ResetGithubBuildNum()
		return nil
	}
	return fmt.Errorf("unknown Network field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *NetworkMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.environment != nil {
		edges = append(edges, network.EdgeEnvironment)
	}
	if m.nodes != nil {
		edges = append(edges, network.EdgeNodes)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *NetworkMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case network.EdgeEnvironment:
		if id := m.environment; id != nil {
			return []ent.Value{*id}
		}
	case network.EdgeNodes:
		ids := make([]ent.Value, 0, len(m.nodes))
		for id := range m.nodes {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *NetworkMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removednodes != nil {
		edges = append(edges, network.EdgeNodes)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *NetworkMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case network.EdgeNodes:
		ids := make([]ent.Value, 0, len(m.removednodes))
		for id := range m.removednodes {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *NetworkMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedenvironment {
		edges = append(edges, network.EdgeEnvironment)
	}
	if m.clearednodes {
		edges = append(edges, network.EdgeNodes)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *NetworkMutation) EdgeCleared(name string) bool {
	switch name {
	case network.EdgeEnvironment:
		return m.clearedenvironment
	case network.EdgeNodes:
		return m.clearednodes
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *NetworkMutation) ClearEdge(name string) error {
	switch name {
	case network.EdgeEnvironment:
		m.ClearEnvironment()
		return nil
	}
	return fmt.Errorf("unknown Network unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *NetworkMutation) ResetEdge(name string) error {
	switch name {
	case network.EdgeEnvironment:
		m.ResetEnvironment()
		return nil
	case network.EdgeNodes:
		m.ResetNodes()
		return nil
	}
	return fmt.Errorf("unknown Network edge %s", name)
}

// NodeMutation represents an operation that mutates the Node nodes in the graph.
type NodeMutation struct {
	config
	op             Op
	typ            string
	id             *int
	clearedFields  map[string]struct{}
	network        *int
	clearednetwork bool
	done           bool
	oldValue       func(context.Context) (*Node, error)
	predicates     []predicate.Node
}

var _ ent.Mutation = (*NodeMutation)(nil)

// nodeOption allows management of the mutation configuration using functional options.
type nodeOption func(*NodeMutation)

// newNodeMutation creates new mutation for the Node entity.
func newNodeMutation(c config, op Op, opts ...nodeOption) *NodeMutation {
	m := &NodeMutation{
		config:        c,
		op:            op,
		typ:           TypeNode,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withNodeID sets the ID field of the mutation.
func withNodeID(id int) nodeOption {
	return func(m *NodeMutation) {
		var (
			err   error
			once  sync.Once
			value *Node
		)
		m.oldValue = func(ctx context.Context) (*Node, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Node.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withNode sets the old Node of the mutation.
func withNode(node *Node) nodeOption {
	return func(m *NodeMutation) {
		m.oldValue = func(context.Context) (*Node, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m NodeMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m NodeMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *NodeMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *NodeMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Node.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetNetworkID sets the "network" edge to the Network entity by id.
func (m *NodeMutation) SetNetworkID(id int) {
	m.network = &id
}

// ClearNetwork clears the "network" edge to the Network entity.
func (m *NodeMutation) ClearNetwork() {
	m.clearednetwork = true
}

// NetworkCleared reports if the "network" edge to the Network entity was cleared.
func (m *NodeMutation) NetworkCleared() bool {
	return m.clearednetwork
}

// NetworkID returns the "network" edge ID in the mutation.
func (m *NodeMutation) NetworkID() (id int, exists bool) {
	if m.network != nil {
		return *m.network, true
	}
	return
}

// NetworkIDs returns the "network" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// NetworkID instead. It exists only for internal usage by the builders.
func (m *NodeMutation) NetworkIDs() (ids []int) {
	if id := m.network; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetNetwork resets all changes to the "network" edge.
func (m *NodeMutation) ResetNetwork() {
	m.network = nil
	m.clearednetwork = false
}

// Where appends a list predicates to the NodeMutation builder.
func (m *NodeMutation) Where(ps ...predicate.Node) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the NodeMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *NodeMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Node, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *NodeMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *NodeMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Node).
func (m *NodeMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *NodeMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *NodeMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *NodeMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown Node field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *NodeMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Node field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *NodeMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *NodeMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *NodeMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown Node numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *NodeMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *NodeMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *NodeMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Node nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *NodeMutation) ResetField(name string) error {
	return fmt.Errorf("unknown Node field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *NodeMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.network != nil {
		edges = append(edges, node.EdgeNetwork)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *NodeMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case node.EdgeNetwork:
		if id := m.network; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *NodeMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *NodeMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *NodeMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearednetwork {
		edges = append(edges, node.EdgeNetwork)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *NodeMutation) EdgeCleared(name string) bool {
	switch name {
	case node.EdgeNetwork:
		return m.clearednetwork
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *NodeMutation) ClearEdge(name string) error {
	switch name {
	case node.EdgeNetwork:
		m.ClearNetwork()
		return nil
	}
	return fmt.Errorf("unknown Node unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *NodeMutation) ResetEdge(name string) error {
	switch name {
	case node.EdgeNetwork:
		m.ResetNetwork()
		return nil
	}
	return fmt.Errorf("unknown Node edge %s", name)
}
