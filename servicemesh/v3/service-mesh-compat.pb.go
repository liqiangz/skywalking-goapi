//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// service-mesh-probe/service-mesh-compat.proto is a deprecated file.

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_mesh_probe_service_mesh_compat_proto protoreflect.FileDescriptor

var file_service_mesh_probe_service_mesh_compat_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x73,
	0x68, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x6f, 0x0a, 0x18, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4d, 0x65, 0x73, 0x68, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x53, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x73,
	0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x68, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x1a, 0x22,
	0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d,
	0x65, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x22, 0x00, 0x28, 0x01, 0x42, 0x81, 0x01, 0x0a, 0x37, 0x6f, 0x72, 0x67, 0x2e, 0x61,
	0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0x50, 0x01, 0x5a, 0x21, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x33, 0xb8, 0x01, 0x01, 0xaa, 0x02, 0x1d, 0x53, 0x6b, 0x79,
	0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_service_mesh_probe_service_mesh_compat_proto_goTypes = []interface{}{
	(*ServiceMeshMetric)(nil),   // 0: skywalking.v3.ServiceMeshMetric
	(*MeshProbeDownstream)(nil), // 1: skywalking.v3.MeshProbeDownstream
}
var file_service_mesh_probe_service_mesh_compat_proto_depIdxs = []int32{
	0, // 0: ServiceMeshMetricService.collect:input_type -> skywalking.v3.ServiceMeshMetric
	1, // 1: ServiceMeshMetricService.collect:output_type -> skywalking.v3.MeshProbeDownstream
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_mesh_probe_service_mesh_compat_proto_init() }
func file_service_mesh_probe_service_mesh_compat_proto_init() {
	if File_service_mesh_probe_service_mesh_compat_proto != nil {
		return
	}
	file_service_mesh_probe_service_mesh_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_mesh_probe_service_mesh_compat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_mesh_probe_service_mesh_compat_proto_goTypes,
		DependencyIndexes: file_service_mesh_probe_service_mesh_compat_proto_depIdxs,
	}.Build()
	File_service_mesh_probe_service_mesh_compat_proto = out.File
	file_service_mesh_probe_service_mesh_compat_proto_rawDesc = nil
	file_service_mesh_probe_service_mesh_compat_proto_goTypes = nil
	file_service_mesh_probe_service_mesh_compat_proto_depIdxs = nil
}
