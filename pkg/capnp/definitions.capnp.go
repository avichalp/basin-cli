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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return DealInfo(st), err
}

func NewRootDealInfo(s *capnp.Segment) (DealInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
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

func (s DealInfo) Timestamp() int64 {
	return int64(capnp.Struct(s).Uint64(8))
}

func (s DealInfo) SetTimestamp(v int64) {
	capnp.Struct(s).SetUint64(8, uint64(v))
}

func (s DealInfo) Archived() bool {
	return capnp.Struct(s).Bit(32)
}

func (s DealInfo) SetArchived(v bool) {
	capnp.Struct(s).SetBit(32, v)
}

func (s DealInfo) ExpiresAt() (string, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.Text(), err
}

func (s DealInfo) HasExpiresAt() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s DealInfo) ExpiresAtBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.TextBytes(), err
}

func (s DealInfo) SetExpiresAt(v string) error {
	return capnp.Struct(s).SetText(1, v)
}

// DealInfo_List is a list of DealInfo.
type DealInfo_List = capnp.StructList[DealInfo]

// NewDealInfo creates a new list of DealInfo.
func NewDealInfo_List(s *capnp.Segment, sz int32) (DealInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2}, sz)
	return capnp.StructList[DealInfo](l), err
}

// DealInfo_Future is a wrapper for a DealInfo promised by a client call.
type DealInfo_Future struct{ *capnp.Future }

func (f DealInfo_Future) Struct() (DealInfo, error) {
	p, err := f.Future.Ptr()
	return DealInfo(p.Struct()), err
}

const schema_8c49da2775b6e7db = "x\xda\x8c\x94_h\x1cU\x14\xc6\xcfw\xef\xcc&\x81" +
	"\xc4t\x98E\xa5DVk\xb0\xb54i\x93\xf8\xe2\x82" +
	"l5\x15\xb2\x9a\xd4=\xdb\x824Xt:\x99\xb8\xa3" +
	";\xbb\xeb\xeel\x9b\x88R\x85*E*\xf8`\xf1?" +
	"\xf6A\xa1\x82E\x11\xf1E\x91*e\xdf\xc4'\x1f\xac" +
	"\x05\xe9\x83\x96\x0a\x8a>\x88\x18\x84\x91\xb3\x9b\x99\x1d\xa3" +
	"\x0d\xbe\xed\x9c9;\xe7\xfb\x9d\xef\xbbw\xcf\xb4\xdak" +
	"L\x8d|\xadI\xf1.3\xf3\xfd\xdb?\x9d\xbf~v" +
	"\xdb+\xd6m\x88\xe6\xd6>\x9e\xffq|\xed\x1c\x99j" +
	"\x80h\xe6+\xbc\x00\xfb\x07\x0c\x10\xd9\x97q\x8cR\xef" +
	"\xad\x1c\xa2\x87\xffl\xdf4p\xd8\xbeJfF\x9a\xef" +
	"R\xd3\xb0\x17\xd4v\xa2\x19O=\x08B4\xdb\xf9\xac" +
	"sz\xba\xf3\x11Yc\x88\xbe\xbb\xf2I{\xfb\xc5\xe2" +
	")2\xe5{3\xdf\xea\xad\xb0\xaf\xea\x1b\x88\xec\xdft" +
	"\x81\x10=\xbf\xe3\xd3\xf6\x85?\xde\xe8\x10\x8fA\xa5\xba" +
	"\xbbBn1v\xc2\x9e2\xe4\xe7\x84\x91\x93o/n" +
	"\xfd\xbcti\xe6\x97\x0e\xf1\xadHM\xea\xb5\x17\xcd2" +
	"\xec\xc3\xa6\xe8>d^!D\x7f\x9d<\xf4\xe5\xfc7" +
	"\xbf^\xa4\x7fBji\xbe3\xf3\x18\xec\x05A\xb0\x8b" +
	"\x99\x0f(\xc5\xc5c\xd8\xa8\xdb\xbe\x9cY\xb3\x7f\xce\x88" +
	"\xec\xdf3\xc7h\"j<\xfe\xe8n\xd7i\xd4\xcc\xc6" +
	"\xee%o\xd9\xaf\xf9\xa1_\xaf\xb5&\xa5\xd4\xc8\x1f\\" +
	"\x99,{n\xbd\xb94Yj\xfa\x81\xd3\\\xbd\xdf[" +
	"%*\x01<\xa8\x0d\"\x03D\xd6\xed;\x89x\\\x83" +
	"\xf7(X@\x16R\x9c\x90\xe2\x0e\x0d\xbeCa\xb4\xe6" +
	"\x04\x1e\x86Ia\x980\x1a\xae6\x92\x87d\xb8\xbe\xd6" +
	"\xf0\\w:\x0f\"\xbd\x82\xa1|b\xb9\xb9X\x98\xad" +
	"W\xdbA-\x8a\x05\x92\xf6V\xf9\xc6D\xddky\"" +
	"~Y\x83\xcf\xa4\xd4\xbdU&\xe275\xf8\xac\x82\xa5" +
	"T\x16\x8a\xc8zW:\xcfh\xf0\xfb\x0a\x96\xd6Yh" +
	"\"\xeb\xbdi\"~G\x83?T\xb0\x0c#\x0b\x83\xc8" +
	":w\x0f\x11\x9f\xd5\xe0\x0b\x0a\x96ifa\x12Y_" +
	",\x12\xf1y\x0d\xbe\xa4Pp\\\x01I0C?\xf0" +
	"Z\xa1\x13\x10\x1aq\xad\xd0r+^\xe0\xc4\x8f\xb9\xd0" +
	"9RM\xf6r\xdc\xed2\xb5p\x1d\xa1\xa4\x81-}" +
	"x\x82\x14\xa3F\x8a6\xe9\x8a\xb7\x12\xf7l\xb6\xdc\x03" +
	"n\xc5\xd3\x81\xc3\x06RQ\xb4\x90__'\x1b\xc9\x06" +
	"G\x84vP\x83\xc7\xd5\x7f\x08K\xfe\xfb\x7f\x86\xee\xf3" +
	"\x9c\xeah\xb1\xb6\\\x97\x04e\x93\x09Oo#\xe2\x15" +
	"\x0d>\xa1\x10[\xf4\xac\x04\xe8)\x0d>)\x16\xa1g" +
	"\xd1s\xe2\xdb\x09\x0d~I,\xba\xb9g\xd1\x8b\xf7\x11" +
	"\xf1)\x0d~U,B\xcf\xa2\xd3\xe5\xbe\xed\x03\xae\xbf" +
	"\x94\xc4\xaf\xe5?\xe9a\x90\x14\x067\xf8b\x92\x82I" +
	"\x88\x9c\xa6[\xf1\x8fzK$D\xa4 g\xd5[i" +
	"\xf8M\xafu7!\xfcWt\x8dkm7p&\xbb" +
	"\xbbDMh\xb7$\xb4\x8e\x90=\xa4\xc1\x95T\"=" +
	")>\xa2\xc1U\x05\xac\x07\xd2\x97DU48\x14Z" +
	"\xf4h\x9fx\x9d\x88C\x0d~f\xd3\x83\xe5\xb7\xf6\xb7" +
	"\xabU\xe7\x08\xe9\xaa\x97`\xf8\xad\x92\xd3\x0c\x1fXF" +
	"|V$<\xf1\xcbMy\xfa\xf7\xc0l}T\x12 " +
	"H\xc3\x09\xd2\xbd\xa2~\xaf\x06\xcf\xa7\x90\x8aR\xdc\xa7" +
	"\xc1\xa5\xd4![\x90\xf34\xa7\xc1\x077\x93\x9f;\xea" +
	"T\xdb\x1eFHa$%Mm\x94V\xe8i\xeb\x86" +
	"\xb8\x7f\xb3#_(\xaf_\x1b\xfdK\xaa\xdc\xbf\x8fb" +
	"\x81S\x12\xec]\x1a<\xa7\x10\xb9\xf5 \xf0\xc3\xf9\x03" +
	"\x84\xfd\x18\"\x85!\xc2\xf1f\xf7+\xa9\xb0'3z" +
	"a\xff;\x00\x00\xff\xff\x88\xeb\x9b\xe1"

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
