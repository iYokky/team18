// Code generated by entc, DO NOT EDIT.

package furniture

const (
	// Label holds the string label denoting the furniture type in the database.
	Label = "furniture"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFurnitureName holds the string denoting the furniture_name field in the database.
	FieldFurnitureName = "furniture_name"

	// EdgeDetails holds the string denoting the details edge name in mutations.
	EdgeDetails = "details"

	// Table holds the table name of the furniture in the database.
	Table = "furnitures"
	// DetailsTable is the table the holds the details relation/edge.
	DetailsTable = "furniture_details"
	// DetailsInverseTable is the table name for the FurnitureDetail entity.
	// It exists in this package in order to avoid circular dependency with the "furnituredetail" package.
	DetailsInverseTable = "furniture_details"
	// DetailsColumn is the table column denoting the details relation/edge.
	DetailsColumn = "furniture_id"
)

// Columns holds all SQL columns for furniture fields.
var Columns = []string{
	FieldID,
	FieldFurnitureName,
}

var (
	// FurnitureNameValidator is a validator for the "furniture_name" field. It is called by the builders before save.
	FurnitureNameValidator func(string) error
)
