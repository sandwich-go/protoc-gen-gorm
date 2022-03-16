// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm v0.0.0
// 	protoc          (unknown)
// source: crud/crud.proto

package crud

import (
	context "context"
	fmt "fmt"
	_ "github.com/complex64/protoc-gen-gorm/gormpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	gorm "gorm.io/gorm"
)

// CrudModel is the GORM model for crud.Crud.
type CrudModel struct {
	Uuid        string `gorm:"primaryKey"`
	StringField string
	Int32Field  int32
	BoolField   bool `gorm:"column:enabled"`
}

// AsProto converts a CrudModel to its protobuf representation.
func (m *CrudModel) AsProto() (*Crud, error) {
	x := new(Crud)
	x.Uuid = m.Uuid
	x.StringField = m.StringField
	x.Int32Field = m.Int32Field
	x.BoolField = m.BoolField
	return x, nil
}

// AsModel converts a Crud to its GORM model.
func (x *Crud) AsModel() (*CrudModel, error) {
	m := new(CrudModel)
	m.Uuid = x.Uuid
	m.StringField = x.StringField
	m.Int32Field = x.Int32Field
	m.BoolField = x.BoolField
	return m, nil
}

type CrudGetOption func(tx *gorm.DB) *gorm.DB
type CrudListOption func(tx *gorm.DB) *gorm.DB

type CrudWithDB struct {
	x  *Crud
	db *gorm.DB
}

func (x *Crud) WithDB(db *gorm.DB) CrudWithDB {
	return CrudWithDB{x: x, db: db}
}

func (c CrudWithDB) Create(ctx context.Context) (*Crud, error) {
	if c.x == nil {
		return nil, nil
	}
	m, err := c.x.AsModel()
	if err != nil {
		return nil, err
	}
	db := c.db.WithContext(ctx)
	if err := db.Create(m).Error; err != nil {
		return nil, err
	}
	if y, err := m.AsProto(); err != nil {
		return nil, err
	} else {
		return y, nil
	}
}

func (c CrudWithDB) Get(ctx context.Context, opts ...CrudGetOption) (*Crud, error) {
	if c.x == nil {
		return nil, nil
	}
	var zero string
	if c.x.Uuid == zero {
		return nil, fmt.Errorf("empty primary key")
	}
	m, err := c.x.AsModel()
	if err != nil {
		return nil, err
	}
	db := c.db.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	out := CrudModel{}
	if err := db.Where(m).First(&out).Error; err != nil {
		return nil, err
	}
	if y, err := out.AsProto(); err != nil {
		return nil, err
	} else {
		return y, nil
	}
}

func (c CrudWithDB) List(ctx context.Context, opts ...CrudListOption) ([]*Crud, error) {
	if c.x == nil {
		return nil, nil
	}
	db := c.db.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	var ms []CrudModel
	if err := db.Find(&ms).Error; err != nil {
		return nil, err
	}
	xs := make([]*Crud, 0, len(ms))
	for _, m := range ms {
		if x, err := m.AsProto(); err != nil {
			return nil, err
		} else {
			xs = append(xs, x)
		}
	}
	return xs, nil
}

func (c CrudWithDB) Update(ctx context.Context) (*Crud, error) {
	if c.x == nil {
		return nil, nil
	}
	m, err := c.x.AsModel()
	if err != nil {
		return nil, err
	}
	db := c.db.WithContext(ctx)
	if err := db.Save(m).Error; err != nil {
		return nil, err
	}
	return c.Get(ctx)
}

func (c CrudWithDB) Patch(ctx context.Context, mask *fieldmaskpb.FieldMask) error {
	if c.x == nil {
		return nil
	}
	if mask == nil {
		_, err := c.Update(ctx)
		return err
	}
	if !mask.IsValid(c.x) {
		return fmt.Errorf("invalid field mask")
	}
	paths := mask.Paths
	if len(paths) == 0 {
		_, err := c.Update(ctx)
		return err
	}
	var zero string
	if c.x.Uuid == zero {
		return fmt.Errorf("empty primary key")
	}
	m, err := c.x.AsModel()
	if err != nil {
		return err
	}
	target := CrudModel{Uuid: m.Uuid}
	cols := LookupCrudModelColumns(paths)
	db := c.db.WithContext(ctx)
	if err := db.Model(&target).Select(cols).Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (c CrudWithDB) Delete(ctx context.Context) error {
	if c.x == nil {
		return nil
	}
	var zero string
	if c.x.Uuid == zero {
		return fmt.Errorf("empty primary key")
	}
	m, err := c.x.AsModel()
	if err != nil {
		return err
	}
	db := c.db.WithContext(ctx)
	if err := db.Where(m).Delete(&CrudModel{}).Error; err != nil {
		return err
	}
	return nil
}

func WithCrudGetFieldMask(mask *fieldmaskpb.FieldMask) CrudGetOption {
	return func(tx *gorm.DB) *gorm.DB {
		cols := LookupCrudModelColumns(mask.Paths)
		tx = tx.Select(cols)
		return tx
	}
}

func WithCrudListFilter(filter string) CrudListOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx
	}
}

func WithCrudListLimit(n int) CrudListOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(n)
	}
}

func WithCrudListFieldMask(mask *fieldmaskpb.FieldMask) CrudListOption {
	return func(tx *gorm.DB) *gorm.DB {
		cols := LookupCrudModelColumns(mask.Paths)
		tx = tx.Select(cols)
		return tx
	}
}

func WithCrudListOffset(n int) CrudListOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(n)
	}
}

func WithCrudListOrder(order string) CrudListOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Order(order)
	}
}

var fieldColumnsCrudModel = map[string]string{
	"uuid":         "Uuid",
	"string_field": "StringField",
	"int32_field":  "Int32Field",
	"bool_field":   "enabled",
}

func LookupCrudModelColumn(field string) string {
	if col, ok := fieldColumnsCrudModel[field]; ok {
		return col
	} else {
		panic(field)
	}
}

func LookupCrudModelColumns(paths []string) (cols []string) {
	for _, p := range paths {
		cols = append(cols, LookupCrudModelColumn(p))
	}
	return
}
