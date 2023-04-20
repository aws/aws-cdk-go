//go:build !no_runtime_type_checking

package awsevents

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (e *jsiiProxy_EventField) validateResolveParameters(_ctx awscdk.IResolveContext) error {
	if _ctx == nil {
		return fmt.Errorf("parameter _ctx is required, but nil was provided")
	}

	return nil
}

func validateEventField_FromPathParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

