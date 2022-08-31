//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (b *jsiiProxy_Backend) validateBindParameters(_scope awscdk.Construct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateBackend_VirtualServiceParameters(virtualService IVirtualService, props *VirtualServiceBackendOptions) error {
	if virtualService == nil {
		return fmt.Errorf("parameter virtualService is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

