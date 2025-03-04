// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: vector_values.gen.go.tmpl

package values

import (
	"fmt"
	"regexp"

	arrow "github.com/influxdata/flux/array"
	"github.com/influxdata/flux/memory"
	"github.com/influxdata/flux/semantic"
)

func NewVectorValue(arr arrow.Array, typ semantic.MonoType) Vector {
	switch typ {

	case semantic.BasicInt:
		return NewIntVectorValue(arr.(*arrow.Int))

	case semantic.BasicUint:
		return NewUintVectorValue(arr.(*arrow.Uint))

	case semantic.BasicFloat:
		return NewFloatVectorValue(arr.(*arrow.Float))

	case semantic.BasicBool:
		return NewBooleanVectorValue(arr.(*arrow.Boolean))

	case semantic.BasicString:
		return NewStringVectorValue(arr.(*arrow.String))

	case semantic.BasicTime:
		return NewTimeVectorValue(arr.(*arrow.Int))

	default:
		panic(fmt.Errorf("unsupported column data type: %s", typ))
	}
}

// A convenience method for unit testing
func NewVectorFromElements(mem memory.Allocator, es ...interface{}) Vector {
	var typ semantic.MonoType
	switch es[0].(type) {

	case int64:
		typ = semantic.BasicInt

	case uint64:
		typ = semantic.BasicUint

	case float64:
		typ = semantic.BasicFloat

	case bool:
		typ = semantic.BasicBool

	case string:
		typ = semantic.BasicString

	case Time:
		typ = semantic.BasicTime

	default:
		panic(fmt.Errorf("unsupported data type"))
	}

	vs := make([]Value, len(es))
	for i, e := range es {
		vs[i] = New(e)
	}
	return newVectorFromSlice(vs, typ, mem)
}

func newVectorFromSlice(values []Value, typ semantic.MonoType, mem memory.Allocator) Vector {
	switch typ {

	case semantic.BasicInt:
		b := arrow.NewIntBuilder(mem)
		for _, v := range values {
			b.Append(v.Int())
		}
		arr := b.NewIntArray()
		return NewIntVectorValue(arr)

	case semantic.BasicUint:
		b := arrow.NewUintBuilder(mem)
		for _, v := range values {
			b.Append(v.UInt())
		}
		arr := b.NewUintArray()
		return NewUintVectorValue(arr)

	case semantic.BasicFloat:
		b := arrow.NewFloatBuilder(mem)
		for _, v := range values {
			b.Append(v.Float())
		}
		arr := b.NewFloatArray()
		return NewFloatVectorValue(arr)

	case semantic.BasicBool:
		b := arrow.NewBooleanBuilder(mem)
		for _, v := range values {
			b.Append(v.Bool())
		}
		arr := b.NewBooleanArray()
		return NewBooleanVectorValue(arr)

	case semantic.BasicString:
		b := arrow.NewStringBuilder(mem)
		for _, v := range values {
			b.Append(v.Str())
		}
		arr := b.NewStringArray()
		return NewStringVectorValue(arr)

	case semantic.BasicTime:
		b := arrow.NewIntBuilder(mem)
		for _, v := range values {
			b.Append(v.Int())
		}
		arr := b.NewIntArray()
		return NewTimeVectorValue(arr)

	default:
		panic(fmt.Errorf("unsupported column data type: %s", typ))
	}
}

func NewVectorRepeatValue(v Value) Vector {
	typ := v.Type()
	switch typ {

	case semantic.BasicInt:
		return &VectorRepeatValue{
			val: v,
			typ: semantic.NewVectorType(semantic.BasicInt),
		}

	case semantic.BasicUint:
		return &VectorRepeatValue{
			val: v,
			typ: semantic.NewVectorType(semantic.BasicUint),
		}

	case semantic.BasicFloat:
		return &VectorRepeatValue{
			val: v,
			typ: semantic.NewVectorType(semantic.BasicFloat),
		}

	case semantic.BasicBool:
		return &VectorRepeatValue{
			val: v,
			typ: semantic.NewVectorType(semantic.BasicBool),
		}

	case semantic.BasicString:
		return &VectorRepeatValue{
			val: v,
			typ: semantic.NewVectorType(semantic.BasicString),
		}

	case semantic.BasicTime:
		return &VectorRepeatValue{
			val: v,
			typ: semantic.NewVectorType(semantic.BasicTime),
		}

	default:
		panic(fmt.Errorf("unsupported column data type: %s", typ))
	}
}

type VectorRepeatValue struct {
	val Value
	typ semantic.MonoType
}

func (v *VectorRepeatValue) Value() Value { return v.val }

func (v *VectorRepeatValue) ElementType() semantic.MonoType {
	t, err := v.typ.ElemType()
	if err != nil {
		panic("could not get element type of vector value")
	}
	return t
}
func (v *VectorRepeatValue) Arr() arrow.Array {
	panic("unsupported - VectorRepeatValue has no array backing")
}

func (v *VectorRepeatValue) IsRepeat() bool { return true }
func (v *VectorRepeatValue) Retain() {
	v.val.Retain()
}
func (v *VectorRepeatValue) Release() {
	v.val.Release()
}
func (v *VectorRepeatValue) Type() semantic.MonoType { return v.typ }
func (v *VectorRepeatValue) IsNull() bool            { return false }
func (v *VectorRepeatValue) Str() string             { panic(UnexpectedKind(semantic.Vector, semantic.String)) }
func (v *VectorRepeatValue) Bytes() []byte           { panic(UnexpectedKind(semantic.Vector, semantic.Bytes)) }
func (v *VectorRepeatValue) Int() int64              { panic(UnexpectedKind(semantic.Vector, semantic.Int)) }
func (v *VectorRepeatValue) UInt() uint64            { panic(UnexpectedKind(semantic.Vector, semantic.UInt)) }
func (v *VectorRepeatValue) Float() float64          { panic(UnexpectedKind(semantic.Vector, semantic.Float)) }
func (v *VectorRepeatValue) Bool() bool              { panic(UnexpectedKind(semantic.Vector, semantic.Bool)) }
func (v *VectorRepeatValue) Time() Time              { panic(UnexpectedKind(semantic.Vector, semantic.Time)) }
func (v *VectorRepeatValue) Duration() Duration {
	panic(UnexpectedKind(semantic.Vector, semantic.Duration))
}
func (v *VectorRepeatValue) Regexp() *regexp.Regexp {
	panic(UnexpectedKind(semantic.Vector, semantic.Regexp))
}
func (v *VectorRepeatValue) Array() Array   { panic(UnexpectedKind(semantic.Vector, semantic.Array)) }
func (v *VectorRepeatValue) Object() Object { panic(UnexpectedKind(semantic.Vector, semantic.Object)) }
func (v *VectorRepeatValue) Function() Function {
	panic(UnexpectedKind(semantic.Vector, semantic.Function))
}
func (v *VectorRepeatValue) Dict() Dictionary {
	panic(UnexpectedKind(semantic.Vector, semantic.Dictionary))
}
func (v *VectorRepeatValue) Vector() Vector {
	return v
}
func (v *VectorRepeatValue) Equal(other Value) bool {
	panic("cannot compare two vectors for equality")
}

var _ Value = &VectorRepeatValue{}
var _ Vector = &VectorRepeatValue{}

var _ Value = &IntVectorValue{}
var _ Vector = &IntVectorValue{}
var _ arrow.Array = &arrow.Int{}

type IntVectorValue struct {
	arr *arrow.Int
	typ semantic.MonoType
}

func NewIntVectorValue(arr *arrow.Int) Vector {
	return &IntVectorValue{
		arr: arr,
		typ: semantic.NewVectorType(semantic.BasicInt),
	}
}

func (v *IntVectorValue) ElementType() semantic.MonoType {
	t, err := v.typ.ElemType()
	if err != nil {
		panic("could not get element type of vector value")
	}
	return t
}
func (v *IntVectorValue) Arr() arrow.Array { return v.arr }
func (v *IntVectorValue) IsRepeat() bool   { return false }
func (v *IntVectorValue) Retain() {
	v.arr.Retain()
}
func (v *IntVectorValue) Release() {
	v.arr.Release()
}

func (v *IntVectorValue) Type() semantic.MonoType { return v.typ }
func (v *IntVectorValue) IsNull() bool            { return false }
func (v *IntVectorValue) Str() string             { panic(UnexpectedKind(semantic.Vector, semantic.String)) }
func (v *IntVectorValue) Bytes() []byte           { panic(UnexpectedKind(semantic.Vector, semantic.Bytes)) }
func (v *IntVectorValue) Int() int64              { panic(UnexpectedKind(semantic.Vector, semantic.Int)) }
func (v *IntVectorValue) UInt() uint64            { panic(UnexpectedKind(semantic.Vector, semantic.UInt)) }
func (v *IntVectorValue) Float() float64          { panic(UnexpectedKind(semantic.Vector, semantic.Float)) }
func (v *IntVectorValue) Bool() bool              { panic(UnexpectedKind(semantic.Vector, semantic.Bool)) }
func (v *IntVectorValue) Time() Time              { panic(UnexpectedKind(semantic.Vector, semantic.Time)) }
func (v *IntVectorValue) Duration() Duration {
	panic(UnexpectedKind(semantic.Vector, semantic.Duration))
}
func (v *IntVectorValue) Regexp() *regexp.Regexp {
	panic(UnexpectedKind(semantic.Vector, semantic.Regexp))
}
func (v *IntVectorValue) Array() Array   { panic(UnexpectedKind(semantic.Vector, semantic.Array)) }
func (v *IntVectorValue) Object() Object { panic(UnexpectedKind(semantic.Vector, semantic.Object)) }
func (v *IntVectorValue) Function() Function {
	panic(UnexpectedKind(semantic.Vector, semantic.Function))
}
func (v *IntVectorValue) Dict() Dictionary {
	panic(UnexpectedKind(semantic.Vector, semantic.Dictionary))
}
func (v *IntVectorValue) Vector() Vector {
	return v
}

func (v *IntVectorValue) Equal(other Value) bool {
	panic("cannot compare two vectors for equality")
}

var _ Value = &UintVectorValue{}
var _ Vector = &UintVectorValue{}
var _ arrow.Array = &arrow.Uint{}

type UintVectorValue struct {
	arr *arrow.Uint
	typ semantic.MonoType
}

func NewUintVectorValue(arr *arrow.Uint) Vector {
	return &UintVectorValue{
		arr: arr,
		typ: semantic.NewVectorType(semantic.BasicUint),
	}
}

func (v *UintVectorValue) ElementType() semantic.MonoType {
	t, err := v.typ.ElemType()
	if err != nil {
		panic("could not get element type of vector value")
	}
	return t
}
func (v *UintVectorValue) Arr() arrow.Array { return v.arr }
func (v *UintVectorValue) IsRepeat() bool   { return false }
func (v *UintVectorValue) Retain() {
	v.arr.Retain()
}
func (v *UintVectorValue) Release() {
	v.arr.Release()
}

func (v *UintVectorValue) Type() semantic.MonoType { return v.typ }
func (v *UintVectorValue) IsNull() bool            { return false }
func (v *UintVectorValue) Str() string             { panic(UnexpectedKind(semantic.Vector, semantic.String)) }
func (v *UintVectorValue) Bytes() []byte           { panic(UnexpectedKind(semantic.Vector, semantic.Bytes)) }
func (v *UintVectorValue) Int() int64              { panic(UnexpectedKind(semantic.Vector, semantic.Int)) }
func (v *UintVectorValue) UInt() uint64            { panic(UnexpectedKind(semantic.Vector, semantic.UInt)) }
func (v *UintVectorValue) Float() float64          { panic(UnexpectedKind(semantic.Vector, semantic.Float)) }
func (v *UintVectorValue) Bool() bool              { panic(UnexpectedKind(semantic.Vector, semantic.Bool)) }
func (v *UintVectorValue) Time() Time              { panic(UnexpectedKind(semantic.Vector, semantic.Time)) }
func (v *UintVectorValue) Duration() Duration {
	panic(UnexpectedKind(semantic.Vector, semantic.Duration))
}
func (v *UintVectorValue) Regexp() *regexp.Regexp {
	panic(UnexpectedKind(semantic.Vector, semantic.Regexp))
}
func (v *UintVectorValue) Array() Array   { panic(UnexpectedKind(semantic.Vector, semantic.Array)) }
func (v *UintVectorValue) Object() Object { panic(UnexpectedKind(semantic.Vector, semantic.Object)) }
func (v *UintVectorValue) Function() Function {
	panic(UnexpectedKind(semantic.Vector, semantic.Function))
}
func (v *UintVectorValue) Dict() Dictionary {
	panic(UnexpectedKind(semantic.Vector, semantic.Dictionary))
}
func (v *UintVectorValue) Vector() Vector {
	return v
}

func (v *UintVectorValue) Equal(other Value) bool {
	panic("cannot compare two vectors for equality")
}

var _ Value = &FloatVectorValue{}
var _ Vector = &FloatVectorValue{}
var _ arrow.Array = &arrow.Float{}

type FloatVectorValue struct {
	arr *arrow.Float
	typ semantic.MonoType
}

func NewFloatVectorValue(arr *arrow.Float) Vector {
	return &FloatVectorValue{
		arr: arr,
		typ: semantic.NewVectorType(semantic.BasicFloat),
	}
}

func (v *FloatVectorValue) ElementType() semantic.MonoType {
	t, err := v.typ.ElemType()
	if err != nil {
		panic("could not get element type of vector value")
	}
	return t
}
func (v *FloatVectorValue) Arr() arrow.Array { return v.arr }
func (v *FloatVectorValue) IsRepeat() bool   { return false }
func (v *FloatVectorValue) Retain() {
	v.arr.Retain()
}
func (v *FloatVectorValue) Release() {
	v.arr.Release()
}

func (v *FloatVectorValue) Type() semantic.MonoType { return v.typ }
func (v *FloatVectorValue) IsNull() bool            { return false }
func (v *FloatVectorValue) Str() string             { panic(UnexpectedKind(semantic.Vector, semantic.String)) }
func (v *FloatVectorValue) Bytes() []byte           { panic(UnexpectedKind(semantic.Vector, semantic.Bytes)) }
func (v *FloatVectorValue) Int() int64              { panic(UnexpectedKind(semantic.Vector, semantic.Int)) }
func (v *FloatVectorValue) UInt() uint64            { panic(UnexpectedKind(semantic.Vector, semantic.UInt)) }
func (v *FloatVectorValue) Float() float64          { panic(UnexpectedKind(semantic.Vector, semantic.Float)) }
func (v *FloatVectorValue) Bool() bool              { panic(UnexpectedKind(semantic.Vector, semantic.Bool)) }
func (v *FloatVectorValue) Time() Time              { panic(UnexpectedKind(semantic.Vector, semantic.Time)) }
func (v *FloatVectorValue) Duration() Duration {
	panic(UnexpectedKind(semantic.Vector, semantic.Duration))
}
func (v *FloatVectorValue) Regexp() *regexp.Regexp {
	panic(UnexpectedKind(semantic.Vector, semantic.Regexp))
}
func (v *FloatVectorValue) Array() Array   { panic(UnexpectedKind(semantic.Vector, semantic.Array)) }
func (v *FloatVectorValue) Object() Object { panic(UnexpectedKind(semantic.Vector, semantic.Object)) }
func (v *FloatVectorValue) Function() Function {
	panic(UnexpectedKind(semantic.Vector, semantic.Function))
}
func (v *FloatVectorValue) Dict() Dictionary {
	panic(UnexpectedKind(semantic.Vector, semantic.Dictionary))
}
func (v *FloatVectorValue) Vector() Vector {
	return v
}

func (v *FloatVectorValue) Equal(other Value) bool {
	panic("cannot compare two vectors for equality")
}

var _ Value = &BooleanVectorValue{}
var _ Vector = &BooleanVectorValue{}
var _ arrow.Array = &arrow.Boolean{}

type BooleanVectorValue struct {
	arr *arrow.Boolean
	typ semantic.MonoType
}

func NewBooleanVectorValue(arr *arrow.Boolean) Vector {
	return &BooleanVectorValue{
		arr: arr,
		typ: semantic.NewVectorType(semantic.BasicBool),
	}
}

func (v *BooleanVectorValue) ElementType() semantic.MonoType {
	t, err := v.typ.ElemType()
	if err != nil {
		panic("could not get element type of vector value")
	}
	return t
}
func (v *BooleanVectorValue) Arr() arrow.Array { return v.arr }
func (v *BooleanVectorValue) IsRepeat() bool   { return false }
func (v *BooleanVectorValue) Retain() {
	v.arr.Retain()
}
func (v *BooleanVectorValue) Release() {
	v.arr.Release()
}

func (v *BooleanVectorValue) Type() semantic.MonoType { return v.typ }
func (v *BooleanVectorValue) IsNull() bool            { return false }
func (v *BooleanVectorValue) Str() string             { panic(UnexpectedKind(semantic.Vector, semantic.String)) }
func (v *BooleanVectorValue) Bytes() []byte           { panic(UnexpectedKind(semantic.Vector, semantic.Bytes)) }
func (v *BooleanVectorValue) Int() int64              { panic(UnexpectedKind(semantic.Vector, semantic.Int)) }
func (v *BooleanVectorValue) UInt() uint64            { panic(UnexpectedKind(semantic.Vector, semantic.UInt)) }
func (v *BooleanVectorValue) Float() float64          { panic(UnexpectedKind(semantic.Vector, semantic.Float)) }
func (v *BooleanVectorValue) Bool() bool              { panic(UnexpectedKind(semantic.Vector, semantic.Bool)) }
func (v *BooleanVectorValue) Time() Time              { panic(UnexpectedKind(semantic.Vector, semantic.Time)) }
func (v *BooleanVectorValue) Duration() Duration {
	panic(UnexpectedKind(semantic.Vector, semantic.Duration))
}
func (v *BooleanVectorValue) Regexp() *regexp.Regexp {
	panic(UnexpectedKind(semantic.Vector, semantic.Regexp))
}
func (v *BooleanVectorValue) Array() Array   { panic(UnexpectedKind(semantic.Vector, semantic.Array)) }
func (v *BooleanVectorValue) Object() Object { panic(UnexpectedKind(semantic.Vector, semantic.Object)) }
func (v *BooleanVectorValue) Function() Function {
	panic(UnexpectedKind(semantic.Vector, semantic.Function))
}
func (v *BooleanVectorValue) Dict() Dictionary {
	panic(UnexpectedKind(semantic.Vector, semantic.Dictionary))
}
func (v *BooleanVectorValue) Vector() Vector {
	return v
}

func (v *BooleanVectorValue) Equal(other Value) bool {
	panic("cannot compare two vectors for equality")
}

var _ Value = &StringVectorValue{}
var _ Vector = &StringVectorValue{}
var _ arrow.Array = &arrow.String{}

type StringVectorValue struct {
	arr *arrow.String
	typ semantic.MonoType
}

func NewStringVectorValue(arr *arrow.String) Vector {
	return &StringVectorValue{
		arr: arr,
		typ: semantic.NewVectorType(semantic.BasicString),
	}
}

func (v *StringVectorValue) ElementType() semantic.MonoType {
	t, err := v.typ.ElemType()
	if err != nil {
		panic("could not get element type of vector value")
	}
	return t
}
func (v *StringVectorValue) Arr() arrow.Array { return v.arr }
func (v *StringVectorValue) IsRepeat() bool   { return false }
func (v *StringVectorValue) Retain() {
	v.arr.Retain()
}
func (v *StringVectorValue) Release() {
	v.arr.Release()
}

func (v *StringVectorValue) Type() semantic.MonoType { return v.typ }
func (v *StringVectorValue) IsNull() bool            { return false }
func (v *StringVectorValue) Str() string             { panic(UnexpectedKind(semantic.Vector, semantic.String)) }
func (v *StringVectorValue) Bytes() []byte           { panic(UnexpectedKind(semantic.Vector, semantic.Bytes)) }
func (v *StringVectorValue) Int() int64              { panic(UnexpectedKind(semantic.Vector, semantic.Int)) }
func (v *StringVectorValue) UInt() uint64            { panic(UnexpectedKind(semantic.Vector, semantic.UInt)) }
func (v *StringVectorValue) Float() float64          { panic(UnexpectedKind(semantic.Vector, semantic.Float)) }
func (v *StringVectorValue) Bool() bool              { panic(UnexpectedKind(semantic.Vector, semantic.Bool)) }
func (v *StringVectorValue) Time() Time              { panic(UnexpectedKind(semantic.Vector, semantic.Time)) }
func (v *StringVectorValue) Duration() Duration {
	panic(UnexpectedKind(semantic.Vector, semantic.Duration))
}
func (v *StringVectorValue) Regexp() *regexp.Regexp {
	panic(UnexpectedKind(semantic.Vector, semantic.Regexp))
}
func (v *StringVectorValue) Array() Array   { panic(UnexpectedKind(semantic.Vector, semantic.Array)) }
func (v *StringVectorValue) Object() Object { panic(UnexpectedKind(semantic.Vector, semantic.Object)) }
func (v *StringVectorValue) Function() Function {
	panic(UnexpectedKind(semantic.Vector, semantic.Function))
}
func (v *StringVectorValue) Dict() Dictionary {
	panic(UnexpectedKind(semantic.Vector, semantic.Dictionary))
}
func (v *StringVectorValue) Vector() Vector {
	return v
}

func (v *StringVectorValue) Equal(other Value) bool {
	panic("cannot compare two vectors for equality")
}

var _ Value = &TimeVectorValue{}
var _ Vector = &TimeVectorValue{}
var _ arrow.Array = &arrow.Int{}

type TimeVectorValue struct {
	arr *arrow.Int
	typ semantic.MonoType
}

func NewTimeVectorValue(arr *arrow.Int) Vector {
	return &TimeVectorValue{
		arr: arr,
		typ: semantic.NewVectorType(semantic.BasicTime),
	}
}

func (v *TimeVectorValue) ElementType() semantic.MonoType {
	t, err := v.typ.ElemType()
	if err != nil {
		panic("could not get element type of vector value")
	}
	return t
}
func (v *TimeVectorValue) Arr() arrow.Array { return v.arr }
func (v *TimeVectorValue) IsRepeat() bool   { return false }
func (v *TimeVectorValue) Retain() {
	v.arr.Retain()
}
func (v *TimeVectorValue) Release() {
	v.arr.Release()
}

func (v *TimeVectorValue) Type() semantic.MonoType { return v.typ }
func (v *TimeVectorValue) IsNull() bool            { return false }
func (v *TimeVectorValue) Str() string             { panic(UnexpectedKind(semantic.Vector, semantic.String)) }
func (v *TimeVectorValue) Bytes() []byte           { panic(UnexpectedKind(semantic.Vector, semantic.Bytes)) }
func (v *TimeVectorValue) Int() int64              { panic(UnexpectedKind(semantic.Vector, semantic.Int)) }
func (v *TimeVectorValue) UInt() uint64            { panic(UnexpectedKind(semantic.Vector, semantic.UInt)) }
func (v *TimeVectorValue) Float() float64          { panic(UnexpectedKind(semantic.Vector, semantic.Float)) }
func (v *TimeVectorValue) Bool() bool              { panic(UnexpectedKind(semantic.Vector, semantic.Bool)) }
func (v *TimeVectorValue) Time() Time              { panic(UnexpectedKind(semantic.Vector, semantic.Time)) }
func (v *TimeVectorValue) Duration() Duration {
	panic(UnexpectedKind(semantic.Vector, semantic.Duration))
}
func (v *TimeVectorValue) Regexp() *regexp.Regexp {
	panic(UnexpectedKind(semantic.Vector, semantic.Regexp))
}
func (v *TimeVectorValue) Array() Array   { panic(UnexpectedKind(semantic.Vector, semantic.Array)) }
func (v *TimeVectorValue) Object() Object { panic(UnexpectedKind(semantic.Vector, semantic.Object)) }
func (v *TimeVectorValue) Function() Function {
	panic(UnexpectedKind(semantic.Vector, semantic.Function))
}
func (v *TimeVectorValue) Dict() Dictionary {
	panic(UnexpectedKind(semantic.Vector, semantic.Dictionary))
}
func (v *TimeVectorValue) Vector() Vector {
	return v
}

func (v *TimeVectorValue) Equal(other Value) bool {
	panic("cannot compare two vectors for equality")
}
