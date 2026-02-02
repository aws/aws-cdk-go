package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a security group.
//
// You must specify ingress rules to allow inbound traffic. By default, no inbound traffic is allowed.
//
// When you create a security group, if you do not add egress rules, we add egress rules that allow all outbound IPv4 and IPv6 traffic. Otherwise, we do not add them. After the security group is created, if you remove all egress rules that you added, we do not add egress rules, so no outbound traffic is allowed.
//
// If you modify a rule, CloudFormation removes the existing rule and then adds a new rule. There is a brief period when neither the original rule or the new rule exists, so the corresponding traffic is dropped.
//
// This type supports updates. For more information about updating stacks, see [AWS CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html) .
//
// > To cross-reference two security groups in the ingress and egress rules of those security groups, use the [AWS::EC2::SecurityGroupEgress](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html) and [AWS::EC2::SecurityGroupIngress](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-ingress.html) resources to define your rules. Do not use the embedded ingress and egress rules in the `AWS::EC2::SecurityGroup` . Doing so creates a circular dependency, which CloudFormation doesn't allow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSecurityGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnSecurityGroupPropsMixin(&CfnSecurityGroupMixinProps{
//   	GroupDescription: jsii.String("groupDescription"),
//   	GroupName: jsii.String("groupName"),
//   	SecurityGroupEgress: []interface{}{
//   		&EgressProperty{
//   			CidrIp: jsii.String("cidrIp"),
//   			CidrIpv6: jsii.String("cidrIpv6"),
//   			Description: jsii.String("description"),
//   			DestinationPrefixListId: jsii.String("destinationPrefixListId"),
//   			DestinationSecurityGroupId: jsii.String("destinationSecurityGroupId"),
//   			FromPort: jsii.Number(123),
//   			IpProtocol: jsii.String("ipProtocol"),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	SecurityGroupIngress: []interface{}{
//   		&IngressProperty{
//   			CidrIp: jsii.String("cidrIp"),
//   			CidrIpv6: jsii.String("cidrIpv6"),
//   			Description: jsii.String("description"),
//   			FromPort: jsii.Number(123),
//   			IpProtocol: jsii.String("ipProtocol"),
//   			SourcePrefixListId: jsii.String("sourcePrefixListId"),
//   			SourceSecurityGroupId: jsii.String("sourceSecurityGroupId"),
//   			SourceSecurityGroupName: jsii.String("sourceSecurityGroupName"),
//   			SourceSecurityGroupOwnerId: jsii.String("sourceSecurityGroupOwnerId"),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroup.html
//
type CfnSecurityGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSecurityGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSecurityGroupPropsMixin
type jsiiProxy_CfnSecurityGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSecurityGroupPropsMixin) Props() *CfnSecurityGroupMixinProps {
	var returns *CfnSecurityGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::SecurityGroup`.
func NewCfnSecurityGroupPropsMixin(props *CfnSecurityGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSecurityGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSecurityGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSecurityGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSecurityGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::SecurityGroup`.
func NewCfnSecurityGroupPropsMixin_Override(c CfnSecurityGroupPropsMixin, props *CfnSecurityGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSecurityGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSecurityGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSecurityGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSecurityGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecurityGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnSecurityGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSecurityGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSecurityGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

