//go:build !no_runtime_type_checking

package awssnssubscriptions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

func (s *jsiiProxy_SqsSubscription) validateBindParameters(topic awssns.ITopic) error {
	if topic == nil {
		return fmt.Errorf("parameter topic is required, but nil was provided")
	}

	return nil
}

func validateNewSqsSubscriptionParameters(queue awssqs.IQueue, props *SqsSubscriptionProps) error {
	if queue == nil {
		return fmt.Errorf("parameter queue is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

