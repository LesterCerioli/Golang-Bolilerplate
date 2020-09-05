package mysql

import (
	"context"
	"database/sql"

	"github.com/eldad87/go-boilerplate/src/app"
	"github.com/eldad87/go-boilerplate/src/app/mysql/models"
	"github.com/eldad87/go-boilerplate/src/pkg/validator"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"
)

func NewVisitService(db *sql.DB, sv validator.StructValidator) *visitService {
	return &visitService{db, sv}
}

type Tenant struct {
	ID        uint       `json:"id" validate:"omitempty,gte=0"`
	Name      string     `validate:"required_without=ID"`
	Accounts  []*Account `json:"accounts" validate:"omitempty,dive,required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
type visitService struct {
	db *sql.DB
	sv validator.StructValidator
}

func (vs *visitService) Get(c context.Context, id *uint) (*app.Visit, error) {
	bVisit, err := models.FindVisit(c, vs.db, *id)

	// No record found
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return sqlBoilerToVisit(bVisit), nil
}

func (vs *visitService) Set(c context.Context, v *app.Visit) (*app.Visit, error) {
	bVisit := models.Visit{
		ID:        v.ID,
		FirstName: null.StringFrom(v.FirstName),
		LastName:  null.StringFrom(v.LastName),
	}

	err := vs.sv.StructCtx(c, v)
	if err != nil {
		return nil, err
	}

	if bVisit.ID == 0 {
		err = bVisit.Insert(c, vs.db, boil.Infer())
	} else {
		_, err = bVisit.Update(c, vs.db, boil.Infer())
	}

	if err != nil {
		return nil, err
	}

	return sqlBoilerToVisit(&bVisit), nil
}

func sqlBoilerToVisit(bVisit *models.Visit) *app.Visit {
	return &app.Visit{
		ID:        bVisit.ID,
		FirstName: bVisit.FirstName.String,
		LastName:  bVisit.LastName.String,
		CreatedAt: bVisit.CreatedAt,
		UpdatedAt: bVisit.UpdatedAt,
	}
}
