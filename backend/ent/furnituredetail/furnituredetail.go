// Code generated by entc, DO NOT EDIT.

package furnituredetail

const (
	// Label holds the string label denoting the furnituredetail type in the database.
	Label = "furniture_detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDateAdd holds the string denoting the date_add field in the database.
	FieldDateAdd = "date_add"

	// EdgeFixs holds the string denoting the fixs edge name in mutations.
	EdgeFixs = "fixs"
	// EdgeFurnitures holds the string denoting the furnitures edge name in mutations.
	EdgeFurnitures = "furnitures"
	// EdgeTypes holds the string denoting the types edge name in mutations.
	EdgeTypes = "types"
	// EdgeRooms holds the string denoting the rooms edge name in mutations.
	EdgeRooms = "rooms"

	// Table holds the table name of the furnituredetail in the database.
	Table = "furniture_details"
	// FixsTable is the table the holds the fixs relation/edge.
	FixsTable = "fix_rooms"
	// FixsInverseTable is the table name for the FixRoom entity.
	// It exists in this package in order to avoid circular dependency with the "fixroom" package.
	FixsInverseTable = "fix_rooms"
	// FixsColumn is the table column denoting the fixs relation/edge.
	FixsColumn = "object_id"
	// FurnituresTable is the table the holds the furnitures relation/edge.
	FurnituresTable = "furniture_details"
	// FurnituresInverseTable is the table name for the Furniture entity.
	// It exists in this package in order to avoid circular dependency with the "furniture" package.
	FurnituresInverseTable = "furnitures"
	// FurnituresColumn is the table column denoting the furnitures relation/edge.
	FurnituresColumn = "furniture_id"
	// TypesTable is the table the holds the types relation/edge.
	TypesTable = "furniture_details"
	// TypesInverseTable is the table name for the FurnitureType entity.
	// It exists in this package in order to avoid circular dependency with the "furnituretype" package.
	TypesInverseTable = "furniture_types"
	// TypesColumn is the table column denoting the types relation/edge.
	TypesColumn = "type_id"
	// RoomsTable is the table the holds the rooms relation/edge.
	RoomsTable = "furniture_details"
	// RoomsInverseTable is the table name for the DataRoom entity.
	// It exists in this package in order to avoid circular dependency with the "dataroom" package.
	RoomsInverseTable = "data_rooms"
	// RoomsColumn is the table column denoting the rooms relation/edge.
	RoomsColumn = "room_id"
)

// Columns holds all SQL columns for furnituredetail fields.
var Columns = []string{
	FieldID,
	FieldDateAdd,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the FurnitureDetail type.
var ForeignKeys = []string{
	"room_id",
	"furniture_id",
	"type_id",
}