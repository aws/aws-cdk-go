package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an entry in a transit gateway metering policy to define traffic measurement rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransitGatewayMeteringPolicyEntryPropsMixin := awscdkmixinspreview.Mixins.NewCfnTransitGatewayMeteringPolicyEntryPropsMixin(&CfnTransitGatewayMeteringPolicyEntryMixinProps{
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	DestinationPortRange: jsii.String("destinationPortRange"),
//   	DestinationTransitGatewayAttachmentId: jsii.String("destinationTransitGatewayAttachmentId"),
//   	DestinationTransitGatewayAttachmentType: jsii.String("destinationTransitGatewayAttachmentType"),
//   	MeteredAccount: jsii.String("meteredAccount"),
//   	PolicyRuleNumber: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	SourceCidrBlock: jsii.String("sourceCidrBlock"),
//   	SourcePortRange: jsii.String("sourcePortRange"),
//   	SourceTransitGatewayAttachmentId: jsii.String("sourceTransitGatewayAttachmentId"),
//   	SourceTransitGatewayAttachmentType: jsii.String("sourceTransitGatewayAttachmentType"),
//   	TransitGatewayMeteringPolicyId: jsii.String("transitGatewayMeteringPolicyId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymeteringpolicyentry.html
//
type CfnTransitGatewayMeteringPolicyEntryPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTransitGatewayMeteringPolicyEntryMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTransitGatewayMeteringPolicyEntryPropsMixin
type jsiiProxy_CfnTransitGatewayMeteringPolicyEntryPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntryPropsMixin) Props() *CfnTransitGatewayMeteringPolicyEntryMixinProps {
	var returns *CfnTransitGatewayMeteringPolicyEntryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMeteringPolicyEntryPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::TransitGatewayMeteringPolicyEntry`.
func NewCfnTransitGatewayMeteringPolicyEntryPropsMixin(props *CfnTransitGatewayMeteringPolicyEntryMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTransitGatewayMeteringPolicyEntryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTransitGatewayMeteringPolicyEntryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransitGatewayMeteringPolicyEntryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnTransitGatewayMeteringPolicyEntryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::TransitGatewayMeteringPolicyEntry`.
func NewCfnTransitGatewayMeteringPolicyEntryPropsMixin_Override(c CfnTransitGatewayMeteringPolicyEntryPropsMixin, props *CfnTransitGatewayMeteringPolicyEntryMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnTransitGatewayMeteringPolicyEntryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTransitGatewayMeteringPolicyEntryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransitGatewayMeteringPolicyEntryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnTransitGatewayMeteringPolicyEntryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransitGatewayMeteringPolicyEntryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnTransitGatewayMeteringPolicyEntryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntryPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMeteringPolicyEntryPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

