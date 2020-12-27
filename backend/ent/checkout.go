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
	"github.com/team18/app/ent/status"
)

// Checkout is the model entity for the Checkout schema.
type Checkout struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CheckoutDate holds the value of the "checkout_date" field.
	CheckoutDate time.Time `json:"checkout_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CheckoutQuery when eager-loading is set.
	Edges              CheckoutEdges `json:"edges"`
	check_in_checkouts *int
	staff_id           *int
	status_checkouts   *int
}

// CheckoutEdges holds the relations/edges for other nodes in the graph.
type CheckoutEdges struct {
	// Statuss holds the value of the statuss edge.
	Statuss *Status
	// Counterstaffs holds the value of the counterstaffs edge.
	Counterstaffs *CounterStaff
	// Checkins holds the value of the checkins edge.
	Checkins *CheckIn
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// StatussOrErr returns the Statuss value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckoutEdges) StatussOrErr() (*Status, error) {
	if e.loadedTypes[0] {
		if e.Statuss == nil {
			// The edge statuss was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: status.Label}
		}
		return e.Statuss, nil
	}
	return nil, &NotLoadedError{edge: "statuss"}
}

// CounterstaffsOrErr returns the Counterstaffs value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckoutEdges) CounterstaffsOrErr() (*CounterStaff, error) {
	if e.loadedTypes[1] {
		if e.Counterstaffs == nil {
			// The edge counterstaffs was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: counterstaff.Label}
		}
		return e.Counterstaffs, nil
	}
	return nil, &NotLoadedError{edge: "counterstaffs"}
}

// CheckinsOrErr returns the Checkins value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckoutEdges) CheckinsOrErr() (*CheckIn, error) {
	if e.loadedTypes[2] {
		if e.Checkins == nil {
			// The edge checkins was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: checkin.Label}
		}
		return e.Checkins, nil
	}
	return nil, &NotLoadedError{edge: "checkins"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Checkout) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // checkout_date
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Checkout) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // check_in_checkouts
		&sql.NullInt64{}, // staff_id
		&sql.NullInt64{}, // status_checkouts
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Checkout fields.
func (c *Checkout) assignValues(values ...interface{}) error {
	if m, n := len(values), len(checkout.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field checkout_date", values[0])
	} else if value.Valid {
		c.CheckoutDate = value.Time
	}
	values = values[1:]
	if len(values) == len(checkout.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field check_in_checkouts", value)
		} else if value.Valid {
			c.check_in_checkouts = new(int)
			*c.check_in_checkouts = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field staff_id", value)
		} else if value.Valid {
			c.staff_id = new(int)
			*c.staff_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field status_checkouts", value)
		} else if value.Valid {
			c.status_checkouts = new(int)
			*c.status_checkouts = int(value.Int64)
		}
	}
	return nil
}

// QueryStatuss queries the statuss edge of the Checkout.
func (c *Checkout) QueryStatuss() *StatusQuery {
	return (&CheckoutClient{config: c.config}).QueryStatuss(c)
}

// QueryCounterstaffs queries the counterstaffs edge of the Checkout.
func (c *Checkout) QueryCounterstaffs() *CounterStaffQuery {
	return (&CheckoutClient{config: c.config}).QueryCounterstaffs(c)
}

// QueryCheckins queries the checkins edge of the Checkout.
func (c *Checkout) QueryCheckins() *CheckInQuery {
	return (&CheckoutClient{config: c.config}).QueryCheckins(c)
}

// Update returns a builder for updating this Checkout.
// Note that, you need to call Checkout.Unwrap() before calling this method, if this Checkout
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Checkout) Update() *CheckoutUpdateOne {
	return (&CheckoutClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Checkout) Unwrap() *Checkout {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Checkout is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Checkout) String() string {
	var builder strings.Builder
	builder.WriteString("Checkout(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", checkout_date=")
	builder.WriteString(c.CheckoutDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Checkouts is a parsable slice of Checkout.
type Checkouts []*Checkout

func (c Checkouts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
