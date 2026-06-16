package awsrtbfabric

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsrtbfabric/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::RTBFabric::LinkRoutingRule.
//
// A routing rule on a link within RTB Fabric that controls request routing based on conditions such as host headers, path matching, and query string parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnLinkRoutingRulePropsMixin := awscdkcfnpropertymixins.Aws_rtbfabric.NewCfnLinkRoutingRulePropsMixin(&CfnLinkRoutingRuleMixinProps{
//   	Conditions: &RuleConditionProperty{
//   		HostHeader: jsii.String("hostHeader"),
//   		HostHeaderWildcard: jsii.String("hostHeaderWildcard"),
//   		PathExact: jsii.String("pathExact"),
//   		PathPrefix: jsii.String("pathPrefix"),
//   		QueryStringEquals: &QueryStringKeyValuePairProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   		QueryStringExists: jsii.String("queryStringExists"),
//   	},
//   	GatewayId: jsii.String("gatewayId"),
//   	LinkId: jsii.String("linkId"),
//   	Priority: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-linkroutingrule.html
//
type CfnLinkRoutingRulePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLinkRoutingRuleMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLinkRoutingRulePropsMixin
type jsiiProxy_CfnLinkRoutingRulePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLinkRoutingRulePropsMixin) Props() *CfnLinkRoutingRuleMixinProps {
	var returns *CfnLinkRoutingRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLinkRoutingRulePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RTBFabric::LinkRoutingRule`.
func NewCfnLinkRoutingRulePropsMixin(props *CfnLinkRoutingRuleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLinkRoutingRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLinkRoutingRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLinkRoutingRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkRoutingRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RTBFabric::LinkRoutingRule`.
func NewCfnLinkRoutingRulePropsMixin_Override(c CfnLinkRoutingRulePropsMixin, props *CfnLinkRoutingRuleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkRoutingRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLinkRoutingRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLinkRoutingRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkRoutingRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLinkRoutingRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_rtbfabric.CfnLinkRoutingRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLinkRoutingRulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLinkRoutingRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

