//go:build !no_runtime_type_checking

package awsautoscalinghooktargets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (q *jsiiProxy_QueueHook) validateBindParameters(_scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewQueueHookParameters(queue awssqs.IQueue) error {
	if queue == nil {
		return fmt.Errorf("parameter queue is required, but nil was provided")
	}

	return nil
}

