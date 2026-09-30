package awslightsail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Lightsail::KeyPair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnKeyPairPropsMixin := awscdkcfnpropertymixins.Aws_lightsail.NewCfnKeyPairPropsMixin(&CfnKeyPairMixinProps{
//   	KeyPairName: jsii.String("keyPairName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-keypair.html
//
type CfnKeyPairPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnKeyPairMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnKeyPairPropsMixin
type jsiiProxy_CfnKeyPairPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnKeyPairPropsMixin) Props() *CfnKeyPairMixinProps {
	var returns *CfnKeyPairMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyPairPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lightsail::KeyPair`.
func NewCfnKeyPairPropsMixin(props *CfnKeyPairMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnKeyPairPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnKeyPairPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnKeyPairPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnKeyPairPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lightsail::KeyPair`.
func NewCfnKeyPairPropsMixin_Override(c CfnKeyPairPropsMixin, props *CfnKeyPairMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnKeyPairPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnKeyPairPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnKeyPairPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnKeyPairPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnKeyPairPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnKeyPairPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnKeyPairPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnKeyPairPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

