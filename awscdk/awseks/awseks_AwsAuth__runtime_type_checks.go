//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awseks

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AwsAuth) validateAddAccountParameters(accountId *string) error {
	if accountId == nil {
		return fmt.Errorf("parameter accountId is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsAuth) validateAddMastersRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsAuth) validateAddRoleMappingParameters(role awsiam.IRole, mapping *AwsAuthMapping) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	if mapping == nil {
		return fmt.Errorf("parameter mapping is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(mapping, func() string { return "parameter mapping" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AwsAuth) validateAddUserMappingParameters(user awsiam.IUser, mapping *AwsAuthMapping) error {
	if user == nil {
		return fmt.Errorf("parameter user is required, but nil was provided")
	}

	if mapping == nil {
		return fmt.Errorf("parameter mapping is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(mapping, func() string { return "parameter mapping" }); err != nil {
		return err
	}

	return nil
}

func validateAwsAuth_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewAwsAuthParameters(scope constructs.Construct, id *string, props *AwsAuthProps) error {
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

