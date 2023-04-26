//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IEcsEc2ContainerDefinition) validateAddUlimitParameters(ulimit *Ulimit) error {
	if ulimit == nil {
		return fmt.Errorf("parameter ulimit is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(ulimit, func() string { return "parameter ulimit" }); err != nil {
		return err
	}

	return nil
}

