package previewawssesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsses/previewawssesmixins/internal"
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
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDedicatedIpPoolPropsMixin := awscdkmixinspreview.Mixins.NewCfnDedicatedIpPoolPropsMixin(&CfnDedicatedIpPoolMixinProps{
//   	PoolName: jsii.String("poolName"),
//   	ScalingMode: jsii.String("scalingMode"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-dedicatedippool.html
//
type CfnDedicatedIpPoolPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDedicatedIpPoolMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDedicatedIpPoolPropsMixin
type jsiiProxy_CfnDedicatedIpPoolPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
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

func (j *jsiiProxy_CfnDedicatedIpPoolPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SES::DedicatedIpPool`.
func NewCfnDedicatedIpPoolPropsMixin(props *CfnDedicatedIpPoolMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDedicatedIpPoolPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDedicatedIpPoolPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDedicatedIpPoolPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnDedicatedIpPoolPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SES::DedicatedIpPool`.
func NewCfnDedicatedIpPoolPropsMixin_Override(c CfnDedicatedIpPoolPropsMixin, props *CfnDedicatedIpPoolMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnDedicatedIpPoolPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDedicatedIpPoolPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDedicatedIpPoolPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnDedicatedIpPoolPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnDedicatedIpPoolPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDedicatedIpPoolPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

