// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team18/app/ent/statuscheckin"
)

// StatusCheckIn is the model entity for the StatusCheckIn schema.
type StatusCheckIn struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StatusName holds the value of the "status_name" field.
	StatusName string `json:"status_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StatusCheckInQuery when eager-loading is set.
	Edges StatusCheckInEdges `json:"edges"`
}

// StatusCheckInEdges holds the relations/edges for other nodes in the graph.
type StatusCheckInEdges struct {
	// Checkins holds the value of the checkins edge.
	Checkins []*CheckIn
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CheckinsOrErr returns the Checkins value or an error if the edge
// was not loaded in eager-loading.
func (e StatusCheckInEdges) CheckinsOrErr() ([]*CheckIn, error) {
	if e.loadedTypes[0] {
		return e.Checkins, nil
	}
	return nil, &NotLoadedError{edge: "checkins"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StatusCheckIn) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // status_name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StatusCheckIn fields.
func (sci *StatusCheckIn) assignValues(values ...interface{}) error {
	if m, n := len(values), len(statuscheckin.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	sci.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field status_name", values[0])
	} else if value.Valid {
		sci.StatusName = value.String
	}
	return nil
}

// QueryCheckins queries the checkins edge of the StatusCheckIn.
func (sci *StatusCheckIn) QueryCheckins() *CheckInQuery {
	return (&StatusCheckInClient{config: sci.config}).QueryCheckins(sci)
}

// Update returns a builder for updating this StatusCheckIn.
// Note that, you need to call StatusCheckIn.Unwrap() before calling this method, if this StatusCheckIn
// was returned from a transaction, and the transaction was committed or rolled back.
func (sci *StatusCheckIn) Update() *StatusCheckInUpdateOne {
	return (&StatusCheckInClient{config: sci.config}).UpdateOne(sci)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (sci *StatusCheckIn) Unwrap() *StatusCheckIn {
	tx, ok := sci.config.driver.(*txDriver)
	if !ok {
		panic("ent: StatusCheckIn is not a transactional entity")
	}
	sci.config.driver = tx.drv
	return sci
}

// String implements the fmt.Stringer.
func (sci *StatusCheckIn) String() string {
	var builder strings.Builder
	builder.WriteString("StatusCheckIn(")
	builder.WriteString(fmt.Sprintf("id=%v", sci.ID))
	builder.WriteString(", status_name=")
	builder.WriteString(sci.StatusName)
	builder.WriteByte(')')
	return builder.String()
}

// StatusCheckIns is a parsable slice of StatusCheckIn.
type StatusCheckIns []*StatusCheckIn

func (sci StatusCheckIns) config(cfg config) {
	for _i := range sci {
		sci[_i].config = cfg
	}
}
