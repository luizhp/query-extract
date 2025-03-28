package entity

import "reflect"

type Column struct {
	position         int
	name             string
	databaseTypeName string
	scanType         reflect.Type
	length           int64
	precision        int64
	scale            int64
	nullable         bool
}

func NewColumn(position int, name, databaseTypeName string, scanType reflect.Type, length, precision, scale int64, nullable bool) *Column {
	return &Column{
		position:         position,
		name:             name,
		databaseTypeName: databaseTypeName,
		scanType:         scanType,
		length:           length,
		precision:        precision,
		scale:            scale,
		nullable:         nullable,
	}
}

func (r *Column) GetPosition() int {
	return r.position
}

func (r *Column) GetName() string {
	return r.name
}

func (r *Column) GetDatabaseTypeName() string {
	return r.databaseTypeName
}

func (r *Column) GetScanType() interface{} {
	return r.scanType
}

func (r *Column) GetLength() int64 {
	return r.length
}

func (r *Column) GetPrecision() int64 {
	return r.precision
}

func (r *Column) GetScale() int64 {
	return r.scale
}

func (r *Column) GetNullable() bool {
	return r.nullable
}
