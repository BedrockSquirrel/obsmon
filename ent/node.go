// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/bedrocksquirrel/obsmon/ent/network"
	"github.com/bedrocksquirrel/obsmon/ent/node"
)

// Node is the model entity for the Node schema.
type Node struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NodeQuery when eager-loading is set.
	Edges         NodeEdges `json:"edges"`
	network_nodes *int
	selectValues  sql.SelectValues
}

// NodeEdges holds the relations/edges for other nodes in the graph.
type NodeEdges struct {
	// Network holds the value of the network edge.
	Network *Network `json:"network,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// NetworkOrErr returns the Network value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NodeEdges) NetworkOrErr() (*Network, error) {
	if e.loadedTypes[0] {
		if e.Network == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: network.Label}
		}
		return e.Network, nil
	}
	return nil, &NotLoadedError{edge: "network"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Node) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case node.FieldID:
			values[i] = new(sql.NullInt64)
		case node.ForeignKeys[0]: // network_nodes
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Node fields.
func (n *Node) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case node.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case node.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field network_nodes", value)
			} else if value.Valid {
				n.network_nodes = new(int)
				*n.network_nodes = int(value.Int64)
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Node.
// This includes values selected through modifiers, order, etc.
func (n *Node) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryNetwork queries the "network" edge of the Node entity.
func (n *Node) QueryNetwork() *NetworkQuery {
	return NewNodeClient(n.config).QueryNetwork(n)
}

// Update returns a builder for updating this Node.
// Note that you need to call Node.Unwrap() before calling this method if this Node
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Node) Update() *NodeUpdateOne {
	return NewNodeClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Node entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Node) Unwrap() *Node {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Node is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Node) String() string {
	var builder strings.Builder
	builder.WriteString("Node(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Nodes is a parsable slice of Node.
type Nodes []*Node