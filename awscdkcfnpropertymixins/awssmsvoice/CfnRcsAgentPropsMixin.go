package awssmsvoice

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AWS RCS Agent in AWS End User Messaging SMS.
//
// The RCS agent is the top-level resource representing your brand for RCS messaging, and serves as an origination identity for sending RCS messages. A newly created agent is in CREATED status; it advances through TESTING, PARTIAL and ACTIVE only as testing and country launch registrations are submitted and approved by carriers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnRcsAgentPropsMixin := awscdkcfnpropertymixins.Aws_smsvoice.NewCfnRcsAgentPropsMixin(&CfnRcsAgentMixinProps{
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	OptOutListName: jsii.String("optOutListName"),
//   	SelfManagedOptOutsEnabled: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TwoWayChannelArn: jsii.String("twoWayChannelArn"),
//   	TwoWayChannelRole: jsii.String("twoWayChannelRole"),
//   	TwoWayEnabled: jsii.Boolean(false),
//   	TwoWayMediaS3BucketName: jsii.String("twoWayMediaS3BucketName"),
//   	TwoWayMediaS3KeyPrefix: jsii.String("twoWayMediaS3KeyPrefix"),
//   	TwoWayMediaS3Role: jsii.String("twoWayMediaS3Role"),
//   	TwoWayRcsEventsEnabled: []*string{
//   		jsii.String("twoWayRcsEventsEnabled"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html
//
type CfnRcsAgentPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRcsAgentMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRcsAgentPropsMixin
type jsiiProxy_CfnRcsAgentPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRcsAgentPropsMixin) Props() *CfnRcsAgentMixinProps {
	var returns *CfnRcsAgentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgentPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SMSVOICE::RcsAgent`.
func NewCfnRcsAgentPropsMixin(props *CfnRcsAgentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRcsAgentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRcsAgentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRcsAgentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_smsvoice.CfnRcsAgentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SMSVOICE::RcsAgent`.
func NewCfnRcsAgentPropsMixin_Override(c CfnRcsAgentPropsMixin, props *CfnRcsAgentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_smsvoice.CfnRcsAgentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRcsAgentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRcsAgentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_smsvoice.CfnRcsAgentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRcsAgentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_smsvoice.CfnRcsAgentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRcsAgentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRcsAgentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

