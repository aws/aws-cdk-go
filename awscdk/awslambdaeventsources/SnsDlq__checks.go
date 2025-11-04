//go:build !no_runtime_type_checking

package awslambdaeventsources

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

func (s *jsiiProxy_SnsDlq) validateBindParameters(target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if targetHandler == nil {
		return fmt.Errorf("parameter targetHandler is required, but nil was provided")
	}

	return nil
}

func validateNewSnsDlqParameters(topic awssns.ITopic) error {
	if topic == nil {
		return fmt.Errorf("parameter topic is required, but nil was provided")
	}

	return nil
}

