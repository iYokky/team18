// Code generated by entc, DO NOT EDIT.

package status

const (
	// Label holds the string label denoting the status type in the database.
	Label = "status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"

	// EdgeCheckouts holds the string denoting the checkouts edge name in mutations.
	EdgeCheckouts = "checkouts"

	// Table holds the table name of the status in the database.
	Table = "status"
	// CheckoutsTable is the table the holds the checkouts relation/edge.
	CheckoutsTable = "checkouts"
	// CheckoutsInverseTable is the table name for the Checkout entity.
	// It exists in this package in order to avoid circular dependency with the "checkout" package.
	CheckoutsInverseTable = "checkouts"
	// CheckoutsColumn is the table column denoting the checkouts relation/edge.
	CheckoutsColumn = "status_checkouts"
)

// Columns holds all SQL columns for status fields.
var Columns = []string{
	FieldID,
	FieldDescription,
}

var (
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)
