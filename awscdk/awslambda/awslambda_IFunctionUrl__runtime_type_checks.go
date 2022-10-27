//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

func (i *jsiiProxy_IFunctionUrl) validateGrantInvokeUrlParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

