// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/team18/app/ent/checkin"
	"github.com/team18/app/ent/counterstaff"
	"github.com/team18/app/ent/customer"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/furniture"
	"github.com/team18/app/ent/promotion"
	"github.com/team18/app/ent/reserveroom"
	"github.com/team18/app/ent/schema"
	"github.com/team18/app/ent/statusroom"
	"github.com/team18/app/ent/typeroom"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	checkinFields := schema.CheckIn{}.Fields()
	_ = checkinFields
	// checkinDescCheckinDate is the schema descriptor for checkin_date field.
	checkinDescCheckinDate := checkinFields[0].Descriptor()
	// checkin.DefaultCheckinDate holds the default value on creation for the checkin_date field.
	checkin.DefaultCheckinDate = checkinDescCheckinDate.Default.(func() time.Time)
	counterstaffFields := schema.CounterStaff{}.Fields()
	_ = counterstaffFields
	// counterstaffDescName is the schema descriptor for name field.
	counterstaffDescName := counterstaffFields[0].Descriptor()
	// counterstaff.NameValidator is a validator for the "name" field. It is called by the builders before save.
	counterstaff.NameValidator = counterstaffDescName.Validators[0].(func(string) error)
	// counterstaffDescEmail is the schema descriptor for email field.
	counterstaffDescEmail := counterstaffFields[1].Descriptor()
	// counterstaff.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	counterstaff.EmailValidator = counterstaffDescEmail.Validators[0].(func(string) error)
	// counterstaffDescPassword is the schema descriptor for password field.
	counterstaffDescPassword := counterstaffFields[2].Descriptor()
	// counterstaff.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	counterstaff.PasswordValidator = counterstaffDescPassword.Validators[0].(func(string) error)
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescName is the schema descriptor for name field.
	customerDescName := customerFields[0].Descriptor()
	// customer.NameValidator is a validator for the "name" field. It is called by the builders before save.
	customer.NameValidator = customerDescName.Validators[0].(func(string) error)
	// customerDescEmail is the schema descriptor for email field.
	customerDescEmail := customerFields[1].Descriptor()
	// customer.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	customer.EmailValidator = customerDescEmail.Validators[0].(func(string) error)
	// customerDescPassword is the schema descriptor for password field.
	customerDescPassword := customerFields[2].Descriptor()
	// customer.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	customer.PasswordValidator = customerDescPassword.Validators[0].(func(string) error)
	dataroomFields := schema.DataRoom{}.Fields()
	_ = dataroomFields
	// dataroomDescPrice is the schema descriptor for price field.
	dataroomDescPrice := dataroomFields[0].Descriptor()
	// dataroom.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	dataroom.PriceValidator = dataroomDescPrice.Validators[0].(func(float64) error)
	// dataroomDescRoomnumber is the schema descriptor for roomnumber field.
	dataroomDescRoomnumber := dataroomFields[1].Descriptor()
	// dataroom.RoomnumberValidator is a validator for the "roomnumber" field. It is called by the builders before save.
	dataroom.RoomnumberValidator = dataroomDescRoomnumber.Validators[0].(func(string) error)
	furnitureFields := schema.Furniture{}.Fields()
	_ = furnitureFields
	// furnitureDescFurnitureName is the schema descriptor for furniture_name field.
	furnitureDescFurnitureName := furnitureFields[0].Descriptor()
	// furniture.FurnitureNameValidator is a validator for the "furniture_name" field. It is called by the builders before save.
	furniture.FurnitureNameValidator = furnitureDescFurnitureName.Validators[0].(func(string) error)
	promotionFields := schema.Promotion{}.Fields()
	_ = promotionFields
	// promotionDescPromotionName is the schema descriptor for promotion_name field.
	promotionDescPromotionName := promotionFields[0].Descriptor()
	// promotion.PromotionNameValidator is a validator for the "promotion_name" field. It is called by the builders before save.
	promotion.PromotionNameValidator = promotionDescPromotionName.Validators[0].(func(string) error)
	// promotionDescDiscount is the schema descriptor for discount field.
	promotionDescDiscount := promotionFields[1].Descriptor()
	// promotion.DiscountValidator is a validator for the "discount" field. It is called by the builders before save.
	promotion.DiscountValidator = promotionDescDiscount.Validators[0].(func(float64) error)
	reserveroomFields := schema.ReserveRoom{}.Fields()
	_ = reserveroomFields
	// reserveroomDescNetPrice is the schema descriptor for net_price field.
	reserveroomDescNetPrice := reserveroomFields[2].Descriptor()
	// reserveroom.NetPriceValidator is a validator for the "net_price" field. It is called by the builders before save.
	reserveroom.NetPriceValidator = reserveroomDescNetPrice.Validators[0].(func(float64) error)
	statusroomFields := schema.StatusRoom{}.Fields()
	_ = statusroomFields
	// statusroomDescStatusName is the schema descriptor for status_name field.
	statusroomDescStatusName := statusroomFields[0].Descriptor()
	// statusroom.StatusNameValidator is a validator for the "status_name" field. It is called by the builders before save.
	statusroom.StatusNameValidator = statusroomDescStatusName.Validators[0].(func(string) error)
	typeroomFields := schema.TypeRoom{}.Fields()
	_ = typeroomFields
	// typeroomDescTypeName is the schema descriptor for type_name field.
	typeroomDescTypeName := typeroomFields[0].Descriptor()
	// typeroom.TypeNameValidator is a validator for the "type_name" field. It is called by the builders before save.
	typeroom.TypeNameValidator = typeroomDescTypeName.Validators[0].(func(string) error)
}
