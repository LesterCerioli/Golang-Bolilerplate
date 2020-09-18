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
	ID        null.Int    `json:"id" validate:"omitempty,gte=0"`
	Name      null.String `validate:"required_without=ID"`
	Accounts  []*Account  `json:"accounts" validate:"omitempty,dive,required"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type Account struct {
	ID        null.Int    `json:"id" validate:"omitempty,gte=0"`
	Name      null.String `validate:"required_without=ID"`
	Tenant    *Tenant     `json:"tenant" validate:"required_without=ID,omitempty,dive,required"`
	Users     []*User     `json:"users" validate:"required_without=ID,omitempty,dive,required"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type User struct {
	ID      null.Int `json:"id" validate:"omitempty,gte=0"`
	Account *Account `json:"account" validate:"required_without=ID,omitempty,dive,required"`

	FirstName null.String `json:"first_name" validate:"required_without=ID,omitempty,gte=4,lte=254"`
	LastName  null.String `json:"last_name" validate:"required_without=ID,omitempty,gte=2,lte=254"`
	Email     null.String `json:"email" validate:"required_without=ID,omitempty,email,isUniqueEmail,lte=254"`
	Password  null.String `json:"password" validate:"required_without=ID,omitempty,lte=254"`
	Active    null.Bool   `json:"required_without=ID,omitempty,active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type IdentityService interface {
	// Create Tenant
	CreateTenant(c context.Context, t Tenant, a Account, u User) (User, error)
	// Get Tenant
	GetTenant(c context.Context, id uint) (Tenant, error)
	// Update Tennat
	UpdateTenantDetails(c context.Context, t Tenant) (Tenant, error)

	// Create Account
	CreateAccount(c context.Context, t Tenant, a Account, u User) (User, error)
	// Get Account
	GetAccount(c context.Context, id uint) (Account, error)
	// Update Account
	UpdateAccountDetails(c context.Context, a Account) (Account, error)

	// Create User
	CreateUser(c context.Context, user User, account Account) (User, error)
	// Get User
	GetUser(c context.Context, id uint) (User, error)
	// Update User
	UpdateUserDetails(c context.Context, u User) (User, error)
	// Find by email
	FindUsersByEmail(c context.Context, t *Tenant, email string) ([]User, error)
}

func RegisterIdentiyServiceValidators(bs IdentityService, v *validator.Validate, l *zap.Logger, allowMultiAccount bool) error {
	bsv := UserServiceValidator{bs: bs, sv: v, logger: l, allowMultiAccount: allowMultiAccount}
	return v.RegisterValidation("isUniqueEmail", bsv.isUniqueEmail)
}

type UserServiceValidator struct {
	bs                IdentityService
	sv                *validator.Validate
	logger            *zap.Logger
	allowMultiAccount bool
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
	accountField := reflect.Indirect(p).FieldByName("Account")
	if !accountField.IsValid() {
		return false
	}

	account, ok := accountField.Interface().(*Account)
	if !ok || account == nil {
		return false
	}

	// Find in DB
	c := context.Background()
	users, err := bsv.bs.FindUsersByEmail(c, account.Tenant, email)
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
	id, ok := idField.Interface().(null.Int)
	if !ok {
		return false
	}

	// Check existing records
	for _, user := range users {
		// Check if email already in use by a different account, within the same tenant
		if account.ID.Int64 != user.Account.ID.Int64 && false == bsv.allowMultiAccount {
			return false
		}

		// Check if email is used by a different user, within the same account
		if id.Int64 != user.ID.Int64 {
			return false
		}
	}

	return true
}

// TODO: Nil relation (isUniqueEmail accept allowMultiAccount - if )
