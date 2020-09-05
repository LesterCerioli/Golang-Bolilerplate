package app

import (
	"context"
	"time"

	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gopkg.in/guregu/null.v4"
)

// ErrNoRows is returned when a corresponding record is not found
var ErrNotFound = errors.New("identity: not found")

type Tenant struct {
	ID        uint        `json:"id" validate:"omitempty,gte=0"`
	Name      null.String `validate:"required_without=ID"`
	Accounts  []*Account  `json:"accounts" validate:"omitempty,dive,required"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type Account struct {
	ID        uint        `json:"id" validate:"omitempty,gte=0"`
	Name      null.String `validate:"required_without=ID"`
	Tenant    Tenant      `json:"tenant" validate:"omitempty,required_without=ID,dive,required"`
	Users     []*User     `json:"users" validate:"omitempty,required_without=ID,dive,required"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type User struct {
	ID      uint    `json:"id" validate:"omitempty,gte=0"`
	Account Account `json:"account" validate:"omitempty,required_without=ID,dive,required"`

	FirstName null.String `json:"first_name" validate:"omitempty,required_without=ID,gte=4,lte=254"`
	LastName  null.String `json:"last_name" validate:"omitempty,required_without=ID,gte=2,lte=254"`
	Email     null.String `json:"email" validate:"omitempty,required_without=ID,email,isUniqueEmail,lte=254"`
	Password  null.String `json:"password" validate:"omitempty,required_without=ID,lte=254"`
	Active    null.Bool   `json:"omitempty,required_without=ID,active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type IdentityService interface {
	// Create Account, User and Tenant and associations
	CreateIdentity(c context.Context, b *Account) (Account, error)

	GetAccount(c context.Context, id uint) (Account, error)
	// Update Account
	UpdateAccount(c context.Context, b *Account) (Account, error)

	// Find by email
	FindUserByEmail(c context.Context, t *Tenant, email string) (Account, error)
}

func RegisterIdentiyServiceValidators(bs IdentityService, v *validator.Validate, l *zap.Logger) error {
	bsv := UserServiceValidator{bs: bs, sv: v, logger: l}
	return v.RegisterValidation("isUniqueEmail", bsv.isUniqueEmail)
}

type UserServiceValidator struct {
	bs     IdentityService
	sv     *validator.Validate
	logger *zap.Logger
}

func (bsv *UserServiceValidator) isUniqueEmail(fl validator.FieldLevel) bool {
	// Get current email
	emailField := fl.Field()
	if !emailField.IsValid() {
		return false
	}
	email := emailField.String()

	// Get current Tenant
	p := fl.Parent()
	tenantField := reflect.Indirect(p).FieldByName("Tenant")
	if !tenantField.IsValid() {
		return false
	}

	tenant, ok := tenantField.Interface().(Tenant)
	if !ok {
		return false
	}

	// Find in DB
	c := context.Background()
	acc, err := bsv.bs.FindUserByEmail(c, &tenant, email)
	if err == ErrNotFound {
		return true
	} else if err != nil {
		// Got an error and there is no way to express that back to our client.
		return false
	}

	// Get ID
	p = fl.Parent()
	idField := reflect.Indirect(p).FieldByName("ID")
	if !idField.IsValid() {
		return false
	}

	// Return true if this email is used by the SAME employee
	return idField.Uint() == uint64(acc.ID)
}
