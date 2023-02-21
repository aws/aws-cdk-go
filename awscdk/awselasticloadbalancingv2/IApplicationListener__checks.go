//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func (i *jsiiProxy_IApplicationListener) validateAddActionParameters(id *string, props *AddApplicationActionProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IApplicationListener) validateAddCertificatesParameters(id *string, certificates *[]IListenerCertificate) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if certificates == nil {
		return fmt.Errorf("parameter certificates is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IApplicationListener) validateAddTargetGroupsParameters(id *string, props *AddApplicationTargetGroupsProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IApplicationListener) validateAddTargetsParameters(id *string, props *AddApplicationTargetsProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IApplicationListener) validateRegisterConnectableParameters(connectable awsec2.IConnectable, portRange awsec2.Port) error {
	if connectable == nil {
		return fmt.Errorf("parameter connectable is required, but nil was provided")
	}

	if portRange == nil {
		return fmt.Errorf("parameter portRange is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IApplicationListener) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

