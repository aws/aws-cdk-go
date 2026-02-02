package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds an inbound (ingress) rule to a security group.
//
// An inbound rule permits instances to receive traffic from the specified IPv4 or IPv6 address range, the IP addresses that are specified by a prefix list, or the instances that are associated with a source security group. For more information, see [Security group rules](https://docs.aws.amazon.com/vpc/latest/userguide/security-group-rules.html) .
//
// You must specify exactly one of the following sources: an IPv4 address range, an IPv6 address range, a prefix list, or a security group.
//
// You must specify a protocol for each rule (for example, TCP). If the protocol is TCP or UDP, you must also specify a port or port range. If the protocol is ICMP or ICMPv6, you must also specify the ICMP/ICMPv6 type and code.
//
// Rule changes are propagated to instances associated with the security group as quickly as possible. However, a small delay might occur.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSecurityGroupIngressPropsMixin := awscdkmixinspreview.Mixins.NewCfnSecurityGroupIngressPropsMixin(&CfnSecurityGroupIngressMixinProps{
//   	CidrIp: jsii.String("cidrIp"),
//   	CidrIpv6: jsii.String("cidrIpv6"),
//   	Description: jsii.String("description"),
//   	FromPort: jsii.Number(123),
//   	GroupId: jsii.String("groupId"),
//   	GroupName: jsii.String("groupName"),
//   	IpProtocol: jsii.String("ipProtocol"),
//   	SourcePrefixListId: jsii.String("sourcePrefixListId"),
//   	SourceSecurityGroupId: jsii.String("sourceSecurityGroupId"),
//   	SourceSecurityGroupName: jsii.String("sourceSecurityGroupName"),
//   	SourceSecurityGroupOwnerId: jsii.String("sourceSecurityGroupOwnerId"),
//   	ToPort: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html
//
type CfnSecurityGroupIngressPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSecurityGroupIngressMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSecurityGroupIngressPropsMixin
type jsiiProxy_CfnSecurityGroupIngressPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSecurityGroupIngressPropsMixin) Props() *CfnSecurityGroupIngressMixinProps {
	var returns *CfnSecurityGroupIngressMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityGroupIngressPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::SecurityGroupIngress`.
func NewCfnSecurityGroupIngressPropsMixin(props *CfnSecurityGroupIngressMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSecurityGroupIngressPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSecurityGroupIngressPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSecurityGroupIngressPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSecurityGroupIngressPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::SecurityGroupIngress`.
func NewCfnSecurityGroupIngressPropsMixin_Override(c CfnSecurityGroupIngressPropsMixin, props *CfnSecurityGroupIngressMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSecurityGroupIngressPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSecurityGroupIngressPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSecurityGroupIngressPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSecurityGroupIngressPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecurityGroupIngressPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSecurityGroupIngressPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSecurityGroupIngressPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSecurityGroupIngressPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

