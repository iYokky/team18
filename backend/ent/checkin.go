// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team18/app/ent/checkin"
	"github.com/team18/app/ent/checkout"
	"github.com/team18/app/ent/counterstaff"
	"github.com/team18/app/ent/customer"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/reserveroom"
	"github.com/team18/app/ent/statuscheckin"
)

// CheckIn is the model entity for the CheckIn schema.
type CheckIn struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CheckinDate holds the value of the "checkin_date" field.
	CheckinDate time.Time `json:"checkin_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CheckInQuery when eager-loading is set.
	Edges       CheckInEdges `json:"edges"`
	staff_id    *int
	customer_id *int
	room_id     *int
	reserves_id *int
	status_id   *int
}

// CheckInEdges holds the relations/edges for other nodes in the graph.
type CheckInEdges struct {
	// Customer holds the value of the customer edge.
	Customer *Customer
	// Counter holds the value of the counter edge.
	Counter *CounterStaff
	// Reserveroom holds the value of the reserveroom edge.
	Reserveroom *ReserveRoom
	// Dataroom holds the value of the dataroom edge.
	Dataroom *DataRoom
	// Status holds the value of the status edge.
	Status *StatusCheckIn
	// Checkouts holds the value of the checkouts edge.
	Checkouts *Checkout
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckInEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[0] {
		if e.Customer == nil {
			// The edge customer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// CounterOrErr returns the Counter value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckInEdges) CounterOrErr() (*CounterStaff, error) {
	if e.loadedTypes[1] {
		if e.Counter == nil {
			// The edge counter was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: counterstaff.Label}
		}
		return e.Counter, nil
	}
	return nil, &NotLoadedError{edge: "counter"}
}

// ReserveroomOrErr returns the Reserveroom value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckInEdges) ReserveroomOrErr() (*ReserveRoom, error) {
	if e.loadedTypes[2] {
		if e.Reserveroom == nil {
			// The edge reserveroom was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: reserveroom.Label}
		}
		return e.Reserveroom, nil
	}
	return nil, &NotLoadedError{edge: "reserveroom"}
}

// DataroomOrErr returns the Dataroom value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckInEdges) DataroomOrErr() (*DataRoom, error) {
	if e.loadedTypes[3] {
		if e.Dataroom == nil {
			// The edge dataroom was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: dataroom.Label}
		}
		return e.Dataroom, nil
	}
	return nil, &NotLoadedError{edge: "dataroom"}
}

// StatusOrErr returns the Status value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckInEdges) StatusOrErr() (*StatusCheckIn, error) {
	if e.loadedTypes[4] {
		if e.Status == nil {
			// The edge status was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: statuscheckin.Label}
		}
		return e.Status, nil
	}
	return nil, &NotLoadedError{edge: "status"}
}

// CheckoutsOrErr returns the Checkouts value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckInEdges) CheckoutsOrErr() (*Checkout, error) {
	if e.loadedTypes[5] {
		if e.Checkouts == nil {
			// The edge checkouts was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: checkout.Label}
		}
		return e.Checkouts, nil
	}
	return nil, &NotLoadedError{edge: "checkouts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CheckIn) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // checkin_date
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*CheckIn) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // staff_id
		&sql.NullInt64{}, // customer_id
		&sql.NullInt64{}, // room_id
		&sql.NullInt64{}, // reserves_id
		&sql.NullInt64{}, // status_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CheckIn fields.
func (ci *CheckIn) assignValues(values ...interface{}) error {
	if m, n := len(values), len(checkin.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ci.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field checkin_date", values[0])
	} else if value.Valid {
		ci.CheckinDate = value.Time
	}
	values = values[1:]
	if len(values) == len(checkin.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field staff_id", value)
		} else if value.Valid {
			ci.staff_id = new(int)
			*ci.staff_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field customer_id", value)
		} else if value.Valid {
			ci.customer_id = new(int)
			*ci.customer_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field room_id", value)
		} else if value.Valid {
			ci.room_id = new(int)
			*ci.room_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field reserves_id", value)
		} else if value.Valid {
			ci.reserves_id = new(int)
			*ci.reserves_id = int(value.Int64)
		}
		if value, ok := values[4].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field status_id", value)
		} else if value.Valid {
			ci.status_id = new(int)
			*ci.status_id = int(value.Int64)
		}
	}
	return nil
}

// QueryCustomer queries the customer edge of the CheckIn.
func (ci *CheckIn) QueryCustomer() *CustomerQuery {
	return (&CheckInClient{config: ci.config}).QueryCustomer(ci)
}

// QueryCounter queries the counter edge of the CheckIn.
func (ci *CheckIn) QueryCounter() *CounterStaffQuery {
	return (&CheckInClient{config: ci.config}).QueryCounter(ci)
}

// QueryReserveroom queries the reserveroom edge of the CheckIn.
func (ci *CheckIn) QueryReserveroom() *ReserveRoomQuery {
	return (&CheckInClient{config: ci.config}).QueryReserveroom(ci)
}

// QueryDataroom queries the dataroom edge of the CheckIn.
func (ci *CheckIn) QueryDataroom() *DataRoomQuery {
	return (&CheckInClient{config: ci.config}).QueryDataroom(ci)
}

// QueryStatus queries the status edge of the CheckIn.
func (ci *CheckIn) QueryStatus() *StatusCheckInQuery {
	return (&CheckInClient{config: ci.config}).QueryStatus(ci)
}

// QueryCheckouts queries the checkouts edge of the CheckIn.
func (ci *CheckIn) QueryCheckouts() *CheckoutQuery {
	return (&CheckInClient{config: ci.config}).QueryCheckouts(ci)
}

// Update returns a builder for updating this CheckIn.
// Note that, you need to call CheckIn.Unwrap() before calling this method, if this CheckIn
// was returned from a transaction, and the transaction was committed or rolled back.
func (ci *CheckIn) Update() *CheckInUpdateOne {
	return (&CheckInClient{config: ci.config}).UpdateOne(ci)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ci *CheckIn) Unwrap() *CheckIn {
	tx, ok := ci.config.driver.(*txDriver)
	if !ok {
		panic("ent: CheckIn is not a transactional entity")
	}
	ci.config.driver = tx.drv
	return ci
}

// String implements the fmt.Stringer.
func (ci *CheckIn) String() string {
	var builder strings.Builder
	builder.WriteString("CheckIn(")
	builder.WriteString(fmt.Sprintf("id=%v", ci.ID))
	builder.WriteString(", checkin_date=")
	builder.WriteString(ci.CheckinDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CheckIns is a parsable slice of CheckIn.
type CheckIns []*CheckIn

func (ci CheckIns) config(cfg config) {
	for _i := range ci {
		ci[_i].config = cfg
	}
}
