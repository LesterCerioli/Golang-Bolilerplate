// Code generated by SQLBoiler 4.1.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testOffices(t *testing.T) {
	t.Parallel()

	query := Offices()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOfficesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Offices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOfficesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Offices().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Offices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOfficesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OfficeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Offices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOfficesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OfficeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Office exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OfficeExists to return true, but got false.")
	}
}

func testOfficesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	officeFound, err := FindOffice(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if officeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOfficesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Offices().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOfficesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Offices().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOfficesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	officeOne := &Office{}
	officeTwo := &Office{}
	if err = randomize.Struct(seed, officeOne, officeDBTypes, false, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}
	if err = randomize.Struct(seed, officeTwo, officeDBTypes, false, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = officeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = officeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Offices().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOfficesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	officeOne := &Office{}
	officeTwo := &Office{}
	if err = randomize.Struct(seed, officeOne, officeDBTypes, false, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}
	if err = randomize.Struct(seed, officeTwo, officeDBTypes, false, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = officeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = officeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Offices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func officeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Office) error {
	*o = Office{}
	return nil
}

func officeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Office) error {
	*o = Office{}
	return nil
}

func officeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Office) error {
	*o = Office{}
	return nil
}

func officeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Office) error {
	*o = Office{}
	return nil
}

func officeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Office) error {
	*o = Office{}
	return nil
}

func officeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Office) error {
	*o = Office{}
	return nil
}

func officeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Office) error {
	*o = Office{}
	return nil
}

func officeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Office) error {
	*o = Office{}
	return nil
}

func officeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Office) error {
	*o = Office{}
	return nil
}

func testOfficesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Office{}
	o := &Office{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, officeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Office object: %s", err)
	}

	AddOfficeHook(boil.BeforeInsertHook, officeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	officeBeforeInsertHooks = []OfficeHook{}

	AddOfficeHook(boil.AfterInsertHook, officeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	officeAfterInsertHooks = []OfficeHook{}

	AddOfficeHook(boil.AfterSelectHook, officeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	officeAfterSelectHooks = []OfficeHook{}

	AddOfficeHook(boil.BeforeUpdateHook, officeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	officeBeforeUpdateHooks = []OfficeHook{}

	AddOfficeHook(boil.AfterUpdateHook, officeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	officeAfterUpdateHooks = []OfficeHook{}

	AddOfficeHook(boil.BeforeDeleteHook, officeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	officeBeforeDeleteHooks = []OfficeHook{}

	AddOfficeHook(boil.AfterDeleteHook, officeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	officeAfterDeleteHooks = []OfficeHook{}

	AddOfficeHook(boil.BeforeUpsertHook, officeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	officeBeforeUpsertHooks = []OfficeHook{}

	AddOfficeHook(boil.AfterUpsertHook, officeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	officeAfterUpsertHooks = []OfficeHook{}
}

func testOfficesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Offices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOfficesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(officeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Offices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOfficeToOneBusinessUsingBusiness(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Office
	var foreign Business

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, businessDBTypes, false, businessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Business struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.BusinessID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Business().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := OfficeSlice{&local}
	if err = local.L.LoadBusiness(ctx, tx, false, (*[]*Office)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Business == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Business = nil
	if err = local.L.LoadBusiness(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Business == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOfficeToOneSetOpBusinessUsingBusiness(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Office
	var b, c Business

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, officeDBTypes, false, strmangle.SetComplement(officePrimaryKeyColumns, officeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, businessDBTypes, false, strmangle.SetComplement(businessPrimaryKeyColumns, businessColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, businessDBTypes, false, strmangle.SetComplement(businessPrimaryKeyColumns, businessColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Business{&b, &c} {
		err = a.SetBusiness(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Business != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Offices[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.BusinessID, x.ID) {
			t.Error("foreign key was wrong value", a.BusinessID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BusinessID))
		reflect.Indirect(reflect.ValueOf(&a.BusinessID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.BusinessID, x.ID) {
			t.Error("foreign key was wrong value", a.BusinessID, x.ID)
		}
	}
}

func testOfficeToOneRemoveOpBusinessUsingBusiness(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Office
	var b Business

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, officeDBTypes, false, strmangle.SetComplement(officePrimaryKeyColumns, officeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, businessDBTypes, false, strmangle.SetComplement(businessPrimaryKeyColumns, businessColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetBusiness(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveBusiness(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Business().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Business != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.BusinessID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Offices) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testOfficesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOfficesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OfficeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOfficesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Offices().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	officeDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`, `Zipcode`: `varchar`, `State`: `varchar`, `City`: `varchar`, `Address`: `varchar`, `Country`: `varchar`, `Phone`: `varchar`, `Email`: `varchar`, `BusinessID`: `int`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`}
	_             = bytes.MinRead
)

func testOfficesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(officePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(officeAllColumns) == len(officePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Offices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, officeDBTypes, true, officePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOfficesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(officeAllColumns) == len(officePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Office{}
	if err = randomize.Struct(seed, o, officeDBTypes, true, officeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Offices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, officeDBTypes, true, officePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(officeAllColumns, officePrimaryKeyColumns) {
		fields = officeAllColumns
	} else {
		fields = strmangle.SetComplement(
			officeAllColumns,
			officePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := OfficeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOfficesUpsert(t *testing.T) {
	t.Parallel()

	if len(officeAllColumns) == len(officePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLOfficeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Office{}
	if err = randomize.Struct(seed, &o, officeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Office: %s", err)
	}

	count, err := Offices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, officeDBTypes, false, officePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Office struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Office: %s", err)
	}

	count, err = Offices().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
