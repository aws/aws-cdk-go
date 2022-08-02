package awsautoscalinghooktargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscalinghooktargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
	"github.com/aws/constructs-go/constructs/v3"
)

// Use a Lambda Function as a hook target.
//
// Internally creates a Topic to make the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//   var key key
//
//   functionHook := awscdk.Aws_autoscaling_hooktargets.NewFunctionHook(function_, key)
//
// Experimental.
type FunctionHook interface {
	awsautoscaling.ILifecycleHookTarget
	// If the `IRole` does not exist in `options`, will create an `IRole` and an SNS Topic and attach both to the lifecycle hook.
	//
	// If the `IRole` does exist in `options`, will only create an SNS Topic and attach it to the lifecycle hook.
	// Experimental.
	Bind(_scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) *awsautoscaling.LifecycleHookTargetConfig
}

// The jsii proxy struct for FunctionHook
type jsiiProxy_FunctionHook struct {
	internal.Type__awsautoscalingILifecycleHookTarget
}

// Experimental.
func NewFunctionHook(fn awslambda.IFunction, encryptionKey awskms.IKey) FunctionHook {
	_init_.Initialize()

	j := jsiiProxy_FunctionHook{}

	_jsii_.Create(
		"monocdk.aws_autoscaling_hooktargets.FunctionHook",
		[]interface{}{fn, encryptionKey},
		&j,
	)

	return &j
}

// Experimental.
func NewFunctionHook_Override(f FunctionHook, fn awslambda.IFunction, encryptionKey awskms.IKey) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_autoscaling_hooktargets.FunctionHook",
		[]interface{}{fn, encryptionKey},
		f,
	)
}

func (f *jsiiProxy_FunctionHook) Bind(_scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) *awsautoscaling.LifecycleHookTargetConfig {
	var returns *awsautoscaling.LifecycleHookTargetConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

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

// Use an SNS topic as a hook target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic topic
//
//   topicHook := awscdk.Aws_autoscaling_hooktargets.NewTopicHook(topic)
//
// Experimental.
type TopicHook interface {
	awsautoscaling.ILifecycleHookTarget
	// If an `IRole` is found in `options`, grant it topic publishing permissions.
	//
	// Otherwise, create a new `IRole` and grant it topic publishing permissions.
	//
	// Returns: the `IRole` with topic publishing permissions and the ARN of the topic it has publishing permission to.
	// Experimental.
	Bind(_scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) *awsautoscaling.LifecycleHookTargetConfig
}

// The jsii proxy struct for TopicHook
type jsiiProxy_TopicHook struct {
	internal.Type__awsautoscalingILifecycleHookTarget
}

// Experimental.
func NewTopicHook(topic awssns.ITopic) TopicHook {
	_init_.Initialize()

	j := jsiiProxy_TopicHook{}

	_jsii_.Create(
		"monocdk.aws_autoscaling_hooktargets.TopicHook",
		[]interface{}{topic},
		&j,
	)

	return &j
}

// Experimental.
func NewTopicHook_Override(t TopicHook, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_autoscaling_hooktargets.TopicHook",
		[]interface{}{topic},
		t,
	)
}

func (t *jsiiProxy_TopicHook) Bind(_scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) *awsautoscaling.LifecycleHookTargetConfig {
	var returns *awsautoscaling.LifecycleHookTargetConfig

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

