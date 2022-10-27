//go:build !no_runtime_type_checking

package awsssm

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

func (i *jsiiProxy_IParameter) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IParameter) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

