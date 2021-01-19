// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/team18/app/ent/counterstaff"
	"github.com/team18/app/ent/customer"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/furniture"
	"github.com/team18/app/ent/promotion"
	"github.com/team18/app/ent/reserveroom"
	"github.com/team18/app/ent/schema"
	"github.com/team18/app/ent/status"
	"github.com/team18/app/ent/statuscheckin"
	"github.com/team18/app/ent/statusreserve"
	"github.com/team18/app/ent/statusroom"
	"github.com/team18/app/ent/typeroom"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
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
	dataroom.PriceValidator = func() func(float64) error {
		validators := dataroomDescPrice.Validators
		fns := [...]func(float64) error{
			validators[0].(func(float64) error),
			validators[1].(func(float64) error),
		}
		return func(price float64) error {
			for _, fn := range fns {
				if err := fn(price); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// dataroomDescRoomnumber is the schema descriptor for roomnumber field.
	dataroomDescRoomnumber := dataroomFields[1].Descriptor()
	// dataroom.RoomnumberValidator is a validator for the "roomnumber" field. It is called by the builders before save.
	dataroom.RoomnumberValidator = dataroomDescRoomnumber.Validators[0].(func(string) error)
	// dataroomDescRoomdetail is the schema descriptor for roomdetail field.
	dataroomDescRoomdetail := dataroomFields[2].Descriptor()
	// dataroom.RoomdetailValidator is a validator for the "roomdetail" field. It is called by the builders before save.
	dataroom.RoomdetailValidator = dataroomDescRoomdetail.Validators[0].(func(string) error)
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
	// reserveroomDescDuration is the schema descriptor for duration field.
	reserveroomDescDuration := reserveroomFields[1].Descriptor()
	// reserveroom.DurationValidator is a validator for the "duration" field. It is called by the builders before save.
	reserveroom.DurationValidator = func() func(int) error {
		validators := reserveroomDescDuration.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(duration int) error {
			for _, fn := range fns {
				if err := fn(duration); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// reserveroomDescProvince is the schema descriptor for province field.
	reserveroomDescProvince := reserveroomFields[2].Descriptor()
	// reserveroom.ProvinceValidator is a validator for the "province" field. It is called by the builders before save.
	reserveroom.ProvinceValidator = func() func(string) error {
		validators := reserveroomDescProvince.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(province string) error {
			for _, fn := range fns {
				if err := fn(province); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// reserveroomDescAmount is the schema descriptor for amount field.
	reserveroomDescAmount := reserveroomFields[3].Descriptor()
	// reserveroom.AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	reserveroom.AmountValidator = func() func(int) error {
		validators := reserveroomDescAmount.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
			validators[2].(func(int) error),
		}
		return func(amount int) error {
			for _, fn := range fns {
				if err := fn(amount); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// reserveroomDescTel is the schema descriptor for tel field.
	reserveroomDescTel := reserveroomFields[4].Descriptor()
	// reserveroom.TelValidator is a validator for the "tel" field. It is called by the builders before save.
	reserveroom.TelValidator = reserveroomDescTel.Validators[0].(func(string) error)
	// reserveroomDescNetPrice is the schema descriptor for net_price field.
	reserveroomDescNetPrice := reserveroomFields[5].Descriptor()
	// reserveroom.NetPriceValidator is a validator for the "net_price" field. It is called by the builders before save.
	reserveroom.NetPriceValidator = reserveroomDescNetPrice.Validators[0].(func(float64) error)
	statusFields := schema.Status{}.Fields()
	_ = statusFields
	// statusDescDescription is the schema descriptor for description field.
	statusDescDescription := statusFields[0].Descriptor()
	// status.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	status.DescriptionValidator = statusDescDescription.Validators[0].(func(string) error)
	statuscheckinFields := schema.StatusCheckIn{}.Fields()
	_ = statuscheckinFields
	// statuscheckinDescStatusName is the schema descriptor for status_name field.
	statuscheckinDescStatusName := statuscheckinFields[0].Descriptor()
	// statuscheckin.StatusNameValidator is a validator for the "status_name" field. It is called by the builders before save.
	statuscheckin.StatusNameValidator = statuscheckinDescStatusName.Validators[0].(func(string) error)
	statusreserveFields := schema.StatusReserve{}.Fields()
	_ = statusreserveFields
	// statusreserveDescStatusName is the schema descriptor for status_name field.
	statusreserveDescStatusName := statusreserveFields[0].Descriptor()
	// statusreserve.StatusNameValidator is a validator for the "status_name" field. It is called by the builders before save.
	statusreserve.StatusNameValidator = statusreserveDescStatusName.Validators[0].(func(string) error)
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
