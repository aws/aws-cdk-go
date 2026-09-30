package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Route53::TrafficPolicyInstance.
//
// Applies a version of a Route 53 traffic policy to a DNS name in a public hosted zone, and Route 53 creates and owns the resource record sets that answer queries for that name.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTrafficPolicyInstancePropsMixin := awscdkcfnpropertymixins.Aws_route53.NewCfnTrafficPolicyInstancePropsMixin(&CfnTrafficPolicyInstanceMixinProps{
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	Name: jsii.String("name"),
//   	TrafficPolicyId: jsii.String("trafficPolicyId"),
//   	TrafficPolicyVersion: jsii.Number(123),
//   	Ttl: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicyinstance.html
//
type CfnTrafficPolicyInstancePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTrafficPolicyInstanceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTrafficPolicyInstancePropsMixin
type jsiiProxy_CfnTrafficPolicyInstancePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTrafficPolicyInstancePropsMixin) Props() *CfnTrafficPolicyInstanceMixinProps {
	var returns *CfnTrafficPolicyInstanceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrafficPolicyInstancePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53::TrafficPolicyInstance`.
func NewCfnTrafficPolicyInstancePropsMixin(props *CfnTrafficPolicyInstanceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTrafficPolicyInstancePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTrafficPolicyInstancePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTrafficPolicyInstancePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnTrafficPolicyInstancePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53::TrafficPolicyInstance`.
func NewCfnTrafficPolicyInstancePropsMixin_Override(c CfnTrafficPolicyInstancePropsMixin, props *CfnTrafficPolicyInstanceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnTrafficPolicyInstancePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTrafficPolicyInstancePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTrafficPolicyInstancePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnTrafficPolicyInstancePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTrafficPolicyInstancePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnTrafficPolicyInstancePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTrafficPolicyInstancePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTrafficPolicyInstancePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

