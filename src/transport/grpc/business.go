package grpc

import (
	"context"
	"github.com/eldad87/go-boilerplate/src/app"
	"github.com/eldad87/go-boilerplate/src/pkg/grpc/wrappers/guregu/null-v4"
	"github.com/eldad87/go-boilerplate/src/transport/grpc/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type BusinessServer struct {
	BusinessService app.BusinessService
	Logger          *zap.Logger
}

func (bs *BusinessServer) Get(c context.Context, id *pb.ID) (*pb.BusinessResponse, error) {
	i := uint(id.GetID())
	b, err := bs.BusinessService.Get(c, &i)
	if err != nil {
		return nil, err
	} else if b == nil {
		return &pb.BusinessResponse{}, nil
	}

	res := &pb.BusinessResponse{}
	res.ID = uint32(b.ID)
	res.Website = null_v4.NullToStringValue(&b.Website)
	if t, err := ptypes.TimestampProto(b.CreatedAt); err == nil {
		res.CreatedAt = t
	}
	if t, err := ptypes.TimestampProto(b.UpdatedAt); err == nil {
		res.UpdatedAt = t
	}

	return res, nil
}

// Create a business
func (bs *BusinessServer) Create(c context.Context, bcr *pb.BusinessCreateRequest) (*pb.BusinessCreateResponse, error) {
	// Business
	b := &app.Business{Name: bcr.Name}
	null_v4.StringValueToNull(bcr.Website, &b.Website)
	// Add Offices
	for _, ocr := range bcr.Offices {
		o := &app.Office{Name: ocr.Name}
		null_v4.StringValueToNull(ocr.Email, &o.Email)
		null_v4.StringValueToNull(ocr.Phone, &o.Phone)
		null_v4.StringValueToNull(ocr.Country, &o.Country)
		null_v4.StringValueToNull(ocr.State, &o.State)
		null_v4.StringValueToNull(ocr.City, &o.City)
		null_v4.StringValueToNull(ocr.Address, &o.Address)
		null_v4.StringValueToNull(ocr.Zipcode, &o.Zipcode)
		b.Offices = append(b.Offices, o)
	}

	// Add Employees
	for _, ecr := range bcr.Employees {
		e := &app.Employee{FirstName: ecr.FirstName}
		e.Email.Scan(ecr.Email)
		e.Password.Scan(ecr.Password)
		null_v4.StringValueToNull(ecr.LastName, &e.LastName)
		b.Employees = append(b.Employees, e)
	}

	b, err := bs.BusinessService.Create(c, b)
	if err != nil {
		return nil, err
	}

	// Build response
	res := &pb.BusinessCreateResponse{}
	copier.Copy(&res, &b) //Copy native values & Initialize our relationships

	res.ID = uint32(b.ID)
	res.Website = null_v4.NullToStringValue(&b.Website)
	if t, err := ptypes.TimestampProto(b.CreatedAt); err == nil {
		res.CreatedAt = t
	}
	if t, err := ptypes.TimestampProto(b.UpdatedAt); err == nil {
		res.UpdatedAt = t
	}

	// Add offices
	for idx, o := range b.Offices {
		res.Offices[idx].ID = uint32(b.ID)
		res.Offices[idx].Email = null_v4.NullToStringValue(&o.Email)
		res.Offices[idx].Phone = null_v4.NullToStringValue(&o.Phone)
		res.Offices[idx].Country = null_v4.NullToStringValue(&o.Country)
		res.Offices[idx].State = null_v4.NullToStringValue(&o.State)
		res.Offices[idx].City = null_v4.NullToStringValue(&o.City)
		res.Offices[idx].Address = null_v4.NullToStringValue(&o.Address)
		res.Offices[idx].Zipcode = null_v4.NullToStringValue(&o.Zipcode)
		if t, err := ptypes.TimestampProto(o.CreatedAt); err == nil {
			res.Offices[idx].CreatedAt = t
		}
		if t, err := ptypes.TimestampProto(o.UpdatedAt); err == nil {
			res.Offices[idx].UpdatedAt = t
		}
	}

	// Add Employees
	for idx, o := range b.Employees {
		res.Employees[idx].ID = uint32(b.ID)
		res.Employees[idx].Email = o.Email.String
		res.Employees[idx].LastName = null_v4.NullToStringValue(&o.LastName)
		if t, err := ptypes.TimestampProto(o.CreatedAt); err == nil {
			res.Employees[idx].CreatedAt = t
		}
		if t, err := ptypes.TimestampProto(o.UpdatedAt); err == nil {
			res.Employees[idx].UpdatedAt = t
		}
	}

	return res, nil
}

// Update
func (bs *BusinessServer) Update(c context.Context, bur *pb.BusinessUpdateRequest) (*pb.BusinessResponse, error) {
	b := &app.Business{ID: uint(bur.ID), Name: bur.Name}
	null_v4.StringValueToNull(bur.Website, &b.Website)

	b, err := bs.BusinessService.Update(c, b)
	if err != nil {
		return nil, err
	}

	res := &pb.BusinessResponse{}
	res.ID = uint32(b.ID)
	res.Website = null_v4.NullToStringValue(&b.Website)
	if t, err := ptypes.TimestampProto(b.CreatedAt); err == nil {
		res.CreatedAt = t
	}
	if t, err := ptypes.TimestampProto(b.UpdatedAt); err == nil {
		res.UpdatedAt = t
	}
	return res, nil
}
