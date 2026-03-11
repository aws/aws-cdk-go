package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a new pool of dedicated IP addresses.
//
// A pool can include one or more dedicated IP addresses that are associated with your AWS account . You can associate a pool with a configuration set. When you send an email that uses that configuration set, the message is sent from one of the addresses in the associated pool.
//
// > You can't delete dedicated IP pools that have a `STANDARD` scaling mode with one or more dedicated IP addresses. This constraint doesn't apply to dedicated IP pools that have a `MANAGED` scaling mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDedicatedIpPoolPropsMixin := awscdkcfnpropertymixins.Aws_ses.NewCfnDedicatedIpPoolPropsMixin(&CfnDedicatedIpPoolMixinProps{
//   	PoolName: jsii.String("poolName"),
//   	ScalingMode: jsii.String("scalingMode"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-dedicatedippool.html
//
type CfnDedicatedIpPoolPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDedicatedIpPoolMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDedicatedIpPoolPropsMixin
type jsiiProxy_CfnDedicatedIpPoolPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDedicatedIpPoolPropsMixin) Props() *CfnDedicatedIpPoolMixinProps {
	var returns *CfnDedicatedIpPoolMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDedicatedIpPoolPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SES::DedicatedIpPool`.
func NewCfnDedicatedIpPoolPropsMixin(props *CfnDedicatedIpPoolMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDedicatedIpPoolPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDedicatedIpPoolPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDedicatedIpPoolPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnDedicatedIpPoolPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SES::DedicatedIpPool`.
func NewCfnDedicatedIpPoolPropsMixin_Override(c CfnDedicatedIpPoolPropsMixin, props *CfnDedicatedIpPoolMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnDedicatedIpPoolPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDedicatedIpPoolPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDedicatedIpPoolPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnDedicatedIpPoolPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDedicatedIpPoolPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnDedicatedIpPoolPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDedicatedIpPoolPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDedicatedIpPoolPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

