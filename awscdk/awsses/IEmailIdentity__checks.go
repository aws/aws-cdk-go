//go:build !no_runtime_type_checking

package awsses

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func (i *jsiiProxy_IEmailIdentity) validateGrantParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IEmailIdentity) validateGrantSendEmailParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

