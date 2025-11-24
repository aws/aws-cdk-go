//go:build !no_runtime_type_checking

package awslambdaeventsources

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

func (k *jsiiProxy_KafkaDlq) validateBindParameters(target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if targetHandler == nil {
		return fmt.Errorf("parameter targetHandler is required, but nil was provided")
	}

	return nil
}

func validateNewKafkaDlqParameters(topicName *string) error {
	if topicName == nil {
		return fmt.Errorf("parameter topicName is required, but nil was provided")
	}

	return nil
}

