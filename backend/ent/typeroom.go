// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team18/app/ent/typeroom"
)

// TypeRoom is the model entity for the TypeRoom schema.
type TypeRoom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TypeName holds the value of the "type_name" field.
	TypeName string `json:"type_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TypeRoomQuery when eager-loading is set.
	Edges TypeRoomEdges `json:"edges"`
}

// TypeRoomEdges holds the relations/edges for other nodes in the graph.
type TypeRoomEdges struct {
	// Datarooms holds the value of the datarooms edge.
	Datarooms []*DataRoom
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DataroomsOrErr returns the Datarooms value or an error if the edge
// was not loaded in eager-loading.
func (e TypeRoomEdges) DataroomsOrErr() ([]*DataRoom, error) {
	if e.loadedTypes[0] {
		return e.Datarooms, nil
	}
	return nil, &NotLoadedError{edge: "datarooms"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TypeRoom) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // type_name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TypeRoom fields.
func (tr *TypeRoom) assignValues(values ...interface{}) error {
	if m, n := len(values), len(typeroom.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	tr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type_name", values[0])
	} else if value.Valid {
		tr.TypeName = value.String
	}
	return nil
}

// QueryDatarooms queries the datarooms edge of the TypeRoom.
func (tr *TypeRoom) QueryDatarooms() *DataRoomQuery {
	return (&TypeRoomClient{config: tr.config}).QueryDatarooms(tr)
}

// Update returns a builder for updating this TypeRoom.
// Note that, you need to call TypeRoom.Unwrap() before calling this method, if this TypeRoom
// was returned from a transaction, and the transaction was committed or rolled back.
func (tr *TypeRoom) Update() *TypeRoomUpdateOne {
	return (&TypeRoomClient{config: tr.config}).UpdateOne(tr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (tr *TypeRoom) Unwrap() *TypeRoom {
	tx, ok := tr.config.driver.(*txDriver)
	if !ok {
		panic("ent: TypeRoom is not a transactional entity")
	}
	tr.config.driver = tx.drv
	return tr
}

// String implements the fmt.Stringer.
func (tr *TypeRoom) String() string {
	var builder strings.Builder
	builder.WriteString("TypeRoom(")
	builder.WriteString(fmt.Sprintf("id=%v", tr.ID))
	builder.WriteString(", type_name=")
	builder.WriteString(tr.TypeName)
	builder.WriteByte(')')
	return builder.String()
}

// TypeRooms is a parsable slice of TypeRoom.
type TypeRooms []*TypeRoom

func (tr TypeRooms) config(cfg config) {
	for _i := range tr {
		tr[_i].config = cfg
	}
}