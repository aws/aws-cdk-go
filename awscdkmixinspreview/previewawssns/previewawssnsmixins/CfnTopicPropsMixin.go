package previewawssnsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssns/previewawssnsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SNS::Topic` resource creates a topic to which notifications can be published.
//
// > One account can create a maximum of 100,000 standard topics and 1,000 FIFO topics. For more information, see [Amazon  endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/sns.html) in the *AWS General Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var archivePolicy interface{}
//   var dataProtectionPolicy interface{}
//
//   cfnTopicPropsMixin := awscdkmixinspreview.Mixins.NewCfnTopicPropsMixin(&CfnTopicMixinProps{
//   	ArchivePolicy: archivePolicy,
//   	ContentBasedDeduplication: jsii.Boolean(false),
//   	DataProtectionPolicy: dataProtectionPolicy,
//   	DeliveryStatusLogging: []interface{}{
//   		&LoggingConfigProperty{
//   			FailureFeedbackRoleArn: jsii.String("failureFeedbackRoleArn"),
//   			Protocol: jsii.String("protocol"),
//   			SuccessFeedbackRoleArn: jsii.String("successFeedbackRoleArn"),
//   			SuccessFeedbackSampleRate: jsii.String("successFeedbackSampleRate"),
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	FifoThroughputScope: jsii.String("fifoThroughputScope"),
//   	FifoTopic: jsii.Boolean(false),
//   	KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	SignatureVersion: jsii.String("signatureVersion"),
//   	Subscription: []interface{}{
//   		&SubscriptionProperty{
//   			Endpoint: jsii.String("endpoint"),
//   			Protocol: jsii.String("protocol"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TopicName: jsii.String("topicName"),
//   	TracingConfig: jsii.String("tracingConfig"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html
//
type CfnTopicPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTopicMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTopicPropsMixin
type jsiiProxy_CfnTopicPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTopicPropsMixin) Props() *CfnTopicMixinProps {
	var returns *CfnTopicMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SNS::Topic`.
func NewCfnTopicPropsMixin(props *CfnTopicMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTopicPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTopicPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTopicPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sns.mixins.CfnTopicPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SNS::Topic`.
func NewCfnTopicPropsMixin_Override(c CfnTopicPropsMixin, props *CfnTopicMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sns.mixins.CfnTopicPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTopicPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTopicPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sns.mixins.CfnTopicPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTopicPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sns.mixins.CfnTopicPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTopicPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTopicPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

