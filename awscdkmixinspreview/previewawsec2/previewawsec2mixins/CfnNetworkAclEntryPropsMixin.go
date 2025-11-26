package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an entry, known as a rule, in a network ACL with a rule number you specify.
//
// Each network ACL has a set of numbered ingress rules and a separate set of numbered egress rules.
//
// To create the network ACL, see [AWS::EC2::NetworkAcl](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkacl.html) .
//
// For information about the protocol value, see [Protocol Numbers](https://docs.aws.amazon.com/https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNetworkAclEntryPropsMixin := awscdkmixinspreview.Mixins.NewCfnNetworkAclEntryPropsMixin(&CfnNetworkAclEntryMixinProps{
//   	CidrBlock: jsii.String("cidrBlock"),
//   	Egress: jsii.Boolean(false),
//   	Icmp: &IcmpProperty{
//   		Code: jsii.Number(123),
//   		Type: jsii.Number(123),
//   	},
//   	Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   	NetworkAclId: jsii.String("networkAclId"),
//   	PortRange: &PortRangeProperty{
//   		From: jsii.Number(123),
//   		To: jsii.Number(123),
//   	},
//   	Protocol: jsii.Number(123),
//   	RuleAction: jsii.String("ruleAction"),
//   	RuleNumber: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkaclentry.html
//
type CfnNetworkAclEntryPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnNetworkAclEntryMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNetworkAclEntryPropsMixin
type jsiiProxy_CfnNetworkAclEntryPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnNetworkAclEntryPropsMixin) Props() *CfnNetworkAclEntryMixinProps {
	var returns *CfnNetworkAclEntryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNetworkAclEntryPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::NetworkAclEntry`.
func NewCfnNetworkAclEntryPropsMixin(props *CfnNetworkAclEntryMixinProps, options *mixins.CfnPropertyMixinOptions) CfnNetworkAclEntryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNetworkAclEntryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNetworkAclEntryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkAclEntryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::NetworkAclEntry`.
func NewCfnNetworkAclEntryPropsMixin_Override(c CfnNetworkAclEntryPropsMixin, props *CfnNetworkAclEntryMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkAclEntryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnNetworkAclEntryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNetworkAclEntryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkAclEntryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNetworkAclEntryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNetworkAclEntryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNetworkAclEntryPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnNetworkAclEntryPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

