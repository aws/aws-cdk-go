//go:build !no_runtime_type_checking

package awscdkcloud9alpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func validateOwner_AccountRootParameters(accountId *string) error {
	if accountId == nil {
		return fmt.Errorf("parameter accountId is required, but nil was provided")
	}

	return nil
}

func validateOwner_AssumedRoleParameters(accountId *string, roleName *string) error {
	if accountId == nil {
		return fmt.Errorf("parameter accountId is required, but nil was provided")
	}

	if roleName == nil {
		return fmt.Errorf("parameter roleName is required, but nil was provided")
	}

	return nil
}

func validateOwner_FederatedUserParameters(accountId *string, userName *string) error {
	if accountId == nil {
		return fmt.Errorf("parameter accountId is required, but nil was provided")
	}

	if userName == nil {
		return fmt.Errorf("parameter userName is required, but nil was provided")
	}

	return nil
}

func validateOwner_UserParameters(user awsiam.IUser) error {
	if user == nil {
		return fmt.Errorf("parameter user is required, but nil was provided")
	}

	return nil
}

