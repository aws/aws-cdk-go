package awssns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssns/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SNS::TopicInlinePolicy` resource associates one Amazon  topic with one policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var policyDocument interface{}
//
//   cfnTopicInlinePolicyPropsMixin := awscdkcfnpropertymixins.Aws_sns.NewCfnTopicInlinePolicyPropsMixin(&CfnTopicInlinePolicyMixinProps{
//   	PolicyDocument: policyDocument,
//   	TopicArn: jsii.String("topicArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topicinlinepolicy.html
//
type CfnTopicInlinePolicyPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTopicInlinePolicyMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTopicInlinePolicyPropsMixin
type jsiiProxy_CfnTopicInlinePolicyPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTopicInlinePolicyPropsMixin) Props() *CfnTopicInlinePolicyMixinProps {
	var returns *CfnTopicInlinePolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopicInlinePolicyPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SNS::TopicInlinePolicy`.
func NewCfnTopicInlinePolicyPropsMixin(props *CfnTopicInlinePolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTopicInlinePolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTopicInlinePolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTopicInlinePolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sns.CfnTopicInlinePolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SNS::TopicInlinePolicy`.
func NewCfnTopicInlinePolicyPropsMixin_Override(c CfnTopicInlinePolicyPropsMixin, props *CfnTopicInlinePolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sns.CfnTopicInlinePolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTopicInlinePolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTopicInlinePolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_sns.CfnTopicInlinePolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTopicInlinePolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_sns.CfnTopicInlinePolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTopicInlinePolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTopicInlinePolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

