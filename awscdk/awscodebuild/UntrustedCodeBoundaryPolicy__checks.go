//go:build !no_runtime_type_checking

package awscodebuild

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) validateAttachToGroupParameters(group awsiam.IGroup) error {
	if group == nil {
		return fmt.Errorf("parameter group is required, but nil was provided")
	}

	return nil
}

func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) validateAttachToRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) validateAttachToUserParameters(user awsiam.IUser) error {
	if user == nil {
		return fmt.Errorf("parameter user is required, but nil was provided")
	}

	return nil
}

func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func validateUntrustedCodeBoundaryPolicy_FromAwsManagedPolicyNameParameters(managedPolicyName *string) error {
	if managedPolicyName == nil {
		return fmt.Errorf("parameter managedPolicyName is required, but nil was provided")
	}

	return nil
}

func validateUntrustedCodeBoundaryPolicy_FromManagedPolicyArnParameters(scope constructs.Construct, id *string, managedPolicyArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if managedPolicyArn == nil {
		return fmt.Errorf("parameter managedPolicyArn is required, but nil was provided")
	}

	return nil
}

func validateUntrustedCodeBoundaryPolicy_FromManagedPolicyNameParameters(scope constructs.Construct, id *string, managedPolicyName *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if managedPolicyName == nil {
		return fmt.Errorf("parameter managedPolicyName is required, but nil was provided")
	}

	return nil
}

func validateUntrustedCodeBoundaryPolicy_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateUntrustedCodeBoundaryPolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateUntrustedCodeBoundaryPolicy_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewUntrustedCodeBoundaryPolicyParameters(scope constructs.Construct, id *string, props *UntrustedCodeBoundaryPolicyProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

