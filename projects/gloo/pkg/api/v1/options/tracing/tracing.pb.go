// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/tracing/tracing.proto

package tracing

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/trace/v3"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Contains settings for configuring Envoy's tracing capabilities at the listener level.
// See here for additional information on Envoy's tracing capabilities: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/tracing.html
// See here for additional information about configuring tracing with Gloo: https://gloo.solo.io/user_guides/setup_options/observability/#tracing
type ListenerTracingSettings struct {
	// Optional. If specified, Envoy will include the headers and header values for any matching request headers.
	RequestHeadersForTags []string `protobuf:"bytes,1,rep,name=request_headers_for_tags,json=requestHeadersForTags,proto3" json:"request_headers_for_tags,omitempty"`
	// Optional. If true, Envoy will include logs for streaming events. Default: false.
	Verbose bool `protobuf:"varint,2,opt,name=verbose,proto3" json:"verbose,omitempty"`
	// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
	// TracePercentages defines the limits for random, forced, and overall tracing percentages.
	TracePercentages *TracePercentages `protobuf:"bytes,3,opt,name=trace_percentages,json=tracePercentages,proto3" json:"trace_percentages,omitempty"`
	// Optional. If not specified, no tracing will be performed
	// ProviderConfig defines the configuration for an external tracing provider.
	//
	// Types that are valid to be assigned to ProviderConfig:
	//	*ListenerTracingSettings_ZipkinConfig
	//	*ListenerTracingSettings_DatadogConfig
	ProviderConfig       isListenerTracingSettings_ProviderConfig `protobuf_oneof:"provider_config"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *ListenerTracingSettings) Reset()         { *m = ListenerTracingSettings{} }
func (m *ListenerTracingSettings) String() string { return proto.CompactTextString(m) }
func (*ListenerTracingSettings) ProtoMessage()    {}
func (*ListenerTracingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f134b8947c6e68, []int{0}
}
func (m *ListenerTracingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerTracingSettings.Unmarshal(m, b)
}
func (m *ListenerTracingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerTracingSettings.Marshal(b, m, deterministic)
}
func (m *ListenerTracingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerTracingSettings.Merge(m, src)
}
func (m *ListenerTracingSettings) XXX_Size() int {
	return xxx_messageInfo_ListenerTracingSettings.Size(m)
}
func (m *ListenerTracingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerTracingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerTracingSettings proto.InternalMessageInfo

type isListenerTracingSettings_ProviderConfig interface {
	isListenerTracingSettings_ProviderConfig()
	Equal(interface{}) bool
}

type ListenerTracingSettings_ZipkinConfig struct {
	ZipkinConfig *v3.ZipkinConfig `protobuf:"bytes,4,opt,name=zipkin_config,json=zipkinConfig,proto3,oneof" json:"zipkin_config,omitempty"`
}
type ListenerTracingSettings_DatadogConfig struct {
	DatadogConfig *v3.DatadogConfig `protobuf:"bytes,5,opt,name=datadog_config,json=datadogConfig,proto3,oneof" json:"datadog_config,omitempty"`
}

func (*ListenerTracingSettings_ZipkinConfig) isListenerTracingSettings_ProviderConfig()  {}
func (*ListenerTracingSettings_DatadogConfig) isListenerTracingSettings_ProviderConfig() {}

func (m *ListenerTracingSettings) GetProviderConfig() isListenerTracingSettings_ProviderConfig {
	if m != nil {
		return m.ProviderConfig
	}
	return nil
}

func (m *ListenerTracingSettings) GetRequestHeadersForTags() []string {
	if m != nil {
		return m.RequestHeadersForTags
	}
	return nil
}

func (m *ListenerTracingSettings) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

func (m *ListenerTracingSettings) GetTracePercentages() *TracePercentages {
	if m != nil {
		return m.TracePercentages
	}
	return nil
}

func (m *ListenerTracingSettings) GetZipkinConfig() *v3.ZipkinConfig {
	if x, ok := m.GetProviderConfig().(*ListenerTracingSettings_ZipkinConfig); ok {
		return x.ZipkinConfig
	}
	return nil
}

func (m *ListenerTracingSettings) GetDatadogConfig() *v3.DatadogConfig {
	if x, ok := m.GetProviderConfig().(*ListenerTracingSettings_DatadogConfig); ok {
		return x.DatadogConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListenerTracingSettings) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListenerTracingSettings_ZipkinConfig)(nil),
		(*ListenerTracingSettings_DatadogConfig)(nil),
	}
}

// Contains settings for configuring Envoy's tracing capabilities at the route level.
// Note: must also specify ListenerTracingSettings for the associated listener.
// See here for additional information on Envoy's tracing capabilities: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/tracing.html
// See here for additional information about configuring tracing with Gloo: https://gloo.solo.io/user_guides/setup_options/observability/#tracing
type RouteTracingSettings struct {
	// Optional. If set, will be used to identify the route that produced the trace.
	// Note that this value will be overridden if the "x-envoy-decorator-operation" header is passed.
	RouteDescriptor string `protobuf:"bytes,1,opt,name=route_descriptor,json=routeDescriptor,proto3" json:"route_descriptor,omitempty"`
	// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
	// TracePercentages defines the limits for random, forced, and overall tracing percentages.
	TracePercentages *TracePercentages `protobuf:"bytes,2,opt,name=trace_percentages,json=tracePercentages,proto3" json:"trace_percentages,omitempty"`
	// Optional. Default is true, If set to false, the tracing headers will not propagate to the upstream.
	Propagate            *types.BoolValue `protobuf:"bytes,3,opt,name=propagate,proto3" json:"propagate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RouteTracingSettings) Reset()         { *m = RouteTracingSettings{} }
func (m *RouteTracingSettings) String() string { return proto.CompactTextString(m) }
func (*RouteTracingSettings) ProtoMessage()    {}
func (*RouteTracingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f134b8947c6e68, []int{1}
}
func (m *RouteTracingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTracingSettings.Unmarshal(m, b)
}
func (m *RouteTracingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTracingSettings.Marshal(b, m, deterministic)
}
func (m *RouteTracingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTracingSettings.Merge(m, src)
}
func (m *RouteTracingSettings) XXX_Size() int {
	return xxx_messageInfo_RouteTracingSettings.Size(m)
}
func (m *RouteTracingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTracingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTracingSettings proto.InternalMessageInfo

func (m *RouteTracingSettings) GetRouteDescriptor() string {
	if m != nil {
		return m.RouteDescriptor
	}
	return ""
}

func (m *RouteTracingSettings) GetTracePercentages() *TracePercentages {
	if m != nil {
		return m.TracePercentages
	}
	return nil
}

func (m *RouteTracingSettings) GetPropagate() *types.BoolValue {
	if m != nil {
		return m.Propagate
	}
	return nil
}

// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
// TracePercentages defines the limits for random, forced, and overall tracing percentages.
type TracePercentages struct {
	// Percentage of requests that should produce traces when the `x-client-trace-id` header is provided.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	ClientSamplePercentage *types.FloatValue `protobuf:"bytes,1,opt,name=client_sample_percentage,json=clientSamplePercentage,proto3" json:"client_sample_percentage,omitempty"`
	// Percentage of requests that should produce traces by random sampling.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	RandomSamplePercentage *types.FloatValue `protobuf:"bytes,2,opt,name=random_sample_percentage,json=randomSamplePercentage,proto3" json:"random_sample_percentage,omitempty"`
	// Overall percentage of requests that should produce traces.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	OverallSamplePercentage *types.FloatValue `protobuf:"bytes,3,opt,name=overall_sample_percentage,json=overallSamplePercentage,proto3" json:"overall_sample_percentage,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}          `json:"-"`
	XXX_unrecognized        []byte            `json:"-"`
	XXX_sizecache           int32             `json:"-"`
}

func (m *TracePercentages) Reset()         { *m = TracePercentages{} }
func (m *TracePercentages) String() string { return proto.CompactTextString(m) }
func (*TracePercentages) ProtoMessage()    {}
func (*TracePercentages) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f134b8947c6e68, []int{2}
}
func (m *TracePercentages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TracePercentages.Unmarshal(m, b)
}
func (m *TracePercentages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TracePercentages.Marshal(b, m, deterministic)
}
func (m *TracePercentages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TracePercentages.Merge(m, src)
}
func (m *TracePercentages) XXX_Size() int {
	return xxx_messageInfo_TracePercentages.Size(m)
}
func (m *TracePercentages) XXX_DiscardUnknown() {
	xxx_messageInfo_TracePercentages.DiscardUnknown(m)
}

var xxx_messageInfo_TracePercentages proto.InternalMessageInfo

func (m *TracePercentages) GetClientSamplePercentage() *types.FloatValue {
	if m != nil {
		return m.ClientSamplePercentage
	}
	return nil
}

func (m *TracePercentages) GetRandomSamplePercentage() *types.FloatValue {
	if m != nil {
		return m.RandomSamplePercentage
	}
	return nil
}

func (m *TracePercentages) GetOverallSamplePercentage() *types.FloatValue {
	if m != nil {
		return m.OverallSamplePercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ListenerTracingSettings)(nil), "tracing.options.gloo.solo.io.ListenerTracingSettings")
	proto.RegisterType((*RouteTracingSettings)(nil), "tracing.options.gloo.solo.io.RouteTracingSettings")
	proto.RegisterType((*TracePercentages)(nil), "tracing.options.gloo.solo.io.TracePercentages")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/tracing/tracing.proto", fileDescriptor_30f134b8947c6e68)
}

var fileDescriptor_30f134b8947c6e68 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xd1, 0x4e, 0xd4, 0x4c,
	0x14, 0xc7, 0xbf, 0x2e, 0x7c, 0x2a, 0x83, 0x08, 0x34, 0x28, 0x65, 0x35, 0x64, 0x83, 0x5e, 0xac,
	0x17, 0xce, 0x44, 0xb8, 0xd0, 0x4b, 0xb3, 0x10, 0x42, 0x88, 0x26, 0xa6, 0xa0, 0x26, 0x78, 0xd1,
	0xcc, 0xb6, 0x67, 0x87, 0x91, 0xd2, 0x33, 0xce, 0xcc, 0x56, 0xe0, 0x4d, 0x7c, 0x03, 0x1f, 0xc1,
	0x47, 0xf1, 0xda, 0x77, 0x30, 0xde, 0x9a, 0xce, 0xb4, 0x40, 0x96, 0x5d, 0x25, 0xc6, 0xab, 0xf6,
	0xfc, 0xcf, 0x39, 0xbf, 0x33, 0xfd, 0x9f, 0x66, 0xc8, 0xae, 0x90, 0xf6, 0x70, 0xd8, 0xa7, 0x29,
	0x1e, 0x33, 0x83, 0x39, 0x3e, 0x91, 0xc8, 0x44, 0x8e, 0xc8, 0x94, 0xc6, 0x0f, 0x90, 0x5a, 0xe3,
	0x23, 0xae, 0x24, 0x2b, 0x9f, 0x32, 0x54, 0x56, 0x62, 0x61, 0x98, 0xd5, 0x3c, 0x95, 0x85, 0x68,
	0x9e, 0x54, 0x69, 0xb4, 0x18, 0x3e, 0x68, 0xc2, 0xba, 0x8c, 0x56, 0xad, 0xb4, 0xa2, 0x52, 0x89,
	0xed, 0xde, 0x04, 0x2c, 0x9c, 0x58, 0xd0, 0x05, 0xcf, 0x19, 0x14, 0x25, 0x9e, 0xb2, 0x14, 0x8b,
	0x81, 0xf4, 0x64, 0x60, 0xe5, 0x06, 0x3b, 0x93, 0xea, 0x48, 0x16, 0x7e, 0x42, 0x7b, 0xf3, 0xef,
	0x18, 0x19, 0xb7, 0x3c, 0xc3, 0xfa, 0x98, 0xed, 0x25, 0x81, 0x02, 0xdd, 0x2b, 0xab, 0xde, 0x6a,
	0x75, 0x55, 0x20, 0x8a, 0x1c, 0x98, 0x8b, 0xfa, 0xc3, 0x01, 0xfb, 0xa4, 0xb9, 0x52, 0xa0, 0xcd,
	0xa4, 0x7c, 0x36, 0xd4, 0xbc, 0xfa, 0xcc, 0x3a, 0xbf, 0x32, 0x9a, 0xe7, 0xc5, 0x69, 0x9d, 0x0a,
	0xe1, 0xc4, 0xfa, 0x79, 0x70, 0x62, 0x9b, 0x72, 0x67, 0xf6, 0x91, 0xb4, 0x8d, 0xb5, 0x1a, 0x06,
	0x3e, 0xb5, 0xf6, 0xb3, 0x45, 0x96, 0x5f, 0x4a, 0x63, 0xa1, 0x00, 0xbd, 0xef, 0x1d, 0xdd, 0x03,
	0x6b, 0x65, 0x21, 0x4c, 0xf8, 0x8c, 0x44, 0x1a, 0x3e, 0x0e, 0xc1, 0xd8, 0xe4, 0x10, 0x78, 0x06,
	0xda, 0x24, 0x03, 0xd4, 0x89, 0xe5, 0xc2, 0x44, 0x41, 0x67, 0xaa, 0x3b, 0x13, 0xdf, 0xad, 0xf3,
	0x3b, 0x3e, 0xbd, 0x8d, 0x7a, 0x9f, 0x0b, 0x13, 0x46, 0xe4, 0x66, 0x09, 0xba, 0x8f, 0x06, 0xa2,
	0x56, 0x27, 0xe8, 0xde, 0x8a, 0x9b, 0x30, 0x7c, 0x4f, 0x16, 0x9d, 0x51, 0x89, 0x02, 0x9d, 0x42,
	0x61, 0xb9, 0x00, 0x13, 0x4d, 0x75, 0x82, 0xee, 0xec, 0x3a, 0xa5, 0xbf, 0xdb, 0x28, 0xad, 0x0e,
	0x07, 0xaf, 0x2f, 0xba, 0xe2, 0x05, 0x3b, 0xa2, 0x84, 0xbb, 0x64, 0xce, 0x2f, 0x30, 0xf1, 0x3b,
	0x89, 0xa6, 0x1d, 0xf8, 0x21, 0x75, 0x8b, 0xa2, 0x5e, 0x74, 0x53, 0x80, 0x96, 0x1b, 0xf4, 0xc0,
	0xd5, 0x6e, 0x3a, 0x75, 0xe7, 0xbf, 0xf8, 0xf6, 0xd9, 0xa5, 0x38, 0x7c, 0x45, 0xee, 0xd4, 0x8b,
	0x6c, 0x60, 0xff, 0x3b, 0xd8, 0xa3, 0x09, 0xb0, 0x2d, 0x5f, 0x7c, 0x4e, 0x9b, 0xcb, 0x2e, 0x0b,
	0xbd, 0x45, 0x32, 0xaf, 0x34, 0x96, 0x32, 0x03, 0x5d, 0xf3, 0xd6, 0xbe, 0x05, 0x64, 0x29, 0xc6,
	0xa1, 0x85, 0x51, 0xdb, 0x1f, 0x93, 0x05, 0x5d, 0xe9, 0x49, 0x06, 0x26, 0xd5, 0x52, 0x59, 0xd4,
	0x51, 0xd0, 0x09, 0xba, 0x33, 0xf1, 0xbc, 0xd3, 0xb7, 0xce, 0xe5, 0xf1, 0x76, 0xb6, 0xfe, 0x91,
	0x9d, 0xcf, 0xc9, 0x8c, 0xd2, 0xa8, 0xb8, 0xe0, 0x16, 0xea, 0x1d, 0xb5, 0xa9, 0xff, 0xf1, 0x68,
	0xf3, 0xe3, 0xd1, 0x1e, 0x62, 0xfe, 0x96, 0xe7, 0x43, 0x88, 0x2f, 0x8a, 0xd7, 0x3e, 0xb7, 0xc8,
	0xc2, 0xe8, 0x80, 0xf0, 0x0d, 0x89, 0xd2, 0x5c, 0x42, 0x61, 0x13, 0xc3, 0x8f, 0x55, 0x7e, 0xf9,
	0xcc, 0xee, 0xf3, 0x66, 0xd7, 0xef, 0x5f, 0xa1, 0x6f, 0xe7, 0xc8, 0xad, 0xc7, 0xdf, 0xf3, 0xcd,
	0x7b, 0xae, 0xf7, 0x82, 0x5b, 0x61, 0x35, 0x2f, 0x32, 0x3c, 0x1e, 0x83, 0x6d, 0x5d, 0x03, 0xeb,
	0x9b, 0xaf, 0x60, 0xdf, 0x91, 0x15, 0x2c, 0x41, 0xf3, 0x3c, 0x1f, 0xc3, 0x9d, 0xfa, 0x33, 0x77,
	0xb9, 0xee, 0x1e, 0x05, 0xf7, 0x76, 0xbf, 0xfe, 0x98, 0x0e, 0xbe, 0x7c, 0x5f, 0x0d, 0x0e, 0x5e,
	0x5c, 0xef, 0x36, 0x54, 0x47, 0x62, 0xc2, 0x8d, 0xd8, 0xbf, 0xe1, 0x26, 0x6f, 0xfc, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0xc8, 0x97, 0xf4, 0x4b, 0x58, 0x05, 0x00, 0x00,
}

func (this *ListenerTracingSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListenerTracingSettings)
	if !ok {
		that2, ok := that.(ListenerTracingSettings)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.RequestHeadersForTags) != len(that1.RequestHeadersForTags) {
		return false
	}
	for i := range this.RequestHeadersForTags {
		if this.RequestHeadersForTags[i] != that1.RequestHeadersForTags[i] {
			return false
		}
	}
	if this.Verbose != that1.Verbose {
		return false
	}
	if !this.TracePercentages.Equal(that1.TracePercentages) {
		return false
	}
	if that1.ProviderConfig == nil {
		if this.ProviderConfig != nil {
			return false
		}
	} else if this.ProviderConfig == nil {
		return false
	} else if !this.ProviderConfig.Equal(that1.ProviderConfig) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ListenerTracingSettings_ZipkinConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListenerTracingSettings_ZipkinConfig)
	if !ok {
		that2, ok := that.(ListenerTracingSettings_ZipkinConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ZipkinConfig.Equal(that1.ZipkinConfig) {
		return false
	}
	return true
}
func (this *ListenerTracingSettings_DatadogConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListenerTracingSettings_DatadogConfig)
	if !ok {
		that2, ok := that.(ListenerTracingSettings_DatadogConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.DatadogConfig.Equal(that1.DatadogConfig) {
		return false
	}
	return true
}
func (this *RouteTracingSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteTracingSettings)
	if !ok {
		that2, ok := that.(RouteTracingSettings)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.RouteDescriptor != that1.RouteDescriptor {
		return false
	}
	if !this.TracePercentages.Equal(that1.TracePercentages) {
		return false
	}
	if !this.Propagate.Equal(that1.Propagate) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TracePercentages) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TracePercentages)
	if !ok {
		that2, ok := that.(TracePercentages)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ClientSamplePercentage.Equal(that1.ClientSamplePercentage) {
		return false
	}
	if !this.RandomSamplePercentage.Equal(that1.RandomSamplePercentage) {
		return false
	}
	if !this.OverallSamplePercentage.Equal(that1.OverallSamplePercentage) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
