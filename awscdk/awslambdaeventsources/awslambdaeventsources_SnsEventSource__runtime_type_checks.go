//go:build !no_runtime_type_checking

package awslambdaeventsources

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

func (s *jsiiProxy_SnsEventSource) validateBindParameters(target awslambda.IFunction) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func validateNewSnsEventSourceParameters(topic awssns.ITopic, props *SnsEventSourceProps) error {
	if topic == nil {
		return fmt.Errorf("parameter topic is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

