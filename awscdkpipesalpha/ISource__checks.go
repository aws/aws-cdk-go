//go:build !no_runtime_type_checking

package awscdkpipesalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func (i *jsiiProxy_ISource) validateBindParameters(pipe IPipe) error {
	if pipe == nil {
		return fmt.Errorf("parameter pipe is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ISource) validateGrantReadParameters(grantee awsiam.IRole) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

