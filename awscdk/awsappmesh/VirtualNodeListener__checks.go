//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (v *jsiiProxy_VirtualNodeListener) validateBindParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateVirtualNodeListener_GrpcParameters(props *GrpcVirtualNodeListenerOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateVirtualNodeListener_HttpParameters(props *HttpVirtualNodeListenerOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateVirtualNodeListener_Http2Parameters(props *Http2VirtualNodeListenerOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateVirtualNodeListener_TcpParameters(props *TcpVirtualNodeListenerOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

