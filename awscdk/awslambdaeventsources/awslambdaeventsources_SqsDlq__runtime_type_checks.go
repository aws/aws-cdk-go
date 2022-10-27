//go:build !no_runtime_type_checking

package awslambdaeventsources

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

func (s *jsiiProxy_SqsDlq) validateBindParameters(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) error {
	if _target == nil {
		return fmt.Errorf("parameter _target is required, but nil was provided")
	}

	if targetHandler == nil {
		return fmt.Errorf("parameter targetHandler is required, but nil was provided")
	}

	return nil
}

func validateNewSqsDlqParameters(queue awssqs.IQueue) error {
	if queue == nil {
		return fmt.Errorf("parameter queue is required, but nil was provided")
	}

	return nil
}

