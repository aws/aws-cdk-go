package awsoam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsoam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates a *sink* in the current account, so that it can be used as a monitoring account in CloudWatch cross-account observability.
//
// A sink is a resource that represents an attachment point in a monitoring account, which source accounts can link to to be able to send observability data.
//
// After you create a sink, you must create a sink policy that allows source accounts to attach to it. For more information, see [PutSinkPolicy](https://docs.aws.amazon.com/OAM/latest/APIReference/API_PutSinkPolicy.html) .
//
// An account can have one sink.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var policy interface{}
//
//   cfnSinkPropsMixin := awscdkcfnpropertymixins.Aws_oam.NewCfnSinkPropsMixin(&CfnSinkMixinProps{
//   	Name: jsii.String("name"),
//   	Policy: policy,
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-oam-sink.html
//
type CfnSinkPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnSinkMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSinkPropsMixin
type jsiiProxy_CfnSinkPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnSinkPropsMixin) Props() *CfnSinkMixinProps {
	var returns *CfnSinkMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSinkPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Oam::Sink`.
func NewCfnSinkPropsMixin(props *CfnSinkMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnSinkPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSinkPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSinkPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_oam.CfnSinkPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Oam::Sink`.
func NewCfnSinkPropsMixin_Override(c CfnSinkPropsMixin, props *CfnSinkMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_oam.CfnSinkPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnSinkPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSinkPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_oam.CfnSinkPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSinkPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_oam.CfnSinkPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSinkPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSinkPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

