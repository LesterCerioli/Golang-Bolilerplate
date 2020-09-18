package mysql

import (
	"context"
	"database/sql"

	"github.com/eldad87/go-boilerplate/src/app"
	"github.com/eldad87/go-boilerplate/src/app/mysql/models"
	"github.com/eldad87/go-boilerplate/src/pkg/crypto"
	"github.com/eldad87/go-boilerplate/src/pkg/validator"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
	gnull "gopkg.in/guregu/null.v4"
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
	currentTenant, err := idt.findTenant(c, uint(t.ID.Int64))
	if err != nil {
		return app.Tenant{}, err
	}

	// Set
	if t.Name.Valid {
		currentTenant.Name = t.Name.String
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
	a.TenantID.SetValid(uint(tenant.ID.Int64))
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
	currentAccount, err := idt.findAccount(c, uint(a.ID.Int64))
	if err != nil {
		return app.Account{}, err
	}

	// Set
	if a.Name.Valid {
		currentAccount.Name = a.Name.String
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

	// Password
	hash, err := idt.passHandler.Generate(user.Password.String)
	if err != nil {
		return app.User{}, err
	}

	// Insert
	u := idt.userToBoiler(user)
	u.Password.Scan(hash)
	u.TenantID.SetValid(uint(account.Tenant.ID.Int64))
	u.AccountID.SetValid(uint(account.ID.Int64))
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
func (idt *identityService) FindUsersByEmail(c context.Context, t Tenant, email string) ([]app.User, error) {
	whereEmail := models.UserWhere.Email.EQ(null.StringFrom(email))
	whereTenant := models.UserWhere.TenantID.EQ(null.UintFrom(t.ID))

	eagerAccount := qm.Load(models.UserRels.Account)
	eagerTenant := qm.Load(models.AccountRels.Tenant)

	users, err := models.Users(whereTenant, whereEmail, eagerAccount, eagerTenant).All(c, idt.db)

	if err != nil {
		if err == sql.ErrNoRows {
			return []app.User{}, app.ErrNotFound
		}
		return []app.User{}, err
	}

	var response []app.User
	for _, user := range users {
		u := idt.userFromBoiler(*user)
		a := idt.accountFromBoiler(*user.R.Account)
		t := idt.tenantFromBoiler(*user.R.Tenant)
		u.Account = &a
		u.Account.Tenant = &t
		response = append(response, u)
	}

	return response, nil
}

func (idt *identityService) UpdateUserDetails(c context.Context, u app.User) (app.User, error) {
	// Validate
	err := idt.sv.StructCtx(c, u)
	if err != nil {
		return app.User{}, err
	}

	// Find
	currentUser, err := idt.findUser(c, uint(u.ID.Int64))
	if err != nil {
		return app.User{}, err
	}

	// Set
	if u.FirstName.Valid {
		currentUser.FirstName = u.FirstName.String
	}
	currentUser.LastName = null.NewString(u.LastName.String, u.LastName.Valid)
	currentUser.Email = null.NewString(u.Email.String, u.Email.Valid)
	currentUser.Active = null.NewBool(u.Active.Bool, u.Active.Valid)

	// Password
	if u.Password.Valid {
		hash, err := idt.passHandler.Generate(u.Password.String)
		if err != nil {
			return app.User{}, err
		}
		currentUser.Password.Scan(hash)
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
func (idt *identityService) tenantFromBoiler(t models.Tenant) app.Tenant {
	return app.Tenant{
		ID:        gnull.IntFrom(int64(t.ID)),
		Name:      gnull.StringFrom(t.Name),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func (idt *identityService) tenantToBoiler(t app.Tenant) models.Tenant {
	return models.Tenant{
		ID:        uint(t.ID.Int64),
		Name:      t.Name.String,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func (idt *identityService) accountFromBoiler(a models.Account) app.Account {
	t := idt.tenantFromBoiler(*a.R.Tenant)
	return app.Account{
		ID:        gnull.IntFrom(int64(a.ID)),
		Name:      gnull.StringFrom(a.Name),
		Tenant:    &t,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func (idt *identityService) accountToBoiler(a app.Account) models.Account {
	return models.Account{
		ID:        uint(a.ID.Int64),
		Name:      a.Name.String,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func (idt *identityService) userFromBoiler(u models.User) app.User {
	a := idt.accountFromBoiler(*u.R.Account)
	return app.User{
		ID:        gnull.IntFrom(int64(u.ID)),
		FirstName: gnull.StringFrom(u.FirstName),
		LastName:  gnull.StringFromPtr(u.LastName.Ptr()),
		Email:     gnull.StringFromPtr(u.Email.Ptr()),
		Password:  gnull.StringFromPtr(u.Password.Ptr()),
		Active:    gnull.BoolFromPtr(u.Active.Ptr()),
		Account:   &a,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (idt *identityService) userToBoiler(u app.User) models.User {
	return models.User{
		ID:        uint(u.ID.Int64),
		FirstName: u.FirstName.String,
		LastName:  null.StringFromPtr(u.LastName.Ptr()),
		Email:     null.StringFromPtr(u.Email.Ptr()),
		Password:  null.StringFromPtr(u.Password.Ptr()),
		Active:    null.BoolFromPtr(u.Active.Ptr()),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
