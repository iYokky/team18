// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// CheckInsColumns holds the columns for the "check_ins" table.
	CheckInsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "checkin_date", Type: field.TypeTime},
		{Name: "staff_id", Type: field.TypeInt, Nullable: true},
		{Name: "customer_id", Type: field.TypeInt, Nullable: true},
		{Name: "reserves_id", Type: field.TypeInt, Nullable: true},
	}
	// CheckInsTable holds the schema information for the "check_ins" table.
	CheckInsTable = &schema.Table{
		Name:       "check_ins",
		Columns:    CheckInsColumns,
		PrimaryKey: []*schema.Column{CheckInsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "check_ins_counter_staffs_checkins",
				Columns: []*schema.Column{CheckInsColumns[2]},

				RefColumns: []*schema.Column{CounterStaffsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "check_ins_customers_checkins",
				Columns: []*schema.Column{CheckInsColumns[3]},

				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "check_ins_reserve_rooms_checkins",
				Columns: []*schema.Column{CheckInsColumns[4]},

				RefColumns: []*schema.Column{ReserveRoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CheckoutsColumns holds the columns for the "checkouts" table.
	CheckoutsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "checkout_date", Type: field.TypeTime},
		{Name: "check_in_checkouts", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "staff_id", Type: field.TypeInt, Nullable: true},
		{Name: "status_checkouts", Type: field.TypeInt, Nullable: true},
	}
	// CheckoutsTable holds the schema information for the "checkouts" table.
	CheckoutsTable = &schema.Table{
		Name:       "checkouts",
		Columns:    CheckoutsColumns,
		PrimaryKey: []*schema.Column{CheckoutsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "checkouts_check_ins_checkouts",
				Columns: []*schema.Column{CheckoutsColumns[2]},

				RefColumns: []*schema.Column{CheckInsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "checkouts_counter_staffs_checkouts",
				Columns: []*schema.Column{CheckoutsColumns[3]},

				RefColumns: []*schema.Column{CounterStaffsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "checkouts_status_checkouts",
				Columns: []*schema.Column{CheckoutsColumns[4]},

				RefColumns: []*schema.Column{StatusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CounterStaffsColumns holds the columns for the "counter_staffs" table.
	CounterStaffsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString, Unique: true},
	}
	// CounterStaffsTable holds the schema information for the "counter_staffs" table.
	CounterStaffsTable = &schema.Table{
		Name:        "counter_staffs",
		Columns:     CounterStaffsColumns,
		PrimaryKey:  []*schema.Column{CounterStaffsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString, Unique: true},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:        "customers",
		Columns:     CustomersColumns,
		PrimaryKey:  []*schema.Column{CustomersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DataRoomsColumns holds the columns for the "data_rooms" table.
	DataRoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "roomnumber", Type: field.TypeString},
		{Name: "promotion_id", Type: field.TypeInt, Nullable: true},
		{Name: "statusroom_id", Type: field.TypeInt, Nullable: true},
		{Name: "typeroom_id", Type: field.TypeInt, Nullable: true},
	}
	// DataRoomsTable holds the schema information for the "data_rooms" table.
	DataRoomsTable = &schema.Table{
		Name:       "data_rooms",
		Columns:    DataRoomsColumns,
		PrimaryKey: []*schema.Column{DataRoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "data_rooms_promotions_datarooms",
				Columns: []*schema.Column{DataRoomsColumns[3]},

				RefColumns: []*schema.Column{PromotionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "data_rooms_status_rooms_datarooms",
				Columns: []*schema.Column{DataRoomsColumns[4]},

				RefColumns: []*schema.Column{StatusRoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "data_rooms_type_rooms_datarooms",
				Columns: []*schema.Column{DataRoomsColumns[5]},

				RefColumns: []*schema.Column{TypeRoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FixRoomsColumns holds the columns for the "fix_rooms" table.
	FixRoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "fix_detail", Type: field.TypeString},
		{Name: "customer_id", Type: field.TypeInt, Nullable: true},
		{Name: "room_id", Type: field.TypeInt, Nullable: true},
		{Name: "object_id", Type: field.TypeInt, Nullable: true},
	}
	// FixRoomsTable holds the schema information for the "fix_rooms" table.
	FixRoomsTable = &schema.Table{
		Name:       "fix_rooms",
		Columns:    FixRoomsColumns,
		PrimaryKey: []*schema.Column{FixRoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "fix_rooms_customers_fixs",
				Columns: []*schema.Column{FixRoomsColumns[2]},

				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "fix_rooms_data_rooms_fixs",
				Columns: []*schema.Column{FixRoomsColumns[3]},

				RefColumns: []*schema.Column{DataRoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "fix_rooms_furniture_details_fixs",
				Columns: []*schema.Column{FixRoomsColumns[4]},

				RefColumns: []*schema.Column{FurnitureDetailsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FurnituresColumns holds the columns for the "furnitures" table.
	FurnituresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "furniture_name", Type: field.TypeString, Unique: true},
	}
	// FurnituresTable holds the schema information for the "furnitures" table.
	FurnituresTable = &schema.Table{
		Name:        "furnitures",
		Columns:     FurnituresColumns,
		PrimaryKey:  []*schema.Column{FurnituresColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// FurnitureDetailsColumns holds the columns for the "furniture_details" table.
	FurnitureDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date_add", Type: field.TypeTime},
		{Name: "room_id", Type: field.TypeInt, Nullable: true},
		{Name: "furniture_id", Type: field.TypeInt, Nullable: true},
		{Name: "type_id", Type: field.TypeInt, Nullable: true},
	}
	// FurnitureDetailsTable holds the schema information for the "furniture_details" table.
	FurnitureDetailsTable = &schema.Table{
		Name:       "furniture_details",
		Columns:    FurnitureDetailsColumns,
		PrimaryKey: []*schema.Column{FurnitureDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "furniture_details_data_rooms_details",
				Columns: []*schema.Column{FurnitureDetailsColumns[2]},

				RefColumns: []*schema.Column{DataRoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "furniture_details_furnitures_details",
				Columns: []*schema.Column{FurnitureDetailsColumns[3]},

				RefColumns: []*schema.Column{FurnituresColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "furniture_details_furniture_types_details",
				Columns: []*schema.Column{FurnitureDetailsColumns[4]},

				RefColumns: []*schema.Column{FurnitureTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FurnitureTypesColumns holds the columns for the "furniture_types" table.
	FurnitureTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "furniture_type", Type: field.TypeString},
	}
	// FurnitureTypesTable holds the schema information for the "furniture_types" table.
	FurnitureTypesTable = &schema.Table{
		Name:        "furniture_types",
		Columns:     FurnitureTypesColumns,
		PrimaryKey:  []*schema.Column{FurnitureTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PromotionsColumns holds the columns for the "promotions" table.
	PromotionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "promotion_name", Type: field.TypeString, Unique: true},
		{Name: "discount", Type: field.TypeFloat64},
	}
	// PromotionsTable holds the schema information for the "promotions" table.
	PromotionsTable = &schema.Table{
		Name:        "promotions",
		Columns:     PromotionsColumns,
		PrimaryKey:  []*schema.Column{PromotionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ReserveRoomsColumns holds the columns for the "reserve_rooms" table.
	ReserveRoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "reserve_date", Type: field.TypeTime},
		{Name: "date_out", Type: field.TypeTime},
		{Name: "net_price", Type: field.TypeFloat64},
		{Name: "customer_id", Type: field.TypeInt, Nullable: true},
		{Name: "room_id", Type: field.TypeInt, Nullable: true},
		{Name: "promotion_id", Type: field.TypeInt, Nullable: true},
	}
	// ReserveRoomsTable holds the schema information for the "reserve_rooms" table.
	ReserveRoomsTable = &schema.Table{
		Name:       "reserve_rooms",
		Columns:    ReserveRoomsColumns,
		PrimaryKey: []*schema.Column{ReserveRoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "reserve_rooms_customers_reserves",
				Columns: []*schema.Column{ReserveRoomsColumns[4]},

				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "reserve_rooms_data_rooms_reserves",
				Columns: []*schema.Column{ReserveRoomsColumns[5]},

				RefColumns: []*schema.Column{DataRoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "reserve_rooms_promotions_reserves",
				Columns: []*schema.Column{ReserveRoomsColumns[6]},

				RefColumns: []*schema.Column{PromotionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StatusColumns holds the columns for the "status" table.
	StatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "description", Type: field.TypeString},
	}
	// StatusTable holds the schema information for the "status" table.
	StatusTable = &schema.Table{
		Name:        "status",
		Columns:     StatusColumns,
		PrimaryKey:  []*schema.Column{StatusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// StatusRoomsColumns holds the columns for the "status_rooms" table.
	StatusRoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status_name", Type: field.TypeString},
	}
	// StatusRoomsTable holds the schema information for the "status_rooms" table.
	StatusRoomsTable = &schema.Table{
		Name:        "status_rooms",
		Columns:     StatusRoomsColumns,
		PrimaryKey:  []*schema.Column{StatusRoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TypeRoomsColumns holds the columns for the "type_rooms" table.
	TypeRoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type_name", Type: field.TypeString},
	}
	// TypeRoomsTable holds the schema information for the "type_rooms" table.
	TypeRoomsTable = &schema.Table{
		Name:        "type_rooms",
		Columns:     TypeRoomsColumns,
		PrimaryKey:  []*schema.Column{TypeRoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CheckInsTable,
		CheckoutsTable,
		CounterStaffsTable,
		CustomersTable,
		DataRoomsTable,
		FixRoomsTable,
		FurnituresTable,
		FurnitureDetailsTable,
		FurnitureTypesTable,
		PromotionsTable,
		ReserveRoomsTable,
		StatusTable,
		StatusRoomsTable,
		TypeRoomsTable,
	}
)

func init() {
	CheckInsTable.ForeignKeys[0].RefTable = CounterStaffsTable
	CheckInsTable.ForeignKeys[1].RefTable = CustomersTable
	CheckInsTable.ForeignKeys[2].RefTable = ReserveRoomsTable
	CheckoutsTable.ForeignKeys[0].RefTable = CheckInsTable
	CheckoutsTable.ForeignKeys[1].RefTable = CounterStaffsTable
	CheckoutsTable.ForeignKeys[2].RefTable = StatusTable
	DataRoomsTable.ForeignKeys[0].RefTable = PromotionsTable
	DataRoomsTable.ForeignKeys[1].RefTable = StatusRoomsTable
	DataRoomsTable.ForeignKeys[2].RefTable = TypeRoomsTable
	FixRoomsTable.ForeignKeys[0].RefTable = CustomersTable
	FixRoomsTable.ForeignKeys[1].RefTable = DataRoomsTable
	FixRoomsTable.ForeignKeys[2].RefTable = FurnitureDetailsTable
	FurnitureDetailsTable.ForeignKeys[0].RefTable = DataRoomsTable
	FurnitureDetailsTable.ForeignKeys[1].RefTable = FurnituresTable
	FurnitureDetailsTable.ForeignKeys[2].RefTable = FurnitureTypesTable
	ReserveRoomsTable.ForeignKeys[0].RefTable = CustomersTable
	ReserveRoomsTable.ForeignKeys[1].RefTable = DataRoomsTable
	ReserveRoomsTable.ForeignKeys[2].RefTable = PromotionsTable
}
