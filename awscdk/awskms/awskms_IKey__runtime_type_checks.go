//go:build !no_runtime_type_checking

package awskms

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

func (i *jsiiProxy_IKey) validateAddAliasParameters(alias *string) error {
	if alias == nil {
		return fmt.Errorf("parameter alias is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IKey) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IKey) validateGrantParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IKey) validateGrantDecryptParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IKey) validateGrantEncryptParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IKey) validateGrantEncryptDecryptParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

