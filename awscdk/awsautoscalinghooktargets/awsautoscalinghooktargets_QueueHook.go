package awsautoscalinghooktargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscalinghooktargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
	"github.com/aws/constructs-go/constructs/v3"
)

// Use an SQS queue as a hook target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//
//   queueHook := awscdk.Aws_autoscaling_hooktargets.NewQueueHook(queue)
//
// Experimental.
type QueueHook interface {
	awsautoscaling.ILifecycleHookTarget
	// If an `IRole` is found in `options`, grant it access to send messages.
	//
	// Otherwise, create a new `IRole` and grant it access to send messages.
	//
	// Returns: the `IRole` with access to send messages and the ARN of the queue it has access to send messages to.
	// Experimental.
	Bind(_scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) *awsautoscaling.LifecycleHookTargetConfig
}

// The jsii proxy struct for QueueHook
type jsiiProxy_QueueHook struct {
	internal.Type__awsautoscalingILifecycleHookTarget
}

// Experimental.
func NewQueueHook(queue awssqs.IQueue) QueueHook {
	_init_.Initialize()

	j := jsiiProxy_QueueHook{}

	_jsii_.Create(
		"monocdk.aws_autoscaling_hooktargets.QueueHook",
		[]interface{}{queue},
		&j,
	)

	return &j
}

// Experimental.
func NewQueueHook_Override(q QueueHook, queue awssqs.IQueue) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_autoscaling_hooktargets.QueueHook",
		[]interface{}{queue},
		q,
	)
}

func (q *jsiiProxy_QueueHook) Bind(_scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) *awsautoscaling.LifecycleHookTargetConfig {
	var returns *awsautoscaling.LifecycleHookTargetConfig

	_jsii_.Invoke(
		q,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

