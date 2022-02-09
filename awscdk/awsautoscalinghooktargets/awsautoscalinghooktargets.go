package awsautoscalinghooktargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscalinghooktargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a Lambda Function as a hook target.
//
// Internally creates a Topic to make the connection.
//
// TODO: EXAMPLE
//
type FunctionHook interface {
	awsautoscaling.ILifecycleHookTarget
	Bind(_scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) *awsautoscaling.LifecycleHookTargetConfig
}

// The jsii proxy struct for FunctionHook
type jsiiProxy_FunctionHook struct {
	internal.Type__awsautoscalingILifecycleHookTarget
}

func NewFunctionHook(fn awslambda.IFunction, encryptionKey awskms.IKey) FunctionHook {
	_init_.Initialize()

	j := jsiiProxy_FunctionHook{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling_hooktargets.FunctionHook",
		[]interface{}{fn, encryptionKey},
		&j,
	)

	return &j
}

func NewFunctionHook_Override(f FunctionHook, fn awslambda.IFunction, encryptionKey awskms.IKey) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling_hooktargets.FunctionHook",
		[]interface{}{fn, encryptionKey},
		f,
	)
}

// If the `IRole` does not exist in `options`, will create an `IRole` and an SNS Topic and attach both to the lifecycle hook.
//
// If the `IRole` does exist in `options`, will only create an SNS Topic and attach it to the lifecycle hook.
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
// TODO: EXAMPLE
//
type QueueHook interface {
	awsautoscaling.ILifecycleHookTarget
	Bind(_scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) *awsautoscaling.LifecycleHookTargetConfig
}

// The jsii proxy struct for QueueHook
type jsiiProxy_QueueHook struct {
	internal.Type__awsautoscalingILifecycleHookTarget
}

func NewQueueHook(queue awssqs.IQueue) QueueHook {
	_init_.Initialize()

	j := jsiiProxy_QueueHook{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling_hooktargets.QueueHook",
		[]interface{}{queue},
		&j,
	)

	return &j
}

func NewQueueHook_Override(q QueueHook, queue awssqs.IQueue) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling_hooktargets.QueueHook",
		[]interface{}{queue},
		q,
	)
}

// If an `IRole` is found in `options`, grant it access to send messages.
//
// Otherwise, create a new `IRole` and grant it access to send messages.
//
// Returns: the `IRole` with access to send messages and the ARN of the queue it has access to send messages to.
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
// TODO: EXAMPLE
//
type TopicHook interface {
	awsautoscaling.ILifecycleHookTarget
	Bind(_scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) *awsautoscaling.LifecycleHookTargetConfig
}

// The jsii proxy struct for TopicHook
type jsiiProxy_TopicHook struct {
	internal.Type__awsautoscalingILifecycleHookTarget
}

func NewTopicHook(topic awssns.ITopic) TopicHook {
	_init_.Initialize()

	j := jsiiProxy_TopicHook{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling_hooktargets.TopicHook",
		[]interface{}{topic},
		&j,
	)

	return &j
}

func NewTopicHook_Override(t TopicHook, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling_hooktargets.TopicHook",
		[]interface{}{topic},
		t,
	)
}

// If an `IRole` is found in `options`, grant it topic publishing permissions.
//
// Otherwise, create a new `IRole` and grant it topic publishing permissions.
//
// Returns: the `IRole` with topic publishing permissions and the ARN of the topic it has publishing permission to.
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

