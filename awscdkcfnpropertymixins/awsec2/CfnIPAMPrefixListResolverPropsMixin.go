package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::EC2::IPAMPrefixListResolver.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnIPAMPrefixListResolverPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnIPAMPrefixListResolverPropsMixin(&CfnIPAMPrefixListResolverMixinProps{
//   	AddressFamily: jsii.String("addressFamily"),
//   	Description: jsii.String("description"),
//   	IpamId: jsii.String("ipamId"),
//   	Rules: []interface{}{
//   		&IpamPrefixListResolverRuleProperty{
//   			Conditions: []interface{}{
//   				&IpamPrefixListResolverRuleConditionProperty{
//   					Cidr: jsii.String("cidr"),
//   					IpamPoolId: jsii.String("ipamPoolId"),
//   					Operation: jsii.String("operation"),
//   					ResourceId: jsii.String("resourceId"),
//   					ResourceOwner: jsii.String("resourceOwner"),
//   					ResourceRegion: jsii.String("resourceRegion"),
//   					ResourceTag: &CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			IpamScopeId: jsii.String("ipamScopeId"),
//   			ResourceType: jsii.String("resourceType"),
//   			RuleType: jsii.String("ruleType"),
//   			StaticCidr: jsii.String("staticCidr"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamprefixlistresolver.html
//
type CfnIPAMPrefixListResolverPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnIPAMPrefixListResolverMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIPAMPrefixListResolverPropsMixin
type jsiiProxy_CfnIPAMPrefixListResolverPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnIPAMPrefixListResolverPropsMixin) Props() *CfnIPAMPrefixListResolverMixinProps {
	var returns *CfnIPAMPrefixListResolverMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPrefixListResolverPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::IPAMPrefixListResolver`.
func NewCfnIPAMPrefixListResolverPropsMixin(props *CfnIPAMPrefixListResolverMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnIPAMPrefixListResolverPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIPAMPrefixListResolverPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIPAMPrefixListResolverPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnIPAMPrefixListResolverPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::IPAMPrefixListResolver`.
func NewCfnIPAMPrefixListResolverPropsMixin_Override(c CfnIPAMPrefixListResolverPropsMixin, props *CfnIPAMPrefixListResolverMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnIPAMPrefixListResolverPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnIPAMPrefixListResolverPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIPAMPrefixListResolverPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnIPAMPrefixListResolverPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIPAMPrefixListResolverPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnIPAMPrefixListResolverPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIPAMPrefixListResolverPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIPAMPrefixListResolverPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

