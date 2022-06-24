// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/connection.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *ConnectionConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ConnectionConfig)
	if !ok {
		that2, ok := that.(ConnectionConfig)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetMaxRequestsPerConnection() != target.GetMaxRequestsPerConnection() {
		return false
	}

	if h, ok := interface{}(m.GetConnectTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConnectTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConnectTimeout(), target.GetConnectTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTcpKeepalive()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTcpKeepalive()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTcpKeepalive(), target.GetTcpKeepalive()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPerConnectionBufferLimitBytes()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPerConnectionBufferLimitBytes()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPerConnectionBufferLimitBytes(), target.GetPerConnectionBufferLimitBytes()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCommonHttpProtocolOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCommonHttpProtocolOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCommonHttpProtocolOptions(), target.GetCommonHttpProtocolOptions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHttp1ProtocolOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHttp1ProtocolOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHttp1ProtocolOptions(), target.GetHttp1ProtocolOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ConnectionConfig_TcpKeepAlive) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ConnectionConfig_TcpKeepAlive)
	if !ok {
		that2, ok := that.(ConnectionConfig_TcpKeepAlive)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetKeepaliveProbes() != target.GetKeepaliveProbes() {
		return false
	}

	if h, ok := interface{}(m.GetKeepaliveTime()).(equality.Equalizer); ok {
		if !h.Equal(target.GetKeepaliveTime()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetKeepaliveTime(), target.GetKeepaliveTime()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetKeepaliveInterval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetKeepaliveInterval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetKeepaliveInterval(), target.GetKeepaliveInterval()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ConnectionConfig_HttpProtocolOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ConnectionConfig_HttpProtocolOptions)
	if !ok {
		that2, ok := that.(ConnectionConfig_HttpProtocolOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetIdleTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIdleTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIdleTimeout(), target.GetIdleTimeout()) {
			return false
		}
	}

	if m.GetMaxHeadersCount() != target.GetMaxHeadersCount() {
		return false
	}

	if h, ok := interface{}(m.GetMaxStreamDuration()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxStreamDuration()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxStreamDuration(), target.GetMaxStreamDuration()) {
			return false
		}
	}

	if m.GetHeadersWithUnderscoresAction() != target.GetHeadersWithUnderscoresAction() {
		return false
	}

	return true
}

// Equal function
func (m *ConnectionConfig_Http1ProtocolOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ConnectionConfig_Http1ProtocolOptions)
	if !ok {
		that2, ok := that.(ConnectionConfig_Http1ProtocolOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetEnableTrailers() != target.GetEnableTrailers() {
		return false
	}

	if h, ok := interface{}(m.GetOverrideStreamErrorOnInvalidHttpMessage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOverrideStreamErrorOnInvalidHttpMessage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOverrideStreamErrorOnInvalidHttpMessage(), target.GetOverrideStreamErrorOnInvalidHttpMessage()) {
			return false
		}
	}

	switch m.HeaderFormat.(type) {

	case *ConnectionConfig_Http1ProtocolOptions_ProperCaseHeaderKeyFormat:
		if _, ok := target.HeaderFormat.(*ConnectionConfig_Http1ProtocolOptions_ProperCaseHeaderKeyFormat); !ok {
			return false
		}

		if m.GetProperCaseHeaderKeyFormat() != target.GetProperCaseHeaderKeyFormat() {
			return false
		}

	case *ConnectionConfig_Http1ProtocolOptions_PreserveCaseHeaderKeyFormat:
		if _, ok := target.HeaderFormat.(*ConnectionConfig_Http1ProtocolOptions_PreserveCaseHeaderKeyFormat); !ok {
			return false
		}

		if m.GetPreserveCaseHeaderKeyFormat() != target.GetPreserveCaseHeaderKeyFormat() {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.HeaderFormat != target.HeaderFormat {
			return false
		}
	}

	return true
}
