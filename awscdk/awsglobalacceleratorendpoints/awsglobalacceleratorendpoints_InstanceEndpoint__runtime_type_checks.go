//go:build !no_runtime_type_checking

package awsglobalacceleratorendpoints

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

func validateNewInstanceEndpointParameters(instance awsec2.IInstance, options *InstanceEndpointProps) error {
	if instance == nil {
		return fmt.Errorf("parameter instance is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

