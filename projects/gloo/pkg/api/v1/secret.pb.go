// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/secret.proto

package v1

import (
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

//
//Certain features such as the AWS Lambda option require the use of secrets for authentication, configuration of SSL Certificates, and other data that should not be stored in plaintext configuration.
//
//Gloo runs an independent (goroutine) controller to monitor secrets. Secrets are stored in their own secret storage layer. Gloo can monitor secrets stored in the following secret storage services:
//
//- Kubernetes Secrets
//- Hashicorp Vault
//- Plaintext files (recommended only for testing)
//- Secrets must adhere to a structure, specified by the option that requires them.
//
//Gloo's secret backend can be configured in Gloo's bootstrap options
type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//	*Secret_Aws
	//	*Secret_Azure
	//	*Secret_Tls
	//	*Secret_Oauth
	//	*Secret_ApiKey
	//	*Secret_Header
	//	*Secret_Extensions
	Kind isSecret_Kind `protobuf_oneof:"kind"`
	// Metadata contains the object metadata for this resource
	Metadata *core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDescGZIP(), []int{0}
}

func (m *Secret) GetKind() isSecret_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Secret) GetAws() *AwsSecret {
	if x, ok := x.GetKind().(*Secret_Aws); ok {
		return x.Aws
	}
	return nil
}

func (x *Secret) GetAzure() *AzureSecret {
	if x, ok := x.GetKind().(*Secret_Azure); ok {
		return x.Azure
	}
	return nil
}

func (x *Secret) GetTls() *TlsSecret {
	if x, ok := x.GetKind().(*Secret_Tls); ok {
		return x.Tls
	}
	return nil
}

func (x *Secret) GetOauth() *v1.OauthSecret {
	if x, ok := x.GetKind().(*Secret_Oauth); ok {
		return x.Oauth
	}
	return nil
}

func (x *Secret) GetApiKey() *v1.ApiKeySecret {
	if x, ok := x.GetKind().(*Secret_ApiKey); ok {
		return x.ApiKey
	}
	return nil
}

func (x *Secret) GetHeader() *HeaderSecret {
	if x, ok := x.GetKind().(*Secret_Header); ok {
		return x.Header
	}
	return nil
}

func (x *Secret) GetExtensions() *Extensions {
	if x, ok := x.GetKind().(*Secret_Extensions); ok {
		return x.Extensions
	}
	return nil
}

func (x *Secret) GetMetadata() *core.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type isSecret_Kind interface {
	isSecret_Kind()
}

type Secret_Aws struct {
	// AWS credentials
	Aws *AwsSecret `protobuf:"bytes,1,opt,name=aws,proto3,oneof"`
}

type Secret_Azure struct {
	// Azure credentials
	Azure *AzureSecret `protobuf:"bytes,2,opt,name=azure,proto3,oneof"`
}

type Secret_Tls struct {
	// TLS secret specification
	Tls *TlsSecret `protobuf:"bytes,3,opt,name=tls,proto3,oneof"`
}

type Secret_Oauth struct {
	// Enterprise-only: OAuth secret configuration
	Oauth *v1.OauthSecret `protobuf:"bytes,5,opt,name=oauth,proto3,oneof"`
}

type Secret_ApiKey struct {
	// Enterprise-only: ApiKey secret configuration
	ApiKey *v1.ApiKeySecret `protobuf:"bytes,6,opt,name=api_key,json=apiKey,proto3,oneof"`
}

type Secret_Header struct {
	// Secrets for use in header payloads (e.g. in the Envoy healthcheck API)
	Header *HeaderSecret `protobuf:"bytes,8,opt,name=header,proto3,oneof"`
}

type Secret_Extensions struct {
	// Extensions will be passed along from Listeners, Gateways, VirtualServices, Routes, and Route tables to the
	// underlying Proxy, making them useful for controllers, validation tools, etc. which interact with kubernetes yaml.
	//
	// Some sample use cases:
	// * controllers, deployment pipelines, helm charts, etc. which wish to use extensions as a kind of opaque metadata.
	// * In the future, Gloo may support gRPC-based plugins which communicate with the Gloo translator out-of-process.
	// Opaque Extensions enables development of out-of-process plugins without requiring recompiling & redeploying Gloo's API.
	Extensions *Extensions `protobuf:"bytes,4,opt,name=extensions,proto3,oneof"`
}

func (*Secret_Aws) isSecret_Kind() {}

func (*Secret_Azure) isSecret_Kind() {}

func (*Secret_Tls) isSecret_Kind() {}

func (*Secret_Oauth) isSecret_Kind() {}

func (*Secret_ApiKey) isSecret_Kind() {}

func (*Secret_Header) isSecret_Kind() {}

func (*Secret_Extensions) isSecret_Kind() {}

//
//
//There are two ways of providing AWS secrets:
//
//- Method 1: `glooctl create secret aws`
//
// ```
// glooctl create secret aws --name aws-secret-from-glooctl \
//     --namespace default \
//     --access-key $ACC \
//     --secret-key $SEC
// ```
//
//will produce a Kubernetes resource similar to this (note the `aws` field and `resource_kind` annotation):
//
// ```
// apiVersion: v1
// data:
//   aws: base64EncodedStringForMachineConsumption
// kind: Secret
// metadata:
//   annotations:
//     resource_kind: '*v1.Secret'
//   creationTimestamp: "2019-08-23T15:10:20Z"
//   name: aws-secret-from-glooctl
//   namespace: default
//   resourceVersion: "592637"
//   selfLink: /api/v1/namespaces/default/secrets/secret-e2e
//   uid: 1f8c147f-c5b8-11e9-bbf3-42010a8001bc
// type: Opaque
// ```
//
// - Method 2: `kubectl apply -f resource-file.yaml`
//   - If using a git-ops flow, or otherwise creating secrets from yaml files, you may prefer to provide AWS credentials
//   using the format below, with `aws_access_key_id` and `aws_secret_access_key` fields.
//   - This circumvents the need for the annotation, which are not supported by some tools such as
//   [godaddy/kubernetes-external-secrets](https://github.com/godaddy/kubernetes-external-secrets)
//
// ```yaml
// # a sample aws secret resource-file.yaml
// apiVersion: v1
// data:
//   aws_access_key_id: some-id
//   aws_secret_access_key: some-secret
// kind: Secret
// metadata:
//   name: aws-secret-abcd
//   namespace: default
// ```
//
type AwsSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// provided by `glooctl create secret aws`
	AccessKey string `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// provided by `glooctl create secret aws`
	SecretKey string `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	// provided by `glooctl create secret aws`
	SessionToken string `protobuf:"bytes,3,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
}

func (x *AwsSecret) Reset() {
	*x = AwsSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsSecret) ProtoMessage() {}

func (x *AwsSecret) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsSecret.ProtoReflect.Descriptor instead.
func (*AwsSecret) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDescGZIP(), []int{1}
}

func (x *AwsSecret) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *AwsSecret) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *AwsSecret) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type AzureSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// provided by `glooctl create secret azure`
	ApiKeys map[string]string `protobuf:"bytes,1,rep,name=api_keys,json=apiKeys,proto3" json:"api_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AzureSecret) Reset() {
	*x = AzureSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AzureSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureSecret) ProtoMessage() {}

func (x *AzureSecret) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureSecret.ProtoReflect.Descriptor instead.
func (*AzureSecret) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDescGZIP(), []int{2}
}

func (x *AzureSecret) GetApiKeys() map[string]string {
	if x != nil {
		return x.ApiKeys
	}
	return nil
}

type TlsSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// provided by `glooctl create secret tls`
	CertChain string `protobuf:"bytes,1,opt,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	// provided by `glooctl create secret tls`
	PrivateKey string `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// provided by `glooctl create secret tls`
	RootCa string `protobuf:"bytes,3,opt,name=root_ca,json=rootCa,proto3" json:"root_ca,omitempty"`
}

func (x *TlsSecret) Reset() {
	*x = TlsSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TlsSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TlsSecret) ProtoMessage() {}

func (x *TlsSecret) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TlsSecret.ProtoReflect.Descriptor instead.
func (*TlsSecret) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDescGZIP(), []int{3}
}

func (x *TlsSecret) GetCertChain() string {
	if x != nil {
		return x.CertChain
	}
	return ""
}

func (x *TlsSecret) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *TlsSecret) GetRootCa() string {
	if x != nil {
		return x.RootCa
	}
	return ""
}

type HeaderSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A collection of header name to header value mappings, each representing an additional header that could be added to a request.
	// Provided by `glooctl create secret header`
	Headers map[string]string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HeaderSecret) Reset() {
	*x = HeaderSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderSecret) ProtoMessage() {}

func (x *HeaderSecret) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderSecret.ProtoReflect.Descriptor instead.
func (*HeaderSecret) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDescGZIP(), []int{4}
}

func (x *HeaderSecret) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x58, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78,
	0x74, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x03, 0x0a,
	0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2b, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x48, 0x00, 0x52,
	0x03, 0x61, 0x77, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x48, 0x00,
	0x52, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x6c, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x48, 0x00, 0x52,
	0x03, 0x74, 0x6c, 0x73, 0x12, 0x3c, 0x0a, 0x05, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x61,
	0x75, 0x74, 0x68, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x12, 0x40, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x61, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x16, 0x82, 0xf1, 0x04, 0x05,
	0x0a, 0x03, 0x73, 0x65, 0x63, 0x82, 0xf1, 0x04, 0x09, 0x12, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x6e, 0x0a, 0x09, 0x41, 0x77,
	0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8c, 0x01, 0x0a, 0x0b, 0x41,
	0x7a, 0x75, 0x72, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x41, 0x0a, 0x08, 0x61, 0x70,
	0x69, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x7a, 0x75, 0x72,
	0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x1a, 0x3a, 0x0a,
	0x0c, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x64, 0x0a, 0x09, 0x54, 0x6c, 0x73,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x65, 0x72, 0x74,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x63,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x74, 0x43, 0x61, 0x22,
	0x8d, 0x01, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x12, 0x41, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x3e, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_goTypes = []interface{}{
	(*Secret)(nil),          // 0: gloo.solo.io.Secret
	(*AwsSecret)(nil),       // 1: gloo.solo.io.AwsSecret
	(*AzureSecret)(nil),     // 2: gloo.solo.io.AzureSecret
	(*TlsSecret)(nil),       // 3: gloo.solo.io.TlsSecret
	(*HeaderSecret)(nil),    // 4: gloo.solo.io.HeaderSecret
	nil,                     // 5: gloo.solo.io.AzureSecret.ApiKeysEntry
	nil,                     // 6: gloo.solo.io.HeaderSecret.HeadersEntry
	(*v1.OauthSecret)(nil),  // 7: enterprise.gloo.solo.io.OauthSecret
	(*v1.ApiKeySecret)(nil), // 8: enterprise.gloo.solo.io.ApiKeySecret
	(*Extensions)(nil),      // 9: gloo.solo.io.Extensions
	(*core.Metadata)(nil),   // 10: core.solo.io.Metadata
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_depIdxs = []int32{
	1,  // 0: gloo.solo.io.Secret.aws:type_name -> gloo.solo.io.AwsSecret
	2,  // 1: gloo.solo.io.Secret.azure:type_name -> gloo.solo.io.AzureSecret
	3,  // 2: gloo.solo.io.Secret.tls:type_name -> gloo.solo.io.TlsSecret
	7,  // 3: gloo.solo.io.Secret.oauth:type_name -> enterprise.gloo.solo.io.OauthSecret
	8,  // 4: gloo.solo.io.Secret.api_key:type_name -> enterprise.gloo.solo.io.ApiKeySecret
	4,  // 5: gloo.solo.io.Secret.header:type_name -> gloo.solo.io.HeaderSecret
	9,  // 6: gloo.solo.io.Secret.extensions:type_name -> gloo.solo.io.Extensions
	10, // 7: gloo.solo.io.Secret.metadata:type_name -> core.solo.io.Metadata
	5,  // 8: gloo.solo.io.AzureSecret.api_keys:type_name -> gloo.solo.io.AzureSecret.ApiKeysEntry
	6,  // 9: gloo.solo.io.HeaderSecret.headers:type_name -> gloo.solo.io.HeaderSecret.HeadersEntry
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_extensions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AwsSecret); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AzureSecret); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TlsSecret); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderSecret); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Secret_Aws)(nil),
		(*Secret_Azure)(nil),
		(*Secret_Tls)(nil),
		(*Secret_Oauth)(nil),
		(*Secret_ApiKey)(nil),
		(*Secret_Header)(nil),
		(*Secret_Extensions)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_secret_proto_depIdxs = nil
}
