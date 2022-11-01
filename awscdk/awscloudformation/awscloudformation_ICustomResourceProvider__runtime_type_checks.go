//go:build !no_runtime_type_checking

package awscloudformation

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (i *jsiiProxy_ICustomResourceProvider) validateBindParameters(scope awscdk.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

