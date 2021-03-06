// Code generated by entc, DO NOT EDIT.

package customer

const (
	// Label holds the string label denoting the customer type in the database.
	Label = "customer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"

	// EdgeReserves holds the string denoting the reserves edge name in mutations.
	EdgeReserves = "reserves"
	// EdgeFixs holds the string denoting the fixs edge name in mutations.
	EdgeFixs = "fixs"
	// EdgeCheckins holds the string denoting the checkins edge name in mutations.
	EdgeCheckins = "checkins"

	// Table holds the table name of the customer in the database.
	Table = "customers"
	// ReservesTable is the table the holds the reserves relation/edge.
	ReservesTable = "reserve_rooms"
	// ReservesInverseTable is the table name for the ReserveRoom entity.
	// It exists in this package in order to avoid circular dependency with the "reserveroom" package.
	ReservesInverseTable = "reserve_rooms"
	// ReservesColumn is the table column denoting the reserves relation/edge.
	ReservesColumn = "customer_id"
	// FixsTable is the table the holds the fixs relation/edge.
	FixsTable = "fix_rooms"
	// FixsInverseTable is the table name for the FixRoom entity.
	// It exists in this package in order to avoid circular dependency with the "fixroom" package.
	FixsInverseTable = "fix_rooms"
	// FixsColumn is the table column denoting the fixs relation/edge.
	FixsColumn = "customer_id"
	// CheckinsTable is the table the holds the checkins relation/edge.
	CheckinsTable = "check_ins"
	// CheckinsInverseTable is the table name for the CheckIn entity.
	// It exists in this package in order to avoid circular dependency with the "checkin" package.
	CheckinsInverseTable = "check_ins"
	// CheckinsColumn is the table column denoting the checkins relation/edge.
	CheckinsColumn = "customer_id"
)

// Columns holds all SQL columns for customer fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPassword,
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
