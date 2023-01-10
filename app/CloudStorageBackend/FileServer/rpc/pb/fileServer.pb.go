// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: pb/fileServer.proto

package pb

import (
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/model/PojoDB/fileMetaPojo"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileMetaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	CreateGroup string `protobuf:"bytes,2,opt,name=createGroup,proto3" json:"createGroup,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime  string `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Authority   string `protobuf:"bytes,5,opt,name=authority,proto3" json:"authority,omitempty"`
	TypeOf      string `protobuf:"bytes,6,opt,name=typeOf,proto3" json:"typeOf,omitempty"`
	UpdateTime  string `protobuf:"bytes,7,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Size        int64  `protobuf:"varint,8,opt,name=size,proto3" json:"size,omitempty"`
	IsDir       bool   `protobuf:"varint,9,opt,name=isDir,proto3" json:"isDir,omitempty"`
	DeleteTime  string `protobuf:"bytes,10,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
}

func GetFileMetaInfo(filemetatable *fileMetaPojo.Filemetatable)(result *FileMetaInfo){
    result=&FileMetaInfo{

		Creator:       filemetatable.Creator,
		CreateGroup:   filemetatable.CreateGroup,
		Name:          filemetatable.Name,
		CreateTime:    filemetatable.CreateTime.String(),
		Authority:     filemetatable.Authority,
		TypeOf:        filemetatable.TypeOf,
		UpdateTime:    filemetatable.UpdateTime.String(),
		Size:          filemetatable.Size,
		IsDir:         filemetatable.IsDir,
	}
	deleteTime:=filemetatable.DeleteTime
	if(deleteTime.Valid){
		result.DeleteTime=deleteTime.Time.String()
	}
	return result
}

func (x *FileMetaInfo) Reset() {
	*x = FileMetaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_fileServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMetaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetaInfo) ProtoMessage() {}

func (x *FileMetaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fileServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetaInfo.ProtoReflect.Descriptor instead.
func (*FileMetaInfo) Descriptor() ([]byte, []int) {
	return file_pb_fileServer_proto_rawDescGZIP(), []int{0}
}

func (x *FileMetaInfo) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *FileMetaInfo) GetCreateGroup() string {
	if x != nil {
		return x.CreateGroup
	}
	return ""
}

func (x *FileMetaInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileMetaInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *FileMetaInfo) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *FileMetaInfo) GetTypeOf() string {
	if x != nil {
		return x.TypeOf
	}
	return ""
}

func (x *FileMetaInfo) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *FileMetaInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileMetaInfo) GetIsDir() bool {
	if x != nil {
		return x.IsDir
	}
	return false
}

func (x *FileMetaInfo) GetDeleteTime() string {
	if x != nil {
		return x.DeleteTime
	}
	return ""
}

type FindFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FindFileReq) Reset() {
	*x = FindFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_fileServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindFileReq) ProtoMessage() {}

func (x *FindFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fileServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindFileReq.ProtoReflect.Descriptor instead.
func (*FindFileReq) Descriptor() ([]byte, []int) {
	return file_pb_fileServer_proto_rawDescGZIP(), []int{1}
}

func (x *FindFileReq) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *FindFileReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type QueryFileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*FileMetaInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *QueryFileRes) Reset() {
	*x = QueryFileRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_fileServer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryFileRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFileRes) ProtoMessage() {}

func (x *QueryFileRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fileServer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFileRes.ProtoReflect.Descriptor instead.
func (*QueryFileRes) Descriptor() ([]byte, []int) {
	return file_pb_fileServer_proto_rawDescGZIP(), []int{2}
}

func (x *QueryFileRes) GetList() []*FileMetaInfo {
	if x != nil {
		return x.List
	}
	return nil
}

var File_pb_fileServer_proto protoreflect.FileDescriptor

var file_pb_fileServer_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xa1, 0x02, 0x0a, 0x0c, 0x46, 0x69,
	0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x79, 0x70,
	0x65, 0x4f, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x4f,
	0x66, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x44, 0x69, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x44, 0x69, 0x72, 0x12, 0x1f, 0x0a, 0x0b,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x37, 0x0a,
	0x0b, 0x46, 0x69, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x6b, 0x0a, 0x0a,
	0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x07, 0x46, 0x69,
	0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_fileServer_proto_rawDescOnce sync.Once
	file_pb_fileServer_proto_rawDescData = file_pb_fileServer_proto_rawDesc
)

func file_pb_fileServer_proto_rawDescGZIP() []byte {
	file_pb_fileServer_proto_rawDescOnce.Do(func() {
		file_pb_fileServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_fileServer_proto_rawDescData)
	})
	return file_pb_fileServer_proto_rawDescData
}

var file_pb_fileServer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_fileServer_proto_goTypes = []interface{}{
	(*FileMetaInfo)(nil), // 0: pb.FileMetaInfo
	(*FindFileReq)(nil),  // 1: pb.FindFileReq
	(*QueryFileRes)(nil), // 2: pb.QueryFileRes
}
var file_pb_fileServer_proto_depIdxs = []int32{
	0, // 0: pb.QueryFileRes.list:type_name -> pb.FileMetaInfo
	1, // 1: pb.fileServer.FindOne:input_type -> pb.FindFileReq
	1, // 2: pb.fileServer.QueryFiles:input_type -> pb.FindFileReq
	0, // 3: pb.fileServer.FindOne:output_type -> pb.FileMetaInfo
	2, // 4: pb.fileServer.QueryFiles:output_type -> pb.QueryFileRes
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_fileServer_proto_init() }
func file_pb_fileServer_proto_init() {
	if File_pb_fileServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_fileServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMetaInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_fileServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindFileReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_fileServer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryFileRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_fileServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_fileServer_proto_goTypes,
		DependencyIndexes: file_pb_fileServer_proto_depIdxs,
		MessageInfos:      file_pb_fileServer_proto_msgTypes,
	}.Build()
	File_pb_fileServer_proto = out.File
	file_pb_fileServer_proto_rawDesc = nil
	file_pb_fileServer_proto_goTypes = nil
	file_pb_fileServer_proto_depIdxs = nil
}
