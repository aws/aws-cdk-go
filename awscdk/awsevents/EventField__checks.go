//go:build !no_runtime_type_checking

package awsevents

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (e *jsiiProxy_EventField) validateResolveParameters(context awscdk.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func validateEventField_FromPathParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

