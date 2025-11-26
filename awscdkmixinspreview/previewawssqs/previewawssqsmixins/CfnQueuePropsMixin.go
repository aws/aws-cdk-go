package previewawssqsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssqs/previewawssqsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SQS::Queue` resource creates an Amazon SQS standard or FIFO queue.
//
// Keep the following caveats in mind:
//
// - If you don't specify the `FifoQueue` property, Amazon SQS creates a standard queue.
//
// > You can't change the queue type after you create it and you can't convert an existing standard queue into a FIFO queue. You must either create a new FIFO queue for your application or delete your existing standard queue and recreate it as a FIFO queue. For more information, see [Moving from a standard queue to a FIFO queue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues-moving.html) in the *Amazon SQS Developer Guide* .
// - If you don't provide a value for a property, the queue is created with the default value for the property.
// - If you delete a queue, you must wait at least 60 seconds before creating a queue with the same name.
// - To successfully create a new queue, you must provide a queue name that adheres to the [limits related to queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/limits-queues.html) and is unique within the scope of your queues.
//
// For more information about creating FIFO (first-in-first-out) queues, see [Creating an Amazon SQS queue ( CloudFormation )](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/create-queue-cloudformation.html) in the *Amazon SQS Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var redriveAllowPolicy interface{}
//   var redrivePolicy interface{}
//
//   cfnQueuePropsMixin := awscdkmixinspreview.Mixins.NewCfnQueuePropsMixin(&CfnQueueMixinProps{
//   	ContentBasedDeduplication: jsii.Boolean(false),
//   	DeduplicationScope: jsii.String("deduplicationScope"),
//   	DelaySeconds: jsii.Number(123),
//   	FifoQueue: jsii.Boolean(false),
//   	FifoThroughputLimit: jsii.String("fifoThroughputLimit"),
//   	KmsDataKeyReusePeriodSeconds: jsii.Number(123),
//   	KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	MaximumMessageSize: jsii.Number(123),
//   	MessageRetentionPeriod: jsii.Number(123),
//   	QueueName: jsii.String("queueName"),
//   	ReceiveMessageWaitTimeSeconds: jsii.Number(123),
//   	RedriveAllowPolicy: redriveAllowPolicy,
//   	RedrivePolicy: redrivePolicy,
//   	SqsManagedSseEnabled: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VisibilityTimeout: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queue.html
//
type CfnQueuePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnQueueMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnQueuePropsMixin
type jsiiProxy_CfnQueuePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnQueuePropsMixin) Props() *CfnQueueMixinProps {
	var returns *CfnQueueMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueuePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SQS::Queue`.
func NewCfnQueuePropsMixin(props *CfnQueueMixinProps, options *mixins.CfnPropertyMixinOptions) CfnQueuePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnQueuePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnQueuePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sqs.mixins.CfnQueuePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SQS::Queue`.
func NewCfnQueuePropsMixin_Override(c CfnQueuePropsMixin, props *CfnQueueMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sqs.mixins.CfnQueuePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnQueuePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnQueuePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sqs.mixins.CfnQueuePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnQueuePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sqs.mixins.CfnQueuePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnQueuePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnQueuePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

