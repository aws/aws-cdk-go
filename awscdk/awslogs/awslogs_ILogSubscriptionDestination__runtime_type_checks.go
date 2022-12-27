//go:build !no_runtime_type_checking

package awslogs

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_ILogSubscriptionDestination) validateBindParameters(scope constructs.Construct, sourceLogGroup ILogGroup) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if sourceLogGroup == nil {
		return fmt.Errorf("parameter sourceLogGroup is required, but nil was provided")
	}

	return nil
}

