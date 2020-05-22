package app

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gopkg.in/guregu/null.v4"
	"reflect"
)

type Business struct {
	ID uint `json:"id" validate:"omitempty,gte=0"`
	//Name      string      `json:"name" validate:"required_without=ID,gte=2,lte=254"`
	Name      string      `validate:"required_without=ID"`
	Website   null.String `json:"website" validate:"omitempty,required_without=ID,uri,lte=254"`
	Offices   []*Office   `json:"offices" validate:"omitempty,required_without=ID,dive,required"`
	Employees []*Employee `json:"employees" validate:"omitempty,required_without=ID,dive,required"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type Office struct {
	ID   uint   `json:"id" validate:"omitempty,gte=0"`
	Name string `json:"name" validate:"required_without=ID,gte=4,lte=254"`
	//	Business Business`json:"Business" validate:"omitempty|isUpdate|required"`
	Zipcode null.String `json:"zipcode" validate:"omitempty,gte=2,lte=254"`
	State   null.String `json:"state" validate:"omitempty,gte=2,lte=254"`
	City    null.String `json:"city" validate:"omitempty,gte=2,lte=254"`
	Address null.String `json:"address" validate:"omitempty,gte=1,lte=254"`
	Country null.String `json:"country" validate:"omitempty,gte=2,lte=254"`
	Phone   null.String `json:"phone" validate:"omitempty,lte=254"`
	Email   null.String `json:"email" validate:"omitempty,email,lte=254"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Employee struct {
	ID        uint        `json:"id" validate:"omitempty,gte=0"`
	FirstName string      `json:"first_name" validate:"omitempty,required_without=ID,gte=4,lte=254"`
	LastName  null.String `json:"last_name" validate:"omitempty,required_without=ID,gte=2,lte=254"`
	Email     null.String `json:"email" validate:"omitempty,required_without=ID,email,lte=254"`
	Password  null.String `json:"password" validate:"omitempty,required_without=ID,lte=254"`
	Active    null.Bool   `json:"omitempty,required_without=ID,active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BusinessService interface {
	// Create business and associations
	Create(c context.Context, b *Business) (*Business, error)
	// Business only
	Get(c context.Context, id *uint) (*Business, error)
	// Update Business
	Update(c context.Context, b *Business) (*Business, error)

	// Find Emp by email
	GetEmployeeByEmail(c context.Context, email string) (*Employee, error)

	// Check if Employee password is correct
	//ComparePassword(c context.Context, b *Employee) error
	//Update Office
	//Delete Office
	//Member
}

func RegisterBusinessServiceValidators(bs BusinessService, v *validator.Validate, l *zap.Logger) error {
	bsv := businessServiceValidator{bs: bs, sv: v, logger: l}
	return v.RegisterValidation("isUniqueEmployeeEmail", bsv.isUniqueEmployeeEmail)
}

type businessServiceValidator struct {
	bs     BusinessService
	sv     *validator.Validate
	logger *zap.Logger
}

func (bsv *businessServiceValidator) isUniqueEmployeeEmail(fl validator.FieldLevel) bool {
	// Get current email
	emailField := fl.Field()
	if !emailField.IsValid() {
		return false
	}
	email := emailField.String()

	// Find in DB
	c := context.Background()
	emp, err := bsv.bs.GetEmployeeByEmail(c, email)
	if err != nil {
		return false // We can't return error...
	}
	if emp == nil {
		return true
	}

	// Get ID
	p := fl.Parent()
	idField := reflect.Indirect(p).FieldByName("ID")
	if !idField.IsValid() {
		return false
	}

	// Return true if this email is used by the SAME employee
	return idField.Uint() == uint64(emp.ID)
}
