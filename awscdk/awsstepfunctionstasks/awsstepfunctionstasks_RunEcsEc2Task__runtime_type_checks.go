//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

func (r *jsiiProxy_RunEcsEc2Task) validateBindParameters(task awsstepfunctions.Task) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RunEcsEc2Task) validateConfigureAwsVpcNetworkingParameters(vpc awsec2.IVpc, subnetSelection *awsec2.SubnetSelection) error {
	if vpc == nil {
		return fmt.Errorf("parameter vpc is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(subnetSelection, func() string { return "parameter subnetSelection" }); err != nil {
		return err
	}

	return nil
}

func validateNewRunEcsEc2TaskParameters(props *RunEcsEc2TaskProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

