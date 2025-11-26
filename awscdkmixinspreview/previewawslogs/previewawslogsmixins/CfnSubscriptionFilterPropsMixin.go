package previewawslogsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs/previewawslogsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Logs::SubscriptionFilter` resource specifies a subscription filter and associates it with the specified log group.
//
// Subscription filters allow you to subscribe to a real-time stream of log events and have them delivered to a specific destination. Currently, the supported destinations are:
//
// - An Amazon Kinesis data stream belonging to the same account as the subscription filter, for same-account delivery.
// - A logical destination that belongs to a different account, for cross-account delivery.
// - An Amazon Kinesis Firehose delivery stream that belongs to the same account as the subscription filter, for same-account delivery.
// - An AWS Lambda function that belongs to the same account as the subscription filter, for same-account delivery.
//
// There can be as many as two subscription filters associated with a log group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSubscriptionFilterPropsMixin := awscdkmixinspreview.Mixins.NewCfnSubscriptionFilterPropsMixin(&CfnSubscriptionFilterMixinProps{
//   	ApplyOnTransformedLogs: jsii.Boolean(false),
//   	DestinationArn: jsii.String("destinationArn"),
//   	Distribution: jsii.String("distribution"),
//   	EmitSystemFields: []*string{
//   		jsii.String("emitSystemFields"),
//   	},
//   	FieldSelectionCriteria: jsii.String("fieldSelectionCriteria"),
//   	FilterName: jsii.String("filterName"),
//   	FilterPattern: jsii.String("filterPattern"),
//   	LogGroupName: jsii.String("logGroupName"),
//   	RoleArn: jsii.String("roleArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html
//
type CfnSubscriptionFilterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSubscriptionFilterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSubscriptionFilterPropsMixin
type jsiiProxy_CfnSubscriptionFilterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSubscriptionFilterPropsMixin) Props() *CfnSubscriptionFilterMixinProps {
	var returns *CfnSubscriptionFilterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionFilterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Logs::SubscriptionFilter`.
func NewCfnSubscriptionFilterPropsMixin(props *CfnSubscriptionFilterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSubscriptionFilterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSubscriptionFilterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSubscriptionFilterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnSubscriptionFilterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Logs::SubscriptionFilter`.
func NewCfnSubscriptionFilterPropsMixin_Override(c CfnSubscriptionFilterPropsMixin, props *CfnSubscriptionFilterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnSubscriptionFilterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSubscriptionFilterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSubscriptionFilterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnSubscriptionFilterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSubscriptionFilterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnSubscriptionFilterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSubscriptionFilterPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSubscriptionFilterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

