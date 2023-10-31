// Code generated by capnpc-go. DO NOT EDIT.

package capnp

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
)

type Tx capnp.Struct

// Tx_TypeID is the unique identifier for the type Tx.
const Tx_TypeID = 0xe9135d071d75f95f

func NewTx(s *capnp.Segment) (Tx, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Tx(st), err
}

func NewRootTx(s *capnp.Segment) (Tx, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Tx(st), err
}

func ReadRootTx(msg *capnp.Message) (Tx, error) {
	root, err := msg.Root()
	return Tx(root.Struct()), err
}

func (s Tx) String() string {
	str, _ := text.Marshal(0xe9135d071d75f95f, capnp.Struct(s))
	return str
}

func (s Tx) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Tx) DecodeFromPtr(p capnp.Ptr) Tx {
	return Tx(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Tx) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Tx) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Tx) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Tx) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Tx) CommitLSN() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s Tx) SetCommitLSN(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s Tx) Records() (Tx_Record_List, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return Tx_Record_List(p.List()), err
}

func (s Tx) HasRecords() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Tx) SetRecords(v Tx_Record_List) error {
	return capnp.Struct(s).SetPtr(0, v.ToPtr())
}

// NewRecords sets the records field to a newly
// allocated Tx_Record_List, preferring placement in s's segment.
func (s Tx) NewRecords(n int32) (Tx_Record_List, error) {
	l, err := NewTx_Record_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Tx_Record_List{}, err
	}
	err = capnp.Struct(s).SetPtr(0, l.ToPtr())
	return l, err
}

// Tx_List is a list of Tx.
type Tx_List = capnp.StructList[Tx]

// NewTx creates a new list of Tx.
func NewTx_List(s *capnp.Segment, sz int32) (Tx_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return capnp.StructList[Tx](l), err
}

// Tx_Future is a wrapper for a Tx promised by a client call.
type Tx_Future struct{ *capnp.Future }

func (f Tx_Future) Struct() (Tx, error) {
	p, err := f.Future.Ptr()
	return Tx(p.Struct()), err
}

type Tx_Record capnp.Struct

// Tx_Record_TypeID is the unique identifier for the type Tx_Record.
const Tx_Record_TypeID = 0xadfa24e64cb4fa48

func NewTx_Record(s *capnp.Segment) (Tx_Record, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 6})
	return Tx_Record(st), err
}

func NewRootTx_Record(s *capnp.Segment) (Tx_Record, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 6})
	return Tx_Record(st), err
}

func ReadRootTx_Record(msg *capnp.Message) (Tx_Record, error) {
	root, err := msg.Root()
	return Tx_Record(root.Struct()), err
}

func (s Tx_Record) String() string {
	str, _ := text.Marshal(0xadfa24e64cb4fa48, capnp.Struct(s))
	return str
}

func (s Tx_Record) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Tx_Record) DecodeFromPtr(p capnp.Ptr) Tx_Record {
	return Tx_Record(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Tx_Record) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Tx_Record) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Tx_Record) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Tx_Record) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Tx_Record) Action() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Tx_Record) HasAction() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Tx_Record) ActionBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Tx_Record) SetAction(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Tx_Record) Timestamp() (string, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.Text(), err
}

func (s Tx_Record) HasTimestamp() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Tx_Record) TimestampBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.TextBytes(), err
}

func (s Tx_Record) SetTimestamp(v string) error {
	return capnp.Struct(s).SetText(1, v)
}

func (s Tx_Record) Schema() (string, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return p.Text(), err
}

func (s Tx_Record) HasSchema() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s Tx_Record) SchemaBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return p.TextBytes(), err
}

func (s Tx_Record) SetSchema(v string) error {
	return capnp.Struct(s).SetText(2, v)
}

func (s Tx_Record) Table() (string, error) {
	p, err := capnp.Struct(s).Ptr(3)
	return p.Text(), err
}

func (s Tx_Record) HasTable() bool {
	return capnp.Struct(s).HasPtr(3)
}

func (s Tx_Record) TableBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(3)
	return p.TextBytes(), err
}

func (s Tx_Record) SetTable(v string) error {
	return capnp.Struct(s).SetText(3, v)
}

func (s Tx_Record) Columns() (Tx_Record_Column_List, error) {
	p, err := capnp.Struct(s).Ptr(4)
	return Tx_Record_Column_List(p.List()), err
}

func (s Tx_Record) HasColumns() bool {
	return capnp.Struct(s).HasPtr(4)
}

func (s Tx_Record) SetColumns(v Tx_Record_Column_List) error {
	return capnp.Struct(s).SetPtr(4, v.ToPtr())
}

// NewColumns sets the columns field to a newly
// allocated Tx_Record_Column_List, preferring placement in s's segment.
func (s Tx_Record) NewColumns(n int32) (Tx_Record_Column_List, error) {
	l, err := NewTx_Record_Column_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Tx_Record_Column_List{}, err
	}
	err = capnp.Struct(s).SetPtr(4, l.ToPtr())
	return l, err
}
func (s Tx_Record) PrimaryKey() (Tx_Record_PrimaryKey_List, error) {
	p, err := capnp.Struct(s).Ptr(5)
	return Tx_Record_PrimaryKey_List(p.List()), err
}

func (s Tx_Record) HasPrimaryKey() bool {
	return capnp.Struct(s).HasPtr(5)
}

func (s Tx_Record) SetPrimaryKey(v Tx_Record_PrimaryKey_List) error {
	return capnp.Struct(s).SetPtr(5, v.ToPtr())
}

// NewPrimaryKey sets the primaryKey field to a newly
// allocated Tx_Record_PrimaryKey_List, preferring placement in s's segment.
func (s Tx_Record) NewPrimaryKey(n int32) (Tx_Record_PrimaryKey_List, error) {
	l, err := NewTx_Record_PrimaryKey_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Tx_Record_PrimaryKey_List{}, err
	}
	err = capnp.Struct(s).SetPtr(5, l.ToPtr())
	return l, err
}

// Tx_Record_List is a list of Tx_Record.
type Tx_Record_List = capnp.StructList[Tx_Record]

// NewTx_Record creates a new list of Tx_Record.
func NewTx_Record_List(s *capnp.Segment, sz int32) (Tx_Record_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 6}, sz)
	return capnp.StructList[Tx_Record](l), err
}

// Tx_Record_Future is a wrapper for a Tx_Record promised by a client call.
type Tx_Record_Future struct{ *capnp.Future }

func (f Tx_Record_Future) Struct() (Tx_Record, error) {
	p, err := f.Future.Ptr()
	return Tx_Record(p.Struct()), err
}

type Tx_Record_Column capnp.Struct

// Tx_Record_Column_TypeID is the unique identifier for the type Tx_Record_Column.
const Tx_Record_Column_TypeID = 0xdaf0d54cc25988fc

func NewTx_Record_Column(s *capnp.Segment) (Tx_Record_Column, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Tx_Record_Column(st), err
}

func NewRootTx_Record_Column(s *capnp.Segment) (Tx_Record_Column, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Tx_Record_Column(st), err
}

func ReadRootTx_Record_Column(msg *capnp.Message) (Tx_Record_Column, error) {
	root, err := msg.Root()
	return Tx_Record_Column(root.Struct()), err
}

func (s Tx_Record_Column) String() string {
	str, _ := text.Marshal(0xdaf0d54cc25988fc, capnp.Struct(s))
	return str
}

func (s Tx_Record_Column) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Tx_Record_Column) DecodeFromPtr(p capnp.Ptr) Tx_Record_Column {
	return Tx_Record_Column(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Tx_Record_Column) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Tx_Record_Column) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Tx_Record_Column) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Tx_Record_Column) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Tx_Record_Column) Name() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Tx_Record_Column) HasName() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Tx_Record_Column) NameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Tx_Record_Column) SetName(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Tx_Record_Column) Type() (string, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.Text(), err
}

func (s Tx_Record_Column) HasType() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Tx_Record_Column) TypeBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.TextBytes(), err
}

func (s Tx_Record_Column) SetType(v string) error {
	return capnp.Struct(s).SetText(1, v)
}

func (s Tx_Record_Column) Value() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return []byte(p.Data()), err
}

func (s Tx_Record_Column) HasValue() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s Tx_Record_Column) SetValue(v []byte) error {
	return capnp.Struct(s).SetData(2, v)
}

// Tx_Record_Column_List is a list of Tx_Record_Column.
type Tx_Record_Column_List = capnp.StructList[Tx_Record_Column]

// NewTx_Record_Column creates a new list of Tx_Record_Column.
func NewTx_Record_Column_List(s *capnp.Segment, sz int32) (Tx_Record_Column_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return capnp.StructList[Tx_Record_Column](l), err
}

// Tx_Record_Column_Future is a wrapper for a Tx_Record_Column promised by a client call.
type Tx_Record_Column_Future struct{ *capnp.Future }

func (f Tx_Record_Column_Future) Struct() (Tx_Record_Column, error) {
	p, err := f.Future.Ptr()
	return Tx_Record_Column(p.Struct()), err
}

type Tx_Record_PrimaryKey capnp.Struct

// Tx_Record_PrimaryKey_TypeID is the unique identifier for the type Tx_Record_PrimaryKey.
const Tx_Record_PrimaryKey_TypeID = 0x9722004316c0ea9f

func NewTx_Record_PrimaryKey(s *capnp.Segment) (Tx_Record_PrimaryKey, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Tx_Record_PrimaryKey(st), err
}

func NewRootTx_Record_PrimaryKey(s *capnp.Segment) (Tx_Record_PrimaryKey, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Tx_Record_PrimaryKey(st), err
}

func ReadRootTx_Record_PrimaryKey(msg *capnp.Message) (Tx_Record_PrimaryKey, error) {
	root, err := msg.Root()
	return Tx_Record_PrimaryKey(root.Struct()), err
}

func (s Tx_Record_PrimaryKey) String() string {
	str, _ := text.Marshal(0x9722004316c0ea9f, capnp.Struct(s))
	return str
}

func (s Tx_Record_PrimaryKey) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Tx_Record_PrimaryKey) DecodeFromPtr(p capnp.Ptr) Tx_Record_PrimaryKey {
	return Tx_Record_PrimaryKey(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Tx_Record_PrimaryKey) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Tx_Record_PrimaryKey) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Tx_Record_PrimaryKey) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Tx_Record_PrimaryKey) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Tx_Record_PrimaryKey) Name() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Tx_Record_PrimaryKey) HasName() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Tx_Record_PrimaryKey) NameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Tx_Record_PrimaryKey) SetName(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Tx_Record_PrimaryKey) Type() (string, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.Text(), err
}

func (s Tx_Record_PrimaryKey) HasType() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Tx_Record_PrimaryKey) TypeBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.TextBytes(), err
}

func (s Tx_Record_PrimaryKey) SetType(v string) error {
	return capnp.Struct(s).SetText(1, v)
}

// Tx_Record_PrimaryKey_List is a list of Tx_Record_PrimaryKey.
type Tx_Record_PrimaryKey_List = capnp.StructList[Tx_Record_PrimaryKey]

// NewTx_Record_PrimaryKey creates a new list of Tx_Record_PrimaryKey.
func NewTx_Record_PrimaryKey_List(s *capnp.Segment, sz int32) (Tx_Record_PrimaryKey_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return capnp.StructList[Tx_Record_PrimaryKey](l), err
}

// Tx_Record_PrimaryKey_Future is a wrapper for a Tx_Record_PrimaryKey promised by a client call.
type Tx_Record_PrimaryKey_Future struct{ *capnp.Future }

func (f Tx_Record_PrimaryKey_Future) Struct() (Tx_Record_PrimaryKey, error) {
	p, err := f.Future.Ptr()
	return Tx_Record_PrimaryKey(p.Struct()), err
}

type Schema capnp.Struct

// Schema_TypeID is the unique identifier for the type Schema.
const Schema_TypeID = 0xb2c63295c6bcc643

func NewSchema(s *capnp.Segment) (Schema, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Schema(st), err
}

func NewRootSchema(s *capnp.Segment) (Schema, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Schema(st), err
}

func ReadRootSchema(msg *capnp.Message) (Schema, error) {
	root, err := msg.Root()
	return Schema(root.Struct()), err
}

func (s Schema) String() string {
	str, _ := text.Marshal(0xb2c63295c6bcc643, capnp.Struct(s))
	return str
}

func (s Schema) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Schema) DecodeFromPtr(p capnp.Ptr) Schema {
	return Schema(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Schema) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Schema) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Schema) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Schema) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Schema) Columns() (Schema_Column_List, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return Schema_Column_List(p.List()), err
}

func (s Schema) HasColumns() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Schema) SetColumns(v Schema_Column_List) error {
	return capnp.Struct(s).SetPtr(0, v.ToPtr())
}

// NewColumns sets the columns field to a newly
// allocated Schema_Column_List, preferring placement in s's segment.
func (s Schema) NewColumns(n int32) (Schema_Column_List, error) {
	l, err := NewSchema_Column_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Schema_Column_List{}, err
	}
	err = capnp.Struct(s).SetPtr(0, l.ToPtr())
	return l, err
}

// Schema_List is a list of Schema.
type Schema_List = capnp.StructList[Schema]

// NewSchema creates a new list of Schema.
func NewSchema_List(s *capnp.Segment, sz int32) (Schema_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Schema](l), err
}

// Schema_Future is a wrapper for a Schema promised by a client call.
type Schema_Future struct{ *capnp.Future }

func (f Schema_Future) Struct() (Schema, error) {
	p, err := f.Future.Ptr()
	return Schema(p.Struct()), err
}

type Schema_Column capnp.Struct

// Schema_Column_TypeID is the unique identifier for the type Schema_Column.
const Schema_Column_TypeID = 0xc6ee33dc50be1a5a

func NewSchema_Column(s *capnp.Segment) (Schema_Column, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Schema_Column(st), err
}

func NewRootSchema_Column(s *capnp.Segment) (Schema_Column, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Schema_Column(st), err
}

func ReadRootSchema_Column(msg *capnp.Message) (Schema_Column, error) {
	root, err := msg.Root()
	return Schema_Column(root.Struct()), err
}

func (s Schema_Column) String() string {
	str, _ := text.Marshal(0xc6ee33dc50be1a5a, capnp.Struct(s))
	return str
}

func (s Schema_Column) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Schema_Column) DecodeFromPtr(p capnp.Ptr) Schema_Column {
	return Schema_Column(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Schema_Column) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Schema_Column) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Schema_Column) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Schema_Column) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Schema_Column) Name() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Schema_Column) HasName() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Schema_Column) NameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Schema_Column) SetName(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Schema_Column) Type() (string, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.Text(), err
}

func (s Schema_Column) HasType() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Schema_Column) TypeBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.TextBytes(), err
}

func (s Schema_Column) SetType(v string) error {
	return capnp.Struct(s).SetText(1, v)
}

func (s Schema_Column) IsNullable() bool {
	return capnp.Struct(s).Bit(0)
}

func (s Schema_Column) SetIsNullable(v bool) {
	capnp.Struct(s).SetBit(0, v)
}

func (s Schema_Column) IsPartOfPrimaryKey() bool {
	return capnp.Struct(s).Bit(1)
}

func (s Schema_Column) SetIsPartOfPrimaryKey(v bool) {
	capnp.Struct(s).SetBit(1, v)
}

// Schema_Column_List is a list of Schema_Column.
type Schema_Column_List = capnp.StructList[Schema_Column]

// NewSchema_Column creates a new list of Schema_Column.
func NewSchema_Column_List(s *capnp.Segment, sz int32) (Schema_Column_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return capnp.StructList[Schema_Column](l), err
}

// Schema_Column_Future is a wrapper for a Schema_Column promised by a client call.
type Schema_Column_Future struct{ *capnp.Future }

func (f Schema_Column_Future) Struct() (Schema_Column, error) {
	p, err := f.Future.Ptr()
	return Schema_Column(p.Struct()), err
}

type DealInfo capnp.Struct

// DealInfo_TypeID is the unique identifier for the type DealInfo.
const DealInfo_TypeID = 0xc69bf7c475bb2886

func NewDealInfo(s *capnp.Segment) (DealInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return DealInfo(st), err
}

func NewRootDealInfo(s *capnp.Segment) (DealInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return DealInfo(st), err
}

func ReadRootDealInfo(msg *capnp.Message) (DealInfo, error) {
	root, err := msg.Root()
	return DealInfo(root.Struct()), err
}

func (s DealInfo) String() string {
	str, _ := text.Marshal(0xc69bf7c475bb2886, capnp.Struct(s))
	return str
}

func (s DealInfo) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (DealInfo) DecodeFromPtr(p capnp.Ptr) DealInfo {
	return DealInfo(capnp.Struct{}.DecodeFromPtr(p))
}

func (s DealInfo) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s DealInfo) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s DealInfo) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s DealInfo) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s DealInfo) Cid() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s DealInfo) HasCid() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s DealInfo) CidBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s DealInfo) SetCid(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s DealInfo) Size() uint32 {
	return capnp.Struct(s).Uint32(0)
}

func (s DealInfo) SetSize(v uint32) {
	capnp.Struct(s).SetUint32(0, v)
}

func (s DealInfo) Created() (string, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.Text(), err
}

func (s DealInfo) HasCreated() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s DealInfo) CreatedBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.TextBytes(), err
}

func (s DealInfo) SetCreated(v string) error {
	return capnp.Struct(s).SetText(1, v)
}

func (s DealInfo) IsPermament() bool {
	return capnp.Struct(s).Bit(32)
}

func (s DealInfo) SetIsPermament(v bool) {
	capnp.Struct(s).SetBit(32, v)
}

// DealInfo_List is a list of DealInfo.
type DealInfo_List = capnp.StructList[DealInfo]

// NewDealInfo creates a new list of DealInfo.
func NewDealInfo_List(s *capnp.Segment, sz int32) (DealInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return capnp.StructList[DealInfo](l), err
}

// DealInfo_Future is a wrapper for a DealInfo promised by a client call.
type DealInfo_Future struct{ *capnp.Future }

func (f DealInfo_Future) Struct() (DealInfo, error) {
	p, err := f.Future.Ptr()
	return DealInfo(p.Struct()), err
}

const schema_8c49da2775b6e7db = "x\xda\x8c\x94Oh\x1ce\x18\xc6\x9f\xe7\xfbfv7" +
	"\x90\x98\x0e\xb3\xfe+\x95\xd8\x1al\x0dM\xda$\x9e\x16" +
	"d\x8b\xa9\x90`Z\xf7\xdd\x16\xc4\xa0\xe8d31\xa3" +
	";\xbb\xeb\xee\xacm\xbc\xd4\x93\x88xT\xc4?\xc5\x1e" +
	"\x14Z\xb0\xf8\x07\xf1\xa2H\x95\xb27\xf1\xe4\xc1Z\x10" +
	"\x0fZ\xeaA\xf4 b\x10F\xbe\xdd\xec\xec4\xb4!" +
	"\xc7y\xe7\x9dy\x9f\xdf\xf3>\xdfwx\xaf:bM" +
	"\x8f|\xaf\xa1\xe4\xa0\x9d\xf9\xf9\xfd\xdf/\xdd1\xb7\xef" +
	"-\xe7~\xc6\xf3\x1b\x9f/\xfe6\xbeq\x11\xb6\xca\x02" +
	"\xb3\xdf\xf15\xba\xbf2\x0b\xb8\xbf\xf0\x14R\xef\x9d1" +
	"\xc6O\xff\xdb\xbe'\xfb\x94{\x1dv\xc64?\xa4f" +
	"\xe8\x1eS\xfb\x81Y_=N0\x9e\xeb|\xd5ys" +
	"\xa6\xf3\x19\x9c=\x8c\x7f\xba\xf6E{\xff\x95\x85\xd7a" +
	"\x9b\xff\xcd\xfe\xa8w\xd3\xbd\xae\xef\x04\xdc\xbft\x11\x8c" +
	"_9\xf0e\xfb\xf2?\xefv {\x98\xee\xee\x0a\xd9" +
	"kM\xd0\x9d\xb6\x8c\x90I\xeb\x1a\x18/\xed\xfe\xbat" +
	"u\xf6\x8f\x0e\xe4>\xa6\x06\xf5\xbam\xbbL\xf7n\xdb" +
	"t\xdfn\x9b\xee\xff^}\xe2\xdb\xc5\x1f\xfe\xbc\x82\x1b" +
	"\x19\xb5i\xfe\xdb~\x8e\xee\x90!p\xed\xcc\xc7Ha" +
	"m\x11\xd2\xb5\xe1lf\xc3\xbd\x901\xaa?\xcd\x9c\xc2" +
	"d\xdcx\xfe\xd9C\x15\xafQ\xb3\x1b\x87V\xfc\xd5\xa0" +
	"\x16DA\xbd\xd6\x9a2\xa5F\xe1\xe4\xe9\xa9\xb2_\xa9" +
	"7W\xa6J\xcd \xf4\x9a\xeb\x8f\xfa\xeb@\x89\x94\x9c" +
	"\xb6\x00\x8b\x80\xf3\xc0\x04 \xe3\x9arX\xd1!\xf34" +
	"\xc5IS<\xa0)\x0f*\x8e\xd6\xbc\xd0\xe70\x14\x87" +
	"\xc1\xd1h\xbd\x91<$\xc3\xf5\xad\x86\x8fu\xa7K\x8e" +
	"i\x0b\x86\x0a\xc9\xc6\xed\xa5\xe2\\\xbd\xda\x0ekq_" +
	" \xb4\xbf.w%\xea\xde.\x00\xf2\x86\xa6\x9cK\xa9" +
	";[\x06\xe4=M9\xaf\xe8(\x95\xa7\x02\x9c\x0fM" +
	"\xe79M\xf9H\xd1\xd1:O\x0d8\x17f\x00\xf9@" +
	"S>Qt,+O\x0bp.>\x0c\xc8yM\xb9" +
	"\xac\xe8\xd8v\x9e6\xe0|\xb3\x04\xc8%M\xb9\xaaX" +
	"\xf4*\x06$\xc1\x8c\x82\xd0oE^\x086\xfa\xb5b" +
	"\xab\xb2\xe6\x87^\xffq,\xf2\x96\xab\x89/g*]" +
	"\xa6\x16o\x03K\x9a\xdc5\x80\x07M1n\xa4h\x93" +
	"\xae\xbe+\xfd\x9e\xed\xcc=QY\xf3u\xe8\x89\xc5T" +
	"\x14\x1d\x166\xed\x14+qp\xc4\xd0\xe64e\\\xdd" +
	"DX\xf2\xedN\x86\x1e\xf5\xbd\xea\xe8Bm\xb5n\x12" +
	"\xb4+\x99\xe0\xed\x03\xe4IMYS\xec\xaf\xc87\x01" +
	"zFS\xaafE\xec\xad(0RV4\xa5aV" +
	"tooE\xe12 UM9\xad\x98\xad\x04+I" +
	"\xd2Z\xc1K>sP\xcc\x19G\x9b\xbe\x17\xf9\xc9\xcb" +
	"8h\x95\xfcf\xe8\x85\xc8\xfa\xb5\x88\x84\"S\xea\xad" +
	"[Y\x16zS]\x83X\xdb\x8201@pn\xc6" +
	"\xc0\xcd\x94\x05&&k\x9a\x12\x19\x04\xf6\x10^x\x07" +
	"\x90HS^\xde\xf6\xb4\x04\xad\xe3\xedj\xd5[\x86\xae" +
	"\xfa\x89\xe4\xa0U\xf2\x9a\xd1c\xab\xec\x1f\x00\x93\x88\x1d" +
	"\xf1\x0c\x0e\xf7\\}\xd4\xac\xd5 \x0d'H\x8f\x18\xf5" +
	"G4e1\x85\xb4`\x8aG5\xa5\x94:9\xc7\xcc" +
	"!\x99\xd7\x94\x93\xdb\xc9\x1f{\xd1\xab\xb6}\x8e@q" +
	"$%Mm\x95V\xeci\xeb&sp[\xb3P," +
	"o\xde\x05\x83\x9b\xa7<\xb8d\xfa\x02\xa7MD\x0ej" +
	"\xca\xbcb\\\xa9\x87a\x10-\x9e\x00\x8fs\x08\x8aC" +
	"\xe0\x99f\xf7/\xa9\x04'3z\x09\xfe?\x00\x00\xff" +
	"\xff\xaa%\x91\xb2"

func RegisterSchema(reg *schemas.Registry) {
	reg.Register(&schemas.Schema{
		String: schema_8c49da2775b6e7db,
		Nodes: []uint64{
			0x9722004316c0ea9f,
			0xadfa24e64cb4fa48,
			0xb2c63295c6bcc643,
			0xc69bf7c475bb2886,
			0xc6ee33dc50be1a5a,
			0xdaf0d54cc25988fc,
			0xe9135d071d75f95f,
		},
		Compressed: true,
	})
}
