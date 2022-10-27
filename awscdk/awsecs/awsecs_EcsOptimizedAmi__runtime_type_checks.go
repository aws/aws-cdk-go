//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (e *jsiiProxy_EcsOptimizedAmi) validateGetImageParameters(scope awscdk.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateNewEcsOptimizedAmiParameters(props *EcsOptimizedAmiProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

