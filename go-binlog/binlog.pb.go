// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: binlog.proto

/*
	Package binlog is a generated protocol buffer package.

	It is generated from these files:
		binlog.proto

	It has these top-level messages:
		TableMutation
		PrewriteValue
		Binlog
*/
package binlog

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MutationType int32

const (
	MutationType_Insert    MutationType = 0
	MutationType_Update    MutationType = 1
	MutationType_DeleteID  MutationType = 2
	MutationType_DeletePK  MutationType = 3
	MutationType_DeleteRow MutationType = 4
)

var MutationType_name = map[int32]string{
	0: "Insert",
	1: "Update",
	2: "DeleteID",
	3: "DeletePK",
	4: "DeleteRow",
}
var MutationType_value = map[string]int32{
	"Insert":    0,
	"Update":    1,
	"DeleteID":  2,
	"DeletePK":  3,
	"DeleteRow": 4,
}

func (x MutationType) Enum() *MutationType {
	p := new(MutationType)
	*p = x
	return p
}
func (x MutationType) String() string {
	return proto.EnumName(MutationType_name, int32(x))
}
func (x *MutationType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MutationType_value, data, "MutationType")
	if err != nil {
		return err
	}
	*x = MutationType(value)
	return nil
}
func (MutationType) EnumDescriptor() ([]byte, []int) { return fileDescriptorBinlog, []int{0} }

type BinlogType int32

const (
	BinlogType_Prewrite BinlogType = 0
	BinlogType_Commit   BinlogType = 1
	BinlogType_Rollback BinlogType = 2
	BinlogType_PreDDL   BinlogType = 3
	BinlogType_PostDDL  BinlogType = 4
)

var BinlogType_name = map[int32]string{
	0: "Prewrite",
	1: "Commit",
	2: "Rollback",
	3: "PreDDL",
	4: "PostDDL",
}
var BinlogType_value = map[string]int32{
	"Prewrite": 0,
	"Commit":   1,
	"Rollback": 2,
	"PreDDL":   3,
	"PostDDL":  4,
}

func (x BinlogType) Enum() *BinlogType {
	p := new(BinlogType)
	*p = x
	return p
}
func (x BinlogType) String() string {
	return proto.EnumName(BinlogType_name, int32(x))
}
func (x *BinlogType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BinlogType_value, data, "BinlogType")
	if err != nil {
		return err
	}
	*x = BinlogType(value)
	return nil
}
func (BinlogType) EnumDescriptor() ([]byte, []int) { return fileDescriptorBinlog, []int{1} }

// TableMutation contains mutations in a table.
type TableMutation struct {
	TableId int64 `protobuf:"varint,1,opt,name=table_id,json=tableId" json:"table_id"`
	// The inserted row contains all column values.
	InsertedRows [][]byte `protobuf:"bytes,2,rep,name=inserted_rows,json=insertedRows" json:"inserted_rows,omitempty"`
	// The updated row contains old values and new values of the row.
	UpdatedRows [][]byte `protobuf:"bytes,3,rep,name=updated_rows,json=updatedRows" json:"updated_rows,omitempty"`
	// Obsolete field.
	DeletedIds []int64 `protobuf:"varint,4,rep,name=deleted_ids,json=deletedIds" json:"deleted_ids,omitempty"`
	// Obsolete field.
	DeletedPks [][]byte `protobuf:"bytes,5,rep,name=deleted_pks,json=deletedPks" json:"deleted_pks,omitempty"`
	// The row value of the deleted row.
	DeletedRows [][]byte `protobuf:"bytes,6,rep,name=deleted_rows,json=deletedRows" json:"deleted_rows,omitempty"`
	// Used to apply table mutations in original sequence.
	Sequence         []MutationType `protobuf:"varint,7,rep,name=sequence,enum=binlog.MutationType" json:"sequence,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *TableMutation) Reset()                    { *m = TableMutation{} }
func (m *TableMutation) String() string            { return proto.CompactTextString(m) }
func (*TableMutation) ProtoMessage()               {}
func (*TableMutation) Descriptor() ([]byte, []int) { return fileDescriptorBinlog, []int{0} }

func (m *TableMutation) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *TableMutation) GetInsertedRows() [][]byte {
	if m != nil {
		return m.InsertedRows
	}
	return nil
}

func (m *TableMutation) GetUpdatedRows() [][]byte {
	if m != nil {
		return m.UpdatedRows
	}
	return nil
}

func (m *TableMutation) GetDeletedIds() []int64 {
	if m != nil {
		return m.DeletedIds
	}
	return nil
}

func (m *TableMutation) GetDeletedPks() [][]byte {
	if m != nil {
		return m.DeletedPks
	}
	return nil
}

func (m *TableMutation) GetDeletedRows() [][]byte {
	if m != nil {
		return m.DeletedRows
	}
	return nil
}

func (m *TableMutation) GetSequence() []MutationType {
	if m != nil {
		return m.Sequence
	}
	return nil
}

type PrewriteValue struct {
	SchemaVersion    int64           `protobuf:"varint,1,opt,name=schema_version,json=schemaVersion" json:"schema_version"`
	Mutations        []TableMutation `protobuf:"bytes,2,rep,name=mutations" json:"mutations"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *PrewriteValue) Reset()                    { *m = PrewriteValue{} }
func (m *PrewriteValue) String() string            { return proto.CompactTextString(m) }
func (*PrewriteValue) ProtoMessage()               {}
func (*PrewriteValue) Descriptor() ([]byte, []int) { return fileDescriptorBinlog, []int{1} }

func (m *PrewriteValue) GetSchemaVersion() int64 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

func (m *PrewriteValue) GetMutations() []TableMutation {
	if m != nil {
		return m.Mutations
	}
	return nil
}

// Binlog contains all the changes in a transaction, which can be used to reconstruct SQL statement, then export to
// other systems.
type Binlog struct {
	Tp BinlogType `protobuf:"varint,1,opt,name=tp,enum=binlog.BinlogType" json:"tp"`
	// start_ts is used in Prewrite, Commit and Rollback binlog Type.
	// It is used for pairing prewrite log to commit log or rollback log.
	StartTs int64 `protobuf:"varint,2,opt,name=start_ts,json=startTs" json:"start_ts"`
	// commit_ts is used only in binlog type Commit.
	CommitTs int64 `protobuf:"varint,3,opt,name=commit_ts,json=commitTs" json:"commit_ts"`
	// prewrite key is used only in Prewrite binlog type.
	// It is the primary key of the transaction, is used to check that the transaction is
	// commited or not if it failed to pair to commit log or rollback log within a time window.
	PrewriteKey []byte `protobuf:"bytes,4,opt,name=prewrite_key,json=prewriteKey" json:"prewrite_key,omitempty"`
	// prewrite_data is marshalled from PrewriteData type,
	// we do not need to unmarshal prewrite data before the binlog have been successfully paired.
	PrewriteValue []byte `protobuf:"bytes,5,opt,name=prewrite_value,json=prewriteValue" json:"prewrite_value,omitempty"`
	// ddl_query is the original DDL statement query.
	DdlQuery []byte `protobuf:"bytes,6,opt,name=ddl_query,json=ddlQuery" json:"ddl_query,omitempty"`
	// ddl_job_id is used for DDL Binlog.
	// If ddl_job_id is setted, this is a DDL Binlog and ddl_query contains the DDL query, we can query the informations about this job from TiKV.
	DdlJobId int64 `protobuf:"varint,7,opt,name=ddl_job_id,json=ddlJobId" json:"ddl_job_id"`
	// ip is the source IP which this binlog generated from.
	Ip               []byte `protobuf:"bytes,9,opt,name=ip" json:"ip"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Binlog) Reset()                    { *m = Binlog{} }
func (m *Binlog) String() string            { return proto.CompactTextString(m) }
func (*Binlog) ProtoMessage()               {}
func (*Binlog) Descriptor() ([]byte, []int) { return fileDescriptorBinlog, []int{2} }

func (m *Binlog) GetTp() BinlogType {
	if m != nil {
		return m.Tp
	}
	return BinlogType_Prewrite
}

func (m *Binlog) GetStartTs() int64 {
	if m != nil {
		return m.StartTs
	}
	return 0
}

func (m *Binlog) GetCommitTs() int64 {
	if m != nil {
		return m.CommitTs
	}
	return 0
}

func (m *Binlog) GetPrewriteKey() []byte {
	if m != nil {
		return m.PrewriteKey
	}
	return nil
}

func (m *Binlog) GetPrewriteValue() []byte {
	if m != nil {
		return m.PrewriteValue
	}
	return nil
}

func (m *Binlog) GetDdlQuery() []byte {
	if m != nil {
		return m.DdlQuery
	}
	return nil
}

func (m *Binlog) GetDdlJobId() int64 {
	if m != nil {
		return m.DdlJobId
	}
	return 0
}

func (m *Binlog) GetIp() []byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

func init() {
	proto.RegisterType((*TableMutation)(nil), "binlog.TableMutation")
	proto.RegisterType((*PrewriteValue)(nil), "binlog.PrewriteValue")
	proto.RegisterType((*Binlog)(nil), "binlog.Binlog")
	proto.RegisterEnum("binlog.MutationType", MutationType_name, MutationType_value)
	proto.RegisterEnum("binlog.BinlogType", BinlogType_name, BinlogType_value)
}
func (m *TableMutation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TableMutation) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintBinlog(dAtA, i, uint64(m.TableId))
	if len(m.InsertedRows) > 0 {
		for _, b := range m.InsertedRows {
			dAtA[i] = 0x12
			i++
			i = encodeVarintBinlog(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if len(m.UpdatedRows) > 0 {
		for _, b := range m.UpdatedRows {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintBinlog(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if len(m.DeletedIds) > 0 {
		for _, num := range m.DeletedIds {
			dAtA[i] = 0x20
			i++
			i = encodeVarintBinlog(dAtA, i, uint64(num))
		}
	}
	if len(m.DeletedPks) > 0 {
		for _, b := range m.DeletedPks {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintBinlog(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if len(m.DeletedRows) > 0 {
		for _, b := range m.DeletedRows {
			dAtA[i] = 0x32
			i++
			i = encodeVarintBinlog(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if len(m.Sequence) > 0 {
		for _, num := range m.Sequence {
			dAtA[i] = 0x38
			i++
			i = encodeVarintBinlog(dAtA, i, uint64(num))
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PrewriteValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrewriteValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintBinlog(dAtA, i, uint64(m.SchemaVersion))
	if len(m.Mutations) > 0 {
		for _, msg := range m.Mutations {
			dAtA[i] = 0x12
			i++
			i = encodeVarintBinlog(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Binlog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Binlog) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintBinlog(dAtA, i, uint64(m.Tp))
	dAtA[i] = 0x10
	i++
	i = encodeVarintBinlog(dAtA, i, uint64(m.StartTs))
	dAtA[i] = 0x18
	i++
	i = encodeVarintBinlog(dAtA, i, uint64(m.CommitTs))
	if m.PrewriteKey != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintBinlog(dAtA, i, uint64(len(m.PrewriteKey)))
		i += copy(dAtA[i:], m.PrewriteKey)
	}
	if m.PrewriteValue != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintBinlog(dAtA, i, uint64(len(m.PrewriteValue)))
		i += copy(dAtA[i:], m.PrewriteValue)
	}
	if m.DdlQuery != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintBinlog(dAtA, i, uint64(len(m.DdlQuery)))
		i += copy(dAtA[i:], m.DdlQuery)
	}
	dAtA[i] = 0x38
	i++
	i = encodeVarintBinlog(dAtA, i, uint64(m.DdlJobId))
	if m.Ip != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintBinlog(dAtA, i, uint64(len(m.Ip)))
		i += copy(dAtA[i:], m.Ip)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintBinlog(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TableMutation) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovBinlog(uint64(m.TableId))
	if len(m.InsertedRows) > 0 {
		for _, b := range m.InsertedRows {
			l = len(b)
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if len(m.UpdatedRows) > 0 {
		for _, b := range m.UpdatedRows {
			l = len(b)
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if len(m.DeletedIds) > 0 {
		for _, e := range m.DeletedIds {
			n += 1 + sovBinlog(uint64(e))
		}
	}
	if len(m.DeletedPks) > 0 {
		for _, b := range m.DeletedPks {
			l = len(b)
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if len(m.DeletedRows) > 0 {
		for _, b := range m.DeletedRows {
			l = len(b)
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if len(m.Sequence) > 0 {
		for _, e := range m.Sequence {
			n += 1 + sovBinlog(uint64(e))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PrewriteValue) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovBinlog(uint64(m.SchemaVersion))
	if len(m.Mutations) > 0 {
		for _, e := range m.Mutations {
			l = e.Size()
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Binlog) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovBinlog(uint64(m.Tp))
	n += 1 + sovBinlog(uint64(m.StartTs))
	n += 1 + sovBinlog(uint64(m.CommitTs))
	if m.PrewriteKey != nil {
		l = len(m.PrewriteKey)
		n += 1 + l + sovBinlog(uint64(l))
	}
	if m.PrewriteValue != nil {
		l = len(m.PrewriteValue)
		n += 1 + l + sovBinlog(uint64(l))
	}
	if m.DdlQuery != nil {
		l = len(m.DdlQuery)
		n += 1 + l + sovBinlog(uint64(l))
	}
	n += 1 + sovBinlog(uint64(m.DdlJobId))
	if m.Ip != nil {
		l = len(m.Ip)
		n += 1 + l + sovBinlog(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBinlog(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBinlog(x uint64) (n int) {
	return sovBinlog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TableMutation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBinlog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TableMutation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableMutation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableId", wireType)
			}
			m.TableId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TableId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsertedRows", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InsertedRows = append(m.InsertedRows, make([]byte, postIndex-iNdEx))
			copy(m.InsertedRows[len(m.InsertedRows)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedRows", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedRows = append(m.UpdatedRows, make([]byte, postIndex-iNdEx))
			copy(m.UpdatedRows[len(m.UpdatedRows)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBinlog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.DeletedIds = append(m.DeletedIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBinlog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBinlog
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBinlog
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.DeletedIds = append(m.DeletedIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedIds", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedPks", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeletedPks = append(m.DeletedPks, make([]byte, postIndex-iNdEx))
			copy(m.DeletedPks[len(m.DeletedPks)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedRows", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeletedRows = append(m.DeletedRows, make([]byte, postIndex-iNdEx))
			copy(m.DeletedRows[len(m.DeletedRows)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType == 0 {
				var v MutationType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBinlog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (MutationType(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Sequence = append(m.Sequence, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBinlog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBinlog
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v MutationType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBinlog
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (MutationType(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Sequence = append(m.Sequence, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBinlog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBinlog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PrewriteValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBinlog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PrewriteValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrewriteValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaVersion", wireType)
			}
			m.SchemaVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SchemaVersion |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mutations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mutations = append(m.Mutations, TableMutation{})
			if err := m.Mutations[len(m.Mutations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBinlog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBinlog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Binlog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBinlog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Binlog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Binlog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tp |= (BinlogType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTs", wireType)
			}
			m.StartTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTs |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitTs", wireType)
			}
			m.CommitTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommitTs |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrewriteKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrewriteKey = append(m.PrewriteKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PrewriteKey == nil {
				m.PrewriteKey = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrewriteValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrewriteValue = append(m.PrewriteValue[:0], dAtA[iNdEx:postIndex]...)
			if m.PrewriteValue == nil {
				m.PrewriteValue = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DdlQuery", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DdlQuery = append(m.DdlQuery[:0], dAtA[iNdEx:postIndex]...)
			if m.DdlQuery == nil {
				m.DdlQuery = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DdlJobId", wireType)
			}
			m.DdlJobId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DdlJobId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = append(m.Ip[:0], dAtA[iNdEx:postIndex]...)
			if m.Ip == nil {
				m.Ip = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBinlog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBinlog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBinlog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBinlog
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBinlog
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBinlog
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBinlog(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBinlog = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBinlog   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("binlog.proto", fileDescriptorBinlog) }

var fileDescriptorBinlog = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xcd, 0x8e, 0xda, 0x30,
	0x10, 0xc7, 0x49, 0xc2, 0xf2, 0x31, 0x24, 0x28, 0xb2, 0xa8, 0x14, 0xb5, 0x12, 0x50, 0xaa, 0x4a,
	0x68, 0x2b, 0x6d, 0x2b, 0x6e, 0xbd, 0x52, 0x2e, 0x74, 0xbb, 0x12, 0x1b, 0xd1, 0xbd, 0x46, 0x01,
	0x5b, 0xd4, 0x25, 0xc4, 0xde, 0xd8, 0xd9, 0x88, 0xa7, 0xe8, 0xb1, 0x7d, 0xa4, 0x3d, 0xf6, 0x09,
	0xaa, 0x8a, 0xbe, 0x48, 0x65, 0x3b, 0xe1, 0xe3, 0xe6, 0xf9, 0xcf, 0x6f, 0x66, 0xec, 0xff, 0x18,
	0xdc, 0x15, 0x4d, 0x13, 0xb6, 0xb9, 0xe1, 0x19, 0x93, 0x0c, 0x35, 0x4c, 0xf4, 0xb2, 0xb7, 0x61,
	0x1b, 0xa6, 0xa5, 0xf7, 0xea, 0x64, 0xb2, 0xa3, 0x1f, 0x36, 0x78, 0xcb, 0x78, 0x95, 0x90, 0xbb,
	0x5c, 0xc6, 0x92, 0xb2, 0x14, 0x0d, 0xa0, 0x25, 0x95, 0x10, 0x51, 0x1c, 0x58, 0x43, 0x6b, 0xec,
	0x4c, 0xeb, 0xcf, 0x7f, 0x06, 0xb5, 0xb0, 0xa9, 0xd5, 0x39, 0x46, 0x6f, 0xc0, 0xa3, 0xa9, 0x20,
	0x99, 0x24, 0x38, 0xca, 0x58, 0x21, 0x02, 0x7b, 0xe8, 0x8c, 0xdd, 0xd0, 0xad, 0xc4, 0x90, 0x15,
	0x02, 0xbd, 0x06, 0x37, 0xe7, 0x38, 0x3e, 0x32, 0x8e, 0x66, 0x3a, 0xa5, 0xa6, 0x91, 0x01, 0x74,
	0x30, 0x49, 0x88, 0x42, 0x28, 0x16, 0x41, 0x7d, 0xe8, 0x8c, 0x9d, 0x10, 0x4a, 0x69, 0x8e, 0x2f,
	0x00, 0xbe, 0x15, 0xc1, 0x95, 0x6e, 0x51, 0x01, 0x8b, 0xad, 0x1e, 0x52, 0x01, 0x7a, 0x48, 0xc3,
	0x0c, 0x29, 0x35, 0x3d, 0xe4, 0x03, 0xb4, 0x04, 0x79, 0xcc, 0x49, 0xba, 0x26, 0x41, 0x73, 0xe8,
	0x8c, 0xbb, 0x93, 0xde, 0x4d, 0x69, 0x4f, 0xf5, 0xe2, 0xe5, 0x9e, 0x93, 0xf0, 0x48, 0x8d, 0x0a,
	0xf0, 0x16, 0x19, 0x29, 0x32, 0x2a, 0xc9, 0x43, 0x9c, 0xe4, 0x04, 0xbd, 0x83, 0xae, 0x58, 0x7f,
	0x23, 0xbb, 0x38, 0x7a, 0x22, 0x99, 0xa0, 0x2c, 0xbd, 0xb0, 0xc5, 0x33, 0xb9, 0x07, 0x93, 0x42,
	0x1f, 0xa1, 0xbd, 0x2b, 0xfb, 0x1a, 0x63, 0x3a, 0x93, 0x17, 0xd5, 0xc0, 0x0b, 0x9f, 0xcb, 0xf2,
	0x13, 0x3d, 0xfa, 0x69, 0x43, 0x63, 0xaa, 0x49, 0x34, 0x06, 0x5b, 0x72, 0x3d, 0xa6, 0x3b, 0x41,
	0x55, 0xb9, 0xc9, 0xa9, 0xdb, 0x96, 0xb5, 0xb6, 0xe4, 0x6a, 0x5b, 0x42, 0xc6, 0x99, 0x8c, 0xa4,
	0x1a, 0x77, 0xb6, 0x2d, 0xad, 0x2e, 0x95, 0x47, 0xed, 0x35, 0xdb, 0xed, 0xa8, 0x26, 0x9c, 0x33,
	0xa2, 0x65, 0x64, 0x8d, 0xb8, 0xbc, 0x7c, 0x71, 0xb4, 0x25, 0xfb, 0xa0, 0x3e, 0xb4, 0x94, 0x8d,
	0x95, 0x76, 0x4b, 0xf6, 0xe8, 0x2d, 0x74, 0x8f, 0xc8, 0x93, 0x72, 0x25, 0xb8, 0xd2, 0x90, 0xc7,
	0x2f, 0xac, 0x7a, 0x05, 0x6d, 0x8c, 0x93, 0xe8, 0x31, 0x27, 0xd9, 0x3e, 0x68, 0x68, 0xa2, 0x85,
	0x71, 0x72, 0xaf, 0x62, 0x34, 0x02, 0x50, 0xc9, 0xef, 0x6c, 0xa5, 0xbe, 0x56, 0xf3, 0xfc, 0x2a,
	0x18, 0x27, 0x9f, 0xd9, 0x6a, 0x8e, 0x51, 0x0f, 0x6c, 0xca, 0x83, 0xb6, 0xaa, 0xac, 0x1e, 0x49,
	0xf9, 0xf5, 0x3d, 0xb8, 0xe7, 0xcb, 0x42, 0x00, 0x8d, 0xb9, 0xfe, 0x6c, 0x7e, 0x4d, 0x9d, 0xbf,
	0xea, 0x4f, 0xe5, 0x5b, 0xc8, 0x85, 0xd6, 0x4c, 0xef, 0x7e, 0x3e, 0xf3, 0xed, 0x53, 0xb4, 0xb8,
	0xf5, 0x1d, 0xe4, 0x41, 0xdb, 0x44, 0x21, 0x2b, 0xfc, 0xfa, 0xf5, 0x1d, 0xc0, 0xc9, 0x4f, 0x85,
	0x56, 0x3b, 0x37, 0x2d, 0x3f, 0x69, 0x6f, 0x4c, 0xcb, 0x90, 0x25, 0xc9, 0x2a, 0x5e, 0x6f, 0x7d,
	0x5b, 0x65, 0x16, 0x19, 0x99, 0xcd, 0xbe, 0xf8, 0x0e, 0xea, 0x40, 0x73, 0xc1, 0x84, 0x54, 0x41,
	0x7d, 0xea, 0x3f, 0x1f, 0xfa, 0xd6, 0xef, 0x43, 0xdf, 0xfa, 0x7b, 0xe8, 0x5b, 0xbf, 0xfe, 0xf5,
	0x6b, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x11, 0x34, 0x0f, 0x07, 0x85, 0x03, 0x00, 0x00,
}
