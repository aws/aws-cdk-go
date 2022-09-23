//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_IdentityPool) validateAddRoleMappingsParameters(roleMappings *[]*IdentityPoolRoleMapping) error {
	for idx_1686ee, v := range *roleMappings {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter roleMappings[%#v]", idx_1686ee) }); err != nil {
			return err
		}
	}

	return nil
}

func (i *jsiiProxy_IdentityPool) validateAddUserPoolAuthenticationParameters(userPool IUserPoolAuthenticationProvider) error {
	if userPool == nil {
		return fmt.Errorf("parameter userPool is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IdentityPool) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IdentityPool) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (i *jsiiProxy_IdentityPool) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func validateIdentityPool_FromIdentityPoolArnParameters(scope constructs.Construct, id *string, identityPoolArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if identityPoolArn == nil {
		return fmt.Errorf("parameter identityPoolArn is required, but nil was provided")
	}

	return nil
}

func validateIdentityPool_FromIdentityPoolIdParameters(scope constructs.Construct, id *string, identityPoolId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if identityPoolId == nil {
		return fmt.Errorf("parameter identityPoolId is required, but nil was provided")
	}

	return nil
}

func validateIdentityPool_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateIdentityPool_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateIdentityPool_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewIdentityPoolParameters(scope constructs.Construct, id *string, props *IdentityPoolProps) error {
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

