package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/eldad87/go-boilerplate/src/app"
	"github.com/eldad87/go-boilerplate/src/app/mysql/models"
	"github.com/eldad87/go-boilerplate/src/pkg/crypto"
	"github.com/eldad87/go-boilerplate/src/pkg/validator"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

func NewIdentityService(db *sql.DB, sv validator.StructValidator, ph crypto.Hash, l *zap.Logger) *identityService {
	return &identityService{db, sv, ph, l}
}

type identityService struct {
	db          *sql.DB
	sv          validator.StructValidator
	passHandler crypto.Hash
	logger      *zap.Logger
}

func (idt *identityService) CreateTenant(c context.Context, t app.Tenant, a app.Account, u app.User) (app.User, error) {
	tx, err := idt.db.BeginTx(c, nil)
	if err != nil {
		return app.User{}, err
	}

	tenant, errr := idt.createTenant(c, t)
	if errr != nil {
		tx.Rollback()
		return app.User{}, err
	}

	user, errrr := idt.CreateAccount(c, tenant, a, u)
	if errrr != nil {
		tx.Rollback()
		return app.User{}, err
	}

	tx.Commit()
	return user, nil
}

func (idt *identityService) CreateAccount(c context.Context, t app.Tenant, a app.Account, u app.User) (app.User, error) {
	tx, err := idt.db.BeginTx(c, nil)
	if err != nil {
		return app.User{}, err
	}

	account, err := idt.createAccount(c, a, t)
	if err != nil {
		tx.Rollback()
		return app.User{}, err
	}

	user, err := idt.CreateUser(c, u, account)
	if err != nil {
		tx.Rollback()
		return app.User{}, err
	}

	tx.Commit()
	return user, nil
}

//***************** Tenant
func (idt *identityService) createTenant(c context.Context, tenant app.Tenant) (app.Tenant, error) {
	// Validate
	err := idt.sv.StructCtx(c, tenant)
	if err != nil {
		return app.Tenant{}, err
	}

	// Start transaction
	tx, err := idt.db.BeginTx(c, nil)
	if err != nil {
		return app.Tenant{}, err
	}

	// Insert
	t := idt.tenantToBoiler(tenant)
	err = t.Insert(c, idt.db, boil.Infer())
	if err != nil {
		tx.Rollback()
		return app.Tenant{}, err
	}

	tx.Commit()

	return idt.tenantFromBoiler(t), nil
}

func (idt *identityService) GetTenant(c context.Context, id uint) (app.Tenant, error) {
	tenant, err := idt.findTenant(c, id)
	if err != nil {
		return app.Tenant{}, err
	}

	t := idt.tenantFromBoiler(tenant)
	return t, nil
}

func (idt *identityService) UpdateTenantDetails(c context.Context, t app.Tenant) (app.Tenant, error) {
	// Validate
	err := idt.sv.StructCtx(c, t)
	if err != nil {
		return app.Tenant{}, err
	}

	// Find
	currentTenant, err := idt.findTenant(c, *t.ID)
	if err != nil {
		return app.Tenant{}, err
	}

	// Set
	if t.Name != nil {
		currentTenant.Name = *t.Name
	}

	// Save
	_, err = currentTenant.Update(c, idt.db, boil.Infer())
	if err != nil {
		return app.Tenant{}, err
	}

	return idt.tenantFromBoiler(currentTenant), nil
}

//***************** Account
func (idt *identityService) createAccount(c context.Context, account app.Account, tenant app.Tenant) (app.Account, error) {
	// Validate
	err := idt.sv.StructCtx(c, account)
	if err != nil {
		return app.Account{}, err
	}

	// Start transaction
	tx, err := idt.db.BeginTx(c, nil)
	if err != nil {
		return app.Account{}, err
	}

	// Insert
	a := idt.accountToBoiler(account)
	a.TenantID.Scan(tenant.ID)
	err = a.Insert(c, idt.db, boil.Infer())
	if err != nil {
		tx.Rollback()
		return app.Account{}, err
	}

	tx.Commit()

	return idt.accountFromBoiler(a), nil
}

func (idt *identityService) GetAccount(c context.Context, id uint) (app.Account, error) {
	account, err := idt.findAccount(c, id)
	if err != nil {
		return app.Account{}, err
	}

	a := idt.accountFromBoiler(account)
	return a, nil
}

func (idt *identityService) UpdateAccountDetails(c context.Context, a app.Account) (app.Account, error) {
	// Validate
	err := idt.sv.StructCtx(c, a)
	if err != nil {
		return app.Account{}, err
	}

	// Find
	currentAccount, err := idt.findAccount(c, *a.ID)
	if err != nil {
		return app.Account{}, err
	}

	// Set
	if a.Name != nil {
		currentAccount.Name = *a.Name
	}

	// Save
	_, err = currentAccount.Update(c, idt.db, boil.Infer())
	if err != nil {
		return app.Account{}, err
	}

	return idt.accountFromBoiler(currentAccount), nil
}

//***************** User
func (idt *identityService) CreateUser(c context.Context, user app.User, account app.Account) (app.User, error) {
	// Validate
	err := idt.sv.StructCtx(c, user)
	if err != nil {
		return app.User{}, err
	}

	// Start transaction
	tx, err := idt.db.BeginTx(c, nil)
	if err != nil {
		return app.User{}, err
	}

	// Tenant
	if account.Tenant == nil {
		// TODO Translation, i18n, localization
		// This shouldn't happen!
		return app.User{}, errors.New("Invalid account's tenant")
	}

	// Password
	if user.Password == nil {
		// TODO Translation, i18n, localization
		// This shouldn't happen!
		return app.User{}, errors.New("Invalid password")
	}
	hash, err := idt.passHandler.Generate(*user.Password)
	if err != nil {
		return app.User{}, err
	}

	// Insert
	u := idt.userToBoiler(user)
	u.Password.Scan(hash)
	u.TenantID.Scan(account.Tenant.ID)
	u.AccountID.Scan(account.ID)
	err = u.Insert(c, idt.db, boil.Infer())
	if err != nil {
		tx.Rollback()
		return app.User{}, err
	}

	tx.Commit()

	return idt.userFromBoiler(u), nil
}
func (idt *identityService) GetUser(c context.Context, id uint) (app.User, error) {
	user, err := idt.findUser(c, id)
	if err != nil {
		return app.User{}, err
	}

	u := idt.userFromBoiler(user)
	return u, nil
}
func (idt *identityService) FindUserByEmail(c context.Context, t Tenant, email string) (app.User, error) {
	whereEmail := models.UserWhere.Email.EQ(null.StringFrom(email))
	whereTenant := models.UserWhere.TenantID.EQ(null.UintFrom(t.ID))

	eagerAccount := qm.Load(models.UserRels.Account)
	eagerTenant := qm.Load(models.UserRels.Account)

	user, err := models.Users(whereTenant, whereEmail, eagerAccount, eagerTenant).One(c, idt.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return app.User{}, app.ErrNotFound
		}
		return app.User{}, err
	}

	u := idt.userFromBoiler(*user)
	acc := idt.accountFromBoiler(*user.R.Account)
	tenant := idt.tenantFromBoiler(*user.R.Tenant)
	u.Account = &acc
	u.Account.Tenant = &tenant
	return u, nil
}

func (idt *identityService) UpdateUserDetails(c context.Context, u app.User) (app.User, error) {
	// Validate
	err := idt.sv.StructCtx(c, u)
	if err != nil {
		return app.User{}, err
	}

	// Find
	currentUser, err := idt.findUser(c, *u.ID)
	if err != nil {
		return app.User{}, err
	}

	// Set
	currentUser.FirstName = null.StringFromPtr(u.FirstName).String
	currentUser.LastName = null.StringFromPtr(u.LastName)
	currentUser.Email = null.StringFromPtr(u.Email)
	currentUser.Active = null.BoolFromPtr(u.Active)

	// Password
	if u.Password != nil {
		hash, err := idt.passHandler.Generate(*u.Password)
		if err != nil {
			return app.User{}, err
		}
		currentUser.Password = null.StringFrom(hash)
	}

	// Save
	_, err = currentUser.Update(c, idt.db, boil.Infer())
	if err != nil {
		return app.User{}, err
	}

	return idt.userFromBoiler(currentUser), nil
}

//***************** Find helpers
func (idt *identityService) findTenant(c context.Context, id uint) (models.Tenant, error) {
	tenant, err := models.FindTenant(c, idt.db, id)

	// No record found
	if err == sql.ErrNoRows {
		return models.Tenant{}, app.ErrNotFound
	} else if err != nil {
		return models.Tenant{}, err
	}

	return *tenant, nil
}

func (idt *identityService) findAccount(c context.Context, id uint) (models.Account, error) {
	whereUser := models.UserWhere.ID.EQ(id)
	eagerTenant := qm.Load(models.AccountRels.Tenant)

	account, err := models.Accounts(whereUser, eagerTenant).One(c, idt.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Account{}, app.ErrNotFound
		}
		return models.Account{}, err
	}

	return *account, nil
}

func (idt *identityService) findUser(c context.Context, id uint) (models.User, error) {
	whereID := models.UserWhere.ID.EQ(id)
	eagerAccount := qm.Load(models.UserRels.Account)
	eagerTenant := qm.Load(models.AccountRels.Tenant)

	user, err := models.Users(whereID, eagerAccount, eagerTenant).One(c, idt.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, app.ErrNotFound
		}
		return models.User{}, err
	}

	return *user, nil
}

//***************** Type conversion helpers
func (idt *identityService) tenantFromBoiler(a models.Tenant) app.Tenant {
	return app.Tenant{
		ID:        &a.ID,
		Name:      &a.Name,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func (idt *identityService) tenantToBoiler(a app.Tenant) models.Tenant {
	return models.Tenant{
		ID:        *a.ID,
		Name:      *a.Name,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func (idt *identityService) accountFromBoiler(a models.Account) app.Account {
	var tenant *app.Tenant
	if a.R != nil && a.R.Tenant != nil {
		t := idt.tenantFromBoiler(*a.R.Tenant)
		tenant = &t
	}

	return app.Account{
		ID:        &a.ID,
		Name:      &a.Name,
		Tenant:    tenant,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func (idt *identityService) accountToBoiler(a app.Account) models.Account {
	/*	var t app.Tenant
		if a.Tenant != nil {
			t = *a.Tenant
		}*/
	return models.Account{
		ID:   null.UintFromPtr(a.ID).Uint,
		Name: null.StringFromPtr(a.Name).String,
		//		TenantID: 	null.UintFromPtr(t.ID),
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func (idt *identityService) userFromBoiler(u models.User) app.User {
	var acc *app.Account
	if u.R != nil && u.R.Account != nil {
		a := idt.accountFromBoiler(*u.R.Account)
		acc = &a
	}

	return app.User{
		ID:        &u.ID,
		FirstName: &u.FirstName,
		LastName:  u.LastName.Ptr(),
		Email:     u.Email.Ptr(),
		Password:  u.Password.Ptr(),
		Active:    u.Active.Ptr(),
		Account:   acc,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (idt *identityService) userToBoiler(u app.User) models.User {
	return models.User{
		ID:        null.UintFromPtr(u.ID).Uint,
		FirstName: null.StringFromPtr(u.FirstName).String,
		LastName:  null.StringFromPtr(u.LastName),
		Email:     null.StringFromPtr(u.Email),
		Password:  null.StringFromPtr(u.Password),
		Active:    null.BoolFromPtr(u.Active),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
