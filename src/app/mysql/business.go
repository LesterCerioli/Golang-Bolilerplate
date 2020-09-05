package mysql

import (
	"context"
	"database/sql"
	"github.com/eldad87/go-boilerplate/src/app"
	"github.com/eldad87/go-boilerplate/src/app/mysql/models"
	"github.com/eldad87/go-boilerplate/src/pkg/validator"
	"github.com/volatiletech/sqlboiler/v4/boil"
	//	"github.com/jinzhu/copier"
	"encoding/json"
	"github.com/eldad87/go-boilerplate/src/pkg/crypto"
	vnull "github.com/volatiletech/null/v8"
	"go.uber.org/zap"
)

func NewBusinessService(db *sql.DB, sv validator.StructValidator, ph crypto.Hash, l *zap.Logger) *BusinessService {
	return &BusinessService{db, sv, ph, l}
}

type BusinessService struct {
	db          *sql.DB
	sv          validator.StructValidator
	passHandler crypto.Hash
	logger      *zap.Logger
}

func (bs *BusinessService) Get(c context.Context, id *uint) (*app.Business, error) {
	bBusiness, err := models.FindBusiness(c, bs.db, *id)

	// No record found
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return bs.businessFromBoiler(bBusiness)
}

func (bs *BusinessService) Create(c context.Context, b *app.Business) (*app.Business, error) {

	// Validate
	err := bs.sv.StructCtx(c, b)
	if err != nil {
		return nil, err
	}

	tx, err := bs.db.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}

	// Case to Boiler object - business
	b.ID = 0 // Create! not update.
	bBusiness, err := bs.businessToBoiler(b)
	if err != nil {
		return nil, err
	}
	// Insert
	bBusiness.Insert(c, tx, boil.Infer())

	// Cast to Boiler object - offices
	bOffices, err := bs.officesToBoiler(b.Offices)
	if err != nil {
		bs.logger.Error("Failed officesToBoiler")
		tx.Rollback()
		return nil, err
	}

	err = bBusiness.AddOffices(c, tx, true, bOffices...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Cast to Boiler object - employees
	bEmployees, err := bs.employeesToBoiler(b.Employees)
	if err != nil {
		bs.logger.Error("Failed employeesToBoiler")
		tx.Rollback()
		return nil, err
	}

	err = bBusiness.AddEmployees(c, tx, true, bEmployees...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	//Response
	b, err = bs.businessFromBoiler(bBusiness)
	if err != nil {
		return nil, err
	}

	// Add Offices
	offices, err := bs.officesFromBoiler(bOffices)
	if err != nil {
		return nil, err
	}
	b.Offices = offices

	// Add Employees
	employees, err := bs.employeesFromBoiler(bEmployees)
	if err != nil {
		return nil, err
	}
	b.Employees = employees

	return b, err
}

func (bs *BusinessService) Update(c context.Context, b *app.Business) (*app.Business, error) {
	// Validate
	err := bs.sv.StructCtx(c, b)
	if err != nil {
		return nil, err
	}

	// Convert
	bBusiness, err := models.FindBusiness(c, bs.db, b.ID)
	//bBusiness, err := bs.businessToBoiler(b)
	if err != nil {
		return nil, err
	}
	bBusiness.ID = b.ID
	bBusiness.Name = b.Name
	bBusiness.Website.Scan(b.Website)

	// Update
	_, err = bBusiness.Update(c, bs.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	// Reload
	err = bBusiness.Reload(c, bs.db)
	if err != nil {
		return nil, err
	}

	return bs.businessFromBoiler(bBusiness)
}

func (bs *BusinessService) GetEmployeeByEmail(c context.Context, email string) (*app.Employee, error) {
	sbEmployee, err := models.Employees(models.EmployeeWhere.Email.EQ(vnull.StringFrom(email))).One(c, bs.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return bs.employeeFromBoiler(sbEmployee)
}

func (bs *BusinessService) officesFromBoiler(sbOffices []*models.Office) ([]*app.Office, error) {
	offices := make([]*app.Office, len(sbOffices))

	err := marshalUn(sbOffices, &offices)
	if err != nil {
		return nil, err
	}
	return offices, nil
}
func (bs *BusinessService) employeesFromBoiler(sbEmployees []*models.Employee) ([]*app.Employee, error) {
	offices := make([]*app.Employee, len(sbEmployees))

	err := marshalUn(sbEmployees, &offices)
	if err != nil {
		return nil, err
	}
	return offices, nil
}
func (bs *BusinessService) employeeFromBoiler(sbEmployee *models.Employee) (*app.Employee, error) {
	employee := &app.Employee{}
	err := marshalUn(sbEmployee, employee)
	if err != nil {
		return nil, err
	}

	return employee, nil
}
func (bs *BusinessService) officesToBoiler(offices []*app.Office) ([]*models.Office, error) {
	sbOffices := make([]*models.Office, len(offices))

	err := marshalUn(offices, &sbOffices)
	if err != nil {
		return nil, err
	}
	return sbOffices, nil
}

func (bs *BusinessService) employeesToBoiler(employees []*app.Employee) ([]*models.Employee, error) {
	sbEmployees := make([]*models.Employee, len(employees))

	err := marshalUn(employees, &sbEmployees)
	if err != nil {
		return nil, err
	}

	// Encrypt password
	for idx, _ := range sbEmployees {
		hash, err := bs.passHandler.Generate(sbEmployees[idx].Password.String)
		if err != nil {
			return nil, err
		}
		sbEmployees[idx].Password.Scan(hash)
	}

	return sbEmployees, nil
}

func (bs *BusinessService) businessToBoiler(b *app.Business) (*models.Business, error) {
	sBusiness := &models.Business{ID: b.ID}
	sBusiness.Name = b.Name
	sBusiness.Website.Scan(b.Website)
	return sBusiness, nil
	/*
		err := marshalUn(b, sBusiness)
		if err != nil {
			return nil, err
		}
		return sBusiness, nil
	*/
}

func (bs *BusinessService) businessFromBoiler(sBusiness *models.Business) (*app.Business, error) {
	// Marshal
	b := &app.Business{}
	err := marshalUn(sBusiness, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func marshalUn(from interface{}, to interface{}) error {
	bBytes, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(bBytes, to)
}
