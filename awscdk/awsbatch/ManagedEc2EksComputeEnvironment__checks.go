//go:build !no_runtime_type_checking

package awsbatch

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (m *jsiiProxy_ManagedEc2EksComputeEnvironment) validateAddInstanceClassParameters(instanceClass awsec2.InstanceClass) error {
	if instanceClass == "" {
		return fmt.Errorf("parameter instanceClass is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedEc2EksComputeEnvironment) validateAddInstanceTypeParameters(instanceType awsec2.InstanceType) error {
	if instanceType == nil {
		return fmt.Errorf("parameter instanceType is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedEc2EksComputeEnvironment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedEc2EksComputeEnvironment) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_ManagedEc2EksComputeEnvironment) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func validateManagedEc2EksComputeEnvironment_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateManagedEc2EksComputeEnvironment_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateManagedEc2EksComputeEnvironment_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewManagedEc2EksComputeEnvironmentParameters(scope constructs.Construct, id *string, props *ManagedEc2EksComputeEnvironmentProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

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

