//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awslogs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (i *jsiiProxy_ILogSubscriptionDestination) validateBindParameters(scope awscdk.Construct, sourceLogGroup ILogGroup) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if sourceLogGroup == nil {
		return fmt.Errorf("parameter sourceLogGroup is required, but nil was provided")
	}

	return nil
}

