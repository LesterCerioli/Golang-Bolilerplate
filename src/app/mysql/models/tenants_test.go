// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testTenants(t *testing.T) {
	t.Parallel()

	query := Tenants()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTenantsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
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

	count, err := Tenants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTenantsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Tenants().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tenants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTenantsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TenantSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tenants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTenantsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TenantExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Tenant exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TenantExists to return true, but got false.")
	}
}

func testTenantsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	tenantFound, err := FindTenant(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if tenantFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTenantsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Tenants().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTenantsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Tenants().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTenantsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tenantOne := &Tenant{}
	tenantTwo := &Tenant{}
	if err = randomize.Struct(seed, tenantOne, tenantDBTypes, false, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}
	if err = randomize.Struct(seed, tenantTwo, tenantDBTypes, false, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tenantOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tenantTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Tenants().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTenantsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tenantOne := &Tenant{}
	tenantTwo := &Tenant{}
	if err = randomize.Struct(seed, tenantOne, tenantDBTypes, false, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}
	if err = randomize.Struct(seed, tenantTwo, tenantDBTypes, false, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tenantOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tenantTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tenants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func tenantBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Tenant) error {
	*o = Tenant{}
	return nil
}

func tenantAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Tenant) error {
	*o = Tenant{}
	return nil
}

func tenantAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Tenant) error {
	*o = Tenant{}
	return nil
}

func tenantBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Tenant) error {
	*o = Tenant{}
	return nil
}

func tenantAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Tenant) error {
	*o = Tenant{}
	return nil
}

func tenantBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Tenant) error {
	*o = Tenant{}
	return nil
}

func tenantAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Tenant) error {
	*o = Tenant{}
	return nil
}

func tenantBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Tenant) error {
	*o = Tenant{}
	return nil
}

func tenantAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Tenant) error {
	*o = Tenant{}
	return nil
}

func testTenantsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Tenant{}
	o := &Tenant{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, tenantDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Tenant object: %s", err)
	}

	AddTenantHook(boil.BeforeInsertHook, tenantBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	tenantBeforeInsertHooks = []TenantHook{}

	AddTenantHook(boil.AfterInsertHook, tenantAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	tenantAfterInsertHooks = []TenantHook{}

	AddTenantHook(boil.AfterSelectHook, tenantAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	tenantAfterSelectHooks = []TenantHook{}

	AddTenantHook(boil.BeforeUpdateHook, tenantBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	tenantBeforeUpdateHooks = []TenantHook{}

	AddTenantHook(boil.AfterUpdateHook, tenantAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	tenantAfterUpdateHooks = []TenantHook{}

	AddTenantHook(boil.BeforeDeleteHook, tenantBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	tenantBeforeDeleteHooks = []TenantHook{}

	AddTenantHook(boil.AfterDeleteHook, tenantAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	tenantAfterDeleteHooks = []TenantHook{}

	AddTenantHook(boil.BeforeUpsertHook, tenantBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	tenantBeforeUpsertHooks = []TenantHook{}

	AddTenantHook(boil.AfterUpsertHook, tenantAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	tenantAfterUpsertHooks = []TenantHook{}
}

func testTenantsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tenants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTenantsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(tenantColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Tenants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTenantToManyAccounts(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Tenant
	var b, c Account

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.TenantID, a.ID)
	queries.Assign(&c.TenantID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Accounts().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.TenantID, b.TenantID) {
			bFound = true
		}
		if queries.Equal(v.TenantID, c.TenantID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TenantSlice{&a}
	if err = a.L.LoadAccounts(ctx, tx, false, (*[]*Tenant)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Accounts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Accounts = nil
	if err = a.L.LoadAccounts(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Accounts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testTenantToManyUsers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Tenant
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.TenantID, a.ID)
	queries.Assign(&c.TenantID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Users().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.TenantID, b.TenantID) {
			bFound = true
		}
		if queries.Equal(v.TenantID, c.TenantID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TenantSlice{&a}
	if err = a.L.LoadUsers(ctx, tx, false, (*[]*Tenant)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Users = nil
	if err = a.L.LoadUsers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testTenantToManyAddOpAccounts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Tenant
	var b, c, d, e Account

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tenantDBTypes, false, strmangle.SetComplement(tenantPrimaryKeyColumns, tenantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Account{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Account{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAccounts(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.TenantID) {
			t.Error("foreign key was wrong value", a.ID, first.TenantID)
		}
		if !queries.Equal(a.ID, second.TenantID) {
			t.Error("foreign key was wrong value", a.ID, second.TenantID)
		}

		if first.R.Tenant != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Tenant != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Accounts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Accounts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Accounts().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testTenantToManySetOpAccounts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Tenant
	var b, c, d, e Account

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tenantDBTypes, false, strmangle.SetComplement(tenantPrimaryKeyColumns, tenantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Account{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetAccounts(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Accounts().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetAccounts(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Accounts().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.TenantID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.TenantID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.TenantID) {
		t.Error("foreign key was wrong value", a.ID, d.TenantID)
	}
	if !queries.Equal(a.ID, e.TenantID) {
		t.Error("foreign key was wrong value", a.ID, e.TenantID)
	}

	if b.R.Tenant != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Tenant != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Tenant != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Tenant != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Accounts[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Accounts[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testTenantToManyRemoveOpAccounts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Tenant
	var b, c, d, e Account

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tenantDBTypes, false, strmangle.SetComplement(tenantPrimaryKeyColumns, tenantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Account{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddAccounts(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Accounts().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveAccounts(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Accounts().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.TenantID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.TenantID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Tenant != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Tenant != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Tenant != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Tenant != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Accounts) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Accounts[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Accounts[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testTenantToManyAddOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Tenant
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tenantDBTypes, false, strmangle.SetComplement(tenantPrimaryKeyColumns, tenantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*User{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUsers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.TenantID) {
			t.Error("foreign key was wrong value", a.ID, first.TenantID)
		}
		if !queries.Equal(a.ID, second.TenantID) {
			t.Error("foreign key was wrong value", a.ID, second.TenantID)
		}

		if first.R.Tenant != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Tenant != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Users[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Users[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Users().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testTenantToManySetOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Tenant
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tenantDBTypes, false, strmangle.SetComplement(tenantPrimaryKeyColumns, tenantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetUsers(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetUsers(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.TenantID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.TenantID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.TenantID) {
		t.Error("foreign key was wrong value", a.ID, d.TenantID)
	}
	if !queries.Equal(a.ID, e.TenantID) {
		t.Error("foreign key was wrong value", a.ID, e.TenantID)
	}

	if b.R.Tenant != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Tenant != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Tenant != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Tenant != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Users[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Users[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testTenantToManyRemoveOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Tenant
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tenantDBTypes, false, strmangle.SetComplement(tenantPrimaryKeyColumns, tenantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddUsers(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveUsers(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.TenantID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.TenantID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Tenant != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Tenant != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Tenant != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Tenant != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Users) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Users[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Users[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testTenantsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
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

func testTenantsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TenantSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTenantsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Tenants().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	tenantDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`}
	_             = bytes.MinRead
)

func testTenantsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(tenantPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(tenantAllColumns) == len(tenantPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tenants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTenantsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(tenantAllColumns) == len(tenantPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Tenant{}
	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tenants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tenantDBTypes, true, tenantPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(tenantAllColumns, tenantPrimaryKeyColumns) {
		fields = tenantAllColumns
	} else {
		fields = strmangle.SetComplement(
			tenantAllColumns,
			tenantPrimaryKeyColumns,
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

	slice := TenantSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTenantsUpsert(t *testing.T) {
	t.Parallel()

	if len(tenantAllColumns) == len(tenantPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLTenantUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Tenant{}
	if err = randomize.Struct(seed, &o, tenantDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Tenant: %s", err)
	}

	count, err := Tenants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, tenantDBTypes, false, tenantPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Tenant: %s", err)
	}

	count, err = Tenants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
